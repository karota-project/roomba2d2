package roomba2d2

type Docomo struct {
	ApiKey string `json:"apikey"`
}

type Facebook struct {
}

type Julius struct {
	Port string `json:"port"`
}

type Openjtalk struct {
}

type Roomba struct {
	Port string `json:"port"`
}

type Roomba2d2 struct {
}

type Twitter struct {
	ConsumerKey       string `json:"consumerkey"`
	ConsumerSecret    string `json:"consumersecret"`
	AccessToken       string `json:"accesstoken"`
	AccessTokenSecret string `json:"accesstokensecret"`
}

type Config struct {
	Docomo    Docomo    `json:"docomo"`
	Facebook  Facebook  `json:"facebook"`
	Julius    Julius    `json:"julius"`
	Openjtalk Openjtalk `json:"openjtalk"`
	Roomba    Roomba    `json:"roomba"`
	Roomba2d2 Roomba2d2 `json:"roomba2d2"`
	Twitter   Twitter   `json:"twitter"`
}
