install:  

go get -u -v github.com/butoften/often  

usage：   

import (  

	butof "github.com/butoften/often"

)  

var log = butof.Log("./log.log")  

log.Info([]string{"this i a message"})  

and you can use   

tail -f path/to/log/file   

and watch some change or info  
