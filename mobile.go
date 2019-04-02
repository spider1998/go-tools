package routing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rs/xid"
	"io/ioutil"
	"net/http"
	"unsafe"
)

type YunXun struct {
	Sid        string `json:"sid"`        //用户的账号唯一标识
	Token      string `json:"token"`      //用户密钥Auth Token
	Appid      string `json:"appid"`      //应用分配标识
	Templateid string `json:"templateid"` //短信模板
	Param      string `json:"param"`      //参数
	Mobile     string `json:"mobile"`     //手机号
	Uid        string `json:"uid"`        //用户透传ID
}

//--------------------------------------------------------------------------------------------------
/*链接：https://office.ucpaas.com*/
//--------------------------------------------------------------------------------------------------

func GetMobileVerify(uid, mobile string) (code string, err error) {
	var yun YunXun
	vcode := xid.New().String()[14:]
	yun.Sid = "116943babddda930dcd8802a7f6f5bd4"   //用户表示
	yun.Token = "28423c3bc2a1b63b4f432540e5b8cd96" //access token
	yun.Appid = "589910649b5347118abf1888f56a6071" //应用id
	yun.Templateid = "413802"                      //模板id
	yun.Param = vcode + "," + "60"                 //过期时间
	yun.Mobile = mobile                            //手机号
	yun.Uid = uid                                  //服务id
	msg, err := json.Marshal(yun)
	if err != nil {
		fmt.Println(err)
	}
	reader := bytes.NewReader(msg)
	request, err := http.NewRequest("POST", "https://open.ucpaas.com/ol/sms/sendsms", reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	str := (*string)(unsafe.Pointer(&respBytes))

	var dat map[string]string
	if err = json.Unmarshal([]byte(*str), &dat); err != nil {
		return
	}
	//存储redis
	/*err = app.Redis.Cmd("SET", u.getMobileKey(uid), vcode, "EX", u.mobileExp).Err
	if err != nil {
		err = errors.Wrap(err, "fail to set redis")
		return
	}*/
	code = dat["msg"]
	return
}
