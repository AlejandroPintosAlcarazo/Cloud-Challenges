package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// makeAEMETRequest constructs and sends an HTTP GET request to the AEMET API
func getInitialDataURL(endpoint string) (*http.Response, error) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Add headers to the request
	req.Header.Add("Accept", "application/json")

	// Create an HTTP client and make the request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Check if the response status code is OK
	if resp.StatusCode == http.StatusOK {
		return resp, nil
	}

	// Check if the response status code is 429
	if resp.StatusCode == http.StatusTooManyRequests {
		return handleTooManyRequests(endpoint, resp)
	}

	// If response code is neither 200 nor 429
	return nil, fmt.Errorf("error: received non-200 response code %d", resp.StatusCode)
}

// requestData makes a second request to get the actual data
func getActualData(dataURL string) ([]byte, error) {
	// Make a second request to get the actual data
	dataResp, err := http.Get(dataURL)
	if err != nil {
		return nil, err
	}
	defer dataResp.Body.Close()

	// Check if the response status code is OK
	if dataResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received non-200 response code %d for data URL", dataResp.StatusCode)
	}

	// Read the response body
	dataBody, err := io.ReadAll(dataResp.Body)
	if err != nil {
		return nil, err
	}

	// Print the response body for debugging
	//fmt.Println("Data Response Body:", string(dataBody))

	return dataBody, nil
}

// handleTooManyRequests handles the 429 Too Many Requests response code
func handleTooManyRequests(endpoint string, resp *http.Response) (*http.Response, error) {
	fmt.Println("Received 429 Too Many Requests response")
	retryAfter := resp.Header.Get("Retry-After")
	if retryAfter != "" {
		fmt.Printf("Retry-After header value: %s\n", retryAfter)
		retryAfterSeconds, err := strconv.Atoi(retryAfter)
		if err != nil {
			return nil, fmt.Errorf("error parsing Retry-After header: %v", err)
		}
		fmt.Printf("Received 429. Retrying after %d seconds...\n", retryAfterSeconds)
		time.Sleep(time.Duration(retryAfterSeconds) * time.Second)
	} else {
		fmt.Println("Received 429 without Retry-After header. Retrying after 60 seconds...")
		time.Sleep(60 * time.Second)
	}
	return getInitialDataURL(endpoint)
}
