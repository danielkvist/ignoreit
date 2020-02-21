package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// TODO: Add tests
// TODO: End README

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	fmt.Println(".gitignore created successfully!")
}

func run(args []string, stdout io.Writer) error {
	filename := ".gitignore"

	if checkIfExists(filename) {
		return fmt.Errorf("%s", ".gitignore already exists!")
	}

	if len(args) < 1 {
		return fmt.Errorf("not enough arguments received")
	}

	url := createURL("https://gitignore.io/api/", args[1:])
	data, err := fetch(url)
	if err != nil {
		return err
	}

	file, err := createFile(filename)
	if err != nil {
		return fmt.Errorf("while creating %q file: %v", filename, err)
	}
	defer file.Close()

	if err := writeTo(file, data); err != nil {
		return fmt.Errorf("while writing %q content: %v", filename, err)
	}

	return nil
}

func checkIfExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func createURL(url string, params []string) string {
	// The last comma doesn't produce any error
	// on the request.
	return url + strings.Join(params, ",")
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("while making request to %q: %v", url, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll((resp.Body))
	if err != nil {
		return nil, fmt.Errorf("while reading response body from %q: %v", url, err)
	}

	data := make([]byte, len(body))
	_ = copy(data, body)
	return data, nil
}

func createFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("while creating .gitnore file: %v", err)
	}

	return file, nil
}

func writeTo(w io.Writer, data []byte) error {
	_, err := w.Write(data)
	if err != nil {
		return err
	}

	return nil
}
