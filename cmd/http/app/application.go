package app

import (
	"github.com/mendezdev/mytheresa-api/internal/core/services/discountsrv"
	"github.com/mendezdev/mytheresa-api/internal/core/services/productsrv"
	"github.com/mendezdev/mytheresa-api/internal/handlers/producthdl"
	"github.com/mendezdev/mytheresa-api/internal/repositories/discountrepo"
	"github.com/mendezdev/mytheresa-api/internal/repositories/productrepo"
)

type handlers struct {
	productHdl *producthdl.HTTPHandler
}

func buildHandlers() handlers {
	// discount singletons
	discountRepo := discountrepo.NewInMemory()
	discountSrv := discountsrv.New(discountRepo)

	// products singletons
	productRepo := productrepo.NewInMemory()
	productSrv := productsrv.New(productRepo, discountSrv)
	productHdl := producthdl.NewHTTPHandler(productSrv)

	return handlers{
		productHdl: productHdl,
	}
}

func StartApplication() {
	router := routes(buildHandlers())
	router.Run(":8080")
}
