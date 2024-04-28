package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	JWTSecret  string
	BcryptSalt int
	Postgres   Postgres
	AWSConfig  AWSConfig
}

type Postgres struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	Params   string
}

type AWSConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
	Region          string
}

func New() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("[config-file-fail-load] \n", err.Error())
	}

	v := viper.GetViper()
	viper.AutomaticEnv()

	return &Config{
		JWTSecret:  v.GetString("JWT_SECRET"),
		BcryptSalt: v.GetInt("BCRYPT_SALT"),
		Postgres: Postgres{
			Host:     v.GetString("DB_HOST"),
			Port:     v.GetString("DB_PORT"),
			Database: v.GetString("DB_NAME"),
			User:     v.GetString("DB_USERNAME"),
			Password: v.GetString("DB_PASSWORD"),
			Params:   v.GetString("DB_PARAMS"),
		},
		AWSConfig: AWSConfig{
			AccessKeyID:     v.GetString("S3_ID"),
			AccessKeySecret: v.GetString("S3_SECRET_KEY"),
			Bucket:          v.GetString("S3_BUCKET_NAME"),
			Region:          v.GetString("S3_REGION"),
		},
	}
}
