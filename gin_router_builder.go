package firead

import (
	"github.com/gin-gonic/gin"
	"github.com/jonnyorman/fireworks"
)

type GinRouterBuilder[TDocument any] struct {
	requestHandler fireworks.RequestHandler[TDocument]
}

func NewGinRouterBuilder[TDocument any](requestHandler fireworks.RequestHandler[TDocument]) *GinRouterBuilder[TDocument] {
	this := new(GinRouterBuilder[TDocument])

	this.requestHandler = requestHandler

	return this
}

func (this GinRouterBuilder[TDocument]) Build() *gin.Engine {
	router := gin.Default()

	router.GET("/:id", this.requestHandler.Handle)

	return router
}
