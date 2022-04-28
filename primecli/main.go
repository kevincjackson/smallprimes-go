package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kevincjackson/smallprimes-go/primedata"
	"github.com/urfave/cli/v2"
)

const num_err_msg = "Ooops. Not a number"

var large_num_err_msg string

func init() {
	large_num_err_msg = fmt.Sprintf("Ooops. Max input is %d.\n", primedata.MaxInt)
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "is",
				Usage: "ask if a number is prime, returns true or false",
				Action: func(c *cli.Context) error {
					x, err := strconv.Atoi(c.Args().First())
					if err != nil {
						fmt.Println(num_err_msg)
					} else if x > primedata.MaxInt {
						fmt.Println(large_num_err_msg)
					} else {
						fmt.Println(primedata.Is(x))
					}
					return nil
				},
			},
			{
				Name:  "between",
				Usage: "get all primes beween x and y",
				Action: func(c *cli.Context) error {
					x, err := strconv.Atoi(c.Args().Get(0))
					y, err2 := strconv.Atoi(c.Args().Get(1))
					if err != nil || err2 != nil {
						fmt.Println(num_err_msg)
					} else if x > primedata.MaxInt || y > primedata.MaxInt {
						fmt.Println(large_num_err_msg)
					} else {
						fmt.Println(primedata.Between(x, y))
					}
					return nil
				},
			},
			{
				Name:  "upto",
				Usage: "get all primes upto x",
				Action: func(c *cli.Context) error {
					x, err := strconv.Atoi(c.Args().Get(0))
					if err != nil {
						fmt.Println(num_err_msg)
					} else if x > primedata.MaxInt {
						fmt.Println(large_num_err_msg)
					} else {
						fmt.Println(primedata.Upto(x))
					}
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
