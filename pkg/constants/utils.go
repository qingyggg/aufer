package constants

import (
	"net"
	"os"
	"strings"
)

func GetIp(key string) string {
	ip := os.Getenv(key)
	if ip == "" {
		ip = "localhost"
	}
	return ip
}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

// UrlInit 调用该函数前，先要加载env
func UrlInit() {
	MySQLDefaultDSN = os.Getenv("MariaDBUser") + ":" + os.Getenv("MariaDBPwd") + "@tcp(" + os.Getenv("MariaDBUrl") + ")/storybook?charset=utf8mb4&parseTime=True&loc=Local"
	MinioEndPoint = os.Getenv("MinioEndPoint")
	MinioAccessKeyID = os.Getenv("MinioAccessKeyID")
	MinioSecretAccessKey = os.Getenv("MinioSecretAccessKey")
	MinioSSL = false
	RedisAddr = os.Getenv("RedisUrl")
	RedisPassword = os.Getenv("RedisPassword")
	MongoDefaultDSN = "mongodb://" + os.Getenv("MongoUser") + ":" + os.Getenv("MongoPwd") + "@" + os.Getenv("MongoUrl") + "/?connect=direct"
	Host = os.Getenv("Host")
	EtcdAddress = os.Getenv("EtcdAddress")
	SecretKey = os.Getenv("SecretKey")
}
