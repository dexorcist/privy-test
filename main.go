package main

import (
	"os"
	"privy-test/api"
	"privy-test/infra"
	"time"

	"github.com/urfave/cli"
)

const (
	// AppName Application name
	AppName = "Privy Test"
	// AppTagLine Application tagline
	AppTagLine = "Privy Test API"
)

var (
	Namespace    string
	BuildVersion string
)

// @title {{.Title}}
// @description {{.Description}}
// @version {{.Version}}
// @host {{.Host}}
// @schemes {{.Schemes}}
// @basePath {{.BasePath}}
func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Usage = AppTagLine
	app.Version = BuildVersion
	app.ExtraInfo = func() map[string]string {
		return map[string]string{
			"Namespace": Namespace,
		}
	}
	app.HideHelp = true

	app.Commands = []cli.Command{
		API,
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	_ = app.Run(os.Args)
}

// API run API server
var API = cli.Command{
	Name:     "api",
	Usage:    "Run API Server",
	HideHelp: true,
	Action: func(ctx *cli.Context) {
		//TODO sleep waiting db running as well, but this solution not practice
		time.Sleep(3 * time.Second)
		api.NewServer(infra.New(Namespace, BuildVersion)).Run()
	},
}
