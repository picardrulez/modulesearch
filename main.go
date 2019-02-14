package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inModules := checkIfInModules()
	if inModules == false {
		fmt.Println("Modulesearch must be run within a puppet modules root directory.")
		return
	}
	moduleList := getContents("*")
	for _, s := range moduleList {
		checkModule(s, os.Args[1])
	}
}

func getContents(path string) []string {
	files, err := filepath.Glob(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func hasString(file string, grep string) bool {
	f, err := os.Open(file)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), grep) {
			return true
		}
		line++
	}
	return false
}

func checkModule(module string, grep string) {
	wd := pwd()
	var filesWithString []string
	manifestList := getContents(wd + "/" + module + "/manifests/*")
	for _, s := range manifestList {
		containsString := hasString(s, grep)
		if containsString {
			filesWithString = append(filesWithString, s)
			fmt.Println(s)
		}

	}
}

func pwd() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wd
}

func checkIfInModules() bool {
	wd := pwd()
	wdsplit := strings.Split(wd, "/")
	currentDir := wdsplit[len(wdsplit)-1]
	if currentDir == "modules" {
		return true
	} else {
		return false
	}
}
