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
			Name:    "album_id",
			Value:   "72157712783898977",
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
		{
			Name:  "download",
			Usage: "Download an album",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				photos, err := GetPhotosFromAlbum(c.String("flickr_userid"), c.String("album_id"))

				if err != nil {
					fmt.Println("No Such Album Found")
					return nil
				}

				fmt.Printf("Here is the list of photos in album: \n")
				for index, photo := range photos {
					fmt.Printf("Downloading %v \n", photo)
					links, err := GetPhotoSizes(index)

					if err != nil {
						fmt.Println("Image links not found")
						return nil
					}

					err = DownloadPhoto(links["Original"], photo)
					
					if err != nil {
						fmt.Println("Error in downloading image")
						return nil
					}
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
