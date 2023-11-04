package fiber

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func NewFiberRouter() *fiber.App {
	conf := fiber.Config{
		// For this scenario it is better to turn the pre-fork off
		Prefork: false,

		ServerHeader:  "",
		StrictRouting: false,
		CaseSensitive: false,
		Immutable:     false,
		UnescapePath:  false,
		ETag:          false,
		BodyLimit:     0,

		// default : 256 * 1024
		Concurrency: 256 * 1024,

		Views:                nil,
		ViewsLayout:          "",
		PassLocalsToViews:    false,
		ReadTimeout:          0,
		WriteTimeout:         0,
		IdleTimeout:          0,
		ReadBufferSize:       0,
		WriteBufferSize:      0,
		CompressedFileSuffix: "",
		ProxyHeader:          "",

		// for this scenario, GETOnly has been set to true
		GETOnly: true,

		ErrorHandler: nil,

		// http 1.1 like behavior
		DisableKeepalive: false,

		DisableDefaultDate:           false,
		DisableDefaultContentType:    false,
		DisableHeaderNormalizing:     false,
		DisableStartupMessage:        false,
		AppName:                      "History Service",
		StreamRequestBody:            false,
		DisablePreParseMultipartForm: false,
		ReduceMemoryUsage:            false,

		// For faster json marshalling, github.com/goccy/go-json has been replaced
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,

		XMLEncoder:              nil,
		Network:                 "",
		EnableTrustedProxyCheck: false,
		TrustedProxies:          nil,
		EnableIPValidation:      false,

		EnablePrintRoutes: false,

		ColorScheme:              fiber.Colors{},
		RequestMethods:           nil,
		EnableSplittingOnParsers: false,
	}

	return fiber.New(conf)
}
