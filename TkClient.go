package tkTools

import (
	"net/http"
	. "github.com/Vniex/tkTools/httpUtils"
	"net/url"
	"log"
	"time"
	"strconv"
	"encoding/json"
)

const (
	HEART_BEAT=2
	API_URL="http://127.0.0.1:8888/api/v1/"

	TEST_URI=API_URL+"test"
	ASSET_URI=API_URL+"asset"
	ORDER_URI=API_URL+"order"
	LOGINFO_URI=API_URL+"loginfo"
)

type TkClient struct {
	//clientName string
	apiKey string
	secretKey string

	httpClient *http.Client


}



func NewTkClient(client *http.Client,apikey,secretKey string) *TkClient{
	return &TkClient{
		apiKey:apikey,
		secretKey:secretKey,
		httpClient:client,

	}
}


func (c *TkClient)SendTestMsg(){
	postData:=url.Values{}
	postData.Set("msg","test api")
	c.buildSign(&postData)
	resp,err:=HttpPostForm(c.httpClient,TEST_URI,postData)
	if err!=nil{
		log.Println(err)
	}else{
		respmap:=make(map[string]interface{})
		if err = json.Unmarshal(resp, &respmap);err!=nil{
			log.Println(err)
		}else{
			log.Println(respmap)
		}

	}

}


func (c *TkClient)SendAsset(robot_name,net_asset,time_stamp string)error{
	postData:=url.Values{}
	postData.Set("robot_name",robot_name)
	postData.Set("net_asset",net_asset)
	postData.Set("time_stamp",time_stamp)
	c.buildSign(&postData)
	resp,err:=HttpPostForm(c.httpClient,ASSET_URI,postData)
	if err!=nil{
		log.Println(err)
		return  err
	}else{
		respmap:=make(map[string]interface{})
		if err = json.Unmarshal(resp, &respmap);err!=nil{
			log.Println(err)
			return err
		}else{
			log.Println(respmap)
		}

	}
	return nil

}

func (c *TkClient)SendOrder(robot_name,price,amount,avg_price,fee,order_id,order_time,pair,side string,) error{
	postData:=url.Values{}
	postData.Set("robot_name",robot_name)
	postData.Set("price",price)
	postData.Set("amount",amount)
	postData.Set("avg_price",avg_price)
	postData.Set("fee",fee)
	postData.Set("order_id",order_id)
	postData.Set("order_time",order_time)
	postData.Set("pair",pair)
	postData.Set("side",side)

	c.buildSign(&postData)
	resp,err:=HttpPostForm(c.httpClient,ORDER_URI,postData)
	if err!=nil{
		log.Println(err)
		return err
	}else{
		respmap:=make(map[string]interface{})
		if err = json.Unmarshal(resp, &respmap);err!=nil{
			log.Println(err)
			return err
		}else{
			log.Println(respmap)
		}

	}
	return nil
}


func (c *TkClient)SendLogInfo(robot_name,level,msg,timestamp string ) error{
	postData:=url.Values{}
	postData.Set("robot_name",robot_name)
	postData.Set("level",level)
	postData.Set("msg",msg)
	postData.Set("time_stamp",timestamp)


	c.buildSign(&postData)
	resp,err:=HttpPostForm(c.httpClient,LOGINFO_URI,postData)
	if err!=nil{
		log.Println(err)
		return err
	}else{
		respmap:=make(map[string]interface{})
		if err = json.Unmarshal(resp, &respmap);err!=nil{
			log.Println(err)
			return err
		}else{
			log.Println(respmap)
		}

	}
	return nil
}

func (c *TkClient)buildSign(postData *url.Values){
	now:=time.Now().Unix()
	nowStr:=strconv.Itoa(int(now))
	sign:=GetSHA256(c.apiKey+c.secretKey+nowStr)
	postData.Set("apikey",c.apiKey)
	postData.Set("timestamp",nowStr)
	postData.Set("sign",sign)

}















