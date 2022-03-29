package main

import (
	"alexa-demo/pkg/alexa"
	"context"
	"time"

	ld "alexa-demo/pkg/lambda"
	_ "alexa-demo/pkg/virtual_device"
	"github.com/aws/aws-lambda-go/lambda"
)

// 这是aws-lambda的golang版本demo实现
func main() {
	respBuilder := alexa.NewResponseBuilder()
	//温度直接写死了
	tempReader := &tempReader{20, respBuilder}

	// 这个地方使用alexa request中的 Namespace 来区分请求并转发至不同的处理函数
	r := alexa.NewRouteNameSpace()
	r.HandleFunc(alexa.NamespacePowerController, alexa.PowerControllerHandler(respBuilder))
	r.HandleFunc(alexa.NamespaceDiscovery, alexa.StaticDiscoveryHandler(respBuilder, endpoints()...))
	r.HandleFunc(alexa.NamespaceAlexa, alexa.StateReportHandler(respBuilder, tempReader))
	r.HandleFunc(alexa.NamespaceAuthorization, alexa.AuthorizationHandler(respBuilder))

	lambda.Start(ld.DebugLambdaRequestHandler(r))
}

//提前写好的endpoints，用来模拟数据，开关开闭流程在redis中，这里只是让发现过程可以发现这些设备
func endpoints() []alexa.DiscoverEndpoint {
	return []alexa.DiscoverEndpoint{
		{
			EndpointID:        "it-is-a-temperature-sensor",
			FriendlyName:      "My Temperature",
			Description:       "Temp monitor",
			ManufacturerName:  "admin",
			DisplayCategories: []string{alexa.DisplayCategoryTemperatureSensor},
			Capabilities: []alexa.DiscoverCapability{
				{
					Type:      "AlexaInterface",
					Interface: alexa.InterfaceTemperatureSensor,
					Version:   "3",
					Properties: &alexa.DiscoverProperties{
						Supported: []alexa.DiscoverProperty{
							{
								Name: "temperature",
							},
						},
						ProactivelyReported: false,
						Retrievable:         true,
					},
				},
			},
		},
		{
			EndpointID:        "switch-1",
			FriendlyName:      "Fan",
			Description:       "Power switch for fan",
			ManufacturerName:  "admin",
			DisplayCategories: []string{alexa.DisplayCategorySwitch},
			Capabilities: []alexa.DiscoverCapability{
				{
					Type:      "AlexaInterface",
					Interface: alexa.InterfacePowerController,
					Version:   "3",
					Properties: &alexa.DiscoverProperties{
						Supported: []alexa.DiscoverProperty{
							{
								Name: "powerState",
							},
						},
						ProactivelyReported: true,
						Retrievable:         true,
					},
				},
			},
		},
	}
}

type tempReader struct {
	temperature float32
	respBuilder *alexa.ResponseBuilder
}

func (t *tempReader) HandleRequest(ctx context.Context, req *alexa.Request) (*alexa.Response, error) {
	now := time.Now()

	temp := alexa.TemperatureValue{
		Value: t.temperature,
		Scale: alexa.TemperatureScaleCelsius,
	}

	return t.respBuilder.StateReportResponse(req,
		alexa.ContextProperty{
			Namespace:                 alexa.NamespaceTemperatureSensor,
			Name:                      "temperature",
			Value:                     temp,
			TimeOfSample:              now,
			UncertaintyInMilliseconds: 60000,
		}), nil
}
