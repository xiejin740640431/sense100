package util

import (
	"log"
	"os"
	"encoding/json"
	"sense100/config"
)

var Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

var Conf *Configuration

type Configuration struct {
	Server       string
	ImageServer  string
	StaticServer string
	RuntimeModel string
	LogLevel     string
	Mysql        string
	Port         string
	SwaggerJson  string
}

func LoadConf() {
	//Logger.Println(os.Getwd())
	//bytes, err := ioutil.ReadFile("config.json")
	//if err != nil {
	//	Logger.Fatal("load configuration file failed：" + err.Error())
	//}
	Conf = &Configuration{}

	if err := json.Unmarshal([]byte(config.ConfigJson), Conf); err != nil {
		Logger.Fatal("parse config.json failed：" + err.Error())
	}
}
