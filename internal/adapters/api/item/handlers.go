package item

import (
	"github.com/gin-gonic/gin"
	"github.com/m-zagornyak/rest-api.git/internal/adapters"
)

const (
	getItemById = "/:item_id"
	updateItem  = "/:item_id"
	deleteItem  = "/:item_id"
)

type handler struct {
	itemService Service
}

func NewHandler(service Service) adapters.Handler {
	return &handler{itemService: service}
}

func (h *handler) Register(router *gin.Engine) {
	router.POST(getItemById, h.createItem)
	router.GET("/", h.getAllItems)
	router.GET("/:item_id", h.getItemById)
	router.PUT("/:item_id", h.updateItem)
	router.DELETE("/:item_id", h.deleteItem)

}

func (h *handler) createItem(c *gin.Context) {

}

func (h *handler) getAllItems(c *gin.Context) {

}

func (h *handler) getItemById(c *gin.Context) {

}

func (h *handler) updateItem(c *gin.Context) {

}

func (h *handler) deleteItem(c *gin.Context) {

}
