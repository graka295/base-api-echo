package main

import (
	"api-echo/database/connections/db"
	"log"
	"os"
	"sort"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "db",
		Usage:                "db",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "db:migrate",
				Aliases: []string{"m"},
				Usage:   "migrate db database",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "up",
						Aliases: []string{"u"},
					},
					&cli.BoolFlag{
						Name:    "down",
						Aliases: []string{"d"},
					},
					&cli.BoolFlag{
						Name:    "test",
						Aliases: []string{"t"},
					},
				},
				Action: func(ctx *cli.Context) error {
					if ctx.Bool("test") {
						os.Setenv("ENV", "MIGRATE-TEST")
					}
					db, err := gorm.Open("mysql", db.GetConnection())
					if err != nil {
						log.Fatalf("Error connection to main db %v \n", err)
					}

					defer db.Close()
					db.LogMode(true)

					driver, err := mysql.WithInstance(db.DB(), &mysql.Config{})
					if err != nil {
						log.Fatalf("could not start sql migration... %v", err)
					}

					m, err := migrate.NewWithDatabaseInstance(
						"file://database/migrations",
						"mysql", driver)
					if err != nil {
						log.Fatalf("migration failed... %v", err)
					}
					if ctx.Bool("up") {
						if err := m.Up(); err != nil {
							log.Fatalf("An error occurred while syncing the database.. %v", err)
						}
						log.Println("Database db migrated")
					}
					if ctx.Bool("down") {
						if err := m.Down(); err != nil {
							log.Fatalf("An error occurred while syncing the database.. %v", err)
						}
						log.Println("Database db down")
					}
					return err
				},
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
