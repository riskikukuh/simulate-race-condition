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

func InitJob(userService service.UserService, walletService service.WalletService) {
	// Add 5 data
	users := []models.UserRequest{
		{
			Name:  "user1",
			Email: "user1@app.com",
		},
		{
			Name:  "user2",
			Email: "user2@app.com",
		},
		{
			Name:  "user3",
			Email: "user3@app.com",
		},
		{
			Name:  "user4",
			Email: "user4@app.com",
		},
		{
			Name:  "user5",
			Email: "user5@app.com",
		},
	}

	for _, v := range users {
		user, err := userService.Create(v)
		if err != nil {
			panic(err)
		}

		_, err = walletService.Create(models.WalletRequest{
			UserId:  user.ID,
			Balance: 10_000_000,
		})
		if err != nil {
			panic(err)
		}
	}
}

func ClearData(userService service.UserService, walletService service.WalletService) {
	slog.Info("Clearing data")

	slog.Info("Clearing wallets")
	wallets, err := walletService.FindAll()
	if err != nil {
		panic(err)
	}

	for _, wallet := range wallets {
		walletService.DeleteByWalletId(wallet.ID)
	}

	slog.Info("Clearing users")
	users, err := userService.FindAll()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		userService.DeleteByUserId(user.ID)
	}

}

func StartBalanceJob(env *config.EnvironmentVariable, wrapDb *database.WrapDB) {
	slog.Info("Start Balance Job")
	// Create instance of service
	userService := service.NewUserServiceImpl(wrapDb.PostgreDB)
	walletService := service.NewWalletServiceImpl(wrapDb.PostgreDB)

	// Clear all user and wallet
	ClearData(userService, walletService)

	// Create 5 data user and wallet
	InitJob(userService, walletService)

	// Get all wallet
	wallets, err := walletService.FindAll()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	// Start the Goroutine to run Worker that simulate race condition
	for i := 0; i < 10_000; i++ {
		wg.Add(1)
		go func(wgOuter *sync.WaitGroup) {
			defer wgOuter.Done()

			var wgInner sync.WaitGroup
			wgInner.Add(5)

			go BalanceWorker(&wgInner, int(wallets[0].ID), userService, walletService)
			go BalanceWorker(&wgInner, int(wallets[1].ID), userService, walletService)
			go BalanceWorker(&wgInner, int(wallets[2].ID), userService, walletService)
			go BalanceWorker(&wgInner, int(wallets[3].ID), userService, walletService)
			go BalanceWorker(&wgInner, int(wallets[4].ID), userService, walletService)

			wgInner.Wait()
		}(&wg)
	}
	wg.Wait()

	// Finish
	slog.Info("Finish Balance Job")
}

func BalanceWorker(wg *sync.WaitGroup, walletId int, userService service.UserService, walletService service.WalletService) {
	defer wg.Done()

	wallet, err := walletService.FindById(int64(walletId))
	if err != nil {
		panic(err)
	}

	wallet.Balance = wallet.Balance - 1_000_000

	wallet, err = walletService.Update(wallet)
	if err != nil {
		panic(err)
	}

	wallet.Balance = wallet.Balance - 500_000
	wallet, err = walletService.Update(wallet)
	if err != nil {
		panic(err)
	}

	wallet.Balance = wallet.Balance + 750_000
	wallet, err = walletService.Update(wallet)
	if err != nil {
		panic(err)
	}

	wallet.Balance = wallet.Balance - 5_000_000
	wallet, err = walletService.Update(wallet)
	if err != nil {
		panic(err)
	}

	wallet.Balance = wallet.Balance - 3_500_000
	wallet, err = walletService.Update(wallet)
	if err != nil {
		panic(err)
	}

	wallet.Balance = wallet.Balance + 250_000
	_, err = walletService.Update(wallet)
	if err != nil {
		panic(err)
	}

	fmt.Println("Wallet ", walletId, " Done")

}
