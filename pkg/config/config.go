package config

var (
	// Data is the configuration data for the application.
	Data = &config{BaseURL: "https://api.clashofclans.com/v1"}
)

type config struct {
	BaseURL        string `json:"base_url"`
	ResponseFormat string `json:"response_format"`
	Log            struct {
		Dir   string `json:"dir"`
		Trial struct {
			File  string `json:"file"`
			Level string `json:"level"`
		} `json:"trial"`
	} `json:"log"`
}
