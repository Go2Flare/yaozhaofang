package model

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

//获取图片验证码
func GetImgCode(uuid string)(string,error){
	//获取redis链接
	conn := GlobalRedis.Get()
	//fmt.Println(redis.String(conn.Do("get", uuid)))
	//获取数据
	return redis.String(conn.Do("get", uuid))
}

//存短信验证码
func SaveSmsCode(phone, vcode string)error{
	//获取redis链接
	conn := GlobalRedis.Get()
	//存储短信验证码，有效期5分钟
	_,err := conn.Do("setex", phone+"_code", 60 * 5, vcode)
	return err
}

//存储用户名和密码  mysql
func SaveUser(mobile, password_hash string)error{
	//链接数据库  gorm插入数据
	var user User
	user.Mobile = mobile
	user.Password_hash = password_hash
	user.Name = mobile

	return GlobalDB.Create(&user).Error
}

//返回存入数据库的验证码
func GetSmsCode(phone string)(string,error){
	InitRedis()
	//获取redis链接
	conn := GlobalRedis.Get()
	//获取数据
	res, err := redis.String(conn.Do("get", phone+"_code"))
	fmt.Printf("code =%v ,res = %v, conn.Do err :%v\n",phone+"_code",res, err)
	return res, err
}

//查看用户是否存在
func CheckMobile(mobile string) error{
	var user User
	err := GlobalDB.Where("mobile = ?", mobile).Find(&user).Error
	if err != nil{
		log.Println("用户不存在", mobile)
	}
	return err
}

//校验登录信息
//等短信验证码完善后使用这个，校验用户的密码
func CheckUser(mobile, pwd_hash string) (User,error){
	//连接数据库
	var user User
	err := GlobalDB.Where("mobile = ?", mobile).Where("password_hash = ?", pwd_hash).Find(&user).Error
	return user,err
}