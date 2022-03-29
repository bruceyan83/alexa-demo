package alexa

import (
	"alexa-demo/pkg/virtual_device"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// StaticDiscoveryHandler 模拟设备群
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

// PowerControllerHandler 处理PowerController指令
func PowerControllerHandler(respBuilder *ResponseBuilder) HandlerFunc {
	return func(ctx context.Context, req *Request) (*Response, error) {
		value := ""
		switch req.Directive.Header.Name {
		case "TurnOn":
			value = "ON"
			virtual_device.SetLamp(true)
		case "TurnOff":
			value = "OFF"
			virtual_device.SetLamp(false)
		default:
			return nil, fmt.Errorf("PowerControllerHandler: unexpected name: %s", req.Directive.Header.Name)
		}
		prop := ContextProperty{
			Namespace:                 NamespacePowerController,
			Name:                      "powerState",
			Value:                     value,
			TimeOfSample:              time.Now(),
			UncertaintyInMilliseconds: 0,
		}
		return respBuilder.BasicResponse(req, prop), nil
	}
}

// StateReportHandler 处理StateReport指令
func StateReportHandler(respBuilder *ResponseBuilder, temperatureHandle Handler) HandlerFunc {
	return func(ctx context.Context, req *Request) (*Response, error) {
		value := ""
		if virtual_device.GetLamp() {
			value = "ON"
		} else {
			value = "OFF"
		}
		switch req.Directive.Endpoint.EndpointID {
		case "it-is-a-temperature-sensor":
			return temperatureHandle.HandleRequest(ctx, req)
		default:
			return respBuilder.StateReportResponse(req,
				ContextProperty{
					Namespace:                 NamespacePowerController,
					Name:                      "powerState",
					Value:                     value,
					TimeOfSample:              time.Now(),
					UncertaintyInMilliseconds: 60000,
				}), nil
		}
	}
}
