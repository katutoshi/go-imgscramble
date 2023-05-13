package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/katutoshi/go-imgscramble"
)

func main() {
	command := os.Args[1]

	switch command {
	case "scramble":
		Scramble(os.Args[2])
	case "unscramble":
		seed, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatalf("%s", err)
		}
		Unscramble(os.Args[2], int64(seed))
	default:
		log.Fatalf("unknown command: %s", command)
	}
}

func Scramble(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer file.Close()

	seed, err := imgscramble.Seed(file)
	if err != nil {
		log.Fatalf("%s", err)
	}

	if _, err := file.Seek(0, 0); err != nil {
		log.Fatalf("%s", err)
	}

	scrambled, err := imgscramble.Scramble(file, seed)
	if err != nil {
		log.Fatalf("%s", err)
	}

	baseFileName := filepath.Base(file.Name())[:len(filepath.Base(file.Name()))-len(filepath.Ext(file.Name()))]
	name := fmt.Sprintf("%s_scrambled%s", baseFileName, filepath.Ext(file.Name()))
	if err := os.WriteFile(name, scrambled, 0644); err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("seed: %d", seed)
	log.Printf("saved: %s", name)
}

func Unscramble(path string, seed int64) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer file.Close()

	scrambled, err := imgscramble.Unscramble(file, seed)
	if err != nil {
		log.Fatalf("%s", err)
	}

	baseFileName := filepath.Base(file.Name())[:len(filepath.Base(file.Name()))-len(filepath.Ext(file.Name()))]
	name := fmt.Sprintf("%s_unscrambled%s", baseFileName, filepath.Ext(file.Name()))
	if err := os.WriteFile(name, scrambled, 0644); err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("seed: %d", seed)
	log.Printf("saved: %s", name)
}
