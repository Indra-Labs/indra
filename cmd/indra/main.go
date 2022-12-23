package main

import (
	"fmt"
	"github.com/Indra-Labs/indra"
	"github.com/cybriq/proc/pkg/app"
	"github.com/cybriq/proc/pkg/cmds"
	log2 "github.com/cybriq/proc/pkg/log"
	"github.com/cybriq/proc/pkg/opts/config"
	"github.com/davecgh/go-spew/spew"
	"os"
)

var (
	log      = log2.GetLogger(indra.PathBase)
	check    = log.E.Chk
	commands = &cmds.Command{
		Name:          "indra",
		Description:   "The indra network daemon.",
		Documentation: lorem,
		Entrypoint: func(c *cmds.Command, args []string) error {
			log.I.Ln("running node")
			return nil
		},
		Default: cmds.Tags("help"),
		Configs: config.Opts{
			//"AutoPorts": toggle.New(meta.Data{
			//	Label:         "Automatic Ports",
			//	Tags:          cmds.Tags("node", "wallet"),
			//	Description:   "RPC and controller ports are randomized, use with controller for automatic peer discovery",
			//	Documentation: lorem,
			//	Default:       "false",
			//}),
		},
		Commands: cmds.Commands{
			{
				Name:        "version",
				Description: "print indra version",

				Documentation: lorem,
				Entrypoint: func(c *cmds.Command, args []string) error {
					fmt.Println(indra.SemVer)
					return nil
				},
			},
		},
	}
)

const lorem = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis 
nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu 
fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in 
culpa qui officia deserunt mollit anim id est laborum.`

func main() {

	//log2.SetLogLevel(log2.Debug)

	cmds.GetConfigBase(commands.Configs, commands.Name, false)

	var err error
	if commands, err = cmds.Init(commands, os.Args); check(err) {
		return
	}

	var application *app.App
	if application, err = app.New(commands, os.Args); err != nil {
		spew.Dump(err)
		os.Exit(1)
	}

	//log.I.S(application)

	if err = application.Launch(); err != nil {
		spew.Dump(err)
		os.Exit(1)
	}
}