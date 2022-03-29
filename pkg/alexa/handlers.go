package alexa

import (
	"context"
	"encoding/json"
	"fmt"
)

// StaticDiscoveryHandler handles discovery requests with a hardcoded set
// of endpoints
func StaticDiscoveryHandler(builder *ResponseBuilder, endpoints ...DiscoverEndpoint) HandlerFunc {
	return func(ctx context.Context, req *Request) (*Response, error) {
		resp, err := builder.DiscoverResponse(endpoints...)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

// AuthorizationHandler Auth处理函数
func AuthorizationHandler(respBuilder *ResponseBuilder) HandlerFunc {
	return func(ctx context.Context, req *Request) (*Response, error) {
		var payload AcceptGrantPayload
		if err := json.Unmarshal(req.Directive.Payload, &payload); err != nil {
			return nil, fmt.Errorf("failed to unmarshal payload: %v", err)
		}
		return respBuilder.AcceptGrantResponse(), nil
	}
}

// PercentageControllerHandler routes handling of set & adjust directives
func PercentageControllerHandler(setPct, adjustPct Handler) HandlerFunc {
	return func(ctx context.Context, req *Request) (*Response, error) {
		switch req.Directive.Header.Name {
		case "SetPercentage":
			return setPct.HandleRequest(ctx, req)
		case "AdjustPercentage":
			return adjustPct.HandleRequest(ctx, req)
		default:
			return nil, fmt.Errorf("PercentageControllerHandler: unexpected name: %s", req.Directive.Header.Name)
		}
	}
}

// PowerControllerHandler routes turn on & off requests
func PowerControllerHandler(turnOn, turnOff Handler) HandlerFunc {
	return func(ctx context.Context, req *Request) (*Response, error) {
		switch req.Directive.Header.Name {
		case "TurnOn":
			return turnOn.HandleRequest(ctx, req)
		case "TurnOff":
			return turnOff.HandleRequest(ctx, req)
		default:
			return nil, fmt.Errorf("PowerControllerHandler: unexpected name: %s", req.Directive.Header.Name)
		}
	}
}

// SceneControllerHandler routes activate & deactivate requests
func SceneControllerHandler(activate, deactivate Handler) HandlerFunc {
	return func(ctx context.Context, req *Request) (*Response, error) {
		switch req.Directive.Header.Name {
		case "Activate":
			return activate.HandleRequest(ctx, req)
		case "Deactivate":
			return deactivate.HandleRequest(ctx, req)
		default:
			return nil, fmt.Errorf("SceneControllerHandler: unexpected name: %s", req.Directive.Header.Name)
		}
	}
}
