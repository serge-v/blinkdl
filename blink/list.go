package blink

// List is /api/v2/videos/changed?since={}&page={} response (Videos).
// Or /api/v1/accounts/{}/media/changed?since={}&page={} response (Media).
type List struct {
	Limit        int64 `json:"limit"`
	PurgeID      int64 `json:"purge_id"`
	RefreshCount int64 `json:"refresh_count"`
	Media        []struct {
		AdditionalDevices []interface{} `json:"additional_devices"`
		CreatedAt         string        `json:"created_at"`
		Deleted           bool          `json:"deleted"`
		Device            string        `json:"device"`
		DeviceID          int64         `json:"device_id"`
		DeviceName        string        `json:"device_name"`
		ID                int64         `json:"id"`
		Media             string        `json:"media"`
		NetworkID         int64         `json:"network_id"`
		NetworkName       string        `json:"network_name"`
		Partial           bool          `json:"partial"`
		Source            string        `json:"source"`
		Thumbnail         string        `json:"thumbnail"`
		TimeZone          string        `json:"time_zone"`
		Type              string        `json:"type"`
		UpdatedAt         string        `json:"updated_at"`
		Watched           bool          `json:"watched"`
	} `json:"media"`
	Videos []struct {
		AccountID       int64       `json:"account_id"`
		Address         string      `json:"address"`
		CameraID        int64       `json:"camera_id"`
		CameraName      string      `json:"camera_name"`
		CreatedAt       string      `json:"created_at"`
		Deleted         bool        `json:"deleted"`
		Description     string      `json:"description"`
		Encryption      string      `json:"encryption"`
		EncryptionKey   interface{} `json:"encryption_key"`
		EventID         interface{} `json:"event_id"`
		ID              int64       `json:"id"`
		Length          int64       `json:"length"`
		Locked          bool        `json:"locked"`
		NetworkID       int64       `json:"network_id"`
		NetworkName     string      `json:"network_name"`
		Partial         bool        `json:"partial"`
		Ready           bool        `json:"ready"`
		Size            int64       `json:"size"`
		StorageLocation string      `json:"storage_location"`
		Thumbnail       string      `json:"thumbnail"`
		TimeZone        string      `json:"time_zone"`
		UpdatedAt       string      `json:"updated_at"`
		UploadTime      int64       `json:"upload_time"`
		Viewed          string      `json:"viewed"`
	} `json:"videos"`
}
