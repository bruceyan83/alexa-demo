package alexa

import (
	"context"
	"fmt"
)

// Handler 抽象接口，所有来自Alexa的请求都应该实现这个接口
type Handler interface {
	HandleRequest(ctx context.Context, req *Request) (*Response, error)
}

// HandlerFunc impl function
type HandlerFunc func(ctx context.Context, req *Request) (*Response, error)

// HandleRequest 调用func
func (h HandlerFunc) HandleRequest(ctx context.Context, req *Request) (*Response, error) {
	return h(ctx, req)
}

// RouteNameSpace 根据不同的space name，调用不同的handler
type RouteNameSpace struct {
	handlerMap map[string]Handler
}

// NewRouteNameSpace 创建一个RouteNameSpace
func NewRouteNameSpace() *RouteNameSpace {
	return &RouteNameSpace{make(map[string]Handler)}
}

// HandleRequest 将请求转发给对应的handler，如果没有对应的handler，则返回错误
func (n *RouteNameSpace) HandleRequest(ctx context.Context, req *Request) (*Response, error) {
	handler := n.handlerMap[req.Directive.Header.Namespace]
	if handler == nil {
		return nil, fmt.Errorf("RouteNameSpace: unhandled namespace: %s", req.Directive.Header.Namespace)
	}
	return handler.HandleRequest(ctx, req)
}

// Handle 注册一个handler
func (n *RouteNameSpace) Handle(namespace string, handler Handler) {
	n.handlerMap[namespace] = handler
}

// HandleFunc 注册一个handler函数
func (n *RouteNameSpace) HandleFunc(namespace string, handler HandlerFunc) {
	n.Handle(namespace, handler)
}
