package main

import(
	"flag"
	"log"
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
	"layeh.com/radius/rfc2868"
	"github.com/autoalan/wilson"
)

var url string
var config wilson.Configuration


func init() {
	// Check for the URL first thing; this may get folded into main
	flag.StringVar(&url, "url", "", wilson.HelpURLFlag)
	flag.Parse()
}

func main() {

	config := wilson.Config(url)
	channel := make(chan wilson.Configuration)

	// Refresh the configuration on an interval if loaded from a URL
	if url != "" {
		interval := config.ConfigRefresh
		// Loop forever while running to check for configuration updates
		go func() { for true { channel <- wilson.RefreshConfig(interval, url) } }()
	}
	
	
	// Enter the handler to "handle" RADIUS requests
	handler := func(w radius.ResponseWriter, r *radius.Request) {
	
		// Check the channel for an updated config before serving the request in a non-blocking fashion
		select {
			case config = <- channel: 
				log.Println(wilson.NoticeConfigUpdated)
			default:
		}
			// The CallingStationID should contain the MAC address of the endpoint
			caller := rfc2865.CallingStationID_GetString(r.Packet)
			code, vlan := wilson.Policy(caller, config)
			
			// Effectively discard any requests without a valid code
			if code == 0 {
				log.Printf("wilson: Notice: Failed to match policy for caller %v", caller)
				return
			}
			
			// Return the matched policy
			response := radius.Code(code)
			packet := r.Response(radius.Code(response))
			
			log.Printf("wilson: Notice: Issued %v for %v", response, caller)
	
			var tag byte = 0	
			svlan := strconv.Itoa(vlan)
			
			// Attributes required for dynamic VLAN assignment
			rfc2868.TunnelPrivateGroupID_Set(packet, tag, []byte(svlan))
			rfc2868.TunnelType_Add(packet, 0, 13)
			rfc2868.TunnelMediumType_Add(packet, byte(0), 6)
			
			// Send the response packet
			w.Write(packet)
		
		}
	
	// Initialize the server
	secret := config.ServerSecret
	listen := config.ServerBinding
	server := radius.PacketServer {
		Handler:      radius.HandlerFunc(handler),
		Addr: 		  listen,
		SecretSource: radius.StaticSecretSource([]byte(secret)),
	
	}
	
	log.Printf(wilson.NoticeStarted)
	
	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(wilson.ErrorStartup, err)
	}
}
