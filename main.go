package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func main() {
	var keyName string
	var keyPath string
	var executeablePath string
	if (len(os.Args) == 2) || (len(os.Args) == 1) {
		if len(os.Args) == 1 {
			executeablePath = interfaces()
		} else {
			executeablePath = os.Args[1]
		}
		filename, format := checker(executeablePath)
		if format != ".exe" {
			fmt.Println("Your format is", format, "we expecting .exe extension")
			exitProgram("Please use Application that using .exe etension")
		}
		keyName = filename + format
		keyPath = executeablePath
		keyMaker(keyName, keyPath)
	} else if len(os.Args) == 3 {
		keyName = os.Args[1] + ".exe"
		_, format := checker(os.Args[2])
		if format != ".exe" {
			fmt.Println("Your format is", format, "we expecting .exe extension")
			exitProgram("Please use Application that using .exe etension")
		}
		keyPath = os.Args[2]
		keyMaker(keyName, keyPath)
	} else {
		exitProgram("We need path")
	}
	fmt.Println("All Done")
	pressAnyKey()
}

func interfaces() string {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println(
		`===============================
           RunMe v1.1
================================`)
	fmt.Print("Enter paths : ")
	patssshs, _ := reader.ReadString('\n')
	//paths = strings.TrimSuffix(paths, "\n")
	patssshs = patssshs[0:(len(patssshs) - 2)]

	return patssshs
}

func checker(value string) (filename string, format string) {
	file := filepath.Base(value)
	format = filepath.Ext(value)
	filename = file[0 : len(file)-len(format)]

	return filename, format
}

func exitProgram(isi string) {
	fmt.Println(isi)
	os.Exit(1)
}

func keyMaker(keyName string, keyPath string) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}

	nk, exist, err := registry.CreateKey(k, keyName, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	if exist {
		fmt.Println("key already exist")
	} else {
		fmt.Println("Done")
	}

	if err := nk.SetStringValue("", keyPath); err != nil {
		log.Fatal(err)
	}

	defer k.Close()
	defer nk.Close()
}

func pressAnyKey() {
	var pressAnyKey string
	fmt.Println("Press Any Key to Continue")
	n, _ := fmt.Scanln(&pressAnyKey)
	_ = n
}
