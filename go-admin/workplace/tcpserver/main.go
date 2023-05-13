package main

import (
	// "bufio"
	// "os"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	Dbname := "go-admin"
	timeout := "1000ms"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("数据库连接失败,error=" + err.Error())
	}
	DB = db

}

type FyEquipment struct {
	Id         int
	Equipname  string
	Equipcode  string
	Username   string
	Area       string
	Type       string
	Status     string
	Faultcount string
	// Remark     string
	// CreatedAt string
}

func (FyEquipment) TableName() string {
	return "fy_equipment_copy1"
}

type FyFault struct {
	Id           int
	Equipname    string
	Equipcode    string
	Area         string
	Username     string
	Type         string
	Status       string
	Handlestatus string
	// Reason       string
	Faulttime time.Time
	Faultcode string
	// Remark    string
}

func (FyFault) TableName() string {
	return "fy_fault_copy1"
}

// 接受数据格式示例{"Equipcode":"equipcode16","Status":"1"}
type Acceptmsg struct {
	Equipcode    string
	Status       string
	Monitordata1 string
	Monitordata2 string
	Monitordata3 string
	Monitordata4 string
	Monitordata5 string
	Monitordata6 string
}

// TCP 服务端
func process(conn net.Conn) {
	// 函数执行完之后关闭连接
	defer conn.Close()
	// 输出主函数传递的conn可以发现属于*TCPConn类型, *TCPConn类型那么就可以调用*TCPConn相关类型的方法, 其中可以调用read()方法读取tcp连接中的数据
	fmt.Printf("服务端: %T\n", conn)
	for {
		var buf [512]byte
		// 将tcp连接读取到的数据读取到byte数组中, 返回读取到的byte的数目
		n, err := conn.Read(buf[:])
		if err != nil {
			// 从客户端读取数据的过程中发生错误
			fmt.Println("read from client failed, err:", err)
			break
		}
		// //给客户端返回ok数据
		// okmsg := "连接ok\n"
		// conn.Write([]byte(okmsg))
		recvStr := string(buf[:n])
		fmt.Println("服务端收到客户端发来的数据：", recvStr, "长度为：", n)
		var amsg Acceptmsg
		var equipment FyEquipment
		var fault FyFault
		err = json.Unmarshal([]byte(recvStr), &amsg)
		if err != nil {
			fmt.Println("Error:", err)
		}
		if amsg.Status == "1" {
			amsg.Status = "发生故障"
		} else if amsg.Status == "2" {
			amsg.Status = "存在风险"
		} else {
			amsg.Status = "0"
		}
		//设立返回信息
		backmsg := ""
		// fmt.Println("设备编码为：", amsg.Equipcode, "设备状态为：", amsg.Status)
		if amsg.Status != "0" {
			//创建、赋予设备和故障结构体值

			DB.Take(&equipment, "equipcode = ?", amsg.Equipcode)
			DB.Last(&fault)
			// fmt.Println("amsg.Status:", amsg.Status)
			//获得故障编号和加一
			fault.Id++
			nfaultcode := fault.Id
			// fmt.Println("nnnnnnnnnnnumber:", nfaultcode)
			//正常运行避免故障叠加
			if equipment.Status == "正常运行" {
				fault.Equipname = equipment.Equipname
				fault.Equipcode = equipment.Equipcode
				fault.Area = equipment.Area
				fault.Username = equipment.Username
				fault.Type = equipment.Type
				fault.Status = "发生故障"
				fault.Handlestatus = "未处理"
				fault.Faulttime = time.Now()
				fault.Faultcode = strconv.Itoa(nfaultcode)
				// fmt.Println("equipment", equipment)
				// fmt.Println("fault", fault)
				if amsg.Status == "发生故障" {
					DB.Create(&fault)
					DB.Model(&equipment).Where("equipcode = ?", amsg.Equipcode).Update("faultcount", nfaultcode)
				}

				DB.Model(&equipment).Where("equipcode = ?", amsg.Equipcode).Update("status", amsg.Status)
				backmsg = "成功登记异常"
				fmt.Println(backmsg)
				conn.Write([]byte(backmsg))
			} else {
				backmsg = "故障已登记，请勿重复登记"
				fmt.Println(backmsg)
				conn.Write([]byte(backmsg))
			}
		} else {
			backmsg = "传输格式有误"
			fmt.Println(backmsg)
			conn.Write([]byte(backmsg))
		}
		amsg = Acceptmsg{}
		equipment = FyEquipment{}
		fault = FyFault{}

		// // 由于是tcp连接所以双方都可以发送数据, 下面接收服务端发送的数据这样客户端也可以收到对应的数据
		// inputReader := bufio.NewReader(os.Stdin)
		// s, _ := inputReader.ReadString('\n')
		// t := strings.Trim(s, "\r\n")
		// // 向当前建立的tcp连接发送数据, 客户端就可以收到服务端发送的数据
		// conn.Write([]byte(t))

	}
}

func main() {
	//ip和端口设置
	//默认ip
	listen_ip := "127.0.0.1"
	Listenport := "20005"
	//暂时无法监听公网ip
	external_ip := strings.Replace(get_external(), "\n", "", -1)
	fmt.Println("公网ip:", external_ip)
	public_ip := GetPulicIP()
	fmt.Println("内网ip:", public_ip)
	listen_ip = public_ip
	listen_ipport := listen_ip + ":" + Listenport

	// 监听当前的tcp连接
	listen, err := net.Listen("tcp", listen_ipport)

	// fmt.Printf("服务端: %T=====%v\n", listen, ipport)
	fmt.Printf("=====服务端监听:%v=====\n", listen_ipport)
	if err != nil {
		fmt.Println("监听失败,错误:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		fmt.Println("当前建立了tcp连接")
		if err != nil {
			fmt.Println("连接失败, err:", err)
			continue
		}
		// 对于每一个建立的tcp连接使用go关键字开启一个goroutine处理
		go process(conn)

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
