/*
The following code contains some utility functions used in the project, mainly related to
file IO.

Author: Shravan Asati
Originially Written: 19 June 2021
Last Edited: 29 June 2021
*/

package services 
// package main 

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

// getProbeDir returns the probe home directory, namely `~/.probe`. Also creates the directory if it doesnt exists.
func getProbeDir() string {
	usr, e := user.Current()
	if e != nil {
		Log("error", "unable to get homedir")
		fmt.Println(e)
		return ""
	}
	
	// * determining probe's directory
	dir := filepath.Join(usr.HomeDir, ".probe")

	_, er := os.Stat(dir)
	if os.IsNotExist(er) {
		os.Mkdir(dir, os.ModePerm)
	}

	tempDir := filepath.Join(dir, "temp")
	_, err := os.Stat(tempDir)
	if os.IsNotExist(err) {
		os.Mkdir(tempDir, os.ModePerm)
	}

	return dir
}

// randomFileName returns a random file name with letters and numbers and some valid symbols mixed.
func randomFileName(lang string) string {
	characters := "abcdefghijklmnopqrstuvwxyz-1234567890_ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())

	var filename string
	for len(filename) <= 12 {
		filename += string(characters[rand.Intn(len(characters))])
	}

	if lang == "python" {
		filename += ".py"
	} else if lang == "javascript" {
		filename += ".js"
	} else {
		Log("error", "invalid value for the lang parameter: "+lang)
	}

	return filename
}

// writeToFile writes the given content to the given filepath.
func writeToFile(filename, content string) {
	f, e := os.Create(filename)
	if e != nil {
		Log("error", "unable to open the file")
		fmt.Println(e)
		return
	}

	defer f.Close()

	if _, e := f.WriteString(content); e != nil {
		Log("error", "unable to write to file "+filename)
		fmt.Println(e)
		return
	}
}

// ClearClutter deletes all the files present in probe's directory.
func ClearClutter() {
	files, er := ioutil.ReadDir(filepath.Join(getProbeDir(), "temp"))
	if er != nil {
		Log("error", "unable to get files in the directory")
		fmt.Println(er)
		return
	}

	// * clearing all files in the probe's directory
	for _, f := range files {
		path := filepath.Join(getProbeDir(), "temp", f.Name())
		if e := os.Remove(path); e != nil {
			Log("error", "unable to remove file "+path)
			fmt.Println(e)
		}
	}
}

// readFile reads the given file and returns the string content of the same.
func readFile(file string) string {
	f, ferr := os.Open(file)
	if ferr != nil {
		Log("error", "unable to read file " + file)
		panic(ferr)
	}
	defer f.Close()

	text := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text += scanner.Text()
	}

	return text
}
