# often
Common tools
go get -u -v github.com/butoften/often
import (
	butof "local.go/m/v2/src/butoften"
)
var log = butof.Log("./log.log")
log.Info([]string{"wo ha ni ha 4"})
