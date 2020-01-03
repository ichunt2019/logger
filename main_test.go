package logger

import (
	"testing"
	"time"
)




func TestLog(t *testing.T){
	logConfig := make(map[string]string)
	logConfig["log_path"] = "logs/order/orderSendPurchaseEmail"
	logConfig["log_chan_size"] = "1000"
	InitLogger("file",logConfig)
	Init()

	for i:=0;i<10;i++{
		Warn("warn555555555555555555555555555555")
		Info("info记录日志555555555555555555555555555555")
		Error("Error记录日志555555555555555555555555555555")
		Fatal("Fatal记录日志555555555555555555555555555555")
	}
	time.Sleep(1*time.Second)
}