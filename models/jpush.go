package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	jpushclient "github.com/lucasluowy/jpush-api-go-client"
)

const (
	apnsProduction = true
)

/**
 * registrationIds: []string{"13165ffa4e7cd3fdf21", "190e35f7e03c05a1bb6", "170976fa8afa405cddf"}
 * alert: test[签名]
 * url: mcgj://notifyenter/indexenter?code=tab_my
 * icon: https://pic1.yilu.cn/202008/d90af286632f081f43b71d224aa205f6.jpg
 */
func Push(registrationIds []string, alert string, title string, url string, icon string) (msg string) {
	//Platform
	var pf jpushclient.Platform
	pf.Add(jpushclient.ANDROID)
	pf.Add(jpushclient.IOS)
	//pf.All()

	//Audience
	var ad jpushclient.Audience
	ad.SetID(registrationIds)

	//Notice
	var notice jpushclient.Notice
	notice.SetAlert(alert)

	extras := map[string]interface{}{
		"title":   title,
		"mcgjUrl": url,
		//"my-attachment": icon,
		"type": "other",
		"link": "1000",
	}

	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: alert, Title: title, BuilderId: 3, LargeIcon: icon, Extras: extras})
	iosAlert := map[string]interface{}{
		"title": title,
		"body":  alert,
	}
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: iosAlert, Badge: "1", Sound: "ping1.caf", MutableContent: true, Extras: extras})

	var op jpushclient.Option
	op.SetApns(apnsProduction)

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetNotice(&notice)
	payload.SetOptions(&op)

	bytes, _ := payload.ToBytes()
	logs.Info("Jpush Info:" + string(bytes))

	//push
	appKey := beego.AppConfig.String("jpush_app_key")
	secret := beego.AppConfig.String("jpush_secret")
	pc := jpushclient.NewPushClient(secret, appKey)
	str, err := pc.Send(bytes)
	msg = str
	logs.Info("Jpush Send:" + str)

	if err != nil {
		logs.Info("Jpush Err:" + err.Error())
	}
	return
}
