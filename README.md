# TheDeveloper10/rem

Package `TheDeveloper10/rem` implements a request router and dispatcher for matching incoming
requests to their respective handler.

The name REM comes from the state of sleep called REM. The same way sleep is a rest from 
engineering this package is a rest from swiss knife packages that solve all of your problems
and create a thousand more.

The main features are:
* It implements the `http.Handler` interface so it is compatible with `net/http`
* Requests can be matched based on URL paths and HTTP Methods
* URL paths can have optional variables

___

* [Install](#install)
* [Examples](#examples)

___

## Install
With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/TheDeveloper10/rem
```

## Examples
This is a very simple example that describes the entire package really well.
First we create a new `rem.Router` and then we use it to create new routes.
The process of creating routes is really easy: we just use the GetRoute for
handling `GET` HTTP method, PostRoute for handling `POST` HTTP method, etc. 
for the rest of the methods and then we assign the handlers that we want the 
data to go through. 

In all of the cases in this example we go through `AuthMiddleware`.
That's just a middleware used for authentication of the call. If the 
call passes through the `AuthMiddleware` it goes to the respective handler
(e.g. GetProductsHandle).

```go
func main() {
	// Create a new router
	router := rem.NewRouter()
	
	// Add a new route
	router.
		NewRoute("/products").
		GetRoute(AuthMiddleware, GetProductsHandle).
	    PostRoute(AuthMiddleware, CreateProductsHandle)
	    
	
	// Add a new route
	router.
		NewRoute("/products/:productId").
		GetRoute(AuthMiddleware, GetProductHandle).
		PutRoute(AuthMiddleware, ReplaceProductHandle).
		DeleteRoute(AuthMiddleware, DeleteProductHandle)
	
	// Start the HTTP server
	http.ListenAndServe(":80", router)
}
```

Creating your own handlers
```go
type userResponse struct {
	id string `json:"id"`
}

func main() {
	router := rem.NewRouter()
	
	router.
		NewRoute("/users").
		PostRoute(func(response IResponse, request IRequest) bool {
			response := userResponse{ id: "jxIZp17" }
		    request.Status(http.StatusCreated).JSON(response)

			// In case of a single handler it doesn't matter what you 
			// return (true or false). It only matters when there's more 
			// than one handler because that's how the route decides to 
			// continue passing to the next handler.
			return true
        })
}
```