locals {
    lambda_function_name = "testee"
}


resource "aws_schemas_schema" "testee" {
name          = "_${local.lambda_function_name}-schema"
  registry_name = "lambda-testevent-schemas"
  type          = "OpenApi3"
  description   = "console tests test"

  content = jsonencode(  {
  "openapi": "3.0.0",
  "info": {
    "version": "1.0.0",
    "title": "Event"
  },
  "paths": {},
  "components": {
    "schemas": {
      "Event": {
        "type": "object",
        "required": [
          "key1"
        ],
        "properties": {
          "key1": {
            "type": "string"
          }
        }
      }
    },
    "examples": {
      "Parameter1": {
        "value": {
          "key1": "value1"
        }
      },
      "Parameter2": {
        "value": {
          "key1": "value2"
        }
      }
    }
  }
} )
}