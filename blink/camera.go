package blink

// CameraConfig is generated using gojson from /network/{}/camera/{}/config response.
type CameraConfig struct {
	Camera []struct {
		A1                        bool        `json:"a1"`
		Account                   int64       `json:"account"`
		AccountID                 int64       `json:"account_id"`
		AlertInterval             int64       `json:"alert_interval"`
		AlertRepeat               string      `json:"alert_repeat"`
		AlertToneEnable           bool        `json:"alert_tone_enable"`
		AlertToneVolume           int64       `json:"alert_tone_volume"`
		AutoTest                  bool        `json:"auto_test"`
		BatteryAlarmEnable        bool        `json:"battery_alarm_enable"`
		BatteryAlertCount         int64       `json:"battery_alert_count"`
		BatteryCheckTime          string      `json:"battery_check_time"`
		BatteryState              string      `json:"battery_state"`
		BatteryVoltage            int64       `json:"battery_voltage"`
		BatteryVoltageHysteresis  int64       `json:"battery_voltage_hysteresis"`
		BatteryVoltageInterval    int64       `json:"battery_voltage_interval"`
		BatteryVoltageThreshold   int64       `json:"battery_voltage_threshold"`
		BuzzerOn                  bool        `json:"buzzer_on"`
		CameraKey                 string      `json:"camera_key"`
		CameraSeq                 int64       `json:"camera_seq"`
		ClipBitrate               int64       `json:"clip_bitrate"`
		ClipMaxLength             int64       `json:"clip_max_length"`
		ClipRate                  int64       `json:"clip_rate"`
		ClipWarningThreshold      int64       `json:"clip_warning_threshold"`
		CreatedAt                 string      `json:"created_at"`
		DeletedAt                 interface{} `json:"deleted_at"`
		EarlyTermination          bool        `json:"early_termination"`
		EarlyTerminationSupported bool        `json:"early_termination_supported"`
		Enabled                   bool        `json:"enabled"`
		FlipImage                 bool        `json:"flip_image"`
		FwVersion                 string      `json:"fw_version"`
		ID                        int64       `json:"id"`
		IlluminatorDuration       int64       `json:"illuminator_duration"`
		IlluminatorEnable         int64       `json:"illuminator_enable"`
		IlluminatorIntensity      int64       `json:"illuminator_intensity"`
		InvertImage               bool        `json:"invert_image"`
		IPAddress                 interface{} `json:"ip_address"`
		LastBatteryAlert          interface{} `json:"last_battery_alert"`
		LastConnect               struct {
			AcPower                 bool   `json:"ac_power"`
			AccountID               int64  `json:"account_id"`
			BatteryAlertStatus      bool   `json:"battery_alert_status"`
			BatteryVoltage          int64  `json:"battery_voltage"`
			CameraID                int64  `json:"camera_id"`
			CreatedAt               string `json:"created_at"`
			Dev1                    int64  `json:"dev_1"`
			Dev2                    int64  `json:"dev_2"`
			Dev3                    int64  `json:"dev_3"`
			DhcpFailureCount        int64  `json:"dhcp_failure_count"`
			ErrorCodes              int64  `json:"error_codes"`
			FwVersion               string `json:"fw_version"`
			IPAddress               string `json:"ip_address"`
			Ipv                     string `json:"ipv"`
			Lfr108Wakeups           int64  `json:"lfr_108_wakeups"`
			LfrStrength             int64  `json:"lfr_strength"`
			LfrTbWakeups            int64  `json:"lfr_tb_wakeups"`
			LifetimeCount           int64  `json:"lifetime_count"`
			LifetimeDuration        int64  `json:"lifetime_duration"`
			LightSensorCh0          int64  `json:"light_sensor_ch0"`
			LightSensorCh1          int64  `json:"light_sensor_ch1"`
			LightSensorDataNew      bool   `json:"light_sensor_data_new"`
			LightSensorDataValid    bool   `json:"light_sensor_data_valid"`
			Mac                     string `json:"mac"`
			NetworkID               int64  `json:"network_id"`
			PirRejections           int64  `json:"pir_rejections"`
			Serial                  string `json:"serial"`
			SocketFailureCount      int64  `json:"socket_failure_count"`
			SyncModuleID            int64  `json:"sync_module_id"`
			TempAlertStatus         bool   `json:"temp_alert_status"`
			Temperature             int64  `json:"temperature"`
			Time108Boot             int64  `json:"time_108_boot"`
			TimeDhcpLease           int64  `json:"time_dhcp_lease"`
			TimeDNSResolve          int64  `json:"time_dns_resolve"`
			TimeFirstVideo          int64  `json:"time_first_video"`
			TimeWlanConnect         int64  `json:"time_wlan_connect"`
			Total108Wakeups         int64  `json:"total_108_wakeups"`
			TotalTbWakeups          int64  `json:"total_tb_wakeups"`
			UnitNumber              int64  `json:"unit_number"`
			UpdatedAt               string `json:"updated_at"`
			WifiConnectFailureCount int64  `json:"wifi_connect_failure_count"`
			WifiStrength            int64  `json:"wifi_strength"`
		} `json:"last_connect"`
		LastLfrAlert            interface{} `json:"last_lfr_alert"`
		LastOfflineAlert        interface{} `json:"last_offline_alert"`
		LastTempAlert           interface{} `json:"last_temp_alert"`
		LastWifiAlert           interface{} `json:"last_wifi_alert"`
		LfrAlertCount           int64       `json:"lfr_alert_count"`
		LfrStrength             int64       `json:"lfr_strength"`
		LfrSyncInterval         int64       `json:"lfr_sync_interval"`
		LiveviewBitrate         int64       `json:"liveview_bitrate"`
		LiveviewEnabled         string      `json:"liveview_enabled"`
		LiveviewRate            int64       `json:"liveview_rate"`
		MacAddress              interface{} `json:"mac_address"`
		MaxResolution           string      `json:"max_resolution"`
		MfgMainRange            int64       `json:"mfg_main_range"`
		MfgMainType             string      `json:"mfg_main_type"`
		MfgMezRange             int64       `json:"mfg_mez_range"`
		MfgMezType              string      `json:"mfg_mez_type"`
		MotionAlert             bool        `json:"motion_alert"`
		MotionRegions           int64       `json:"motion_regions"`
		MotionRegionsCompatible bool        `json:"motion_regions_compatible"`
		MotionSensitivity       float64     `json:"motion_sensitivity"`
		Name                    string      `json:"name"`
		Network                 int64       `json:"network"`
		NetworkID               int64       `json:"network_id"`
		OfflineAlertCount       int64       `json:"offline_alert_count"`
		Onboarded               bool        `json:"onboarded"`
		RecordAudio             bool        `json:"record_audio"`
		RecordAudioEnable       bool        `json:"record_audio_enable"`
		RetryCount              int64       `json:"retry_count"`
		Serial                  string      `json:"serial"`
		SirenEnable             bool        `json:"siren_enable"`
		SirenVolume             interface{} `json:"siren_volume"`
		Status                  string      `json:"status"`
		SyncModuleID            int64       `json:"sync_module_id"`
		TempAdjust              int64       `json:"temp_adjust"`
		TempAlarmEnable         bool        `json:"temp_alarm_enable"`
		TempAlertCount          int64       `json:"temp_alert_count"`
		TempAlertState          string      `json:"temp_alert_state"`
		TempHysteresis          interface{} `json:"temp_hysteresis"`
		TempInterval            int64       `json:"temp_interval"`
		TempMax                 interface{} `json:"temp_max"`
		TempMin                 interface{} `json:"temp_min"`
		Temperature             int64       `json:"temperature"`
		Thumbnail               string      `json:"thumbnail"`
		Type                    string      `json:"type"`
		UnitNumber              int64       `json:"unit_number"`
		UpdatedAt               string      `json:"updated_at"`
		Video50_60hz            string      `json:"video_50_60hz"`
		VideoLength             int64       `json:"video_length"`
		VideoQuality            string      `json:"video_quality"`
		VideoQualitySupport     []string    `json:"video_quality_support"`
		WifiAlertCount          int64       `json:"wifi_alert_count"`
		WifiStrength            int64       `json:"wifi_strength"`
		WifiTimeout             int64       `json:"wifi_timeout"`
	} `json:"camera"`
	Signals struct {
		Battery      int64       `json:"battery"`
		BatteryState string      `json:"battery_state"`
		Lfr          int64       `json:"lfr"`
		Temp         int64       `json:"temp"`
		UpdatedAt    interface{} `json:"updated_at"`
		Wifi         int64       `json:"wifi"`
	} `json:"signals"`
}