package logs

import (
	"io"
	"log"
	"os"
	"time"
)

//日志路径
const (
	LOGPATH = "logs/"
	FORMAT  = "20060102"
)

//以天为基准记录日志
var path = LOGPATH + time.Now().Format(FORMAT) + "/"

var (
	Trace   *log.Logger //记录所有日志
	Info    *log.Logger //重要信息
	Warning *log.Logger //需要注意的信息
	Error   *log.Logger //非常严重的问题
)

//初始化日志系统，打开日志文件
func InitLog() error {
	var (
		err1 error
		err2 error
		f1   *os.File //礼品领取日志
		f2   *os.File //错误日志
	)
	if !IsExist(path) {
		err := CreateDir(path)
		if err != nil {
			return err //目录创建失败
		}
	}
	f1, err1 = os.OpenFile(path+"TraceLog.logs", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err1 != nil {
		return err1 //TraceLog文件打开失败
	}
	f2, err2 = os.OpenFile(path+"errorsLog.logs", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err2 != nil {
		return err2 //errorsLog文件打开失败
	}

	Trace = log.New(io.MultiWriter(f1, os.Stderr),
		"Trace: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(io.MultiWriter(f2, os.Stderr),
		"Info: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(io.MultiWriter(f2, os.Stderr),
		"Warning: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(f2, os.Stderr),
		"Error: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

//检查文件夹是否存在
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

//创建文件夹
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	os.Chmod(path, os.ModePerm)
	return nil
}
