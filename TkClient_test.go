package tkTools

import (
	"testing"
	"net/http"
	"time"
)

func TestNewTkClient(t *testing.T) {
	client:=NewTkClient(http.DefaultClient,"","")
	client.SendTestMsg()
	time.Sleep(time.Second)
}


func TestTkClient_SendAsset(t *testing.T) {
	client:=NewTkClient(http.DefaultClient,"","")
	client.SendAsset("api_test","111","1111")
	time.Sleep(time.Second)
}


func TestTkClient_SendOrder(t *testing.T) {
	client:=NewTkClient(http.DefaultClient,"","")
	client.SendOrder("api_test","6.8","100","6.8","0","123",
		"1234","UST_QC","BUY")
	time.Sleep(time.Second)
}

func TestTkClient_SendLogInfo(t *testing.T) {
	client:=NewTkClient(http.DefaultClient,"","")
	client.SendLogInfo("api_test","1","api test","1234")
	time.Sleep(time.Second)
}