package config

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

func (cfg Config) HTTPSignatureAuth(host string) runtime.ClientAuthInfoWriter {
	f := func(req runtime.ClientRequest, reg strfmt.Registry) error {

		if cfg.DebugLogging {
			log.Println("Performing HTTP Signature Auth...")
			log.Println("MerchantId: " + cfg.MerchantId)
			log.Println("KeyId: " + cfg.MerchantKeyId)
		}

		hdrs := req.GetHeaderParams()

		// The response is application/hal+json, and if our Accept
		// header is application/json, we get a 404.
		//
		// Ideally go-swagger would put application/hal+json on
		// the Accept header, but I'm not sure how to accomplish that.
		//
		hdrs.Del("Accept")

		hdrs["v-c-merchant-id"] = []string{cfg.MerchantId}
		hdrs["date"] = []string{time.Now().UTC().Format(time.RFC1123)}
		hdrs["Host"] = []string{host}

		skipDigest := strings.ToUpper(req.GetMethod()) == http.MethodGet

		body := req.GetBody()

		if cfg.DebugLogging {
			log.Printf("Body: `%s`", string(body))
		}

		if !skipDigest {
			digest, err := generateDigest(req.GetBody())
			if err != nil {
				return err
			}
			hdrs["digest"] = []string{digest}
		}

		signature, err := cfg.generateHTTPSignatureHeader(req, skipDigest)
		if err != nil {
			return err
		}
		hdrs["signature"] = []string{signature}

		return nil
	}
	return AuthenticateRequestFunc(f)
}

func generateDigest(body []byte) (string, error) {
	hasher := sha256.New()
	if _, err := hasher.Write(body); err != nil {
		return "", fmt.Errorf("unable to hash data: %v", err)
	}
	b64 := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return "SHA-256=" + b64, nil
}

func (cfg Config) generateHTTPSignatureHeader(req HTTPRequest, skipDigest bool) (string, error) {

	headerNames := []string{"host", "date", "request-target", "digest", "v-c-merchant-id"}
	if skipDigest {
		headerNames = append(headerNames[:3], headerNames[4:]...)
	}

	signatureValue, err := generateHTTPSignatureValue(cfg.MerchantSecretKey, req, headerNames...)
	if err != nil {
		return "", fmt.Errorf("unable to generate HTTP signature header: %v", err)
	}

	keyid := fmt.Sprintf(`keyid="%s"`, cfg.MerchantKeyId)
	algorithm := fmt.Sprintf(`algorithm="%s"`, "HmacSHA256")
	headers := fmt.Sprintf(`headers="%s"`, strings.Join(headerNames, " "))
	signature := fmt.Sprintf(`signature="%s"`, signatureValue)

	return strings.Join([]string{keyid, algorithm, headers, signature}, ", "), nil
}

func generateHTTPSignatureValue(merchantSecretKey string, req HTTPRequest, headerNames ...string) (string, error) {

	headers := req.GetHeaderParams()

	var keyValuePairs [][]string
	for _, header := range headerNames {

		key := header
		switch key {

		case "request-target":
			u := url.URL{Path: req.GetPath(), RawQuery: req.GetQueryParams().Encode()}
			requestTarget := fmt.Sprintf("%s %s", strings.ToLower(req.GetMethod()), u.String())
			keyValuePairs = append(keyValuePairs, []string{key, requestTarget})

		default:
			val := headers.Get(header)
			if val == "" {
				val = getHeaderDirect(headers, key)
				if val == "" {
					return "", fmt.Errorf("Unable to find header '%s' when creating signature value!", header)
				}
			}
			keyValuePairs = append(keyValuePairs, []string{key, val})
		}
	}

	var lines []string
	for _, kvp := range keyValuePairs {
		lines = append(lines, strings.Join(kvp, ": "))
	}

	toSign := strings.Join(lines, "\n")

	merchantDecoded, err := base64.StdEncoding.DecodeString(merchantSecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode the merchant key: %v", err)
	}

	hash := hmac.New(sha256.New, merchantDecoded)

	if _, err := hash.Write([]byte(toSign)); err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}

	b64 := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return strings.Replace(b64, "\n", "", -1), nil
}

func getHeaderDirect(hdrs http.Header, key string) (val string) {
	if vals, set := hdrs[key]; set {
		if len(vals) > 0 {
			val = vals[0]
		}
	}
	return
}
