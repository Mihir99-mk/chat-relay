package config

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"bot/ent/entgen"

	"entgo.io/ent/dialect/sql"
)

var entOnce sync.Once
var entClient *entgen.Client

func GetDB(env IEnv) *entgen.Client {
	entOnce.Do(func() {
		var err error
		entClient, err = InitEntDb(env)
		if err != nil {
			log.Fatalf("Error initializing database: %s", err)
		}
		log.Println("database is up!!")
	})

	return entClient
}

func InitEntDb(env IEnv) (*entgen.Client, error) {
	entDriver, err := GetSqlDriver(env)
	if err != nil {
		return nil, err
	}
	options := []entgen.Option{entgen.Driver(entDriver)}

	// Create the Ent client using the dialect
	client := entgen.NewClient(options...)
	return client, nil
}

func GetConnectionString(env IEnv) (string, error) {
	if env.GetDBUsername() == "" || env.GetDBPassword() == "" || env.GetDBHost() == "" || env.GetDBPort() == "" || env.GetDBName() == "" {
		return "", fmt.Errorf("one or more required DB environment variables are missing")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=UTC", env.GetDBUsername(), env.GetDBPassword(), env.GetDBHost(), env.GetDBPort(), env.GetDBName())
	return dsn, nil
}

func GetSqlDriver(env IEnv) (*sql.Driver, error) {
	url, err := GetConnectionString(env)
	if err != nil {
		return nil, err
	}
	// "root:iam@tcp(127.0.0.1:3304)/iam?parseTime=true&loc=UTC"
	// Open a MySQL connection using standard `sql.Open`
	drv, err := sql.Open("mysql", url)
	if err != nil {
		//  echoerror.WrapToDatabaseError("db connection error (InitEntDb)", err)
		return nil, errors.New("db connection error (InitEntDb)")
	}

	// Create a MySQL dialect from the *sql.DB
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	entDriver := sql.NewDriver("mysql", sql.Conn{ExecQuerier: db})

	return entDriver, nil
}
