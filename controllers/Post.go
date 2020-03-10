package controllers

import (
	"fmt"
	//"debug/dwarf"
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
        "io/ioutil"
	//"net/http"
	//"time"

)

//
//ch:=make(chan int,100)
//
//
//func transfer(i int){
//	ch<-i
//
//}
var tamp string


var data []byte

func RegisterPost(c *gin.Context) {
         

        
         body,_:=ioutil.ReadAll(c.Request.Body)
         tamp=string(body)
         if body!=nil{
              fmt.Printf("request_info:%s",tamp)
         }

         
        

         fmt.Println(tamp)
         
        
         fmt.Println("---------------")

         c.String(http.StatusOK,tamp)
         
	 //c.JSON(http.StatusOK,gin.H{"messages":data})


}

func RegisterGet(c *gin.Context) {

//	c.HTML(http.StatusOK,"datas.html",gin.H{"wechat":"666"})
	c.HTML(http.StatusOK,"test.html",gin.H{})


}
func Get(c *gin.Context){
	c.String(http.StatusOK,tamp)


}
