package utils

import "github.com/ipfs/go-log/v2"

var Logger = log.Logger("cryptomunt")

func InitLogger() {
	log.SetAllLoggers(log.LevelWarn)
	err := log.SetLogLevel("cryptomunt", "info")
	//err := log.SetLogLevel("cryptomunt", "fatal")
	if err != nil {
		return
	}
}
