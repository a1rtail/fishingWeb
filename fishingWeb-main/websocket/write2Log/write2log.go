package write2Log

import (
	"io"
	"os"
)

const (
	//LOGPATH  这里定义一下存放路径
	LOGPATH = "G:\\go\\pkg\\websocket"
	//FORMAT 随机参量罢了,其实不用也行
	FORMAT = "20211111"
	//LineFeed 换行
	LineFeed = "\r\n"
)

//	以天为基准,存日志路径
//var path = LOGPATH + time.Now().Format(FORMAT) + "/"
// 定义存入路径
var path = LOGPATH  + "/"

//WriteLog 写入
func WriteLog(fileName, msg string) error {
	if !IsExist(path) {
		return CreateDir(path)
	}
	var (
		err error
		f   *os.File
	)

	f, err = os.OpenFile(path+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	_, err = io.WriteString(f, LineFeed+msg)

	defer f.Close()
	return err
}

//CreateDir  创建文件夹
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	os.Chmod(path, os.ModePerm)
	return nil
}

//IsExist  判断文件夹/文件是否存在
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}