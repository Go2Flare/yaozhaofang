package handler

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"go_code/yaozhaofang/service/register/model"
	register "go_code/yaozhaofang/service/register/proto/register"
	"go_code/yaozhaofang/service/register/utils"
	"log"
	"math/rand"
	_ "os"
	"time"
)

type Register struct{}

func CreateClient (accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func (e *Register) SmsCode(ctx context.Context, req *register.Request, rsp *register.Response) error  {
	//查询手机号是否注册过
	err := model.CheckMobile(req.Mobile)
	if err == nil{
		rsp.Errno = utils.RECODE_USERONERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_USERONERR)
		return errors.New("用户已注册")
	}
	//从redis中获取到存储的图片验证码
	rnd, err := model.GetImgCode(req.Uuid)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}

	//判断输入的图片验证码是否正确
	if req.Text != rnd {
		rsp.Errno = utils.RECODE_IMGSERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_IMGSERR)
		//返回自定义的error数据
		return errors.New("验证码输入错误")
	}

	//调用API
	AKID, AKSecret := model.InitSmsApi()
	client, err := CreateClient(tea.String(AKID), tea.String(AKSecret))
	if err != nil {
		return err
	}
	//获取6位数随机码
	myRnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06d", myRnd.Int31n(1000000))
	// 生成request
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName: tea.String("耀斑小记"),
		TemplateCode: tea.String("SMS_230640005"),
		PhoneNumbers: tea.String(req.Mobile),
		TemplateParam: tea.String(`{"code":`+vcode+`}`),
	}
	// 调用阿里云API
	_, err = client.SendSms(sendSmsRequest)
	//如果不成功
	if err!=nil{
		rsp.Errno = utils.RECODE_SMSERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		log.Printf("client.SendSms : %v", err)
		return errors.New("发送短信失败")
	}

	//存储短信验证码 存redis中
	err = model.SaveSmsCode(req.Mobile, vcode)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return nil
}


//注册校验
func (e *Register)Register(ctx context.Context, req *register.RegRequest, rsp*register.RegResponse) error{
	//实现具体的业务 把数据存储到mysql中 校验短信验证码是否正确

	//查库校验校验短信验证码会否正确
	smsCode,err := model.GetSmsCode(req.Mobile)
	if err != nil {
		//短信验证码查库错误
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	if smsCode != req.SmsCode{
		//短信验证码校验错误
		rsp.Errno = utils.RECODE_SMSERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		return errors.New("短信验证码错误")
	}
	//存储用户数据到Mysql上
	//给密码加密
	pwdByte := sha256.Sum256([]byte(req.Password))
	pwd_hash := string(pwdByte[:])
	//要把sha256得到的数据转换之后存储转换16进制的
	pwdHash := fmt.Sprintf("%x",pwd_hash)
	err = model.SaveUser(req.Mobile,pwdHash)
	fmt.Printf("model.SaveUser : %v\n", err)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		log.Printf("model.SaveUser : %v", err)
		return err
	}
	fmt.Println("已保存注册用户信息", req.Mobile, pwdHash)
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return nil
}

//登录
func (e *Register)Login(ctx context.Context, req*register.RegRequest, rsp*register.RegResponse) error{
	//注释之后添加了验证码逻辑会解开
	//查询输入手机号和密码是否正确  mysql
	//给密码加密
	pwdByte := sha256.Sum256([]byte(req.Password))
	pwd_hash := string(pwdByte[:])
	//要把sha256得到的数据转换之后存储  转换16进制的
	pwdHash := fmt.Sprintf("%x",pwd_hash)

	user,err := model.CheckUser(req.Mobile, pwdHash)

	if err != nil {
		rsp.Errno = utils.RECODE_LOGINERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_LOGINERR)
		fmt.Println("register微服务查库找不到该用户,err",err)
		return err
	}

	//查询成功  登录成功  把用户名存储到session中  把用户名传给web端
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	rsp.Name = user.Name
	return nil
}

