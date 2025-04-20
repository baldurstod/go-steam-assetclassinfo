package config

type Config struct {
	HTTPS `json:"https"`
	Api   `json:"api"`
}

type HTTPS struct {
	Port          int      `json:"port"`
	HttpsKeyFile  string   `json:"https_key_file"`
	HttpsCertFile string   `json:"https_cert_file"`
	AllowOrigins  []string `json:"allow_origins"`
}

type Api struct {
	SteamApiKey string `json:"steam_api_key"`
}
