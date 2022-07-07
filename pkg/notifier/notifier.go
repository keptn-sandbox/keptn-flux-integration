package notifier

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/keptn-sandbox/keptn-flux-integration/pkg/provider"
)

func PostMessage(event provider.KeptnEvent) error {
	httpClient := retryablehttp.NewClient()

	httpClient.HTTPClient.Timeout = 15 * time.Second
	httpClient.RetryWaitMin = 2 * time.Second
	httpClient.RetryWaitMax = 30 * time.Second
	httpClient.RetryMax = 4
	httpClient.Logger = nil

	// data, err := json.Marshal(event.Body)
	// fmt.Println("*****")
	// fmt.Println(event.Body)
	// fmt.Println("*****")
	// if err != nil {
	// 	return fmt.Errorf("marshalling notification payload failed: %w", err)
	// }

	req, err := retryablehttp.NewRequest(http.MethodPost, event.Address, event.Body)
	if err != nil {
		return fmt.Errorf("failed to create a new request: %w", err)
	}

	for k, v := range event.Headers {
		headerKey := fmt.Sprintf("%s", k)
		headerValue := fmt.Sprintf("%s", v)
		req.Header.Set(headerKey, headerValue)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusCreated {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("unable to read response body, %s", err)
		}
		return fmt.Errorf("request failed with status code %d, %s", resp.StatusCode, string(b))
	}

	return nil
}
