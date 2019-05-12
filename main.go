package main

import (
	"gopkg.in/urfave/cli.v1"
	"net/http"
	"os"
	"shop/logger"
	"shop/repository"
	"shop/server"
)

func main() {
	db := repository.LoadDatabase()
	defer db.Close()

	logger.Init()

	clientApp := cli.NewApp()
	clientApp.Name = "Shop Service"
	clientApp.Commands = []cli.Command{
		{
			Name:        "start",
			Description: "Start Shop server",
			Action: func(c *cli.Context) error {
				logger.Logger.Println("server is running on port 8080")
				router := server.NewRouter(db)
				http.ListenAndServe(":8080", router)
				return nil
			},
		},
		{
			Name:        "migrate",
			Description: "Run database migrations for shop",
			Action: func(c *cli.Context) error {
				return repository.RunDatabaseMigrations()
			},
		},
		{
			Name:        "rollback",
			Description: "Rollback latest database migration",
			Action: func(c *cli.Context) error {

				return repository.RollbackLatestMigration()
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
