package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	Proxy      string
	Prefix     string
	StorageDir string

	Sign     string
	Username string
	Password string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadServer()
	LoadIO()
	LoadJWT()
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadIO() {
	sec, err := Cfg.GetSection("io")
	if err != nil {
		log.Fatalf("Fail to get section 'request': %v", err)
	}
	Proxy = sec.Key("PROXY").MustString("images")
	Prefix = sec.Key("PREFIX").MustString("http://localhost:8000/")
	StorageDir = sec.Key("STORAGE_DIR").MustString("images")
}

func LoadJWT() {
	sec, err := Cfg.GetSection("jwt")
	if err != nil {
		log.Fatalf("Fail to get section 'request': %v", err)
	}

	Sign = sec.Key("SIGN").MustString("easyio-key")
	Username = sec.Key("USERNAME").MustString("root")
	Password = sec.Key("PASSWORD").MustString("123456")
}
