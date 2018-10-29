package tkTools

import (
	"reflect"
	log "github.com/sirupsen/logrus"
	"time"
	"github.com/Vniex/tkTools/httpUtils"
	"net/http"
	"net/url"
	"runtime"
)



func SendToWechat(SERVER_SCKEY,text,desp string) {
	// to convert a float number to a string
	wechatUrl:="https://sc.ftqq.com/"+SERVER_SCKEY+".send"
	params := url.Values{}
	params.Set("text",text)
	params.Set("desp",desp)
	httpUtils.HttpPostForm(http.DefaultClient,wechatUrl,params)
}


/**
  间隔2s进行容错重试
  @method 调用的函数，
  @params 参数,顺序一定要按照实际调用函数入参顺序一样
  @return 返回
*/
func Retry( method interface{}, params ...interface{}) interface{} {

	invokeM := reflect.ValueOf(method)
	funcName:=runtime.FuncForPC(reflect.ValueOf(method).Pointer()).Name()
	if invokeM.Kind() != reflect.Func {
		panic("method not a function")
		return nil
	}

	var value []reflect.Value = make([]reflect.Value, len(params))
	var i int = 0
	for ; i < len(params); i++ {
		value[i] = reflect.ValueOf(params[i])
	}

	var retV interface{}
	var retryC int = 0
_CALL:
	if retryC > 0 {
		//log.Info("sleep....", time.Duration(2*int(time.Second)))
		time.Sleep(time.Duration(2 * int(time.Second)))
	}

	retValues := invokeM.Call(value)

	for _, vl := range retValues {
		if vl.Type().String() == "error" {
			if !vl.IsNil() {
				log.Println(vl)
				retryC++
				log.Infof("[%s] Error , Begin Retry Call [%d] ...", funcName, retryC)
				goto _CALL

			}
		} else {
			retV = vl.Interface()
		}
	}

	return retV
}
