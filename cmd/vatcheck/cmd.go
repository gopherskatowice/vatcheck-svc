package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	app *cli.App

	// Do not set these manually! these variables
	// are meant to be set through ldflags.
	buildTag, buildDate string
)

func init() {
	app = cli.NewApp()
	app.Name = "vatcheck-svc"
	app.Usage = "Check if the given VAT number is valid against the EU VIES service"
	app.Author = "Gophers Katowice"
	app.Version = fmt.Sprintf("%s built %s", buildTag, buildDate)
}

func main() {
	AddCommands()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// AddCommands adds child commands to the root command Cmd.
func AddCommands() {
	AddCommand(ServerCommand())
}

// AddCommand adds a child command.
func AddCommand(cmd cli.Command) {
	app.Commands = append(app.Commands, cmd)
}

// InitializeLogging sets logrus log level.
func InitializeLogging(logLevel string) {
	// If log level cannot be resolved, exit gracefully
	if logLevel == "" {
		log.Warning("Log level could not be resolved, fallback to fatal")
		log.SetLevel(log.FatalLevel)
		return
	}
	// Parse level from string
	lvl, err := log.ParseLevel(logLevel)

	if err != nil {
		log.WithFields(log.Fields{
			"passed":  logLevel,
			"default": "fatal",
		}).Warn("Log level is not valid, fallback to default level")
		log.SetLevel(log.FatalLevel)
		return
	}

	log.SetLevel(lvl)
	log.WithFields(log.Fields{
		"level": logLevel,
	}).Debug("Log level successfully set")
}
