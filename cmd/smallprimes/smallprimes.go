package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kevincjackson/smallprimes-go/internal/format"
	"github.com/kevincjackson/smallprimes-go/pkg/primedata"
	"github.com/urfave/cli/v2"
)

const num_err_msg = "Ooops. Not a number"

var large_num_err_msg = fmt.Sprintf("Ooops. Max input is %d.\n", primedata.MaxInt)

func main() {
	var formatArg string
	app := cli.NewApp()
	app.Name = "smallprimes"
	app.Description = "This app does no calculations. It's designed to quickly return primes based on byte data."
	app.Usage = "get primes up to 10**9."
	app.UsageText = strings.Join([]string{
		"smallprimes is 7                   => true",
		"smallprimes between 7 13           => 7 11 13",
		"smallprimes upto 7                 => 2 3 5 7",
		"smallprimes -f json upto 7         => [2, 3, 5, 7]",
		"smallprimes -f newlines upto 7     => \"2\\n3\\n5\\n7\\n\"",
		"smallprimes upto 100 > primes.txt  => primes.txt",
	}, "\n")
	app.ArgsUsage = "[arrgh]"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "format",
			Aliases:     []string{"f"},
			Value:       "spaces",
			Destination: &formatArg,
			Usage:       "--format=json|newlines|spaces",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:        "is",
			Usage:       "yields true|false",
			UsageText:   "smallprimes is 7",
			Description: "is x prime, yields true|false",

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
			Name:        "between",
			Usage:       "yields list of ints",
			UsageText:   "smallprimes between 10 20",
			Description: "between x y inclusive, yields list of ints",
			Action: func(c *cli.Context) error {
				x, err := strconv.Atoi(c.Args().Get(0))
				y, err2 := strconv.Atoi(c.Args().Get(1))
				if err != nil || err2 != nil {
					cli.ShowCommandHelp(c, "between")
				} else if x > primedata.MaxInt || y > primedata.MaxInt {
					fmt.Println(large_num_err_msg)
				} else {
					format.PrintInts(primedata.Between(x, y), formatArg)
				}
				return nil
			},
		},
		{
			Name:        "upto",
			Usage:       "yields list of ints",
			UsageText:   "smallprimes upto 10",
			Description: "upto x inclusive, yields list of ints",
			Action: func(c *cli.Context) error {
				x, err := strconv.Atoi(c.Args().Get(0))
				if err != nil {
					fmt.Println(num_err_msg)
				} else if x > primedata.MaxInt {
					fmt.Println(large_num_err_msg)
				} else {
					format.PrintInts(primedata.Upto(x), formatArg)
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
