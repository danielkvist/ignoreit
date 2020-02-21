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

	fmt.Println(".gitignore created successfully!")
}

func run(args []string, stdout io.Writer) error {
	filename := ".gitignore"

	if len(args) == 1 {
		return fmt.Errorf("not enough arguments received")
	}

	if checkIfExists(filename) {
		return fmt.Errorf("%s", filename+" already exists!")
	}

	url, err := createURL("https://gitignore.io/api/", args[1:])
	if err != nil {
		return err
	}

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

func createURL(url string, params []string) (string, error) {
	if len(params) == 0 {
		return "", fmt.Errorf("not enough parameters received")
	}

	return url + strings.Join(params, ","), nil
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
		return nil, fmt.Errorf("while creating %q file: %v", filename, err)
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
