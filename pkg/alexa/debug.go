package alexa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"alexa-demo/pkg/schema"
	"github.com/xeipuuv/gojsonschema"
)

// DebugHandler 这个处理器会在处理前和处理后打印日志
// 并且对返回对response进行校验
func DebugHandler(handler Handler) Handler {
	return ResponseDebugHandler(RequestDebugHandler(handler))
}

// RequestDebugHandler 这个处理器会打印请求日志
func RequestDebugHandler(handler Handler) HandlerFunc {
	return func(ctx context.Context, req *Request) (*Response, error) {
		reqJSON, err := json.Marshal(req)
		if err != nil {
			log.Printf("RequestDebugHandler: Failed to marshal request: %v", err)
		} else {
			log.Printf("RequestDebugHandler: Debug request:\n%s\n", string(reqJSON))
		}

		return handler.HandleRequest(ctx, req)
	}
}

// ResponseDebugHandler 这个处理器会打印响应日志并验证响应值.
func ResponseDebugHandler(handler Handler) HandlerFunc {
	return func(ctx context.Context, req *Request) (*Response, error) {
		resp, err := handler.HandleRequest(ctx, req)

		if resp == nil {
			log.Println("Response is null.")
			return resp, err
		}

		respJSON, jsonErr := json.Marshal(resp)
		if jsonErr != nil {
			log.Printf("Failed to marshal debug response: %v\n", jsonErr)
		}
		log.Printf("Debug response:\n%s\n", respJSON)

		if schemaErr := validateSchema(string(respJSON)); schemaErr != nil {
			log.Printf("Failed to validate schema: %v\n", schemaErr)
		} else {
			log.Printf("Schema validated!\n")
		}

		return resp, err
	}
}

func validateSchema(resp string) error {
	schemaLoader := gojsonschema.NewStringLoader(schema.AlexaSmartHome)
	result, err := gojsonschema.Validate(schemaLoader, gojsonschema.NewStringLoader(resp))
	if err != nil {
		return fmt.Errorf("Failed to validate schema: %v", err)
	}
	if !result.Valid() {
		log.Printf("Response is not valid:\n")
		for _, desc := range result.Errors() {
			log.Printf("- %s\n", desc)
		}
		return errors.New("Response is not valid")
	}
	return nil
}
