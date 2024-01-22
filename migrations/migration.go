package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "./migrations", "directory with migration files")
)

func init() {
	println("Starting Migrations + Loading .env file")
	godotenv.Load()
}

func main() {
	flags.Parse(os.Args[1:])

	args := flags.Args()

	for _, arg := range args {
		println(arg)
	}
	print("length of args: ", len(args), "\n")

	if len(args) < 1 {
		flags.Usage()
		return
	}

	dbstring := os.Getenv("DATABASE_URL")

	command := args[0]

	db, err := goose.OpenDBWithDriver("postgres", dbstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.RunContext(context.TODO(), command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
