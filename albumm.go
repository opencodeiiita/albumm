package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:    "albumm",
		Usage:   "Download Photos from a Flickr Album",
		Version: "1.0.0",
	}

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:    "flickr_username",
			Value:   "Acoustics & Media Society, IIIT Allahabad",
			Aliases: []string{"name"},
		},
		&cli.StringFlag{
			Name:    "album_name",
			Value:   "Kite Flying 2020",
			Aliases: []string{"album"},
		},
		&cli.StringFlag{
			Name:    "flickr_userid",
			Value:   "129074767@N06",
			Aliases: []string{"id"},
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "id",
			Usage: "Get User ID of a username",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				userId, err := GetUserIDByUsername(c.String("flickr_username"))

				if err != nil {
					fmt.Println("No Users Found")
					return nil
				}

				fmt.Printf("Username: %s \nUser ID: %s \n", c.String("flickr_username"), userId)
				return nil
			},
		},
		{
			Name:  "albums",
			Usage: "List all Albums of a user",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				albums, err := GetAlbumsFromUserID(c.String("flickr_userid"))

				if err != nil {
					fmt.Println("No Albums Found")
					return nil
				}

				fmt.Printf("Here is a list of albums: \n")
				for index, album := range albums {
					fmt.Printf("%v) %v \n", index, album)
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
