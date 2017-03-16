package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/boltdb/bolt"
	"github.com/urfave/cli"
)

// the bolt database
var DB *bolt.DB

func main() {
	app := cli.NewApp()
	app.Name = "pacnews"
	app.Usage = "Read the arch news"
	app.Description = "An arch news reader"
	app.Version = "0.0.1"
	app.Author = "Kristof Vannotten"

	app.Commands = []cli.Command{
		{
			Name:    "check",
			Aliases: []string{"c"},
			Usage:   "check if there is news",
			Action:  checkNews,
		},
		{
			Name:    "read",
			Aliases: []string{"r"},
			Usage:   "read the news",
			Action:  readNews,
		},
	}

	DB = setupDB()

	defer DB.Close()

	app.Run(os.Args)
}

func setupDB() *bolt.DB {
	path := path.Join("/var/cache", "pacnews.db")

	db, err := bolt.Open(
		path,
		os.FileMode(0666),
		&bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	os.Chmod(path, 0666)

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("news"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return db
}
