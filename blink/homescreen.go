package blink

// Homescreen is generated using gojson from /api/v3/accounts/{}/homescreen response.
type Homescreen struct {
	AppUpdates struct {
		Code            int64  `json:"code"`
		Message         string `json:"message"`
		UpdateAvailable bool   `json:"update_available"`
		UpdateRequired  bool   `json:"update_required"`
	} `json:"app_updates"`
	Cameras []struct {
		Battery   string `json:"battery"`
		CreatedAt string `json:"created_at"`
		Enabled   bool   `json:"enabled"`
		FwVersion string `json:"fw_version"`
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		NetworkID int64  `json:"network_id"`
		Serial    string `json:"serial"`
		Status    string `json:"status"`
		Thumbnail string `json:"thumbnail"`
		Type      string `json:"type"`
		UpdatedAt string `json:"updated_at"`
		UsageRate bool   `json:"usage_rate"`
	} `json:"cameras"`
	Networks []struct {
		Armed     bool   `json:"armed"`
		CreatedAt string `json:"created_at"`
		Dst       bool   `json:"dst"`
		ID        int64  `json:"id"`
		LvSave    bool   `json:"lv_save"`
		Name      string `json:"name"`
		TimeZone  string `json:"time_zone"`
		UpdatedAt string `json:"updated_at"`
	} `json:"networks"`
	SyncModules []struct {
		CreatedAt        string `json:"created_at"`
		EnableTempAlerts bool   `json:"enable_temp_alerts"`
		FwVersion        string `json:"fw_version"`
		ID               int64  `json:"id"`
		LastHb           string `json:"last_hb"`
		Name             string `json:"name"`
		NetworkID        int64  `json:"network_id"`
		Onboarded        bool   `json:"onboarded"`
		Serial           string `json:"serial"`
		Status           string `json:"status"`
		UpdatedAt        string `json:"updated_at"`
		WifiStrength     int64  `json:"wifi_strength"`
	} `json:"sync_modules"`
	VideoStats struct {
		AutoDeleteDays int64 `json:"auto_delete_days"`
		Storage        int64 `json:"storage"`
	} `json:"video_stats"`
}
