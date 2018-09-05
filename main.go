package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func initApp() *cli.App {
	app := cli.NewApp()
	app.Name = "drone-chartmuseum-plugin"
	app.Usage = "drone plugin to upload charts to chartmuseum server"
	app.Version = fmt.Sprintf("1.0.0")

	mainFlag := []cli.Flag{
		cli.StringFlag{
			Name:   "repo-url,u",
			Value:  "",
			Usage:  "ChartMuseum API base URL",
			EnvVar: "PLUGIN_REPO_URL,REPO_URL",
		},
		cli.StringFlag{
			Name:   "username,u",
			Value:  "",
			Usage:  "Chartmuseum HTTP Basic auth username",
			EnvVar: "PLUGIN_REPO_USERNAME",
		},
		cli.StringFlag{
			Name:   "password,u",
			Value:  "",
			Usage:  "Chartmuseum HTTP Basic auth password",
			EnvVar: "PLUGIN_REPO_PASSWORD",
		},
		cli.StringFlag{
			Name:   "chart-path,i",
			Usage:  "Path to chart, relative to charts-dir",
			Value:  "",
			EnvVar: "PLUGIN_CHART_PATH,CHART_PATH",
		},
		cli.StringFlag{
			Name:   "charts-dir,d",
			Value:  "./",
			Usage:  "chart directory",
			EnvVar: "PLUGIN_CHARTS_DIR,CHARTS_DIR",
		},
		cli.StringFlag{
			Name:   "save-dir,o",
			Value:  "uploads/",
			Usage:  "Directory to save chart packages",
			EnvVar: "PLUGIN_SAVE_DIR,SAVE_DIR",
		},
		cli.StringFlag{
			Name:   "previous-commit,p",
			Usage:  "Previous commit id (`COMMIT_SHA`)",
			EnvVar: "PLUGIN_PREVIOUS_COMMIT,PREVIOUS_COMMIT",
		},
		cli.StringFlag{
			Name:   "current-commit,c",
			Usage:  "Current commit id (`COMMIT_SHA`)",
			EnvVar: "PLUGIN_CURRENT_COMMIT,CURRENT_COMMIT",
		},
		cli.StringFlag{
			Name:   "log-level",
			Value:  "error",
			Usage:  "Log level (panic, fatal, error, warn, info, or debug)",
			EnvVar: "PLUGIN_LOG_LEVEL,LOG_LEVEL",
		},
	}

	app.Action = cli.ActionFunc(defaultAction)
	app.Flags = mainFlag

	return app
}

func defaultAction(c *cli.Context) error {
	logLevelString := c.String("log-level")
	logLevel, err := log.ParseLevel(logLevelString)
	if err != nil {
		return err
	}
	log.SetLevel(logLevel)

	plugin := Plugin{
		Config: &Config{
			RepoURL:          c.String("repo-url"),
			Username:         c.String("username"),
			Password:         c.String("password"),
			ChartsDir:        c.String("charts-dir"),
			ChartPath:        c.String("chart-path"),
			PreviousCommitID: c.String("previous-commit"),
			CurrentCommitID:  c.String("current-commit"),
			SaveDir:          c.String("save-dir"),
		},
	}
	return plugin.exec()
}

func main() {
	app := initApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
