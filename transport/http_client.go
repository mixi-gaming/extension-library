package transport

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"time"
)

type (
	httpElement struct {
		method       string
		url          string
		headerKeys   []string
		headerValues []string
		body         *bytes.Buffer
	}
)

type stop struct {
	error
}

func retry(attempts int, sleep time.Duration, elements httpElement, f func(httpElement) ([]byte, error)) ([]byte, error) {
	resp, err := f(elements)
	if err != nil {
		if s, ok := err.(stop); ok {
			fmt.Println(s.error)
			// Return the original error for later checking
			return nil, err
		}

		if attempts--; attempts >= 0 {
			// Add some randomness to prevent creating a Thundering Herd
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			time.Sleep(sleep)
			return retry(attempts, 2*sleep, elements, f)
		}

		if err != nil {
			fmt.Println(err)
		}
		return nil, err
	}

	return resp, nil
}

func makeRequest(element httpElement) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Transport: tr,
		Timeout:   Timeout,
	}

	request, err := http.NewRequest(element.method, element.url, element.body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for i := 0; i < len(element.headerKeys); i++ {
		if element.headerKeys[i] != "" && element.headerValues[i] != "" {
			request.Header.Set(element.headerKeys[i], element.headerValues[i])
		}
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return bodyBytes, nil
}

// MakeHTTPGetRequest - MakeHTTPGetRequest
func MakeHTTPGetRequest(ctx context.Context, URL string, headerKeys, headerValues []string, bodyRequest []byte) ([]byte, error) {
	headerKeys = append(headerKeys, "Content-Type")
	headerValues = append(headerValues, "application/json")
	return retry(5, 1, httpElement{"GET", URL, headerKeys, headerValues, bytes.NewBuffer(bodyRequest)}, makeRequest)
}

// MakeHTTPPostRequest - MakeHTTPPostRequest
func MakeHTTPPostRequest(ctx context.Context, URL string, headerKeys, headerValues []string, bodyRequest []byte) ([]byte, error) {
	headerKeys = append(headerKeys, "Content-Type")
	headerValues = append(headerValues, "application/json")
	return retry(5, 1, httpElement{"POST", URL, headerKeys, headerValues, bytes.NewBuffer(bodyRequest)}, makeRequest)
}

// MakeHTTPPutRequest - MakeHTTPPutRequest
func MakeHTTPPutRequest(ctx context.Context, URL string, headerKeys, headerValues []string, bodyRequest []byte) ([]byte, error) {
	headerKeys = append(headerKeys, "Content-Type")
	headerValues = append(headerValues, "application/json")
	return retry(5, 1, httpElement{"PUT", URL, headerKeys, headerValues, bytes.NewBuffer(bodyRequest)}, makeRequest)
}

// MakeHTTPPatchRequest - MakeHTTPPatchRequest
func MakeHTTPPatchRequest(ctx context.Context, URL string, headerKeys, headerValues []string, bodyRequest []byte) ([]byte, error) {
	headerKeys = append(headerKeys, "Content-Type")
	headerValues = append(headerValues, "application/json")
	return retry(5, 1, httpElement{"PATCH", URL, headerKeys, headerValues, bytes.NewBuffer(bodyRequest)}, makeRequest)
}

// MakeHTTPDeleteRequest - MakeHTTPDeleteRequest
func MakeHTTPDeleteRequest(ctx context.Context, URL string, headerKeys, headerValues []string) ([]byte, error) {
	headerKeys = append(headerKeys, "Content-Type")
	headerValues = append(headerValues, "application/json")
	httpElmItem := httpElement{"DELETE", URL, headerKeys, headerValues, nil}
	fmt.Println("DEBUG httpElmItem:", httpElmItem)
	return retry(5, 1, httpElmItem, makeRequest)
}

// MakeHTTPPostFormRequest - MakeHTTPPostFormRequest
func MakeHTTPPostFormRequest(ctx context.Context, URL string, headerKeys, headerValues []string, multiPartWriter *multipart.Writer, body bytes.Buffer) ([]byte, error) {
	headerKeys = append(headerKeys, "Content-Type")
	headerValues = append(headerValues, multiPartWriter.FormDataContentType())
	return retry(5, 1, httpElement{"POST", URL, headerKeys, headerValues, &body}, makeRequest)
}
