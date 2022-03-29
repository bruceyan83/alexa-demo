package alexa

import (
	"encoding/json"
	"time"
)

// smart home对应结构组，详情请查看源链接
// https://developer.amazon.com/docs/smarthome/smart-home-skill-api-message-reference.html

// Request 基础的request结构体
type Request struct {
	Directive RequestDirective `json:"directive"`
}

type RequestDirective struct {
	Header   Header          `json:"header"`
	Endpoint RequestEndpoint `json:"endpoint"`
	Payload  json.RawMessage `json:"payload"`
}

type Header struct {
	Namespace        string `json:"namespace"`
	Name             string `json:"name"`
	MessageID        string `json:"messageId"`
	CorrelationToken string `json:"correlationToken,omitempty"`
	PayloadVersion   string `json:"payloadVersion"`
}

type RequestEndpoint struct {
	Scope      Scope             `json:"scope,omitempty"`
	EndpointID string            `json:"endpointId,omitempty"`
	Cookie     map[string]string `json:"cookie,omitempty"`
}

type Scope struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

// Response 基础的response结构体
type Response struct {
	Event   Event            `json:"event"`
	Context *ResponseContext `json:"context,omitempty"`
}

type ResponseContext struct {
	Properties []ContextProperty `json:"properties,omitempty"`
}

// Namespace enums
const (
	NamespaceAlexa             = "Alexa"
	NamespaceAuthorization     = "Alexa.Authorization"
	NamespaceDiscovery         = "Alexa.Discovery"
	NamespacePowerController   = "Alexa.PowerController"
	NamespaceTemperatureSensor = "Alexa.TemperatureSensor"
)

type ContextProperty struct {
	Namespace                 string      `json:"namespace"`
	Name                      string      `json:"name"`
	Value                     interface{} `json:"value"`
	TimeOfSample              time.Time   `json:"timeOfSample"`
	UncertaintyInMilliseconds int32       `json:"uncertaintyInMilliseconds"`
}

type Event struct {
	Header   Header            `json:"header"`
	Endpoint *ResponseEndpoint `json:"endpoint,omitempty"`
	Payload  json.RawMessage   `json:"payload"`
}

type ResponseEndpoint struct {
	EndpointID string            `json:"endpointId,omitempty"`
	Cookie     map[string]string `json:"cookie,omitempty"`
	Scope      Scope             `json:"scope,omitempty"`
}

// DisplayCategory enums
const (
	DisplayCategorySwitch            = "SWITCH"
	DisplayCategoryTemperatureSensor = "TEMPERATURE_SENSOR"
)

// Interface enums
const (
	InterfacePowerController   = NamespacePowerController
	InterfaceTemperatureSensor = NamespaceTemperatureSensor
)

// EmptyPayload 空payload
var EmptyPayload = json.RawMessage("{}")

type DiscoverPayload struct {
	Endpoints []DiscoverEndpoint `json:"endpoints"`
}

type DiscoverEndpoint struct {
	EndpointID        string               `json:"endpointId"`
	ManufacturerName  string               `json:"manufacturerName"`
	FriendlyName      string               `json:"friendlyName"`
	Description       string               `json:"description"`
	DisplayCategories []string             `json:"displayCategories"`
	Cookie            map[string]string    `json:"cookie,omitempty"`
	Capabilities      []DiscoverCapability `json:"capabilities"`
}

type DiscoverCapability struct {
	Type                 string              `json:"type"`
	Interface            string              `json:"interface"`
	Version              string              `json:"version"`
	Properties           *DiscoverProperties `json:"properties,omitempty"`
	SupportsDeactivation *bool               `json:"supportsDeactivation,omitempty"`
	ProactivelyReported  *bool               `json:"proactivelyReported,omitempty"`
}

type DiscoverProperties struct {
	Supported           []DiscoverProperty `json:"supported,omitempty"`
	ProactivelyReported bool               `json:"proactivelyReported"`
	Retrievable         bool               `json:"retrievable"`
}

type DiscoverProperty struct {
	Name string `json:"name"`
}

type AcceptGrantPayload struct {
	Grant   AcceptGrantGrant   `json:"grant"`
	Grantee AcceptGrantGrantee `json:"grantee"`
}

type AcceptGrantGrant struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

type AcceptGrantGrantee struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

// TemperatureScale enums
const (
	TemperatureScaleFahrenheit = "FAHRENHEIT"
	TemperatureScaleCelsius    = "CELSIUS"
)

type TemperatureValue struct {
	Value float32 `json:"value"`
	Scale string  `json:"scale"`
}
