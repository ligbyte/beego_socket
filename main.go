package main

import (
	_ "beego_socket/routers"
	_ "beego_socket/sockets"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

