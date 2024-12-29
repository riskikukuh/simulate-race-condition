Jika tidak tahu maka cari tahulah. Jika takut, maka cobalah

Cara mengetahui Race Condition
- Buat 5 User dan 5 Wallet
- Nominal awal 10.000.000
- Seluruh wallet akan menjalankan proses dibawah ini
- Dikurangi 1.000.000 = SISA 9.000.000
- Dikurangi 500.000 = SISA 8.500.000
- Ditambah 750.000 = SISA 9.250.000
- Dikurangi 5.000.000 = SISA 4.250.000
- Dikurangi 3.500.000 = SISA 750.000
- Ditambah 250.000 = SISA 1.000.000
- EKSPEKSTASINYA, seluruh wallet harus bernilai 1.000.000