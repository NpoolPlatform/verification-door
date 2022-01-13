package googlerecaptcha

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	npool "github.com/NpoolPlatform/message/npool/verification"
	"golang.org/x/xerrors"
)

const (
	KeyRecaptchaURL    = "recaptcha_url"
	KeyRecaptchaSecret = "recaptcha_secret"
)

type SiteVerifyResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func VerifyGoogleRecaptcha(in *npool.VerifyGoogleRecaptchaRequest) (*npool.VerifyGoogleRecaptchaResponse, error) {
	hostname := config.GetStringValueWithNameSpace("", config.KeyHostname)
	recaptchaURL := config.GetStringValueWithNameSpace(hostname, KeyRecaptchaURL)
	recaptchaSecret := config.GetStringValueWithNameSpace(hostname, KeyRecaptchaSecret)
	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}
	request := url.Values{"secret": {recaptchaSecret}, "response": {in.GetResponse()}}
	resp, err := httpClient.PostForm(recaptchaURL, request)
	if err != nil {
		return nil, xerrors.Errorf("fail to verify: %v", err)
	}

	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, xerrors.Errorf("read response body error: %v", err)
	}

	result := SiteVerifyResponse{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, xerrors.Errorf("fail to unmarshal response: %v", err)
	}

	if !result.Success {
		return nil, xerrors.Errorf("fail to check google recaptcha: %v", result.ErrorCodes)
	}

	return &npool.VerifyGoogleRecaptchaResponse{
		Info: true,
	}, nil
}
