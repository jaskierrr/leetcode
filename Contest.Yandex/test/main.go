package main

import (
	"bufio"
	"log"
	"os"
)

func read() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	data, _ := reader.ReadString('\n')
	data2, err := reader.ReadString('\n')


	if err != nil {
		return "", "", err
	}

	return data, data2, nil
}

func write(data string) error {
	writer := bufio.NewWriter(os.Stdout)
	_, err := writer.WriteString(data)
	if err != nil {
		return err
	}
	return writer.Flush()
}

func main() {
	data, data2, err := read()
	if err != nil {
		log.Fatal("wrong data input")
	}

	err = write(data)
	if err != nil {
		log.Fatal("wrong data output")
	}
	err = write(data2)
	if err != nil {
		log.Fatal("wrong data output")
	}
}
