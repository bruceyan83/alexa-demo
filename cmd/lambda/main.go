package main

import (
	"alexa-demo/pkg/alexa"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"alexa-demo/pkg/lambda"
	awslambda "github.com/aws/aws-lambda-go/lambda"
)

// 这是aws-lambda的golang版本demo实现
func main() {
	respBuilder := alexa.NewResponseBuilder()
	tempReader := tempReader{75, respBuilder}
	mux := alexa.NewRouteNameSpace()
	//mux.HandleFunc(alexa.NamespacePercentageController, alexa.DeferredRelayHandler(sqsRelay, respBuilder))
	//mux.HandleFunc(alexa.NamespacePowerController, alexa.DeferredRelayHandler(sqsRelay, respBuilder))
	mux.HandleFunc(alexa.NamespaceDiscovery, alexa.StaticDiscoveryHandler(respBuilder, endpoints()...))
	mux.HandleFunc(alexa.NamespaceAlexa, tempReader.GetTemperature)
	mux.HandleFunc(alexa.NamespaceAuthorization,
		alexa.AuthorizationHandler(respBuilder))

	awslambda.Start(lambda.DebugLambdaRequestHandler(mux))
}

func endpoints() []alexa.DiscoverEndpoint {
	return []alexa.DiscoverEndpoint{
		{
			EndpointID:        "temp-sensor-1",
			FriendlyName:      "Home Temperature",
			Description:       "Temp monitor",
			ManufacturerName:  "McTofu",
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
			ManufacturerName:  "McTofu",
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
		{
			EndpointID:        "window-1",
			FriendlyName:      "Window",
			Description:       "Window control",
			ManufacturerName:  "McTofu",
			DisplayCategories: []string{alexa.DisplayCategoryOther},
			Capabilities: []alexa.DiscoverCapability{
				{
					Type:      "AlexaInterface",
					Interface: alexa.InterfacePercentageController,
					Version:   "3",
					Properties: &alexa.DiscoverProperties{
						Supported: []alexa.DiscoverProperty{
							{
								Name: "percentage",
							},
						},
						ProactivelyReported: false,
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

func (t *tempReader) GetTemperature(ctx context.Context, req *alexa.Request) (*alexa.Response, error) {
	now := time.Now()

	temp := alexa.TemperatureValue{
		Value: t.temperature,
		Scale: alexa.TemperatureScaleFahrenheit,
	}

	tempJSON, err := json.Marshal(temp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal temp: %v", err)
	}

	return t.respBuilder.StateReportResponse(req,
		alexa.ContextProperty{
			Namespace:                 alexa.NamespaceTemperatureSensor,
			Name:                      "temperature",
			Value:                     tempJSON,
			TimeOfSample:              now,
			UncertaintyInMilliseconds: 60000,
		}), nil
}
