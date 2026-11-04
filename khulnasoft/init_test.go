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
	if os.Getenv("TF_ACC") == "" {
		log.Println("skipping acceptance test setup because TF_ACC is not set")
		return
	}

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
	if !present || khulnasoftURL == "" {
		panic("KHULNASOFT_URL env is missing or empty, please set it")
	}

	apiKey = os.Getenv("KHULNASOFT_API_KEY")
	secretKey = os.Getenv("KHULNASOFT_API_SECRET")
	useAPIKeyStr = os.Getenv("KHULNASOFT_USE_API_KEY")
	useAPIKey = false

	// Check if KHULNASOFT_USE_API_KEY is explicitly set
	if useAPIKeyStr != "" {
		var err error
		useAPIKey, err = strconv.ParseBool(useAPIKeyStr)
		if err != nil {
			panic(fmt.Sprintf("Invalid boolen for KHULNASOFT_USE_API_KEY: %v", err))
		}
	} else {
		// Auto-detect: use API key auth if API key credentials are provided
		if apiKey != "" && secretKey != "" {
			useAPIKey = true
		}
	}

	// Get credentials based on auth method
	if useAPIKey {
		if apiKey == "" {
			panic("KHULNASOFT_API_KEY env is missing or empty, please set it when using API key authentication")
		}
		if secretKey == "" {
			panic("KHULNASOFT_API_SECRET env is missing or empty, please set it when using API key authentication")
		}
	} else {
		username, present = os.LookupEnv("KHULNASOFT_USER")
		if !present || username == "" {
			panic("KHULNASOFT_USER env is missing or empty, please set it (or use KHULNASOFT_API_KEY and KHULNASOFT_API_SECRET for API key auth)")
		}

		password, present = os.LookupEnv("KHULNASOFT_PASSWORD")
		if !present || password == "" {
			panic("KHULNASOFT_PASSWORD env is missing or empty, please set it (or use KHULNASOFT_API_KEY and KHULNASOFT_API_SECRET for API key auth)")
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
		khulnasoftClient, err = client.NewClientWithAPIKey(khulnasoftURL, apiKey, secretKey, verifyTLS, caCertByte)
		if err != nil {
			panic(fmt.Errorf("failed to create client with api key auth, error: %s", err))
		}
	} else {
		khulnasoftClient, err = client.NewClientWithTokenAuth(khulnasoftURL, username, password, verifyTLS, caCertByte)
		if err != nil {
			panic(fmt.Errorf("failed to create client with token auth, error: %s", err))
		}
	}
	token, url, err := khulnasoftClient.GetAuthToken()

	if err != nil {
		panic(fmt.Errorf("failed to receive token, error: %s", err))
	}

	err = os.Setenv("TESTING_AUTH_TOKEN", token)
	if err != nil {
		panic("Failed to set AUTH_TOKEN env")
	}

	err = os.Setenv("TESTING_URL", url)
	if err != nil {
		panic("Failed to set TESTING_URL env")
	}
	log.Println("Finished to set token")

}
