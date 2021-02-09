package routes

type AppInfo struct {
	OK        bool   `json:"ok"`
	BuildHash string `json:"build_hash"`
	BuildTime string `json:"build_time"`
	AppName   string `json:"app_name"`
}
