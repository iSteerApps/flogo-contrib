package coap

//"encoding/json"
//"net/http"
//"testing"

//"net/http"
//"github.com/TIBCOSoftware/flogo-lib/core/trigger"

const testConfig string = `{
  "id": "tibco-coap",
  "ref": "github.com/TIBCOSoftware/flogo-contrib/trigger/coap",
  "settings": {
    "port": "5683"
  },
  "handlers": [
    {
      "actionId": "my_test_flow",
      "settings": {
        "method": "POST",
        "path": "/device/:id/reset"
      }
    }
  ]
}
`

/*
// TODO Fix this test

func TestHandlerOk(t *testing.T) {
	config := trigger.Config{}
	json.Unmarshal([]byte(testConfig), &config)

	// New  factory
	f := &CoapFactory{}
	tgr := f.New(&config)

	runner := &TestRunner{}

	tgr.Init(runner)

	tgr.Start()
	defer tgr.Stop()

	uri := "http://127.0.0.1:5683/device/12345/reset"

	req, err := http.NewRequest("POST", uri, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Debug("response Status:", resp.Status)

	if resp.StatusCode >= 300 {
		t.Fail()
	}
}
*/
