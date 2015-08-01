
#Using Golang standard library  to implement http GET/POST 
Date: 7/31/2015


## Server
start a server with listen port 9001 ,   
url: http://localhost:9001/api/v1  

## Client-GET
do http.Get, and then use ioutil.ReadAll(res.Body) to read the result

## Client-Post
use http.Post, save the data to a slice, and print the result  

## Reference:
http://blog.csdn.net/typ2004/article/details/38669949  


