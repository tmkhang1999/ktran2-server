package structs

type AwsItem struct {
	Name           string `json:"name"`
	Country        string `json:"country"`
	Region         string `json:"region"`
	Lat            string `json:"lat"`
	Lon            string `json:"lon"`
	TimezoneID     string `json:"timezone_id"`
	Localtime      string `json:"localtime"`
	LocaltimeEpoch int    `json:"localtime_epoch"`
	UtcOffset      string `json:"utc_offset"`

	ObservationTime     string   `json:"observation_time"`
	Temperature         int      `json:"temperature"`
	WeatherCode         int      `json:"weather_code"`
	WeatherIcons        []string `json:"weather_icons"`
	WeatherDescriptions []string `json:"weather_descriptions"`
	WindSpeed           int      `json:"wind_speed"`
	WindDegree          int      `json:"wind_degree"`
	WindDir             string   `json:"wind_dir"`
	Pressure            int      `json:"pressure"`
	Precip              int      `json:"precip"`
	Humidity            int      `json:"humidity"`
	Cloudcover          int      `json:"cloudcover"`
	Feelslike           int      `json:"feelslike"`
	UvIndex             int      `json:"uv_index"`
	Visibility          int      `json:"visibility"`
	IsDay               string   `json:"is_day"`
}
