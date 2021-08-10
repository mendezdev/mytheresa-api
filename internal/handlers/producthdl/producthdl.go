package producthdl

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mendezdev/mytheresa-api/internal/ports"
)

type HTTPHandler struct {
	productService ports.ProductService
}

func NewHTTPHandler(ps ports.ProductService) *HTTPHandler {
	return &HTTPHandler{
		productService: ps,
	}
}

func (hdl *HTTPHandler) GetAll(c *gin.Context) {
	category := c.Query("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, "category query param is mandatory")
		return
	}

	var lessThan *int64
	lessThanQuery := c.Query("category")
	if lessThanQuery != "" {
		n, err := strconv.ParseInt(lessThanQuery, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "less_than query param should be a number and without decimals")
			return
		}
		lessThan = &n
	}

	products, err := hdl.productService.GetProductsByCategory(category, lessThan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}
