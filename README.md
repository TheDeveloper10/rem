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

Creating your own handlers:
```go
type userResponse struct {
	id string `json:"id"`
}

func main() {
	router := rem.NewRouter()
	
	router.
		NewRoute("/users").
		PostRoute(func(response rem.IResponse, request rem.IRequest) bool {
			request.
				Status(http.StatusCreated).
				JSON(useResponse{
					id: "jxIZp17"
				})
        
			// In case of a single handler it doesn't matter what you 
			// return (true or false). It only matters when there's more 
			// than one handler because that's how the route decides to 
			// continue passing to the next handler.
			return true
		})
}
```

Middlewares and handlers:
```go
func AuthMiddleware(response rem.IResponse, request rem.IRequest) bool {
	authHeader, status := request.GetHeaders().Get("Authentication")
	if !status {
		response.Status(http.StatusUnauthorized)
		// Returning false means it will not continue to the next handler
		return false
	}
	
	// Obviously here you won't perform this type of authentication...
	if authHeader != "1234" {
		response.Status(http.StatusUnauthorized)
		return false
	}
	
	return true
}

func GetProductsHandler(response rem.IResponse, request rem.IRequest) bool {
	response.
		JSON([]Product{
			{ id: "nXJf", price: 39.99, name: "Super new shoes" },
			{ id: "zeq2", price: 1999.99, name: "Newest Phone" },
			{ id: "cv41", price: 4.99, name: "New headphones" },
		}).
		Status(http.StatusOK)
	
	return true
}

type Product struct {
	id    string
	price float32
	name  string
}
```

Path and Query parameters in URL:
```go
type Product struct {
	Id    string  `json:"id"`
	Price float32 `json:"price"`
	Name  string  `json:"name"` 
}

func GetProductHandler(response rem.IResponse, request rem.IRequest) bool {
	productId, status1 := request.GetURLParameters().Get("productId")
	if !status1 {
		response.Status(http.StatusBadRequest)
	    return true	
    }
	
	priceStr, status2 := request.GetQueryParameters().Get("price")
	
	priceVal := 3.99
	if status2 {
		priceFloat, err := strconv.ParseFloat(priceStr, 64)
		if err == nil {
		    priceVal = float32(priceFloat)
		}
	}
	
	// e.g. for request GET /products/ZxF?price=40 the result will be:
	// { "id": "ZxF", "price": 40, "Name": "Some Name" }

	// and for request GET /products/qwe the result will be:
	// { "id": "qwe", "price": 3.99, "Name": "Some Name" }

	response.
		Status(http.StatusOK).
		JSON(Product{
		    Id: productId,
			Price: priceVal,
			Name: "Some name"
		})
	
	return true
}

func main() {
	router := rem.NewRouter()
	
	router.
		NewRoute("/products/:productId").
		GetRoute(GetProductHandler)
	
	http.ListenAndServe(":80", router)
}
```