# g4vercel-demo

> Deploy go web server in vercel

The Go Runtime is used by Vercel to compile Go Serverless Functions that expose a single HTTP handler, from a `.go` file within an `/api` directory at your project's root.

For example, define an `index.go` file inside an `/api` directory as follows:

package handler

```go

package handle

import (
	"fmt"
	"github.com/tbxark/g4vercel"
	"net/http"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	server := gee.Default()
	server.GET("/", func(context *gee.Context) {
		context.JSON(200, gee.H{
			"status": "OK",
		})
	})
	server.GET("/hello", func(context *gee.Context) {
		name := context.Query("name")
		if name == "" {
			context.JSON(400, gee.H{
				"error": "name not found",
			})
		} else {
			context.JSON(200, gee.H{
				"data": fmt.Sprintf("Hello %s!", name),
			})
		}
	})
	server.GET("/user/:id", func(context *gee.Context) {
		context.JSON(400, gee.H{
			"data": gee.H{
				"id": context.Param("id"),
			},
		})
	})
	server.GET("/long/long/long/path/*test", func(context *gee.Context) {
		context.JSON(200, gee.H{
			"data": gee.H{
				"url": context.Path,
			},
		})
	})
	server.Handle(w, r)
}

```

An example `index.go` file inside an `/api` directory.