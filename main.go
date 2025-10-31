package main

import (
	"context"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {

	//x := strings.Contains("xxxyyy", "yyy")
	//fmt.Println(x)

	//return
	h := server.Default(server.WithHostPorts(":9000"))

	h.GET("/sleep", func(c context.Context, ctx *app.RequestContext) {
		sParam := ctx.DefaultQuery("s", "1")
		s, err := strconv.Atoi(string(sParam))
		if err != nil || s < 0 {
			ctx.String(400, "Invalid sleep time: %s", sParam)
			return
		}
		time.Sleep(time.Duration(s) * time.Second)
		ctx.String(200, "Slept for %d seconds\n", s)
	})

	h.GET("/echo", func(c context.Context, ctx *app.RequestContext) {
		msg := ctx.DefaultQuery("msg", "")
		ctx.String(200, "Echo: %s\n", msg)
	})

	h.GET("/info", func(c context.Context, ctx *app.RequestContext) {
		hostname, _ := os.Hostname()
		ip := getOutboundIP()
		ctx.JSON(200, map[string]interface{}{
			"hostname": hostname,
			"ip":       ip,
		})
	})

	h.Spin()
}

func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "unknown"
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
