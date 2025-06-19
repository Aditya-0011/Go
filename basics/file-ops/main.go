package main

import (
	"fmt"
	"io"
	"os"
)

func createFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func writeToFile(file *os.File, content string) error {
	byte, err := io.WriteString(file, content+"\n")
	if err != nil {
		return err
	}
	fmt.Printf("Wrote %d bytes to file.\n", byte)
	return nil
}

func createAndWriteFile(fileName string) {
	file, err := createFile(fileName)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File created successfully:", file.Name())

	err = writeToFile(file, "Hello, World!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("'Hello, World!' written to file successfully.")

	err = writeToFile(file, "More Data!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("'More Data!' written to file successfully.")
}

func openReadOnlyFileUsingBuffer(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)

	for {
		data, err := file.Read(buffer)
		if err == io.EOF {
			return
		} else if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println("Read data:", string(buffer[:data]))
	}
}

func openReadOnlyFileUsingOs(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(data))
}

func openAndEditFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644) // linux chmod 644
	if err != nil {
		fmt.Println("Error opening file for editing:", err)
	}
	defer file.Close()

	err = writeToFile(file, "Some more Data!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("'Some more Data!' written to file successfully.")
}

func main() {
	//createAndWriteFile("example.txt")
	//openReadOnlyFileUsingBuffer("example.txt")
	openReadOnlyFileUsingOs("example.txt")
	openAndEditFile("example.txt")
	openReadOnlyFileUsingOs("example.txt")
}
