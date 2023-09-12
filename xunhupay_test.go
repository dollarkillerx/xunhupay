package xunhupay

import (
	"fmt"
	"testing"
)

func TestHuPiClient_Execute(t *testing.T) {
	//Pay()
	/**
	{
	  "openid": 202112602216,
	  "url_qrcode": "https://api.dpweixin.com/payments/wechat/qrcode?id=202112602216&nonce_str=3951642032&time=1694522330&appid=20211114110&hash=267a53829932248a56f8c55bc95c8ac8",
	  "url": "https://api.dpweixin.com/payments/wechat/index?id=202112602216&nonce_str=6253419032&time=1694522330&appid=20211114110&hash=71fc09e380d19fb127b5274640d2038d",
	  "errcode": 0,
	  "errmsg": "success!",
	  "hash": "129a520166aeb62a70363fc243e14187"
	}
	*/

	appId := "20211114110"                          //Appid
	appSecret := "ed722800855944bfef60e2709b609b5d" //密钥

	client := NewHuPi(&appId, &appSecret) //初始化调用

	sign := client.Sign(map[string]string{
		"openid":     "202112602216",
		"url_qrcode": "https://api.dpweixin.com/payments/wechat/qrcode?id=202112602216&nonce_str=3951642032&time=1694522330&appid=20211114110&hash=267a53829932248a56f8c55bc95c8ac8",
		"url":        "https://api.dpweixin.com/payments/wechat/index?id=202112602216&nonce_str=6253419032&time=1694522330&appid=20211114110&hash=71fc09e380d19fb127b5274640d2038d",
		//"errcode":    "0",
		"errmsg": "success!",
		//"hash": "129a520166aeb62a70363fc243e14187"
	})
	fmt.Println(sign)

}

func TestHuPiClient_Execute1(t *testing.T) {
	Pay()
	// &{202112602280 https://api.dpweixin.com/payments/wechat/qrcode?id=202112602280&nonce_str=8291435366&time=1694523368&appid=20211114110&hash=88c274cca0d91028e6869ea697edc541 https://api.dpweixin.com/payments/wechat/index?id=202112602280&nonce_str=2639143568&time=1694523368&appid=20211114110&hash=b7ecc2684810ae827743a459b534967d 0 success! 52dcf7544a83662fbb90d110efa7e8ea}
}

func TestHuPiClient_Execute2(t *testing.T) {
	Query()
}
