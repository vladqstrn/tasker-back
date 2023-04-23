package database

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/vladqstrn/tasker-back/internal/config"
	"github.com/vladqstrn/tasker-back/internal/models"
)

func InitDB() *pg.DB {

	options := &pg.Options{
		User:     config.User,
		Password: config.Password,
		Addr:     config.Host + ":" + config.Port,
		Database: config.DbName,
	}

	db := pg.Connect(options)

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to database")
	return db
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{&models.Task{}} {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
