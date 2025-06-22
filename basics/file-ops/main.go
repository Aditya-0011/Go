package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func example1() {
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("File name:", info.Name())
	fmt.Println("File size:", info.Size())
	fmt.Println("File mode:", info.Mode())
	fmt.Println("Last modified:", info.ModTime())
	fmt.Println("Is directory:", info.IsDir())
	fmt.Println("System interface:", info.Sys())
}

func example2() {
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := make([]byte, 12)

	for {
		i, err := file.Read(buffer)
		if err == io.EOF {
			return
		} else if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println("Read data:", string(buffer[:i]))
	}
}

func example3() {
	file, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File content:", string(file))
}

func example4() {
	dir, err := os.Open("..")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	info, err := dir.ReadDir(-1)
	if err != nil {
		panic(err)
	}

	for _, i := range info {
		fmt.Println("Name:", i.Name())
		fmt.Println("Is Directory:", i.IsDir())
	}
}

func example5() {
	file, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//file.WriteString("Hello, World!\n")
	bytes := []byte("Hello, World!\n")
	_, err = file.Write(bytes)
	if err != nil {
		panic(err)
	}
}

func example6() {
	source, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer source.Close()

	dest, err := os.Create("dest.txt")
	if err != nil {
		panic(err)
	}
	defer dest.Close()

	reader := bufio.NewReader(source)
	writer := bufio.NewWriter(dest)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		err = writer.WriteByte(b)
		if err != nil {
			panic(err)
		}

	}

	writer.Flush()
}

func main() {
	// example1()
	// example2()
	// example3()
	// example4()
	// example5()
	// example6()
}
