package sec

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httputil"
	"reflect"
	"strings"
	"time"
)

var (
	ErrFirstReqLine  = errors.New("first line of the request appears malformed")
	ErrInvalidHost   = errors.New("host header appears invalid")
	ErrMalformedJson = errors.New("json data appears malformed")
)

var httpScannerPayloads = []string{
	"'",
	"\\'",
	"\"",
	"\\\"",
	"`",
	"\\`",
}

const (
	injectionMarker = "<INJECTION>"
)

type ScannerRequest struct {
	Payload     string
	Status      int
	RawResponse string
}

func HttpScanner(request *string, tlsEnabled bool, timeout time.Duration, rpm int) (scannedRequests []ScannerRequest, err error) {
	slicedRequesst := strings.Split(*request, "\r\n\r\n")
	headers := slicedRequesst[0]
	body := strings.Join(slicedRequesst[1:], "\r\n\r\n")
	headerLines := strings.Split(headers, "\r\n")

	method := ""
	pathWithQuery := ""
	// httpVer := ""

	host := ""

	isJson := false

	for lineNo, line := range headerLines {
		lowercaseLine := strings.ToLower(line)

		if lineNo == 0 {
			sliced := strings.Split(line, " ")
			if len(sliced) != 3 {
				return nil, ErrFirstReqLine
			}
			method = sliced[0]
			pathWithQuery = sliced[1]
			// httpVer = sliced[2]
		}

		if len(host) == 0 {
			if strings.HasPrefix(lowercaseLine, "host: ") {
				sliced := strings.Split(lowercaseLine, ": ")
				if len(sliced) != 2 {
					return nil, ErrInvalidHost
				}
				host = sliced[1]
			}
		}

		if !isJson {
			if strings.HasPrefix(lowercaseLine, "content-type:") && strings.Contains(lowercaseLine, "application/json") {
				isJson = true
			}
		}
	}

	slicedPathWithQuery := strings.Split(pathWithQuery, "?")

	path := slicedPathWithQuery[0]

	rawParams := strings.Join(slicedPathWithQuery[1:], "?")
	params := strings.Split(rawParams, "&")
	if len(params) == 1 && params[0] == "" {
		params = []string{}
	}

	bodies := []map[string]any{}

	if isJson {
		var data any

		if err := json.Unmarshal([]byte(body), &data); err != nil {
			return nil, err
		}

		kindOfData := reflect.TypeOf(data).Kind()

		switch kindOfData {
		default:
			return nil, ErrMalformedJson

		case reflect.Slice:
			// TODO:

		case reflect.Map:
			// TODO:
			dataStruct := data.(map[string]any)
			bodies = injectIntoMap(dataStruct)
		}

	}

	c := http.Client{
		Timeout:   timeout,
		Transport: &http.Transport{},
	}

	for _, injectionPayload := range httpScannerPayloads {

		for _, body := range bodies {
			jsonBody, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}

			jsonBody = bytes.ReplaceAll(jsonBody, []byte(injectionMarker), []byte(injectionPayload))

			pathWithParams := path + "?" + rawParams

			targetURL := "://" + host + pathWithParams
			if tlsEnabled {
				targetURL = "https" + targetURL
			} else {
				targetURL = "http" + targetURL
			}

			req, err := http.NewRequest(method, targetURL, bytes.NewBuffer(jsonBody))
			if err != nil {
				return nil, err
			}

			resp, err := c.Do(req)
			if err != nil {
				return nil, err
			}

			rawResponse, err := httputil.DumpResponse(resp, true)
			if err != nil {
				return nil, err
			}

			scannedReq := ScannerRequest{
				Payload:     injectionPayload,
				Status:      resp.StatusCode,
				RawResponse: string(rawResponse),
			}

			scannedRequests = append(scannedRequests, scannedReq)
		}

		for _, paramWithValue := range params {
			pathWithParams := path + "?" + strings.ReplaceAll(rawParams, paramWithValue, paramWithValue+injectionPayload)

			targetURL := "://" + host + pathWithParams
			if tlsEnabled {
				targetURL = "https" + targetURL
			} else {
				targetURL = "http" + targetURL
			}

			req, err := http.NewRequest(method, targetURL, bytes.NewBuffer([]byte(body)))
			if err != nil {
				return nil, err
			}

			resp, err := c.Do(req)
			if err != nil {
				return nil, err
			}

			rawResponse, err := httputil.DumpResponse(resp, true)
			if err != nil {
				return nil, err
			}

			scannedReq := ScannerRequest{
				Payload:     injectionPayload,
				Status:      resp.StatusCode,
				RawResponse: string(rawResponse),
			}

			scannedRequests = append(scannedRequests, scannedReq)
		}
	}

	return scannedRequests, nil
}

func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}

func injectIntoMap(input map[string]any) []map[string]any {
	var result []map[string]any

	for key, value := range input {
		// Deep copy the original map to avoid modifying the original or prior results.
		modifiedMap := CopyMap(input)
		changed := false

		switch v := value.(type) {
		case string:
			// Append <INJECTION> to string values.
			modifiedMap[key] = v + injectionMarker
			changed = true
		case map[string]any:
			// Handle nested maps recursively.
			nestedResults := injectIntoMap(v)
			for _, nestedMap := range nestedResults {
				// Replace the nested map at the current key.
				modifiedCopy := CopyMap(input)
				modifiedCopy[key] = nestedMap
				result = append(result, modifiedCopy)
			}
			continue
		}

		// Only add to the result if a change was made.
		if changed {
			result = append(result, modifiedMap)
		}
	}

	return result
}
