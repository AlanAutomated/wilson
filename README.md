

# wilson

<img src="https://github.com/autoalan/wilson/raw/master/images/wilson.jpg" alt="wilson" align="right" width="10%"/>A Go (golang) [dot1x](https://github.com/layeh/radius) server that runs locally on network switches as secondary authentication server. When a switch becomes stranded from its primary dot1x server, wilson will authenticate endpoints by [OUI](https://en.wikipedia.org/wiki/Organizationally_unique_identifier) using a flexible policy. Wilson was developed to provide supplemental `policy-map type control` support for an [Arista EOS](https://www.arista.com/en/products/eos) campus healthcare environment with colorless ports and a high up-time requirement.



## TODO

This is effectively a prototype. While it servers the intended purpose, it needs tests and perhaps some refactoring. The goal of wilson is to be easily readable and maintainable. 



## Installation & Compilation

```
go get -u github.com/autoalan/wilson
```

Wilson will run once compiled without modification on most platforms. Simply clone this repository and compile wilson.go in the apps folder.  For Arista EOS switches, use the 386 architecture.  

```
# GOARCH=386 go build apps/wilson.go
```



## Usage

When executed for the first time, wilson expects to load its configuration from URL. Subsequent executions will use a defined configuration file (.wilson by default) automatically created in the directory containing the directory if the the server is unreachable or if the URL flag is omitted.


```
./wilson -url https://my-lb-site.internal.org/wilson.json
```

For implementations on Arista EOS, consider using an [event-handler](https://www.arista.com/en/um-eos/eos-section-3-9-command-line-interface-commands#ww1124606) or event perhaps [rc.eos](https://www.arista.com/assets/data/pdf/Whitepapers/Arista_EOS_parser.pdf).


Wilson expects the configuration to conform to a known [JSON schema](https://github.com/autoalan/wilson/blob/35d32639d00e05aeb7c1a301a6d2e96112d53034/schema.go#L23-L101). Below is an example of a configuration file.

```json
   {
   	"configFile": ".wilson",
   	"configRefresh": 5,
   	"configURL": "https://my-lb-site.internal.org/wilson.json",
   	"serverBinding": "127.0.0.1:1812",
   	"serverSecret": "127001",
   	"policies": [{
   			"comment": "Issue an access-accept for trusted Roche analyzers",
   			"clientOui": "B8:78:79",
   			"clientVlan": 5,
   			"radiusCode": 2
   		},
   		{
   			"comment": "Issue an access-reject for unauthorized TP-Link endpoints",
   			"clientOui": "d8-07-b6",
   			"clientVlan": 0,
   			"radiusCode": 2
   		},
   		{
   			"comment": "Ignore all other requests",
   			"clientOui": "0000.00",
   			"clientVlanvlan": 0,
   			"radiusCode": 0
   		}
   	]
   }
```



| Parameter     | Description                                                  |
| ------------- | ------------------------------------------------------------ |
| configFile    | This is the path to the configuration that will be saved locally if the the server hosting the URL config is unavailable. |
| configRefresh | The interval in seconds to poll the server for configuration updates. |
| configURL     | The URL to the initial configuration file.                   |
| serverBinding | The server binding used for requests. Typically this will be localhost for obvious reasons. |
| serverSecret  | The RADIUS secret to authenticate the NAS client.            |
| comment       | Ignored by wilson. This is for humans.                       |
| clientOui     | A 24-bit hexadecimal string representing the OUI of a MAC address. Delimiters (:, -, .) are ignored. |
| clientVlan    | The VLAN to be assigned to the client on access-accept.      |
| radiusCode    | [Standard RADIUS codes supported](https://github.com/layeh/radius/blob/3e43fd4ead922ac65515918994c1e7942d1f0013/code.go#L11-L27) by the underlying radius library. A typically deployment would leverage 2 (Accept), 3 (Reject) and 0 (Ignore or discard the request). |



## License

MPL 2.0



## Author

Alan Haynes (alan@networkautomation.engineer). 

Huge thanks to Tim Cooper for the superb radius implementation.

