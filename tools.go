package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	compendium         = "---\n"
	collection         = []string{}
	raw                = []string{}
	blogs              = []string{}
	server, path, user = "", "", ""
	reader             = bufio.NewReader(os.Stdin)
	simplecsv          = "ID,Name,Blog,URL,Role,Timestamp\n"
)

// Check for errors, print the result if found
func inspect(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Get user input via screen prompt
func solicit(prompt string) string {
	fmt.Print(prompt)
	response, _ := reader.ReadString('\n')
	return strings.TrimSpace(response)
}

// Run a terminal command using flags to customize the output
func execute(variation, task string, args ...string) []byte {
	osCmd := exec.Command(task, args...)
	switch variation {
	case "-e":
		// exec.Command(task, args...).CombinedOutput()
		osCmd.CombinedOutput()
	case "-c":
		both, _ := osCmd.CombinedOutput()
		return both
	case "-v":
		osCmd.Stdout = os.Stdout
		osCmd.Stderr = os.Stderr
		err := osCmd.Run()
		inspect(err)
	}
	return nil
}

// Convert a string slice into a int slice
func transformer(slice []string) []int {
	var converted []int
	for _, element := range slice {
		i, _ := strconv.Atoi(element)
		converted = append(converted, i)
	}
	return converted
}

// Generic function to remove duplicates
func unique[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

// Check if a file is present in the supplied path
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Write a passed variable to a named file
func document(name string, content []byte) {
	inspect(os.WriteFile(name, content, 0644))
}

// Read any file and return the contents as a byte variable
func readit(file string) []byte {
	outcome, err := os.ReadFile(file)
	inspect(err)
	return outcome
}

// Remove files or directories
func cleanup(cut string) {
	inspect(os.Remove(cut))
}
