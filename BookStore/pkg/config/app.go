package config

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
dsn := "shivam:1234@tcp(localhost:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"

    d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("❌ Failed to connect to the database:", err)
        panic(err)
    }

    db = d
    fmt.Println("✅ Database connection established successfully.")
}

func GetDB() *gorm.DB {
    return db
}
