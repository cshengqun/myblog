package env

import (
	"encoding/json"
	"os"
	"github.com/cshengqun/asyncLog"
	"log"
	"time"
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
	MonConf  LogConfig `json:MonConf`
	PageSize int `json:PageSize`
	ProjectPath string `json:ProjectPath`
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
	Monitor *asyncLog.ALog
}

func (Env *TEnv) Report(moduleName string, bizSeqNo string, consumerSeqNo string, logPoint string, statisticText string, msg string) {
	threadID := os.Getpid()
	tcurTime := time.Now()
	curTime := tcurTime.Format("2006-01-02 15:04:05")
	Env.Monitor.Info("[%s][%d][%s][%s][%s][%s][%s][%s]", curTime, threadID, moduleName, bizSeqNo, consumerSeqNo, logPoint, statisticText, msg)
}

var Env TEnv

func init() {
	file, err := os.Open("/root/workshop/go/src/github.com/cshengqun/myblog/conf/myblog.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Env.Conf)
	if err != nil {
		panic(err)
	}

	Env.Logger = asyncLog.NewLogger(Env.Conf.LogConf.FileName, Env.Conf.LogConf.Level, Env.Conf.LogConf.ChanSize, Env.Conf.LogConf.ThreadCnt, log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	if Env.Conf.LogConf.FileCnt != 0 {
		Env.Logger.SetLogCnt(Env.Conf.LogConf.FileCnt)
	}

	if Env.Conf.LogConf.FileSize != 0 {
		Env.Logger.SetFileSize(Env.Conf.LogConf.FileSize)
	}

	if Env.Conf.LogConf.WriterLv != 0 {
		Env.Logger.SetWriterLv(Env.Conf.LogConf.WriterLv)
	}

	Env.Monitor = asyncLog.NewLogger(Env.Conf.MonConf.FileName, Env.Conf.MonConf.Level, Env.Conf.MonConf.ChanSize, Env.Conf.MonConf.ThreadCnt, 0)
	if Env.Conf.MonConf.FileCnt != 0 {
		Env.Logger.SetLogCnt(Env.Conf.MonConf.FileCnt)
	}

	if Env.Conf.MonConf.FileSize != 0 {
		Env.Logger.SetFileSize(Env.Conf.MonConf.FileSize)
	}

	if Env.Conf.MonConf.WriterLv != 0 {
		Env.Logger.SetWriterLv(Env.Conf.MonConf.WriterLv)
	}

}
