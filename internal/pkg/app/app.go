package app

import (
	"fmt"
	"log"
	"os"
	"wallet-app/internal/app/db"
	"wallet-app/internal/app/endpoint"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

type App struct {
	e    *endpoint.Endpoint
	db   *db.DB
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.echo = echo.New()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading.env file: %s", err.Error())
		return nil, err
	}

	dbX, err := db.NewPostgresDB(&db.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("POSTGRES_SSL"),
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
		return nil, err
	}

	a.db = db.NewDB(dbX)
	a.e = endpoint.New(a.db)

	a.echo.GET("/api/v1/wallets/:id", a.e.Balance)
	a.echo.POST("/api/v1/wallet", a.e.DepOrWithdraw)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Running the application...")

	err := a.echo.Start(":8080")

	if err != nil {
		log.Fatal(err)
	}

	return err
}
