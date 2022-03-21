package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gomodule/redigo/redis"
)


//校验图片验证码
func CheckImgCode(uuid, imgCode string) bool {
	//	链接redis
	/*conn,err := redis.Dial("tcp", "localhost:6379")
	if err!=nil{
		fmt.Println("redis Dial err:", err)
		return false
	}*/
	conn := GlobalRedis.Get()
	defer conn.Close()

	//	查询redis数据
	inputCode, err := redis.String(conn.Do("get", uuid))
	if err != nil {
		fmt.Println("查询redis数据出错：", err)
		return false
	}

	//	返回校验结果
	return inputCode == imgCode
}

//存储短信验证码
func SaveSmsCode(phone, code string) error {
	//	从连接池里取一条数据
	conn := GlobalRedis.Get()
	defer conn.Close()
	//	储存验证码到redis中
	_, err := conn.Do("setex", phone+"_code", 60*3, code)
	return err

}

//处理登录业务，根据手机号、密码，获取用户名
func Login(mobile, pwd string) (string, error) {
	var user User
	//md5哈希存储
	m5 := md5.New()
	m5.Write([]byte(pwd))
	pwd_hash := hex.EncodeToString(m5.Sum(nil))

	//相当于sql:select name from user where mobile = xx and password_hash =xx
	err := GlobalDB.Where("mobile = ?", mobile).Select("name").
		Where("password_hash = ?", pwd_hash).Find(&user).Error

	fmt.Printf("model/modelFunc Login gorm查库结果, err:%v\n", err)
	return user.Name, err
}

//自己写的不使用hash加密的数据查找数据库, 因为我注册的逻辑还没完善
func LoginNotUseHash(mobile, pwd string) (string, error) {
	var user User
	err := GlobalDB.Where("mobile = ?", mobile).Select("name").Where("password_hash=?",pwd).Find(&user).Error
	return user.Name, err
}

// 获取用户信息
//上述函数，也可以写成。---- go语法写。
func GetUserInfo(userName string) (user User, err error) {
	// 实现SQL: select * from user where name = userName;
	err = GlobalDB.Where("name = ?", userName).First(&user).Error
	return
}

//更新用户名到mysql中
func UpdateUserName(oldName,newName string) error {
	return GlobalDB.Model(new(User)).Where("name = ?", oldName).Update("name",newName).Error
}

//-------------1.移植到微服务---------------
//根据用户名更新用户头像
func UpdateAvatar(userName, avatar string) error{
	//update user set avatar_url = avatar where name = userName
	return GlobalDB.Model(new(User)).Where("name = ?", userName).Update("avatar_url", avatar).Error
}
