package main

import (
    "log"
    "os"

    "ecommerce/database"
    "ecommerce/router"
    
    "github.com/joho/godotenv"
)

func init() {
    if _, err := os.Stat(".env"); err == nil {
        err = godotenv.Load(".env")
        if err != nil {
            log.Fatal("Error loading .env config file")
        }
        log.Println("Successfully loaded the config file")
    } else {
        log.Println(".env file not found, skipping load")
    }
}

func main() {
    database.ConnectDb()

    router.ClientRoutes()
}