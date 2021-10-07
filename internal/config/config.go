package config

import (
	"log"
	"os"
	"strconv"
)

func GetConfig(param string) (string, bool) {
	if val, exist := os.LookupEnv(param); exist {
		return val, exist
	}
	log.Printf("ERROR: no param %s available", param)
	return "", false
}

func GetConfigAsInt(param string) (int, bool) {
	valStr, exist := GetConfig(param)
	if !exist {
		return 0, false
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Printf("ERROR: unable to convert value %s to integer; %+v", valStr, err)
		return 0, false
	}
	return val, true
}
