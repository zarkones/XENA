package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

var sourceToObfuscatePath string = "xena-bot-apep"

// locateFiles returns a slice of files in a directory.
// If a file is ending in .go
func locateFiles(path string) ([]string, error) {
	var foundFiles []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".go") {
			foundFiles = append(foundFiles, file.Name())
		}
	}

	return foundFiles, nil
}

func main() {
	files, err := locateFiles(sourceToObfuscatePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	filesAndSources := make(map[string]string)
	funcNames := []string{}
	funcNameMap := make(map[string]string)

	// Load files and function names.
	for _, file := range files {
		content, err := ioutil.ReadFile(sourceToObfuscatePath + "/" + file)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		filesAndSources[file] = string(content)

		r := regexp.MustCompile(`func (.*?)\(`)
		matches := r.FindAllString(filesAndSources[file], -1)

		for _, funcName := range matches {
			fun := funcName[5:len(funcName)-1]
			if len(fun) > 0 {
				funcNames = append(funcNames, fun)
			}
		}
	}

	// Map function names to obfuscated names.
	for _, funcName := range funcNames {
		word1 := randomPopularWord(time.Now().UnixMilli())
		word2 := strings.Title(randomPopularWord(time.Now().UnixNano()))
		word3 := strings.Title(randomPopularWord(time.Now().UnixMicro()))
		word4 := strings.Title(randomPopularWord(time.Now().UnixNano()))
		if funcName == "main" {
			funcNameMap[funcName] = funcName
		} else {
			funcNameMap[funcName] = word1 + word2 + word3 + word4
		}
	}

	obfFilesAndSources := make(map[string]string)

	// Iterate over sources and perform modifications.
	for file, source := range filesAndSources {
		obfuscatedSource := source

		for orgFuncName, newFuncName := range funcNameMap {
			r := regexp.MustCompile(orgFuncName)
			obfuscatedSource = r.ReplaceAllString(obfuscatedSource, newFuncName)
		}

		obfFilesAndSources[file] = obfuscatedSource
	}

	// Output the obfuscated source.
	for file, source := range obfFilesAndSources {
		os.Mkdir("output", 0755)
		f, err := os.Create("output/" + file)
		if err != nil {
			fmt.Println(err.Error())
		}
		f.WriteString(source)
	}
}
