package main

import (
	"example.com/dotfiles/pkg/backup"
	"example.com/dotfiles/pkg/stdio"
	"flag"
	"log"
)

func main() {
	inputFilePath := parseFlags()
	if len(inputFilePath) == 0 {
		log.Println("No config file provided, looking in ./config/")
		inputFilePath = "./config/config.json"
	}
	ioSvc := stdio.NewIOSvc()
	backupSvc := backup.NewBackupSvc(ioSvc)
	configOpts, err := backupSvc.ReadCfgFile(inputFilePath)
	err = backupSvc.ProcessCfgFile(configOpts)
	log.Println(err)
}

func parseFlags() string {
	var inputFilePath string
	flag.StringVar(&inputFilePath, "f", "", "Enter the full path to the config.json file")
	flag.Parse()
	return inputFilePath
}
