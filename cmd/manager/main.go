package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"time"
)

func config() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	var (
		logLength = os.Getenv("LOG_FILENAME_LENGTH")
		logDir    = os.Getenv("LOG_DIR")
	)
	logfile, err := os.Create(fmt.Sprintf("%s/log-manager-%d.log", logDir, time.Now().UnixMilli()))
	if err != nil {
		panic(err)
	}
	if logLength == "long" {
		log.SetFlags(log.LstdFlags | log.Llongfile)
	} else {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
	mw := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(mw)
	log.Print("configuration loaded successfully")
}

func main() {
	config()
}
