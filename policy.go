package wilson


import(
	"log"
	"regexp"
	"strings"
)


// Return the first 24 bits of a MAC address
func returnOUI(mac string) (string) {

	// A simple way to account for the various MAC patterns
	expression := regexp.MustCompile("[^a-fA-F0-9]+")

	result := expression.ReplaceAllString(mac, "")
	if len(result) < 6 {
		log.Printf(WarnPolicyBadMAC, mac)
		return ""
	}

	// First six hexademical characters
	oui := strings.ToLower(result[0:6])

	return oui
}


// Return the first matching policy code and VLAN
func Policy(oui string, config Configuration) (int, int) {

	// Should be derived from CallingStationID
	callerOui := returnOUI(oui)
	// Discard requests to match invalid OUIs
	if callerOui == "" {
		log.Printf(WarnPolicyDiscardRequest, oui)
		return 0, 0
	}


	// Loop through the policies for a match
	for _, v := range config.Policies {
		// This will also warn for invalid OUI patterns in the configuration file
		if returnOUI(v.ClientOui) == callerOui {
			return v.RadiusCode, v.ClientVlan
		}
	}

	// If there is not a match at this point, pull the last policy for the default action
	i := len(config.Policies) - 1
	policy := config.Policies[i]

	return policy.RadiusCode, policy.ClientVlan
}