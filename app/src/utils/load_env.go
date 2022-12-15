package utils

import "github.com/joho/godotenv"

func LoadEnv() {
	// Run migration commands here just like php artisan migrate
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
}
