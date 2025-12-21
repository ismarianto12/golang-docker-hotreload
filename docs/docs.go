package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "Scope Api Pembayaran",
        "title": "Scope Api Pembayaran",
        "contact": {},
        "version": "V1.0"
    },
    "host": "localhost:6060",
    "basePath": "{{.BasePath}}",
    "paths": {
	  "/api/v1/product/list":{
	     "get": {"/list"} 
	  }
	}
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
