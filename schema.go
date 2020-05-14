package wilson


type Configuration struct {

	ConfigFile		string		`json:"configFile`
	ConfigRefresh	float64		`json:"configRefresh"`
	ConfigUrl		string		`json:"configUrl"`
	ServerBinding	string		`json:"serverBinding"`
	ServerSecret	string		`json:"serverSecret"`
	Policies		[]Pol		`json:"policies"`
}


type Pol struct {
	ClientOui		string	`json:"clientOui"`
	ClientVlan		int		`json:"clientVlan"`
	RadiusCode		int		`json:"radiusCode"`
}


const Schema string = `
{
	"$schema": "http://json-schema.org/draft-07/schema",
	"$id": "github.com/autoalan/wilson/wilson.json",
	"type": "object",
	"required": [
		"configFile",
		"configRefresh",
		"configUrl",
		"serverBinding",
		"serverSecret",
		"policies"
	],
	"additionalProperties": false,
	"properties": {
		"configFile": {
			"$id": "#/properties/configFile",
			"type": "string",
			"default": "0"
		},
		"configRefresh": {
			"$id": "#/properties/configRefresh",
			"type": "int",
			"default": ""
		},
		"configUrl": {
			"$id": "#/properties/configUrl",
			"type": "string",
			"default": ""
		},
		"serverBinding": {
			"$id": "#/properties/serverBinding",
			"type": "string",
			"default": ""
		},		
		"serverSecret": {
			"$id": "#/properties/serverSecret",
			"type": "string",
			"default": 0
		},
		"policies": {
			"$id": "#/properties/policies",
			"type": "array",
			"additionalItems": true,
			"items": {
				"anyOf": [{
					"$id": "#/properties/policies/items/anyOf/0",
					"type": "object",
					"required": [
						"clientOui",
						"clientVlan",
						"radiusCode"
					],
					"additionalProperties": false,
					"properties": {
						"comment": {
							"$id": "#/properties/policies/items/anyOf/0/properties/comment",
							"type": "string"
						},
						"clientOui": {
							"$id": "#/properties/policies/items/anyOf/0/properties/clientOui",
							"type": "string"
						},
						"clientVlan": {
							"$id": "#/properties/policies/items/anyOf/0/properties/clientVlan",
							"type": "integer",
							"default": 0
						},
						"radiusCode": {
							"$id": "#/properties/policies/items/anyOf/0/properties/radiusCode",
							"type": "integer",
							"default": 3
						}
					}
				}],
				"$id": "#/properties/policies/items"
			}
		}
	}
}
`