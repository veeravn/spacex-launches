package main

type GraphqlResponse struct {
	Data interface{} `json:"data"`
}
type SpacexLauchesPast struct {
	Data Data `json:"data"`
}
type LaunchSite struct {
	SiteNameLong string `json:"site_name_long"`
}
type Links struct {
	ArticleLink interface{} `json:"article_link"`
	VideoLink   string      `json:"video_link"`
}
type Core struct {
	ReuseCount int    `json:"reuse_count"`
	Status     string `json:"status"`
}
type Cores struct {
	Flight int  `json:"flight"`
	Core   Core `json:"core"`
}
type FirstStage struct {
	Cores []Cores `json:"cores"`
}
type Payloads struct {
	PayloadType    string  `json:"payload_type"`
	PayloadMassKg  int     `json:"payload_mass_kg"`
	PayloadMassLbs float64 `json:"payload_mass_lbs"`
}
type SecondStage struct {
	Payloads []Payloads `json:"payloads"`
}
type Rocket struct {
	RocketName  string      `json:"rocket_name"`
	FirstStage  FirstStage  `json:"first_stage"`
	SecondStage SecondStage `json:"second_stage"`
}
type Ships struct {
	Name     string `json:"name"`
	HomePort string `json:"home_port"`
	Image    string `json:"image"`
}
type LaunchesPast struct {
	MissionName     string     `json:"mission_name"`
	LaunchDateLocal string     `json:"launch_date_local"`
	LaunchSite      LaunchSite `json:"launch_site"`
	Links           Links      `json:"links"`
	Rocket          Rocket     `json:"rocket"`
	Ships           []Ships    `json:"ships"`
}
type Data struct {
	LaunchesPast []LaunchesPast `json:"launchesPast"`
}
