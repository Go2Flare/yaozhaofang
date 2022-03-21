package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //"_"代码不直接使用包，底层连接要使用
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
	"time"
)

//创建全局结构体
type Student struct {
	Id    int    //成为默认主键，主键索引，查询速度快
	Name  string `gorm:"size:100;default:'xiaoming'"`
	Age   int
	Class int       `gorm:"not null"`
	Join  time.Time `gorm:"type:timestamp"`
}

/* 用户 table_name = user */
type User struct {
	ID            int           //用户编号
	Name          string        `gorm:"size:32;unique"`  //用户名
	Password_hash string        `gorm:"size:128" `       //用户密码加密的
	Mobile        string        `gorm:"size:11;unique" ` //手机号
	Real_name     string        `gorm:"size:32" `        //真实姓名  实名认证
	Id_card       string        `gorm:"size:20" `        //身份证号  实名认证
	Avatar_url    string        `gorm:"size:256" `       //用户头像路径       通过fastdfs进行图片存储
	Houses        []*House      //用户发布的房屋信息  一个人多套房
	Orders        []*OrderHouse //用户下的订单       一个人多次订单
}

/* 房屋信息 table_name = house */
type House struct {
	gorm.Model                    //房屋编号
	UserId          uint          //房屋主人的用户编号  与用户进行关联
	AreaId          uint          //归属地的区域编号   和地区表进行关联
	Title           string        `gorm:"size:64" `                 //房屋标题
	Address         string        `gorm:"size:512"`                 //地址
	Room_count      int           `gorm:"default:1" `               //房间数目
	Acreage         int           `gorm:"default:0" json:"acreage"` //房屋总面积
	Price           int           `json:"price"`
	Unit            string        `gorm:"size:32;default:''" json:"unit"`               //房屋单元,如 几室几厅
	Capacity        int           `gorm:"default:1" json:"capacity"`                    //房屋容纳的总人数
	Beds            string        `gorm:"size:64;default:''" json:"beds"`               //房屋床铺的配置
	Deposit         int           `gorm:"default:0" json:"deposit"`                     //押金
	Min_days        int           `gorm:"default:1" json:"min_days"`                    //最少入住的天数
	Max_days        int           `gorm:"default:0" json:"max_days"`                    //最多入住的天数 0表示不限制
	Order_count     int           `gorm:"default:0" json:"order_count"`                 //预定完成的该房屋的订单数
	Index_image_url string        `gorm:"size:256;default:''" json:"index_image_url"`   //房屋主图片路径
	Facilities      []*Facility   `gorm:"many2many:house_facilities" json:"facilities"` //房屋设施   与设施表进行关联
	Images          []*HouseImage `json:"img_urls"`                                     //房屋的图片   除主要图片之外的其他图片地址
	Orders          []*OrderHouse `json:"orders"`                                       //房屋的订单    与房屋表进行管理
}

/* 区域信息 table_name = area */ //区域信息是需要我们手动添加到数据库中的
type Area struct {
	Id     int      `json:"aid"`                  //区域编号     1    2
	Name   string   `gorm:"size:32" json:"aname"` //区域名字     昌平 海淀
	Houses []*House `json:"houses"`               //区域所有的房屋   与房屋表进行关联
}

/* 设施信息 table_name = "facility"*/ //设施信息 需要我们提前手动添加的
type Facility struct {
	Id     int      `json:"fid"`     //设施编号
	Name   string   `gorm:"size:32"` //设施名字
	Houses []*House //都有哪些房屋有此设施  与房屋表进行关联的
}

/* 房屋图片 table_name = "house_image"*/
type HouseImage struct {
	Id      int    `json:"house_image_id"`      //图片id
	Url     string `gorm:"size:256" json:"url"` //图片url     存放我们房屋的图片
	HouseId uint   `json:"house_id"`            //图片所属房屋编号
}

/* 订单 table_name = order */
type OrderHouse struct {
	gorm.Model            //订单编号
	UserId      uint      `json:"user_id"`       //下单的用户编号   //与用户表进行关联
	HouseId     uint      `json:"house_id"`      //预定的房间编号   //与房屋信息进行关联
	Begin_date  time.Time `gorm:"type:datetime"` //预定的起始时间
	End_date    time.Time `gorm:"type:datetime"` //预定的结束时间
	Days        int       //预定总天数
	House_price int       //房屋的单价
	Amount      int       //订单总金额
	Status      string    `gorm:"default:'WAIT_ACCEPT'"` //订单状态
	Comment     string    `gorm:"size:512"`              //订单评论
	Credit      bool      //表示个人征信情况 true表示良好
}

//type Student struct {
//	gorm.Model //成为默认主键，主键索引，查询速度快
//	Name string
//	Age  int
//}

//创建全局连接池句柄
var GlobalConn *gorm.DB

//初始化配置, 从vip里取数据
func InitConfig() {
	workDir, err := os.Getwd()
	//workDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/conf")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	//初始化vip
	InitConfig()

	fmt.Println("--------------this is 02mapInterface-----------")

	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	fmt.Println(dsn)
	//	连接数据库，--格式：用户名：密码@协议（IP：port）/数据库名
	conn, err := gorm.Open("mysql", dsn)
	//conn, err := gorm.Open("mysql", "root:4.234.23123@tcp(localhost:3306)/search_house?parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm.Open err:", err)
		return
	}
	defer conn.Close()

	GlobalConn = conn
	//初始数
	GlobalConn.DB().SetMaxIdleConns(10)
	//最大数
	GlobalConn.DB().SetMaxOpenConns(100)

	//不要复数表名
	GlobalConn.SingularTable(true)

	//借助gorm创建数据库表
	//fmt.Println(GlobalConn.AutoMigrate(new(Student)).Error)
	//fmt.Println(GlobalConn.AutoMigrate(new(Area)).Error)
	//fmt.Println(GlobalConn.AutoMigrate(new(Facility)).Error)

	//1.插入数据
	//InsertData1()
	//	2.查询数据
	searchData()
	//	3。更新数据
	//UpdateData()
	//4.删除数据
	//DeleteData()
}

//数据初始化
func InsertData() {
	// 先创建数据
	var stu Student
	stu.Name = "zhangsan"
	stu.Age = 100
	//	插入（创建）数据
	fmt.Println(GlobalConn.Create(&stu).Error)
}
func InsertData1() {
	// 先创建数据
	//var stu Student
	var a Area
	a.Name = "天河区"
	//	插入（创建）数据
	fmt.Println(GlobalConn.Create(&a).Error)
}

func searchData() {
	//var stu Student
	//GlobalConn.First(&stu)
	//
	//GlobalConn.Select("name,age").First(&stu)
	//GlobalConn.First(&stu)
	//var stu []Student
	//GlobalConn.Find(&stu)
	//fmt.Println(stu)
	//GlobalConn.Select("name,age").Find(&stu)

	var user User
	//select name, age from student where name = ‘lisi’;
	GlobalConn.Table("user").Where("mobile = ?", "13790887214").Find(&user)
	//GlobalConn.Where("name = ?", "liuqi").Select("name, age").Find(&stu)
	fmt.Println(user)

	//select name, age from student where name = ‘lisi’ and age = 22;
	//GlobalConn.Select("name, age").Where("name = ?", "lisi").
	//	Where("age = ?", 22).Find(&stu)
	//GlobaelConn.Select("name,age").Where("name = ?", "list").Find(&stu)	GlobalConn.Select("name, age").Where("name = ? and age = ?", "lisi", 22).Find(&stu)

	//fmt.Println(stu)
}

func UpdateData() {
	var stu Student
	stu.Name = "wangwu"
	stu.Age = 89
	//无数据插入
	//fmt.Println(GlobalConn.Save(&stu).Error)

	fmt.Println(GlobalConn.Model(new(Student)).Where("age = ?", 100).
		//Update("name", "lisi").Error)//单条
		Updates(map[string]interface{}{"name": "wocao", "age": 180}).Error) //多条

	//Model(new(Student)//: 指定更新 “student” 表
	//Where("name = ?", "zhaoliu")//： 指定过滤条件。
	//Update("name", "lisi").Error)//：指定 把 “zhaoliu” 更新成 “lisi”
}

func DeleteData() {
	fmt.Println(GlobalConn.Where("name = ?", "liuqi").Delete(new(Student)).Error)
}
