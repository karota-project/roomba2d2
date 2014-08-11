package main

import (
	"./roomba2d2"
	"encoding/json"
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/karota-project/gobot-docomo"
	"github.com/karota-project/gobot-docomo/dialogue"
	"github.com/karota-project/gobot-facebook"
	"github.com/karota-project/gobot-julius"
	"github.com/karota-project/gobot-openjtalk"
	"github.com/karota-project/gobot-roomba"
	"github.com/karota-project/gobot-twitter"
	"io/ioutil"
	"net/url"
)

func main() {
	//config
	config := roomba2d2.Config{}

	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(f, &config)
	if err != nil {
		panic(err)
	}

	//gobot
	master := gobot.NewGobot()
	api.NewAPI(master).Start()

	//docomo
	docomoAdaptor := docomo.NewDocomoAdaptor("docomo-a01", config.Docomo.ApiKey)
	dialogueDriver := dialogue.NewDialogueDriver(docomoAdaptor, "dialogue-d01")

	//facebook
	facebookAdaptor := facebook.NewFacebookAdaptor("facebook-a01")
	facebookDriver := facebook.NewFacebookDriver(facebookAdaptor, "facebook-d01")

	//julius
	juliusAdaptor := julius.NewJuliusAdaptor("julius-a01", config.Julius.Port)
	juliusDriver := julius.NewJuliusDriver(juliusAdaptor, "julius-d01")

	//openjtalk
	openjtalkAdaptor := openjtalk.NewOpenjtalkAdaptor("openjtalk-a01")
	openjtalkDriver := openjtalk.NewOpenjtalkDriver(openjtalkAdaptor, "openjtalk-d01")

	//roomba
	roombaAdaptor := roomba.NewRoombaAdaptor("roomba-a01", config.Roomba.Port)
	roombaDriver := roomba.NewRoombaDriver(roombaAdaptor, "roomba-d01")

	//twitter
	twitterDriver := twitter.NewTwitterDriver("twitter-d01", config.Twitter.ConsumerKey, config.Twitter.ConsumerSecret)
	twitterDriver.SetAccessToken(config.Twitter.AccessToken, config.Twitter.AccessTokenSecret)

	master.AddRobot(
		gobot.NewRobot(
			"roomba2d2",
			[]gobot.Connection{docomoAdaptor, facebookAdaptor, juliusAdaptor, openjtalkAdaptor, roombaAdaptor},
			[]gobot.Device{dialogueDriver, facebookDriver, juliusDriver, openjtalkDriver, roombaDriver, twitterDriver},
			func() {
				fmt.Println("work")

				// 雑談対話
				dialogueResult, err := dialogueDriver.Get(dialogue.RequestBody{
					Utt:            "こちらルンバです",
					Context:        "53e816d98b3b3",
					Nickname:       "光",
					NicknameY:      "ヒカリ",
					Sex:            "女",
					Bloodtype:      "A",
					BirthdateY:     "1997",
					BirthdateM:     "5",
					BirthdateD:     "30",
					Age:            "16",
					Constellations: "双子座",
					Place:          "東京",
					Mode:           "dialog",
				})

				if err == nil {
					fmt.Println(dialogueResult)
				} else {
					fmt.Println(err)
				}

				//twitter
				value := url.Values{}
				tweets, err := twitterDriver.TwitterApi().GetSearch("roomba", value)

				if err == nil {
					for _, tweet := range tweets {
						fmt.Println(tweet.User.Name + "    @" + tweet.User.ScreenName)
						fmt.Println(tweet.Text)
						fmt.Println()
					}
				}

				//julius
				events := []string{
					//julius.START_PROC,
					//julius.END_PROC,
					//julius.START_RECOG,
					//julius.END_RECOG,
					//julius.INPUT,
					//julius.INPUT_PARAM,
					//julius.GMM,
					julius.RECOG_OUT,
					//julius.RECOG_FAIL,
					//julius.REJECTED,
					//julius.GRAPH_OUT,
					//julius.GRAM_INFO,
					//julius.SYS_INFO,
					//julius.ENGINE_INFO,
					//julius.GRAMMER,
					//julius.RECOG_PROCESS,
				}

				for _, event := range events {
					gobot.On(juliusDriver.Event(event), func(data interface{}) {
						fmt.Println("-----")
						for _, whypo := range data.(julius.RecogOut).Shypo.Whypo {
							fmt.Println(whypo.Word)
						}
						fmt.Println("-----")
					})
				}
			}))

	master.Start()
}
