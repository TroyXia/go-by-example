package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Animal struct {
	Name string `gorm:"column:name;primary_key"`
	Age  int    `gorm:"column:age"`
}

func (p *Animal) TableName() string {
	return "animal"
}

var (
	host     = pflag.StringP("host", "H", "127.0.0.1:3306", "MySQL service host address")
	username = pflag.StringP("username", "u", "root", "Username for access to mysql service")
	password = pflag.StringP("password", "p", "Troy@0403", "Password for access to mysql, should be used pair with password")
	database = pflag.StringP("database", "d", "test", "Database name to use")
	help     = pflag.BoolP("help", "h", false, "Print this help message")
)

func main() {
	// Parse command line flags
	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s", *username, *password, *host, *database, true, "Local")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 1. Auto migration the given models
	if err := db.AutoMigrate(&Animal{}); err != nil {
		panic("failed to auto migrate")
	}

	// 2. Insert the value into database
	if err := db.Create(&Animal{Name: "car", Age: 3}).Error; err != nil {
		log.Fatalf("Failed to create product: %v", err)
	}
	PrintProducts(db)

	// 3. Find first record that match given conditions
	animal := &Animal{}
	if err := db.Where("name= ?", "monkey").First(animal).Error; err != nil {
		log.Fatalf("Failed to find product: %v", err)
	}

	// 4. Update value in database, if the value doesn't have primary key, will insert it
	animal.Age = 4
	if err := db.Save(animal).Error; err != nil {
		log.Fatalf("Failed to update product: %v", err)
	}
	PrintProducts(db)

	// 5. Delete value match given conditions
	if err := db.Where("name = ?", "monkey").Delete(&Animal{}).Error; err != nil {
		log.Fatalf("Delete animal error: %v", err)
	}
	PrintProducts(db)
}

func PrintProducts(db *gorm.DB) {
	animals := make([]*Animal, 0)
	var count int64
	d := db.Where("name like ?", "%mon%").Offset(0).Limit(2).Order("name desc").Find(&animals).Offset(-1).Limit(-1).Count(&count)
	if d.Error != nil {
		log.Fatalf("List products error: %v", d.Error)
	}

	log.Printf("totalcount: %d", count)
	for _, animal := range animals {
		log.Printf("\tname: %s, age: %d\n", animal.Name, animal.Age)
	}
}
