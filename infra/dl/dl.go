package dl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var DEFAULT_ENV = make(map[string]string)

func readEnviroment() {
	filePath := ""
	if os.Getenv("APP_MODE") == "debug" {
		filePath = "./../enviroments.txt"
	} else {
		filePath = "./enviroments.txt"

	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := readLine(scanner.Text())
		DEFAULT_ENV[line[0]] = line[1]

		fmt.Println(line)
	}
}

func readLine(line string) []string {
	itens := make([]string, 0)
	indice := strings.Index(line, "=")
	itens = append(itens, line[:indice])
	itens = append(itens, line[indice+1:])
	return itens
}

func GetEnv(key string) string {
	if len(DEFAULT_ENV) == 0 {
		readEnviroment()
	}
	keyFile := DEFAULT_ENV[key]
	value, exists := os.LookupEnv(key)
	if !exists {
		return keyFile
	}
	return value
}
