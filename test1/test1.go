/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 20:57 2019-11-05
 */
package main

import (
	"github.com/dollarkillerx/erguotou"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	app := erguotou.New()

	app.Get("/", func(ctx *erguotou.Context) {
		ctx.String(200, "TestHttp")
	})

	go func() {
		log.Println(http.ListenAndServe(":8081", nil))
	}()

	app.Run(erguotou.SetHost(":8083"))

}
