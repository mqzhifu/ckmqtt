package ckmqtt

import (
	"fmt"
	"os"
	"src/zlib"
	"time"
)

const (
	TYPE_USER = 1
	TYPE_SERV = 2
)

type App struct {
	Id int
}



type DataRecord struct {
	Id 	int
	Name string
	Type int	//1 user 2 service
	Key string
	ServiceName string
}

var Data map[int]DataRecord

func (app *App)GetById(appId int)DataRecord{
	return Data[appId]
}
//通过用户名密码登陆
func LoginGetTokenByUnamePs(username string ,ps string){
	// doing something  auth ok~

}

//通过用户名密码登陆
func (app *App)  LoginGetTokenByKey(){

}

func testJWT(){
	current := time.Now().Second()
	ExpireTime := int(current) + (   2 * 60 * 60 )
	JwtDataPayload := zlib.JwtDataPayload{
		Id:1,
		AppId: 2,
		ATime: current,
		Username : "wang",
		Expire :ExpireTime,
	}

	secretKey := "bbbb"

	jwtString := zlib.CreateJwtToken(secretKey,JwtDataPayload)
	fmt.Println("myself : ",jwtString)
	tokenString := zlib.JwtGoCreateToken(secretKey,JwtDataPayload)
	fmt.Println("jwt-go : ",tokenString)


	myTokenStr :=
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiRXhwaXJlIjo3MjA2LCJBVGltZSI6NiwiQXBwSWQiOjIsIlVzZXJuYW1lIjoid2FuZyJ9.TaDNfdN0RykFfIrzdQu57NqYnK9kU-HO6060s6qspbM"
	data,err := zlib.ParseJwtToken(secretKey,myTokenStr)
	if err != nil{
		zlib.ExitPrint(data,err.Error())
	}

	zlib.MyPrint("check ok")

	os.Exit(-100)



	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//tokenString, err = token.SignedString(SecretKey)
}

func (app *App) loadAppData(){
	Data = make(map[int]DataRecord)
	dataRecord := DataRecord{
		Id:1,
		Name : "ckmqtt",
		Type : 1,
		Key : "123456",
	}
	Data[1] = dataRecord

	dataRecord = DataRecord{
		Id:2,
		Name : "ckmqtt",
		Type : 2,
		Key : "123456",
		ServiceName:"ucenter",
	}
	Data[2] = dataRecord
}

func  NewApp(appId int) *App{
	app := new (App)
	app.loadAppData()
	app.Id = appId
	return app
}