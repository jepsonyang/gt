package gtNet

import (
	"net"
	"strings"
)

/*
* 获取本机IP
* @note 创建Conn后,不实际发送数据,即可拿到本地发送数据使用的ip信息
**/
func GetLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0], nil
}
