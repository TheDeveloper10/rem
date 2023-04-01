package main

import (
	"net/http"

	"github.com/TheDeveloper10/rem"
)

func main() {
	router := rem.NewRouter()

	router.Get("/test/1", func(ctx *rem.Context) rem.IResponse {
		return rem.Success(nil)
	})

	router.
		NewRoute("/test/2").
		Get(func(ctx *rem.Context) rem.IResponse {
			return rem.Success(nil)
		})

	router.Get("/test/3", func(ctx *rem.Context) rem.IResponse {
		ctx.SetData("userID", 1)
		return ctx.Next()
	}, func(ctx *rem.Context) rem.IResponse {
		return rem.Success(ctx.GetData("userID"))
	})

	http.ListenAndServe(":8080", router)
}
