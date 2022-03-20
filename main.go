package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Note struct {
	Id      int    `json:"id,omitempty" gorm:"column:id;"`
	Title   string `json:"title,omitempty" gorm:"column:title;"`
	Content string `json:"content,omitempty" gorm:"column:content;"`
}

func (Note) TableName() string {
	return "notes"
}

func main() {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := os.Getenv("DBConnectionString")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	dbConnectionString := os.Getenv("DB_Connection_String")

	db, err := gorm.Open(mysql.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error open database", err)
	}

	// newNote := Note{Title: "Demo note", Content:  "This is content of demo not"}

	// if result := db.Create(&newNote); result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	var notes []Note
	db.Where("status = 1").Find(&notes)
	fmt.Println(notes)

	db.Table(Note{}.TableName()).Where("id = 3").Delete(nil)

}
