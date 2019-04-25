package main

import (
	"log"
	"os"

	"golang.org/x/sys/windows/registry"
)

func main() {
	keyName := os.Args[1] + ".exe"
	keyPath := os.Args[2]
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}

	nk, _, err := registry.CreateKey(k, keyName, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}

	if err := nk.SetStringValue("", keyPath); err != nil {
		log.Fatal(err)
	}

	defer k.Close()
	defer nk.Close()
}
