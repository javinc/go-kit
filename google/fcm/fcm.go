package fcm

import (
	fcm "github.com/NaySoftware/go-fcm"
	"github.com/javinc/go-kit/config"
)

// Send fcm topic
func Send(topic string, data map[string]string) (result []map[string]string, err error) {
	c := fcm.NewFcmClient(config.GetString("fcm.server_key"))
	c.NewFcmMsgTo("/topics/"+topic, data)

	status, err := c.Send()
	if err != nil {
		return
	}

	result = status.Results
	return
}
