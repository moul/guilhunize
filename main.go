package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	cli "gopkg.in/urfave/cli.v2"
	"moul.io/guilhunize/guilhunize"
)

func main() {
	app := cli.App{Action: run}
	if err := app.Run(os.Args); err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
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
