package ginServer

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// GinRoutingFn is callback function for setting up routers.
type GinRoutingFn func(router *gin.Engine)

// GinServerParams contains all dependencies of ginServer.
type Params struct {
	dig.In
	Routing GinRoutingFn
	Conf    *Config
	Router  *gin.Engine
}

// NewGinServer returns new instance of Server.
func NewGinServer(params Params) Server {
	return &ginServer{
		conf:    params.Conf,
		router:  params.Router,
		routing: params.Routing,
	}
}

type ginServer struct {
	routing GinRoutingFn
	conf    *Config
	router  *gin.Engine
}

func (_this *ginServer) Open() error {
	if _this.routing == nil {
		return ErrNilRoutingFn
	}
	_this.routing(_this.router)

	if err := _this.router.Run(_this.conf.ListenerAddr()); err != nil {
		return err
	}

	return nil
}

func (_this *ginServer) Close() {
	// Blank function
}
