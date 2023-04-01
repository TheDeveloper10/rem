package main

import (
	"net/http"

	"github.com/TheDeveloper10/rem"
)

func main() {
	rem.SetConfig(rem.DefaultConfig())
	router := rem.NewRouter()

	router.Get("/test/1", func(ctx *rem.Context) rem.IResponse {
		return rem.Success("Test")
	})

	router.
		NewRoute("/test/2").
		Get(func(ctx *rem.Context) rem.IResponse {
			return rem.Success(struct {
				Msg string `json:"msg"`
			}{Msg: "Hello world!"})
		})

	router.Get("/test/3", func(ctx *rem.Context) rem.IResponse {
		ctx.SetData("userID", 1)
		return ctx.Next()(ctx)
	}, func(ctx *rem.Context) rem.IResponse {
		return rem.Success(ctx.GetData("userID"))
	})

	http.ListenAndServe(":8080", router)
}
