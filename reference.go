package main

import (
	"strings"
)

const (
	identity string = "  - ID: "
	link     string = "    URL: "
)

// Create a list of all urls in the WordPress environment
func solo() {
	list := source("site", "field=url", server)
	document("sources/urls.txt", []byte(list))
}

// Create a combined list of urls and blog ids in the WordPress environment
func duo() {
	j := "[\n"
	y := "---\n"
	index := 0

	for index < len(blogs) {
		y = y + "- ID: " + blogs[index] + "\n"
		j = j + "    {\n        \"ID\": " + blogs[index] + ",\n"
		index++
		y = y + "  URL: " + blogs[index] + "\n"
		j = j + "        \"URL\": \"" + blogs[index] + "\"\n    },\n"
		index++
	}

	j = strings.TrimSuffix(j, "\n")
	j = strings.TrimSuffix(j, ",")
	j = j + "\n]"
	y = y + "..."
	document("sources/urls-ids.yaml", []byte(y))
	document("sources/urls-ids.json", []byte(j))
}

// Organize the list by type
func category() {
	types := []string{"engage", "events", "forms", "vanity", "workingforyou"}
	blog, engage, events, forms, vanity, workingforyou := "Misc:\n", "Engage:\n", "Events:\n", "Forms:\n", "Vanity:\n", "Workingforyou:\n"
	index := 1
	for index < len(blogs) {
		variety := "other"
		for _, element := range types {
			element = strings.TrimSuffix(element, "\n")
			if strings.Contains(blogs[index], element) {
				variety = element
			}
		}

		switch variety {
		case "forms":
			forms = forms + identity + blogs[index-1] + "\n"
			forms = forms + link + blogs[index] + "\n"
			index += 2
		case "engage":
			engage = engage + identity + blogs[index-1] + "\n"
			engage = engage + link + blogs[index] + "\n"
			index += 2
		case "events":
			events = events + identity + blogs[index-1] + "\n"
			events = events + link + blogs[index] + "\n"
			index += 2
		case "vanity":
			vanity = vanity + identity + blogs[index-1] + "\n"
			vanity = vanity + link + blogs[index] + "\n"
			index += 2
		case "workingforyou":
			workingforyou = workingforyou + identity + blogs[index-1] + "\n"
			workingforyou = workingforyou + link + blogs[index] + "\n"
			index += 2
		default:
			blog = blog + identity + blogs[index-1] + "\n"
			blog = blog + link + blogs[index] + "\n"
			index += 2
		}

		compile := "---\n" + blog + engage + events + forms + vanity + workingforyou + "..."
		document("sources/blog-types.yaml", []byte(compile))
	}
}
