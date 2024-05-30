package initializers

import (
    "database/sql"
    "fmt"
    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
    "os"
    "strconv"
)

var DB *sql.DB

func ConnectDatabse() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }

    host := os.Getenv("HOST")
    port, _ := strconv.Atoi(os.Getenv("PORT"))
    user := os.Getenv("USER")
    password := os.Getenv("PASSWORD")

    psqlSetup := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, "stock")
    db, errSql := sql.Open("postgres", psqlSetup)
    if errSql != nil {
        fmt.Println("An error occured while connecting to the database", err)
        panic((err))
    } else {
        DB = db
        fmt.Println("Successfully connected to the database")
    }
}
