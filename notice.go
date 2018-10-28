package tkTools

import (
	"net/url"
	"github.com/nntaoli-project/GoEx"
	"net/http"
)

func SendToWechat(SERVER_SCKEY,text,desp string) {
	// to convert a float number to a string
	wechatUrl:="https://sc.ftqq.com/"+SERVER_SCKEY+".send"
	params := url.Values{}
	params.Set("text",text)
	params.Set("desp",desp)
	goex.HttpPostForm(http.DefaultClient,wechatUrl,params)
}
