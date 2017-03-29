package weather

type Conditions struct {
	CurrentObservation struct {
		Uv              string `json:"UV"`
		DewpointC       int64  `json:"dewpoint_c"`
		DewpointF       int64  `json:"dewpoint_f"`
		DewpointString  string `json:"dewpoint_string"`
		DisplayLocation struct {
			City           string `json:"city"`
			Country        string `json:"country"`
			CountryIso3166 string `json:"country_iso3166"`
			Elevation      string `json:"elevation"`
			Full           string `json:"full"`
			Latitude       string `json:"latitude"`
			Longitude      string `json:"longitude"`
			Magic          string `json:"magic"`
			State          string `json:"state"`
			StateName      string `json:"state_name"`
			Wmo            string `json:"wmo"`
			Zip            string `json:"zip"`
		} `json:"display_location"`
		Estimated       struct{} `json:"estimated"`
		FeelslikeC      string   `json:"feelslike_c"`
		FeelslikeF      string   `json:"feelslike_f"`
		FeelslikeString string   `json:"feelslike_string"`
		ForecastURL     string   `json:"forecast_url"`
		HeatIndexC      string   `json:"heat_index_c"`
		HeatIndexF      string   `json:"heat_index_f"`
		HeatIndexString string   `json:"heat_index_string"`
		HistoryURL      string   `json:"history_url"`
		Icon            string   `json:"icon"`
		IconURL         string   `json:"icon_url"`
		Image           struct {
			Link  string `json:"link"`
			Title string `json:"title"`
			URL   string `json:"url"`
		} `json:"image"`
		LocalEpoch          string `json:"local_epoch"`
		LocalTimeRfc822     string `json:"local_time_rfc822"`
		LocalTzLong         string `json:"local_tz_long"`
		LocalTzOffset       string `json:"local_tz_offset"`
		LocalTzShort        string `json:"local_tz_short"`
		Nowcast             string `json:"nowcast"`
		ObURL               string `json:"ob_url"`
		ObservationEpoch    string `json:"observation_epoch"`
		ObservationLocation struct {
			City           string `json:"city"`
			Country        string `json:"country"`
			CountryIso3166 string `json:"country_iso3166"`
			Elevation      string `json:"elevation"`
			Full           string `json:"full"`
			Latitude       string `json:"latitude"`
			Longitude      string `json:"longitude"`
			State          string `json:"state"`
		} `json:"observation_location"`
		ObservationTime       string  `json:"observation_time"`
		ObservationTimeRfc822 string  `json:"observation_time_rfc822"`
		Precip1hrIn           string  `json:"precip_1hr_in"`
		Precip1hrMetric       string  `json:"precip_1hr_metric"`
		Precip1hrString       string  `json:"precip_1hr_string"`
		PrecipTodayIn         string  `json:"precip_today_in"`
		PrecipTodayMetric     string  `json:"precip_today_metric"`
		PrecipTodayString     string  `json:"precip_today_string"`
		PressureIn            string  `json:"pressure_in"`
		PressureMb            string  `json:"pressure_mb"`
		PressureTrend         string  `json:"pressure_trend"`
		RelativeHumidity      string  `json:"relative_humidity"`
		Solarradiation        string  `json:"solarradiation"`
		StationID             string  `json:"station_id"`
		TempC                 float64 `json:"temp_c"`
		TempF                 float64 `json:"temp_f"`
		TemperatureString     string  `json:"temperature_string"`
		VisibilityKm          string  `json:"visibility_km"`
		VisibilityMi          string  `json:"visibility_mi"`
		Weather               string  `json:"weather"`
		WindDegrees           int64   `json:"wind_degrees"`
		WindDir               string  `json:"wind_dir"`
		WindGustKph           int64   `json:"wind_gust_kph"`
		WindGustMph           int64   `json:"wind_gust_mph"`
		WindKph               float64 `json:"wind_kph"`
		WindMph               float64 `json:"wind_mph"`
		WindString            string  `json:"wind_string"`
		WindchillC            string  `json:"windchill_c"`
		WindchillF            string  `json:"windchill_f"`
		WindchillString       string  `json:"windchill_string"`
	} `json:"current_observation"`
	Response struct {
		Features struct {
			Conditions int64 `json:"conditions"`
		} `json:"features"`
		TermsofService string `json:"termsofService"`
		Version        string `json:"version"`
	} `json:"response"`
}
