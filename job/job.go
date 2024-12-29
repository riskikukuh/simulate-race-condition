package job

import (
	"fmt"
	"log/slog"
	"simulation-race-condition/config"
	"simulation-race-condition/database"
	"simulation-race-condition/models"
	"simulation-race-condition/service"
	"sync"
)

func StartJob(env *config.EnvironmentVariable, wrapDb *database.WrapDB) {
	slog.Info("Starting Job")
	userService := service.NewUserServiceImpl(wrapDb.PostgreDB)

	var wg sync.WaitGroup

	for i := 0; i < 1_00; i++ {
		wg.Add(1)

		go Worker(&wg, i, userService)
	}
	wg.Wait()
}

func Worker(wg *sync.WaitGroup, groupN int, service service.UserService) {
	defer wg.Done()

	for i := range 1_00 {
		var dummyUser = models.UserRequest{
			Name:  fmt.Sprintf("User-%d-%d", groupN, i),
			Email: fmt.Sprintf("User-%d-%d@app.com", groupN, i),
		}

		user, err := service.Create(dummyUser)
		if err != nil {
			panic(err)
		}

		fmt.Printf("User ID %d: created", user.ID)

		user.Name = fmt.Sprintf("Updated-user-%d-%d", groupN, i)
		user.Email = fmt.Sprintf("Updated-email-%d-%d@app.com", groupN, i)

		user, err = service.Update(user)
		if err != nil {
			panic(err)
		}

		findUser, err := service.FindById(user.ID)
		if err != nil {
			panic(err)
		}

		err = service.DeleteByUserId(findUser.ID)
		if err != nil {
			panic(err)
		}
	}
}
