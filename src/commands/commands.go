package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	columnize "github.com/ryanuber/columnize"
)

const (
	helpHeader = `Usage: dokku docker-options[:COMMAND]

	Display app's Docker options for all phases.

	Additional commands:`

	helpContent = `
    docker-options:add <app> <phase(s)> OPTION, Add Docker option to app for phase (comma separated phase list)
    docker-options:remove <app> <phase(s)> OPTION, Remove Docker option from app for phase (comma separated phase list)
    docker-options:report [<app>] [<flag>], Displays a docker options report for one or more apps
	`
)

func main() {
	flag.Usage = usage
	flag.Parse()

	cmd := flag.Arg(0)
	switch cmd {
	case "docker-options", "docker-options:help":
		usage()
	case "help":
		fmt.Print("\n    docker-option, Display app's Docker options for all phases.\n")
	default:
		dokkuNotImplementExitCode, err := strconv.Atoi(os.Getenv("DOKKU_NOT_IMPLEMENTED_EXIT"))
		if err != nil {
			fmt.Println("failed to retrieve DOKKU_NOT_IMPLEMENTED_EXIT environment variable")
			dokkuNotImplementExitCode = 10
		}
		os.Exit(dokkuNotImplementExitCode)
	}
}

func usage() {
	config := columnize.DefaultConfig()
	config.Delim = ","
	config.Prefix = "    "
	config.Empty = ""
	content := strings.Split(helpContent, "\n")[1:]
	fmt.Println(helpHeader)
	fmt.Println(columnize.Format(content, config))
}
