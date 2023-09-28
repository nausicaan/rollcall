package main

import (
	"fmt"
	"strings"
)

const (
	reset  string = "\033[0m"
	green  string = "\033[32m"
	yellow string = "\033[33m"
	red    string = "\033[41m"
	bv     string = "1.0"
)

func main() {
	hn := execute("-c", "hostname")

	switch strings.TrimSpace(string(hn)) {
	case servers[0]:
		server, path, user = testpack[0], testpack[1], testpack[2]
		greenlight()
	case servers[4]:
		server, path, user = prodpack[0], prodpack[1], prodpack[2]
		greenlight()
	default:
		about()
	}
}

func greenlight() {
	blogs = strings.Split(source("site", "fields=blog_id,url", server), ",")

	if !fileExists("sources/everyone.txt") {
		everyone()
	}

	ids := strings.Split(string(readit("sources/everyone.txt")), "\n")

	for _, element := range ids {
		element = strings.TrimSuffix(element, "\n")
		found := execute("-c", "wp", "user", "meta", "list", element, "--skip-plugins", "--skip-themes", "--url="+server, "--ssh="+user+":"+path, "--format=csv")
		document("found.csv", found)
		cap := execute("-c", "grep", "_capabilities", "found.csv")
		ust := execute("-c", "grep", "user-settings-time", "found.csv")
		if len(ust) > 1 || len(cap) > 1 {
			capabilities(string(cap))
			usertime(string(ust))
			nickname := execute("-c", "wp", "user", "get", collection[0], "--field=login", "--url="+server, "--ssh="+user+":"+path)
			// stitch(strings.TrimSuffix(string(nickname), "\n"))
			csv(strings.TrimSuffix(string(nickname), "\n"))
		}
	}

	compendium = compendium + "..."

	cleanup("found.csv")
	document("results/compendium.csv", []byte(simplecsv))
	// document("results/compendium.yaml", []byte(compendium))
}

// Print the program version to screen
func version() {
	fmt.Println(yellow+"Rollcall", green+bv, reset)
}

// Print help information for using the program
func about() {
	fmt.Println(yellow, "\nUsage:", reset)
	fmt.Println("  [program]")
	fmt.Println(green, "   rollcall")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "   https://github.com/nausicaan/rollcall.git")
	fmt.Println(reset)
}
