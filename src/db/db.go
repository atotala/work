package db

//Sample comment
type Asset struct {
	Asset_name     string `json:"asset_name"`
	Asset_id       string `json:"asset_id"`
	Asset_version  string `json:"asset_version"`
	Latest_version string `json:"latest_version"`
	Asset_state    string `json:"asset_state"`
}

type SystemState struct {
	System_state string `json:"system_sate"`
}

var Assets []Asset
var System_state SystemState

//func comment
func GetAssets() ([]Asset){

	Assets = append(Assets, Asset{Asset_name: "Management switch", Asset_id: "10", Asset_version: "v1", Latest_version: "v4", Asset_state: "Quarantined"})

	Assets = append(Assets, Asset{Asset_name: "Management switch", Asset_id: "11", Asset_version: "v1", Latest_version: "v2", Asset_state: "Amber"})

	Assets = append(Assets, Asset{Asset_name: "Management switch", Asset_id: "12", Asset_version: "v1", Latest_version: "v3", Asset_state: "Red"})

	Assets = append(Assets, Asset{Asset_name: "Management switch", Asset_id: "13", Asset_version: "v1", Latest_version: "v2", Asset_state: "Green"})

	Assets = append(Assets, Asset{Asset_name: "IPI Controller", Asset_id: "23", Asset_version: "v1", Latest_version: "v2", Asset_state: "Amber"})

	Assets = append(Assets, Asset{Asset_name: "IPI Controller", Asset_id: "20", Asset_version: "v1", Latest_version: "v3", Asset_state: "Red"})

	Assets = append(Assets, Asset{Asset_name: "IPI Controlle", Asset_id: "22", Asset_version: "v1", Latest_version: "v4", Asset_state: "Quarantined"})

	Assets = append(Assets, Asset{Asset_name: "IPI Controller", Asset_id: "21", Asset_version: "v1", Latest_version: "v1", Asset_state: "Green"})


	return Assets

}

func GetSystemState() (SystemState){

	System_state = SystemState{System_state: "Red"}
	return System_state
}
