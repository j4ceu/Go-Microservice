package init

import (
	"Go-Microservice/app/employee/models"
	"Go-Microservice/app/user/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDatabase() *gorm.DB {
	postgresConfig := config.LoadPostgresConfig()

	log.Println(postgresConfig)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		postgresConfig.Host,
		postgresConfig.User,
		postgresConfig.Password,
		postgresConfig.Name,
		postgresConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	tx := db.Session(&gorm.Session{PrepareStmt: true})

	if err != nil {
		log.Fatal("Koneksi DB Gagal")
	}

	migrateDatabase(tx)

	DB = tx

	// initJobs()

	return tx

}

func migrateDatabase(db *gorm.DB) {

	errMigrate := db.AutoMigrate(&models.Employee{})

	if errMigrate != nil {
		log.Fatal("Gagal Migrate")
	}

	log.Println("Migrate Berhasil!")

}
