package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Creating input directory")
	cmd := exec.Command("rm", "-rf", "testdata/input")
	cmd.Run()
	cmd = exec.Command("cp", "-R", "testdata/golden-sample", "testdata/input")
	cmd.Run()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestA(t *testing.T) {
	log.Println("TestA running")
	path := "testdata/input"
	fullPath, _ := filepath.Abs(path)
	filesMap := findMatchedFiles(fullPath)
	renameFiles(filesMap)

	fmt.Println("##### Comparing results with golden sample ######")
	cmd := exec.Command("diff", "-r", "testdata/expected", "testdata/input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	fmt.Printf("%s\n", out.String())
	if err != nil {
		log.Fatal(err)
	}
}
