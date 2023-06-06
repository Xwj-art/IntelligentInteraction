package utils

import (
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func init() {
	file, err := ini.Load("config/config.ini")
	HandleErr(err, "初始化失败")
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {
	// file中读取AppMode，绑定默认为debug
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3030")
	JwtKey = file.Section("server").Key("JwtKey").MustString(":i;jefnaiq893u4")
}

func LoadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("Xwj")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("coisini1215")
	DbName = file.Section("database").Key("DbName").MustString("Ai")
}

func LoadQiniu(file *ini.File) {
	// file中读取AppMode，绑定默认为debug
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("KlE2CscMY89GNjOSQL32dmVApbetYlCEdtq5g3gU")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("XHtSYXuo2NIYtdPB-lJKEHtMVYuv2waI6AgkxhWt")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("xwjginblog")
	QiniuServer = file.Section("qiniu").Key("QiniuServer").MustString("http://rvde0rpcd.hd-bkt.clouddn.com/")
}
