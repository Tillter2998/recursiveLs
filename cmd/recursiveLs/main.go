package main

import (
	"log"
	"os"

	fileTree "github.com/Tillter2998/recursiveLs/internal"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	tree, err := fileTree.BuildFileTree(workingDir)

	// tree.PrintTree("", true, true)
	tree.PrintTree("", 0, false)
}
