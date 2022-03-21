package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//	改成直接在终端输入host和password
	View()
	//	TODO 其他测试
}

func cmdClear() {
	//清空CMD终端
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	//清除linux终端
	cmd = exec.Command("cmd", "/c", "clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// View 界面视图
func View() {
	var set1 int

before:
	for {
		fmt.Println("该程序能测试容器中连接msyql或redis：")
		fmt.Println("请输入需要连接的数据库：")
		fmt.Println("0: 退出       1: mysql        2: redis        ")
		fmt.Scanln(&set1)
		switch set1 {
		case 1:
			//mysql 默认值
			var username, password, host, port, database, charset string = "root", "", "127.0.0.1", "3306", "test", "utf8"
		mysqlBefore:
			fmt.Println("输入0重置录入，输入-1返回主菜单")
			fmt.Println("请输入mysql username(默认：root):")
			fmt.Scanln(&username)
			switch username{
			case "0":
				cmdClear()
				goto mysqlBefore
			case "-1":
				cmdClear()
				goto before
			}
			fmt.Println("请输入mysql password(默认：空):")
			fmt.Scanln(&password)
			switch password{
			case "0":
				cmdClear()
				goto mysqlBefore
			case "-1":
				cmdClear()
				goto before
			}
			fmt.Println("请输入mysql host(默认：127.0.0.1):")
			fmt.Scanln(&host)
			switch host{
			case "0":
				cmdClear()
				goto mysqlBefore
			case "-1":
				cmdClear()
				goto before
			}
			fmt.Println("请输入mysql port(默认：3306)")
			fmt.Scanln(&port)
			switch port{
			case "0":
				cmdClear()
				goto mysqlBefore
			case "-1":
				cmdClear()
				goto before
			}
			fmt.Println("请输入mysql database(默认：test):")
			fmt.Scanln(&database)
			switch database{
			case "0":
				cmdClear()
				goto mysqlBefore
			case "-1":
				cmdClear()
				goto before
			}
			fmt.Println("请输入mysql charset(默认：utf8):")
			fmt.Scanln(&charset)
			switch charset{
			case "0":
				cmdClear()
				goto mysqlBefore
			case "-1":
				cmdClear()
				goto before
			}
			//构建配置
			newMysqlConfig := MysqlConfig{
				username, password, host, port, database, charset,
			}
			cmdClear()
			InitMysql(newMysqlConfig)
			defer GlobalDB.Close()
			anything := ""
			fmt.Println("输入任意键重新录入:")
			fmt.Scanln(&anything)
			cmdClear()
		case 2:
			//mysql 默认值
			var network, address, password string = "tcp", "127.0.0.1:6379", ""
		redisBefore:
			fmt.Println("输入0重置录入")
			fmt.Println("请输入redis network(默认：tcp):")
			fmt.Scanln(&network)
			switch network{
			case "0":
				cmdClear()
				goto redisBefore
			case "-1":
				cmdClear()
				goto before
			}
			fmt.Println("请输入redis address(默认：127.0.0.1:6379):")
			fmt.Scanln(&address)
			switch address{
			case "0":
				cmdClear()
				goto redisBefore
			case "-1":
				cmdClear()
				goto before
			}
			fmt.Println("请输入redis password(默认：空):")
			fmt.Scanln(&password)
			switch password{
			case "0":
				cmdClear()
				goto redisBefore
			case "-1":
				cmdClear()
				goto before
			}
			//构建配置
			newRedisConfig := RedisConfig{
				network, address, password,
			}
			cmdClear()
			InitRedis(newRedisConfig)
			anything := ""
			fmt.Println("输入任意键重新录入:")
			fmt.Scanln(&anything)
			cmdClear()
		case 0:
			//退出
			fmt.Println("Bye~~")
			return
		default:
			fmt.Println("输入有误，请重新输入")
			cmdClear()
		}
	}
}


