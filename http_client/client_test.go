package http_client

import (
	"github.com/sirupsen/logrus"
	"testing"
)

var testClient = NewHttpClient(logrus.NewEntry(logrus.New()), nil, nil)

type testJson struct {
	Name string `json:"name"`
	Sort bool   `json:"sort"`
}

type rsJson struct {
	Afk   bool   `json:"afk"`
	Hello string `json:"hello"`
}

func TestHttpClient_GetJson(t *testing.T) {
	send := new(testJson)
	send.Name = "hehe"
	send.Sort = true
	rs := new(rsJson)
	t.Log(testClient.GetJson("http://127.0.0.1/v1/testjson", nil, rs))
	t.Log(rs)
}
