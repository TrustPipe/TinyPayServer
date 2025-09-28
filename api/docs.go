package api

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed openapi.yaml
var openapiSpec string

// SetupSwaggerUI sets up Swagger UI for API documentation
func SetupSwaggerUI(router *gin.Engine) {
	// Serve the OpenAPI spec
	router.GET("/api/openapi.yaml", func(c *gin.Context) {
		c.Header("Content-Type", "application/x-yaml")
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.String(http.StatusOK, openapiSpec)
	})

	// Serve Swagger UI
	router.GET("/docs", func(c *gin.Context) {
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		html := `<!DOCTYPE html>
<html>
<head>
	<title>TinyPay API Documentation</title>
	<link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui.css" />
	<style>
		html {
			box-sizing: border-box;
			overflow: -moz-scrollbars-vertical;
			overflow-y: scroll;
		}
		*, *:before, *:after {
			box-sizing: inherit;
		}
		body {
			margin:0;
			background: #fafafa;
		}
	</style>
</head>
<body>
	<div id="swagger-ui"></div>
	<script src="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui-bundle.js"></script>
	<script src="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui-standalone-preset.js"></script>
	<script>
		window.onload = function() {
			const ui = SwaggerUIBundle({
				url: '/api/openapi.yaml',
				dom_id: '#swagger-ui',
				deepLinking: true,
				presets: [
					SwaggerUIBundle.presets.apis,
					SwaggerUIStandalonePreset
				],
				plugins: [
					SwaggerUIBundle.plugins.DownloadUrl
				],
				layout: "StandaloneLayout"
			});
		};
	</script>
</body>
</html>`
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, html)
	})

	// Redirect /docs/ to /docs
	router.GET("/docs/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/docs")
	})
}

// ServeOpenAPISpec serves the OpenAPI specification
func ServeOpenAPISpec(c *gin.Context) {
	c.Header("Content-Type", "application/x-yaml")
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")

	// Use embedded OpenAPI spec
	c.String(http.StatusOK, openapiSpec)
}

// SetupDocumentationRoutes sets up all documentation related routes
func SetupDocumentationRoutes(router *gin.Engine) {
	// Serve Swagger UI
	router.GET("/docs", func(c *gin.Context) {
		html := `<!DOCTYPE html>
<html>
<head>
	<title>TinyPay API Documentation</title>
	<link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui.css" />
	<style>
		html {
			box-sizing: border-box;
			overflow: -moz-scrollbars-vertical;
			overflow-y: scroll;
		}
		*, *:before, *:after {
			box-sizing: inherit;
		}
		body {
			margin:0;
			background: #fafafa;
		}
	</style>
</head>
<body>
	<div id="swagger-ui"></div>
	<script src="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui-bundle.js"></script>
	<script src="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui-standalone-preset.js"></script>
	<script>
		window.onload = function() {
			const ui = SwaggerUIBundle({
				url: '/openapi.yaml',
				dom_id: '#swagger-ui',
				deepLinking: true,
				presets: [
					SwaggerUIBundle.presets.apis,
					SwaggerUIStandalonePreset
				],
				plugins: [
					SwaggerUIBundle.plugins.DownloadUrl
				],
				layout: "StandaloneLayout"
			});
		};
	</script>
</body>
</html>`
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, html)
	})

	// Serve OpenAPI spec
	router.GET("/openapi.yaml", ServeOpenAPISpec)
	
	// Redirect /docs/ to /docs
	router.GET("/docs/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/docs")
	})
}