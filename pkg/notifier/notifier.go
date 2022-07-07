package notifier

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/keptn-sandbox/keptn-flux-integration/pkg/provider"
)

func PostMessage(event provider.KeptnEvent) error {
	var jsonStr = []byte(event.Body)
	req, err := http.NewRequest("POST", event.Address, bytes.NewBuffer(jsonStr))
	fmt.Println(string(jsonStr))
	for k, v := range event.Headers {
		headerKey := fmt.Sprintf("%s", k)
		headerValue := fmt.Sprintf("%s", v)
		req.Header.Set(headerKey, headerValue)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}
