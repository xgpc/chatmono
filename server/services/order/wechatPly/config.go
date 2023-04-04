/**
 * @Author: smono
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2022/10/10 23:39
 */

package wechatPly

type Config struct {
	AppID                      string `yaml:"appID"`
	Mchid                      string `yaml:"mchid"`
	MchKey                     string `yaml:"mchKey"`
	NotifyUrl                  string `yaml:"notifyUrl"`
	MchCertificateSerialNumber string `yaml:"mchCertificateSerialNumber"`
	MchAPIv3Key                string `yaml:"mchAPIv3Key"`
	MchPrivateKeyPath          string `yaml:"mchPrivateKeyPath"`
}
