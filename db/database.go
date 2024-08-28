package db

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type MySQL struct {
	*gorm.DB
}

type User struct {
	UserID   uuid.UUID `gorm:"primaryKey;not null"`
	Username string    `gorm:"unique;not null"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
}

type Database interface {
	GetUser(email string) (*User, error)
	CreateUser(user User) error
}

func NewMySQL(username, password, host, port, dbName string) Database {
	var err error
	var count int64

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port)
	initDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connection,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})

	initDB.Raw("SELECT count(*) FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = ?", dbName).Scan(&count)
	if count <= 0 {
		if err := initDB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName).Error; err != nil {
			panic("Error creating database: " + err.Error())
		}
	}

	connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connection,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	DB.AutoMigrate(User{})

	return &MySQL{DB}
}

func (db *MySQL) GetUser(email string) (*User, error) {
	var userData User
	err := db.DB.Table("users").Where("email = ?", email).First(&userData).Error
	if err != nil {
		return nil, err
	}
	return &userData, nil
}

func (db *MySQL) CreateUser(user User) error {
	err := db.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
