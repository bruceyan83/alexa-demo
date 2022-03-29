package alexa

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// UUIDMessageID 生成一个UUID
func UUIDMessageID() string {
	return uuid.New().String()
}

// ResponseBuilder 为smart home请求生成响应
// skill api
type ResponseBuilder struct {
	MessageID func() string
}

// NewResponseBuilder 生成一个带有UUID作为MsgID的ResponseBuilder
func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{UUIDMessageID}
}

// DiscoverResponse 创建一个DiscoverResponse
func (r *ResponseBuilder) DiscoverResponse(endpoints ...DiscoverEndpoint) (*Response, error) {
	payload := DiscoverPayload{
		Endpoints: endpoints,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	resp := Response{
		Event: Event{
			Header: Header{
				Namespace:      "Alexa.Discovery",
				Name:           "Discover.Response",
				PayloadVersion: "3",
				MessageID:      r.MessageID(),
			},
			Payload: payloadJSON,
		},
	}

	return &resp, nil
}

// BasicErrorResponse error response
func (r *ResponseBuilder) BasicErrorResponse(req *Request, errorType, msg string) (*Response, error) {
	payload := struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	}{errorType, msg}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}
	return &Response{
		Event: Event{
			Header: Header{
				Namespace:        req.Directive.Header.Namespace,
				Name:             "ErrorResponse",
				PayloadVersion:   "3",
				MessageID:        r.MessageID(),
				CorrelationToken: req.Directive.Header.CorrelationToken,
			},
			Endpoint: &ResponseEndpoint{
				EndpointID: req.Directive.Endpoint.EndpointID,
				Scope:      req.Directive.Endpoint.Scope,
			},
			Payload: payloadJSON,
		},
	}, nil
}

// CustomErrorResponse 自定义error response
func (r *ResponseBuilder) CustomErrorResponse(req *Request, payload json.RawMessage) *Response {
	return &Response{
		Event: Event{
			Header: Header{
				Namespace:        req.Directive.Header.Namespace,
				Name:             "ErrorResponse",
				PayloadVersion:   "3",
				MessageID:        r.MessageID(),
				CorrelationToken: req.Directive.Header.CorrelationToken,
			},
			Endpoint: &ResponseEndpoint{
				EndpointID: req.Directive.Endpoint.EndpointID,
				Scope:      req.Directive.Endpoint.Scope,
			},
			Payload: payload,
		},
	}
}

// StateReportResponse 根据提供的状态生成状态报告响应
func (r *ResponseBuilder) StateReportResponse(req *Request, properties ...ContextProperty) *Response {
	return &Response{
		Event: Event{
			Header: Header{
				Namespace:        NamespaceAlexa,
				Name:             "StateReport",
				PayloadVersion:   "3",
				MessageID:        r.MessageID(),
				CorrelationToken: req.Directive.Header.CorrelationToken,
			},
			Endpoint: &ResponseEndpoint{
				EndpointID: req.Directive.Endpoint.EndpointID,
				Scope:      req.Directive.Endpoint.Scope,
			},
			Payload: EmptyPayload,
		},
		Context: &ResponseContext{
			Properties: properties,
		},
	}
}

// BasicResponse 创建一个基础响应
func (r *ResponseBuilder) BasicResponse(req *Request, properties ...ContextProperty) *Response {
	return &Response{
		Event: Event{
			Header: Header{
				Namespace:        NamespaceAlexa,
				Name:             "Response",
				PayloadVersion:   "3",
				MessageID:        r.MessageID(),
				CorrelationToken: req.Directive.Header.CorrelationToken,
			},
			Endpoint: &ResponseEndpoint{
				EndpointID: req.Directive.Endpoint.EndpointID,
				Scope:      req.Directive.Endpoint.Scope,
			},
			Payload: EmptyPayload,
		},
		Context: &ResponseContext{
			Properties: properties,
		},
	}
}

// AcceptGrantResponse accept grant响应
func (r *ResponseBuilder) AcceptGrantResponse() *Response {
	return &Response{
		Event: Event{
			Header: Header{
				Namespace:      NamespaceAuthorization,
				Name:           "AcceptGrant.Response",
				PayloadVersion: "3",
				MessageID:      r.MessageID(),
			},
			Payload: EmptyPayload,
		},
	}
}
