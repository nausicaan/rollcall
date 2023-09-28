package main

import (
	"fmt"
	"sort"
	"strings"
)

// Query WordPress for a CSV list with variable fields
func source(spec, field, dest string) string {
	action := execute("-c", "wp", spec, "list", "--"+field, "--skip-plugins", "--skip-themes", "--ssh="+user+":"+path, "--url="+dest, "--format=csv")
	result := strings.Replace(string(action), "blog_id,url\n", "", 1)
	result = strings.ReplaceAll(result, "\n", ",")
	result = strings.TrimSuffix(result, ",")
	return result
}

// Gather and document a list of user ids to use for other queries
func everyone() {
	if !fileExists("sources/urls.txt") {
		solo()
	}

	links := strings.Split(string(readit("sources/urls.txt")), ",")
	unfiltered, refiltered := "", ""

	for _, element := range links {
		list := execute("-c", "wp", "user", "list", "--field=ID", "--skip-plugins", "--skip-themes", "--ssh="+user+":"+path, "--url="+element)
		raw := strings.ReplaceAll(string(list), "\n", ",")
		unfiltered = unfiltered + raw
	}

	unfiltered = strings.TrimSuffix(unfiltered, ",")
	prefiltered := strings.Split(unfiltered, ",")
	sorted := transformer(prefiltered)
	sort.Ints(sorted)
	sorted = unique(sorted)

	for _, element := range sorted {
		refiltered = refiltered + fmt.Sprint(element) + "\n"
	}

	refiltered = strings.TrimSuffix(refiltered, "\n")
	document("sources/everyone.txt", []byte(refiltered))
}

// Filter out extraneous information from the cap variable
func capabilities(meta string) {
	deletions := []string{"\"", "wp_", "_capabilities", "a:1:", "a:2:", "s:6:", "s:10:", "s:12:", "s:13:", ";b:1;", "asset-loader", "{", "}"}
	index := 0
	meta = strings.ReplaceAll(meta, "\n", ",")
	meta = strings.ReplaceAll(meta, "wp_capabilities", "1")
	meta = strings.ReplaceAll(meta, "a:0:{}", "role blank")
	for index < len(deletions) {
		meta = strings.ReplaceAll(meta, deletions[index], "")
		index++
	}
	collection = strings.Split(meta, ",")
}

// Filter out extraneous information from the time variable
func usertime(meta string) {
	meta = strings.ReplaceAll(meta, "\n", ",")
	bunch := strings.Split(meta, ",")
	for _, v := range bunch {
		if strings.Contains(v, "wp_user") {
			v = strings.ReplaceAll(v, "wp_user-settings-time", "0")
		} else {
			v = strings.ReplaceAll(v, "wp_", "")
			v = strings.ReplaceAll(v, "_user-settings-time", "")
		}
		v = strings.TrimSuffix(v, "\n")
		raw = append(raw, v)
	}
}

// Find the URL that matches the Blog ID
func matcher(candidate string, seek []string) string {
	response := "0"
	index := 0
	for index < len(seek) {
		if candidate == seek[index] {
			response = seek[index+1]
		}
		index++
	}
	return response
}

// Create a variable used to write a csv file
func csv(nickname string) {
	index := 1
	capacity := len(collection)
	for index < capacity {
		simplecsv = simplecsv + collection[0] + ","
		simplecsv = simplecsv + nickname + ","
		simplecsv = simplecsv + collection[index] + ","
		link := matcher(collection[index], blogs)
		record := matcher(collection[index], raw)
		simplecsv = simplecsv + link + ","
		index++
		if len(collection[index]) > 20 {
			simplecsv = simplecsv + "administrator,"
		} else {
			simplecsv = simplecsv + collection[index] + ","
		}
		simplecsv = simplecsv + record + "\n"
		index += 2
	}
}

// Piece together the results of the wp user meta list commands into a yaml ready variable
func stitch(nickname string) {
	index := 1
	capacity := len(collection)
	compendium = compendium + nickname + ":\n"
	compendium = compendium + "  - ID: " + collection[0] + "\n"
	for index < capacity {
		compendium = compendium + "  - Blog: " + collection[index] + "\n"
		link := matcher(collection[index], blogs)
		record := matcher(collection[index], raw)
		compendium = compendium + "    URL: " + link + "\n"
		index++
		if len(collection[index]) > 20 {
			compendium = compendium + "    Role: administrator\n"
		} else {
			compendium = compendium + "    Role: " + collection[index] + "\n"
		}
		compendium = compendium + "    Timestamp: " + record + "\n"
		index += 2
	}
}
