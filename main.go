package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sys/windows/registry"
)

func main() {

	var keyName string //variable, key that we want called from run
	var keyPath string //variable, the Path of the key

	if len(os.Args) == 1 { //if we not using any argument
		keyPath = interfaces()            //get the Path from func interfaces
		keyName, _ = getFileName(keyPath) //Get the filename
		keyName = keyName + ".exe"        //Automatically add .exe to Key
	} else if len(os.Args) == 2 { //if we using 1 argument (Path)
		keyPath = os.Args[1]              //Assign the Path from argument
		keyName, _ = getFileName(keyPath) //Get the filename
		keyName = keyName + ".exe"        //Automatically add .exe to Key
	} else if len(os.Args) == 3 { //if we using 2 argument (Key Name, Path)
		keyName = os.Args[1] + ".exe" //Automatically add .exe to Key
		keyPath = os.Args[2]          //Assign the Path from second argument
	} else {
		exitProgram("Max argument is 2")
	}

	_, format := getFileName(keyPath) //Get the extecsion of file

	//Check the file source extension that must .exe
	if format != ".exe" {
		if format == "" {
			format = "empty"
		}

		fmt.Println("Your format is\"", format, "\"we expecting .exe extension")
		exitProgram("")
	}

	//Writing key to registry
	keyMaker(keyName, keyPath)

	fmt.Println("All Done, you can use \"", keyName[0:(len(keyName)-4)], "\" in Run")
	pressAnyKey()
}

//interfaces return Path
func interfaces() string {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println(
		`
===============================
           RunMe v1.1
================================`)
	fmt.Print("Enter path : ")
	thePath, _ := reader.ReadString('\n')
	thePath = thePath[0:(len(thePath) - 2)] //removing "\n"

	return thePath
}

func getFileName(value string) (filename string, format string) {
	file := filepath.Base(value)
	format = filepath.Ext(value)
	filename = file[0 : len(file)-len(format)]

	return filename, format
}

func keyMaker(keyName string, keyPath string) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
		time.Sleep(time.Second * 3)
	}

	nk, exist, err := registry.CreateKey(k, keyName, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
		time.Sleep(time.Second * 3)
	}

	if exist {
		fmt.Println("key already exist")
	} else {
		fmt.Println("Done")
	}

	if err := nk.SetStringValue("", keyPath); err != nil {
		log.Fatal(err)
		time.Sleep(time.Second * 3)
	}

	defer k.Close()
	defer nk.Close()
}

func exitProgram(isi string) {
	fmt.Println(isi)
	pressAnyKey()
	os.Exit(1)
}

func pressAnyKey() {
	var pressAnyKey string
	fmt.Println("Press Any Key to Continue")
	n, _ := fmt.Scanln(&pressAnyKey)
	_ = n
}
