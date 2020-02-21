package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	if len(args) < 1 {
		return fmt.Errorf("not enough arguments received")
	}

	// The last comma doesn't produce any error
	// on the request.
	list := strings.Join(args[1:], ",")
	url := "https://gitignore.io/api/" + list
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("while making request to %q: %v", url, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll((resp.Body))
	if err != nil {
		return fmt.Errorf("while reading response body from %q: %v", url, err)
	}

	file, err := os.Create(".gitignore")
	if err != nil {
		return fmt.Errorf("while creating .gitnore file: %v", err)
	}
	defer file.Close()

	if _, err := file.Write(body); err != nil {
		return fmt.Errorf("while writing .gitignore content: %v", err)
	}

	return nil
}
