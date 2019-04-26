package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func main() {
	var keyName string
	var keyPath string
	if len(os.Args) == 2 {
		filename, format := checker(os.Args[1])
		if format != ".exe" {
			fmt.Println("Your format is", format, "we expecting .exe extension")
			exitProgram("Please use Application that using .exe etension")
		}
		keyName = filename + format
		keyPath = os.Args[1]
	} else if len(os.Args) == 3 {
		keyName = os.Args[1] + ".exe"
		_, format := checker(os.Args[2])
		if format != ".exe" {
			fmt.Println("Your format is", format, "we expecting .exe extension")
			exitProgram("Please use Application that using .exe etension")
		}
		keyPath = os.Args[2]
	} else {
		exitProgram("We need path")
	}

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
