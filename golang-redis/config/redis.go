package config

import (
	"log"
	"strconv"
)

type RedisCfg struct {
	Address string
	Password string
	Database int
}

func LoadRedisCfg() *RedisCfg {
	databaseStr := GetEnv("DATABASE")
	database, err := strconv.Atoi(databaseStr)
	if err != nil {
		log.Fatalln("database environmetn variable must be a number")
	}

	return &RedisCfg{
		Address:  GetEnv("ADDRESS"),
		Password: GetEnv("PASSWORD"),
		Database: database,
	}
}

