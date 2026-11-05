package khulnasoft

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
)

func init() {
	log.Println("setup suite")
	var (
		present                                          bool
		username, password, khulnasoftURL                      string
		verifyTLS, useAPIKey                             bool
		verifyTLSString, apiKey, secretKey, useAPIKeyStr string
		caCertPath                                       string
		err                                              error
		caCertByte                                       []byte
	)

	khulnasoftURL, present = os.LookupEnv("KHULNASOFT_URL")
	if !present {
		panic("KHULNASOFT_URL env is missing, please set it")
	}

	apiKey = os.Getenv("KHULNASOFT_API_KEY_ID")
	secretKey = os.Getenv("KHULNASOFT_API_SECRET")
	useAPIKeyStr = os.Getenv("KHULNASOFT_USE_API_KEY")
	useAPIKey = false
	if useAPIKeyStr != "" {
		var err error
		useAPIKey, err = strconv.ParseBool(useAPIKeyStr)
		if err != nil {
			panic(fmt.Sprintf("Invalid boolean for KHULNASOFT_USE_API_KEY: %v", err))
		}
	}

	if !useAPIKey {
		username, present = os.LookupEnv("KHULNASOFT_USER")
		if !present {
			panic("KHULNASOFT_USER env is missing, please set it")
		}

		password, present = os.LookupEnv("KHULNASOFT_PASSWORD")
		if !present {
			panic("KHULNASOFT_PASSWORD env is missing, please set it")
		}
	}

	verifyTLSString, present = os.LookupEnv("KHULNASOFT_TLS_VERIFY")
	if !present {
		verifyTLSString = "true"
	}
	verifyTLS, _ = strconv.ParseBool(verifyTLSString)

	caCertPath, present = os.LookupEnv("KHULNASOFT_CA_CERT_PATH")
	if present {
		if caCertPath != "" {
			caCertByte, err = os.ReadFile(caCertPath)
			if err != nil {
				panic("Unable to read CA certificates")
			}
		}
	}

	var khulnasoftClient *client.Client
	if useAPIKey {
		khulnasoftClient = client.NewClientWithAPIKey(khulnasoftURL, apiKey, secretKey, verifyTLS, caCertByte)
	} else {
		khulnasoftClient = client.NewClientWithTokenAuth(khulnasoftURL, username, password, verifyTLS, caCertByte)
	}
	token, url, err := khulnasoftClient.GetAuthToken()

	if err != nil {
		panic(fmt.Errorf("failed to receive token, error: %s", err))
	}

	err = os.Setenv("TESTING_AUTH_TOKEN", token)
	if err != nil {
		panic("Failed to set TESTING_AUTH_TOKEN env")
	}

	err = os.Setenv("TESTING_URL", url)
	if err != nil {
		panic("Failed to set TESTING_URL env")
	}
	log.Println("Finished to set token")

}