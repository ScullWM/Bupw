package main

import (
	"bufio"
	"log"
	"os"
)

func readFiles(f []string) map[string]bool {
	l := make(map[string]bool)

	for _, fn := range f {
		log.Printf("Import file %s \n", fn)
		l = readFile(fn, l)
		log.Printf("Import words %d \n", len(l))
	}

	return l
}

func readFile(fn string, l map[string]bool) map[string]bool {
	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		log.Fatalln(err.Error())
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		l[scanner.Text()] = true
	}

	return l
}
