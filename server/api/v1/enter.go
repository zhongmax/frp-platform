package v1

type RouterGroup struct {
	BaseRouter
	InitRouter
	UserRouter
}

var RouterGroupApp = new(RouterGroup)
