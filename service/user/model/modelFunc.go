package model

//以下准备转移为register服务
//校验图片验证码
//func CheckImgCode(uuid, imgCode string) bool {
//	//	链接redis
//	/*conn,err := redis.Dial("tcp", "localhost:6379")
//	if err!=nil{
//		fmt.Println("redis Dial err:", err)
//		return false
//	}*/
//	conn := GlobalRedis.Get()
//	defer conn.Close()
//
//	//	查询redis数据
//	inputCode, err := redis.String(conn.Do("get", uuid))
//	if err != nil {
//		fmt.Println("查询redis数据出错：", err)
//		return false
//	}
//
//	//	返回校验结果
//	return inputCode == imgCode
//}
////储存短信验证码
//func SaveSmsCode(phone, code string) error {
//	//	从连接池里取一条数据
//	conn := GlobalRedis.Get()
//	defer conn.Close()
//	//	储存验证码到redis 中
//	_, err := conn.Do("setex", phone+"_code", 60*3, code)
//	return err
//
//}
//
////校验短信验证码
//func CheckSmsCode(phone, code string) error{
////	连接redis
//	conn:=GlobalRedis.Get()
//
////	根据key获取Value
//	smsCode, err :=redis.String(conn.Do("get", phone+"_code"))
//	if err!=nil{
//		fmt.Println("redis get phone_code err",err)
//		return err
//	}
////	验证码匹配
//	if smsCode != code {
//		return errors.New("短信验证码匹配失败")
//	}
//	return nil
//}
//// 注册用户信息,写 MySQL 数据库.
//func RegisterUser(mobile, pwd string) error {
//	var user User
//	user.Name = mobile		// 默认使用手机号作为用户名
//
//	// 使用 md5 对 pwd 加密
//	m5 := md5.New()			// 初始md5对象
//	m5.Write([]byte(pwd))			// 将 pwd 写入缓冲区
//	pwd_hash := hex.EncodeToString(m5.Sum(nil))	// 不使用额外的秘钥
//
//	user.Password_hash = pwd_hash
//
//	// 插入数据到MySQL
//	return GlobalConn.Create(&user).Error
//}


//需要操作数据库的操作

//获取用户信息
func GetUserInfo(userName string)(User,error){
	//连接数据库
	var user User
	err := GlobalDB.Table("user").Where("name = ?",userName).Find(&user).Error
	return user,err
}

//更新用户名
func UpdateUserName(oldName,newName string)error{
	//更新  链式调用
	return GlobalDB.Model(new(User)).
		Where("name = ?",oldName).
		Update("name",newName).Error
	//return GlobalDB.Table("user").
	//	Where("name = ?",oldName).
	//	Update("name",newName).Error
}

////存储用户头像   更新
//func SaveUserAvatar(userName,avatarUrl string)error{
//	return GlobalDB.Model(new(User)).Where("name = ?",userName).Update("avatar_url",avatarUrl).Error
//}
//根据用户名更新用户头像
func UpdateAvatar(userName, avatar string) error{
	//update user set avatar_url = avatar where name = userName
	return GlobalDB.Model(new(User)).Where("name = ?", userName).Update("avatar_url", avatar).Error
}

//存储用户真实姓名
func SaveRealName(userName,realName,idCard string)error{
	return GlobalDB.Model(new(User)).Where("name = ?",userName).
		Updates(map[string]interface{}{"real_name":realName,"id_card":idCard}).Error
}