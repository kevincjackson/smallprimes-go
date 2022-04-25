package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kevincjackson/smallprimes/spri"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "is",
				Usage: "ask if a number is prime, returns true or false",
				Action: func(c *cli.Context) error {
					arg, err := strconv.Atoi(c.Args().First())
					if err != nil {
						fmt.Println(spri.Is(arg))
					} else {
						fmt.Println("Ooops. Not a number")
					}
					return nil
				},
			},
			{
				Name:  "between x y",
				Usage: "get all primes beween x and y",
				Action: func(c *cli.Context) error {
					fmt.Println("-222 -333 -444", c.Args().First())
					return nil
				},
			},
			{
				Name:  "upto x",
				Usage: "get all primes upto x",
				Action: func(c *cli.Context) error {
					fmt.Println("-11 -22 -33", c.Args().First())
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
