# Mir
[![Build Status](https://api.travis-ci.com/dyc92/mir.svg?branch=master)](https://travis-ci.com/dyc92/mir)
[![codecov](https://codecov.io/gh/dyc92/mir/branch/master/graph/badge.svg)](https://codecov.io/gh/dyc92/mir)
[![GoDoc](https://godoc.org/github.com/dyc92/mir?status.svg)](https://godoc.org/github.com/dyc92/mir)
[![Release](https://img.shields.io/github/release/dyc92/mir.svg?style=flat-square)](https://github.com/dyc92/mir/releases)

Mir is used for register handler to http router(eg: [Gin](https://github.com/gin-gonic/gin), [Chi](https://github.com/go-chi/chi), [Echo](https://github.com/labstack/echo), [Iris](https://github.com/kataras/iris), [Macaron](https://github.com/go-macaron/macaron), [Mux](https://github.com/gorilla/mux), [httprouter](https://github.com/julienschmidt/httprouter))
 depends on struct tag string info that defined in logic object's struct type field.

### Usage (eg: gin backend)
* Get Mir.Gin module first

```bash
go get github.com/dyc92/mir/module/gin@master
```

* Then happy in codding enjoy your heart...

```go
package main

import(
	"github.com/dyc92/mir"
	"github.com/gin-gonic/gin"
	"net/http"
	
	mirE "github.com/dyc92/mir/module/gin"
)

type site struct {
	Chain mir.Chain     `mir:"-"`
	Group mir.Group     `mir:"v1"`
	index mir.Get       `mir:"/index/"`
	articles mir.Get    `mir:"/articles/:category/#GetArticles"`
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (h *site) Index(c *gin.Context) {
	c.String(http.StatusOK, "get index data")
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
// Path info is the second or first(if no host info) segment start with '/'(eg: /articles/:category/#GetArticles)
// Handler info is forth info start with '#' that indicate real handler method name(eg: GetArticles).if no handler info will
// use field name capital first char as default handler name(eg: if articles had no #GetArticles then the handler name will
// is Articles) 
func (h *site) GetArticles(c *gin.Context) {
	c.String(http.StatusOK, "get articles data")
}

func main() {
	//Create a new gin engine
	engine := gin.New()
	
	// Register handler to engine by mir
	mirE.Register(engine, &site{Chain: gin.HandlersChain{gin.Logger()}})
	
	// Start gin engine serve
	engine.Run()
}

```

### Sample code use mir build web application
* [github.com/dyc92/mir-music](https://github.com/dyc92/mir-music) : use [mir](https://github.com/dyc92/mir)+[gin](https://github.com/gin-gonic/gin) build a sample music infomation service app like [spring-music](https://github.com/cloudfoundry-samples/spring-music).
* [github.com/dyc92/chi-music](https://github.com/dyc92/chi-music) : use [mir](https://github.com/dyc92/mir)+[go-chi](https://github.com/go-chi/chi) implement [spring-music](https://github.com/cloudfoundry-samples/spring-music) in golang.
