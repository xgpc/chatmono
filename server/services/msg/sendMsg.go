package msg

import (
	"chatmono/services/msg/random"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-redis/redis/v8"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/log"
	"os"
)

type Config struct {
	AccessKeyID     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
}

var Hand *dysmsapi20170525.Client

func Init(conf Config, rdb *redis.Client) {
	hand, err := createClient(&conf.AccessKeyID, &conf.AccessKeySecret)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}
	Hand = hand
	random.Init(rdb)
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func createClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func SendCode(mobile, code string) error {

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(mobile),
		SignName:      tea.String("智慧县域平台"),
		TemplateCode:  tea.String("SMS_228845548"),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	}
	runtime := &util.RuntimeOptions{}

	// 复制代码运行请自行打印 API 的返回值
	r, _err := Hand.SendSmsWithOptions(sendSmsRequest, runtime)
	if _err != nil {
		log.Error(r)
		return _err
	}

	return nil
}

func _main(args []*string) (_err error) {
	client, _err := createClient(tea.String("accessKeyId"), tea.String("accessKeySecret"))
	if _err != nil {
		return _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("smono"),
		TemplateCode:  tea.String("SMS_254310339"),
		TemplateParam: tea.String("{\"code\":\"312123\"}"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
