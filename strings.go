package wilson

const (
	ErrConfigNotFound			= "wilson: Error: The configuration could not be downloaded or read locally from disk."
	ErrConfigNotValid			= "wilson: Error: The configuration failed."
	ErrConfigDecodeFailed		= "wilson: Error: The configuration could not be decoded from JSON."
	ErrorStartup				= "wilson: Error: Wilson is adrift: "
	HelpURLFlag					= "The URL to a JSON configuration"
	NoticeStarted				= "wilson: Notice: Wilson is listening to you"
	NoticeConfigUpdated			= "wilson: Notice: The configuration was refreshed as requested"
	WarnConfigWriteFailed 		= "wilson: Warning: Failed to write the configuration to disk."
	WarnPolicyBadMAC			= "wilson: Warning: Failed to extract a valid OUI from address "
	WarnPolicyDiscardRequest	= "wilson: Warning: Discarding request for invalid OUI "
)