package entity

type SentenceSta struct{
	Uploaded int `json:"uploaded"`
	Approved int `json:"approved"`
	Requested int `json:"requested"`
}

type WallpaperSta struct{
	Approved int `json:"approved"`
	Requested int `json:"requested"`
}

type OtherSta struct{
	SiteServed     int `json:"site_served"`
	WpPluginLatest string `json:"wp_plugin_latest"`
}

type Data struct{
	Sentence SentenceSta `json:"sentence"`
	Wallpaper WallpaperSta `json:"wallpaper"`
	Other OtherSta `json:"other"`
}

type Statistic struct{
	Status string `json:"status"`
	Data Data `json:"data"`
}
