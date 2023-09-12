package xunhupay

import "fmt"

// Pay 支付示例
func Pay() {
	appId := ""                                           //Appid
	appSecret := ""                                       //密钥
	var host = "https://api.dpweixin.com/payment/do.html" //跳转支付页接口URL

	client := NewHuPi(&appId, &appSecret) //初始化调用

	//支付参数，appid、time、nonce_str和hash这四个参数不用传，调用的时候执行方法内部已经处理
	params := map[string]string{
		"version":        "1.1",
		"trade_order_id": "1234567895",         // 商户订单号
		"total_fee":      "0.1",                // 支付金额
		"title":          "测试标题",               // 支付标题
		"notify_url":     "http://xxxxxxx.com", // 通知回调网址 可选
		//"return_url":     "http://xxxx.com",    // 	跳转网址 可选
		"wap_name": "测试店铺",                     // 网站名称 可选
		"wap_url":  "https://api.dpweixin.com", // 商品网址 可选
		"type":     "WAP",                      // 支付类型
	}

	execute, err := client.ExecutePay(host, params) //执行支付操作
	if err != nil {
		panic(err)
	}
	fmt.Println(execute) //打印支付结果
}

/**
request:
1	version	API 版本号	string(24)	必填。目前为1.1
2	appid	APP ID	string(32)	必填。填写虎皮椒的APPID，不是小程序APPID
3	trade_order_id	商户订单号	string(32)	必填。请确保在当前网站内是唯一订单号，只支持数字，大小写英文以及部分特殊字符：!#$'()*+,/:;=?@-._~%
4	total_fee	订单金额(元)	decimal(18,2)	必填。单位为人民币 元，没小数位不用强制保留2位小数
5	title	订单标题	string(128)	必填。商户订单标题（不能超过128个字符，请注意控制下长度）
6	time	当前时间戳	int(11)	必填。PHP示例：time()
7	notify_url	通知回调网址	string(128)	必填。用户支付成功后，我们服务器会主动发送一个post消息到这个网址(注意：当前接口内，SESSION内容无效，手机端不支持中文域名)
8	return_url	跳转网址	string(128)	可选。用户支付成功后，我们会让用户浏览器自动跳转到这个网址
9	callback_url	商品网址	string(128)	可选。用户取消支付后，我们可能引导用户跳转到这个网址上重新进行支付
10	plugins	名称	string(128)	可选。 用于识别对接程序或作者
11	attach	备注	text	可选。备注字段，可以传入一些备注数据，回调时原样返回
12	nonce_str	随机值	string(32)	必填。作用：1.避免服务器页面缓存，2.防止安全密钥被猜测出来
13	hash	签名	string(32)	必填。
14	type	支付通道类型	string(32)	微信H5支付请填"WAP"，微信小程序支付请填"JSAPI" ，请参考小程序demo对接小程序支付，微信内支付请勿填写"JSAPI"，支付网关为：https://api.xunhupay.com 跳转小程序APPID：wx2574b5c5ee8da56b，其他支付网关跳转小程序APPID：wx402faa5bd5eda155，（支付宝不需要此参数）
15	wap_url	网站域名	string(128)	网站域名，H5支付通道请填你网站域名，小程序支付通道请填支付网关（例如：https://api.dpweixin.com）。（支付宝不需要此参数）
16	wap_name	网站名称	string(128)	店铺名称或网站域名，长度32或以内，H5支付通道请求必填。（支付宝不需要此参数）

response:
1	oderid	订单id	int	订单id(此处有个历史遗留错误，返回名称是openid，值是orderid，一般对接不需要这个参数)
2	url_qrcode	二维码地址(PC端使用)	string(156)	PC端可将该参数展示出来进行扫码支付，不用再转二维码，需自己处理跳转
3	url	请求url(移动端使用)	string(155)	只需跳转此参数即可，系统会自动判断是微信端还是手机端，自动返回return_url，不能先显示“url_qrcode”二维码，再跳转“url”链接
4	errcode	错误码	int
5	errmsg	错误信息	string(8)	错误信息具体值
6	hash	签名	string(32)	数据签名，参考下面签名算法
*/

// Query 查询示例
func Query() {
	appId := ""                                              //Appid
	appSecret := ""                                          //密钥
	var host = "https://api.xunhupay.com/payment/query.html" //查询接口URL

	client := NewHuPi(&appId, &appSecret) //初始化调用

	//查询参数，appid、time、nonce_str和hash这四个参数不用传，调用的时候执行方法内部已经处理
	params := map[string]string{
		"open_order_id": "202112602280",
	}

	execute, err := client.ExecuteQuery(host, params) //执行查询操作

	if err != nil {
		panic(err)
	}
	fmt.Println(execute) //打印查询结果
}
