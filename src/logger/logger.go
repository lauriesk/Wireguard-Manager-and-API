package logger

import (
	"fmt"
	"log"
	"os"
	"github.com/spf13/viper"
)

func LoggerSetup() {
	logDirectory := viper.GetString("SERVER.WORKDIR") + "/logs"

	errCreateDir := os.MkdirAll(logDirectory, 0755) //create log directory if does not exist

	if errCreateDir != nil { //if an error on creating directory
		fmt.Println("Error on creating directory for logger \n", errCreateDir)
		os.Exit(1) //exit program
	}

	file, errLog := os.OpenFile(logDirectory+"/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755) //open file
	if errLog != nil {
		fmt.Println("Error when opening log file \n", errLog)
	} else {
		log.SetOutput(file)
	}
}
