install:  

````
$	go get -u -v github.com/butoften/often  
````

usage： 

1.Log
````
import (  
	butof "github.com/butoften/often"
)  
var log = butof.Log("./log.log")  
log.Info([]string{"this i a message"})  
````
and you can use   
````
$	tail -f path/to/log/file   
````
and watch some change or info  

2.cors for gin
````
r := gin.Default()
r.Use(often.Cors("xx.domain.com"))
````