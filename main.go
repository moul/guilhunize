package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	cli "gopkg.in/urfave/cli.v2"
	"moul.io/guilhunize/guilhunize"
)

func main() {
	app := cli.App{
		Action: run,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "quote"},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	if c.Bool("quote") {
		rand.Seed(time.Now().UnixNano())
		fmt.Println(guilhunize.Quote())
		return nil
	}
	message := strings.Join(c.Args().Slice(), " ")
	if message == "" {
		in, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		message = string(in)
	}
	out := guilhunize.Guilhunize(message)
	fmt.Println(out)
	return nil
}
