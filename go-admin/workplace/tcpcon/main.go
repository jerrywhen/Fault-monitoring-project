package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

// 客户端
func main() {
	listen_ip := "127.0.0.1"
	Listenport := "20005"
	//暂时无法监听公网ip
	// external_ip := strings.Replace(get_external(), "\n", "", -1)
	// fmt.Println("公网ip:", external_ip)
	public_ip := GetPulicIP()
	// fmt.Println("内网ip:", public_ip)
	listen_ip = public_ip
	listen_ipport := listen_ip + ":" + Listenport

	// 连接到服务端建立的tcp连接
	conn, err := net.Dial("tcp", listen_ipport)
	// conn, err := net.Dial("tcp", "192.168.3.2:20000") //kt的服务端
	// conn, err := net.Dial("tcp", "172.26.89.254:20005") //校园网
	// 输出当前建Dial函数的返回值类型, 属于*net.TCPConn类型
	fmt.Printf("客户端: %T\n", conn)
	if err != nil {
		// 连接的时候出现错误
		fmt.Println("err :", err)
		return
	}
	// 当函数返回的时候关闭连接
	defer conn.Close()

	// 创建一个JSON对象
	jsonObj := map[string]interface{}{
		"Equipcode": "equipcode16",
		"Status":    "1",
	}
	// 将JSON对象序列化为JSON格式的字节数组
	jsonBytes, err := json.Marshal(jsonObj)
	if err != nil {
		fmt.Println("序列化JSON失败:", err)
		return
	}
	// 输出序列化后的JSON字节数组
	fmt.Println(string(jsonBytes))
	_, err = conn.Write([]byte(string(jsonBytes))) // 发送数据
	if err != nil {
		return
	}
	// 读取服务端发送的数据
	buf := [512]byte{}
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("recv failed, err:", err)
		return
	}
	fmt.Println("客户端接收服务端发送的数据: ", string(buf[:n]))

	// 获取一个标准输入的*Reader结构体指针类型的变量
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 调用*Reader结构体指针类型的读取方法
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		// 去除掉\r \n符号
		inputInfo := strings.Trim(input, "\r\n")
		// 判断输入的是否是Q, 如果是Q则退出
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		// 读取服务端发送的数据
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println("客户端接收服务端发送的数据: ", string(buf[:n]))

	}
}

func get_external() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := io.ReadAll(resp.Body)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	//s := buf.String()
	return string(content)
}

func GetPulicIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}
