package roku

// DeviceInfo provides detailed information about the device the client is connected to
type DeviceInfo struct {
	Udn                      string `xml:"udn"`
	SerialNumber             string `xml:"serial-number"`
	DeviceID                 string `xml:"device-id"`
	AdvertisingID            string `xml:"advertising-id"`
	VendorName               string `xml:"vendor-name"`
	ModelName                string `xml:"model-name"`
	ModelNumber              string `xml:"model-number"`
	ModelRegion              string `xml:"model-region"`
	IsTv                     bool   `xml:"is-tv"`
	IsStick                  bool   `xml:"is-stick"`
	UiResolution             string `xml:"ui-resolution"`
	SupportsEthernet         bool   `xml:"supports-ethernet"`
	WifiMac                  string `xml:"wifi-mac"`
	WifiDriver               string `xml:"wifi-driver"`
	HasWifiExtender          bool   `xml:"has-wifi-extender"`
	HasWifi5GSupport         bool   `xml:"has-wifi-5G-support"`
	CanUseWifiExtender       bool   `xml:"can-use-wifi-extender"`
	NetworkType              string `xml:"network-type"`
	NetworkName              string `xml:"network-name"`
	FriendlyDeviceName       string `xml:"friendly-device-name"`
	FriendlyModelName        string `xml:"friendly-model-name"`
	DefaultDeviceName        string `xml:"default-device-name"`
	UserDeviceName           string `xml:"user-device-name"`
	UserDeviceLocation       string `xml:"user-device-location"`
	BuildNumber              string `xml:"build-number"`
	SoftwareVersion          string `xml:"software-version"`
	SoftwareBuild            string `xml:"software-build"`
	SecureDevice             bool   `xml:"secure-device"`
	Language                 string `xml:"language"`
	Country                  string `xml:"country"`
	Locale                   string `xml:"locale"`
	TimeZoneAuto             bool   `xml:"time-zone-auto"`
	TimeZone                 string `xml:"time-zone"`
	TimeZoneName             string `xml:"time-zone-name"`
	TimeZoneTz               string `xml:"time-zone-tz"`
	TimeZoneOffset           string `xml:"time-zone-offset"`
	ClockFormat              string `xml:"clock-format"`
	Uptime                   string `xml:"uptime"`
	PowerMode                string `xml:"power-mode"`
	SupportsSuspend          bool   `xml:"supports-suspend"`
	SupportsFindRemote       bool   `xml:"supports-find-remote"`
	FindRemoteIsPossible     bool   `xml:"find-remote-is-possible"`
	SupportsAudioGuide       bool   `xml:"supports-audio-guide"`
	SupportsRva              bool   `xml:"supports-rva"`
	DeveloperEnabled         bool   `xml:"developer-enabled"`
	KeyedDeveloperID         string `xml:"keyed-developer-id"`
	SearchEnabled            bool   `xml:"search-enabled"`
	SearchChannelsEnabled    bool   `xml:"search-channels-enabled"`
	VoiceSearchEnabled       bool   `xml:"voice-search-enabled"`
	NotificationsEnabled     bool   `xml:"notifications-enabled"`
	NotificationsFirstUse    bool   `xml:"notifications-first-use"`
	SupportsPrivateListening bool   `xml:"supports-private-listening"`
	HeadphonesConnected      bool   `xml:"headphones-connected"`
	SupportsEcsTextedit      bool   `xml:"supports-ecs-textedit"`
	SupportsEcsMicrophone    bool   `xml:"supports-ecs-microphone"`
	SupportsWakeOnWlan       bool   `xml:"supports-wake-on-wlan"`
	HasPlayOnRoku            bool   `xml:"has-play-on-roku"`
	HasMobileScreensaver     bool   `xml:"has-mobile-screensaver"`
	SupportURL               string `xml:"support-url"`
	GrandcentralVersion      string `xml:"grandcentral-version"`
	TrcVersion               string `xml:"trc-version"`
	TrcChannelVersion        string `xml:"trc-channel-version"`
	DavinciVersion           string `xml:"davinci-version"`
}
