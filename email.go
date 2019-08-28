package tools

import (
	"github.com/astaxie/beego/utils"
	"github.com/rs/xid"
)

func GetVerifyCode(account, req string) (str string, err error) {
	str = xid.New().String()[14:]

	// 创建一个字符串变量，存放相应配置信息
	config :=
		`{"username":"2387805574@qq.com","password":"henuqnarpnucdjci","host":"smtp.qq.com","port":587}`
	// 通过存放配置信息的字符串，创建Email对象
	temail := utils.NewEMail(config)

	temail.To = []string{req}
	temail.From = "2387805574@qq.com"
	temail.Subject = "GoBlog-用户验证"
	temail.HTML = `<html>
		<head>
		</head>
	    	 <body>
			   <div>您的验证码为：` + str + ` </a></div>
	     	</body>
	 	</html>`

	err = temail.Send()
	if err != nil {
		return
	}
	//存入redis
	/*err = app.Redis.Cmd("SET", u.getCachKey(account), str, "EX", u.emailExp).Err
	if err != nil {
		err = errors.Wrap(err, "fail to set redis")
		return
	}*/
	return

}
