package testdata

const EnumWithNoOneOf = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "someChoice": {
			"enum": [
				"FOOa",
				"BAR",
				"FIZZ",
				"BUZZ"
			],
            "type": "string"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
