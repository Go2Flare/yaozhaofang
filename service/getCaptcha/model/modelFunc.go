package model

import "fmt"

//储存图片id到redis数据库
func SaveImgCode(code, uuid string) error {
	//连接reids连接池
	conn := RedisPool.Get()
	//	推荐用Do()函数
	//	redis中设置验证码过期时间5分钟
	res, err := conn.Do("setex", uuid, 60*5 ,code)
	fmt.Println(res)
	return err
}