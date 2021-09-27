package main

//GraphqlResponse Parent type for response from graph api
type GraphqlResponse struct {
	Data interface{} `json:"data"`
}

// LaunchSite for response from graph api
type LaunchSite struct {
	SiteNameLong string `json:"site_name_long"`
}

//Links type
type Links struct {
	ArticleLink interface{} `json:"article_link"`
	VideoLink   string      `json:"video_link"`
}

//Core type
type Core struct {
	ReuseCount int    `json:"reuse_count"`
	Status     string `json:"status"`
}

//Cores type
type Cores struct {
	Flight int  `json:"flight"`
	Core   Core `json:"core"`
}

//FirstStage type
type FirstStage struct {
	Cores []Cores `json:"cores"`
}

//Payloads type
type Payloads struct {
	PayloadType    string  `json:"payload_type"`
	PayloadMassKg  float64 `json:"payload_mass_kg"`
	PayloadMassLbs float64 `json:"payload_mass_lbs"`
}

//SecondStage data type
type SecondStage struct {
	Payloads []Payloads `json:"payloads"`
}

//Rocket information type
type Rocket struct {
	RocketName  string      `json:"rocket_name"`
	FirstStage  FirstStage  `json:"first_stage"`
	SecondStage SecondStage `json:"second_stage"`
}

//Ships type
type Ships struct {
	Name     string `json:"name"`
	HomePort string `json:"home_port"`
	Image    string `json:"image"`
}

//LaunchesPast type
type LaunchesPast struct {
	MissionName     string     `json:"mission_name"`
	LaunchDateLocal string     `json:"launch_date_local"`
	LaunchSite      LaunchSite `json:"launch_site"`
	Links           Links      `json:"links"`
	Rocket          Rocket     `json:"rocket"`
	Ships           []Ships    `json:"ships"`
}

//Data type
type Data struct {
	LaunchesPast []LaunchesPast `json:"launchesPast"`
}

type LaunchesPastTable struct {
	MissionName  string
	SiteNameLong string
	RocketName   string
}
type SpaceXQueryForm struct {
	Limit  int
	Offset int
}
