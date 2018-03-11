package env

import (
	"encoding/json"
	"os"
	"github.com/cshengqun/asyncLog"
)

type Config struct {
	Ip string `json:"ip"`
	Port string `json:port`
	DbUser string `json:dbUser`
	DbPasswd string `json:dbPasswd`
	DbName string `json:dbName`
	DbIp string `json:dbIp`
	DbPort string `json:dbPort`
	Account string `json:account`
	Password string `json:password`
	LogConf  LogConfig `json:LogConf`
}

type LogConfig struct {
	FileName    string
	FileCnt     int
	FileSize    int64
	Level       int
	ChanSize    int
	ThreadCnt   int
	WriterLv    int
}

type TEnv struct {
	Conf   Config
	Logger *asyncLog.ALog
}

var Env TEnv

func init() {
	file, err := os.Open("../conf/myblog.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Env.Conf)
	if err != nil {
		panic(err)
	}

	
	Env.Logger = asyncLog.NewLogger(Env.Conf.LogConf.FileName, Env.Conf.LogConf.Level, Env.Conf.LogConf.ChanSize, Env.Conf.LogConf.ThreadCnt)
	if Env.Conf.LogConf.FileCnt != 0 {
		Env.Logger.SetLogCnt(Env.Conf.LogConf.FileCnt)
	}

	if Env.Conf.LogConf.FileSize != 0 {
		Env.Logger.SetFileSize(Env.Conf.LogConf.FileSize)
	}

	if Env.Conf.LogConf.WriterLv != 0 {
		Env.Logger.SetWriterLv(Env.Conf.LogConf.WriterLv)
	}

}
