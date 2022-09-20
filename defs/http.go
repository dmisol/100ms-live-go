package defs

type Register struct {
	Email string `json:"email"`
}

type RegisterResp struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	TicketNumber string `json:"ticketNumber"`
	CreatedAt    string `json:"createdAt"`
}

type Auth struct {
	LoggedIn bool `json:"loggedIn"`
}

type Token struct {
	Role string `json:"role"`
	Room string `json:"room_id"`
}

type TokenResp struct {
	Token   string `json:"token"`
	Msg     string `json:"msg"`
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Version string `json:"api_version"`
}

// token in "authorization: Bearer"
type InitResp struct {
	Ep        string `json:"endpoint"`
	RtcConfig `json:"rtcConfiguration"`
	Config    `json:"config"`
}

type RtcConfig struct {
	IceSrvs []IceSrv `json:"ice_servers"`
}

type IceSrv struct {
	Urls       []string `json:"urls"`
	Username   string   `json:"username"`
	Credential string   `json:"credential"`
}

type Config struct {
	Flags    []string `json:"enabledFlags"`
	NwHealth `json:"networkHealth"`
}

type NwHealth struct {
	ScoreMap interface{} `json:"scoreMap"`
	Timeout  int         `json:"timeout"`
	Url      string      `json:"url"`
}

/*
{
	"endpoint":"wss://prod-in2.100ms.live/v2/ws",
	"rtcConfiguration":{
		"ice_servers":[
			{
				"urls":[
					"turn:turn-in.100ms.live",
					"turn:turn-in.100ms.live?transport=tcp",
					"turn:turn-in.100ms.live:443",
					"turn:turn-in.100ms.live:443?transport=tcp"
				],
				"username":"1663677392:",
			 	"credential":"mjI2cBOD1N+Odrc13G564wM5ZlM="
			}
		]
	},
	"config":{
		"enabledFlags":[
				"nonWebRTCDisableOffer"
			],
		"networkHealth":{
			"scoreMap":{
				"1":{
					"low":1,
					"high":100
					},
				"2":{"low":100,"high":300},
				"3":{"low":300,"high":500},
				"4":{"low":500,"high":700},
				"5":{"low":700}
			},
			"timeout":10000,
			"url":"https://storage.googleapis.com/100ms-speed-test-download/test1Mb.db"
		}
	}
}
*/
type StagesResp struct {
	Stages []Stage
}

type Stage struct {
	Discord string `json:"discord"`
	IsLive  bool   `json:"isLive"`
	Name    string `json:"name"`
	Room    string `json:"roomId"`
	Slug    string `json:"slug"`
	Stream  string `json:"stream"`

	Schedule []Schedule `json:"schedule"`
}

type Schedule struct {
	Title    string    `json:"title"`
	Start    string    `json:"start"`
	End      string    `json:"end"`
	Speakers []Speaker `json:"speaker"`
}

type Speaker struct {
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Image `json:"image"`
}

type Image struct {
	Blur string `json:"blurDataURL"`
	Url  string `json:"url"`
}
