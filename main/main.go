package main

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"loc-process-auth/authen"
	"os"
)

func main() {
	env := os.Args[1] + "-env"
	fmt.Println("running profile file : ", env)
	viper.SetConfigName(env) // ชื่อ config file
	viper.AddConfigPath(".") // ระบุ path ของ config file
	viper.AutomaticEnv()     // อ่าน value จาก ENV variable
	r := authen.SetupRouter()
	//r.Run(viper.GetString("app.port"))
	r.Run()
}
