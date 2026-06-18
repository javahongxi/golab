package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func convertToBin(n int) string {
	if n == 0 {
		return "0"
	}
	var sb strings.Builder
	sb.Grow(32) // pre-allocate for efficiency
	for ; n > 0; n /= 2 {
		sb.WriteByte(byte('0' + n%2))
	}
	// reverse the string
	bytes := []byte(sb.String())
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println("convertToBin results:")
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(72387885),
		convertToBin(0),
	)

	fmt.Println("abc.txt contents:")
	printFile("lang/basic/branch/abc.txt")

	fmt.Println("printing a string:")
	s := `abc"d"
	kkkk
	123

	p`
	printFileContents(strings.NewReader(s))

	// Uncomment to see it runs forever
	// forever()
}
