package khulnasoft

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/go-homedir"
)

// Config - godoc
type Config struct {
	Username string `json:"tenant"`
	Password string `json:"token"`
	KhulnasoftURL  string `json:"khulnasoft_url"`
}

// Provider -
func Provider(v string) *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("KHULNASOFT_USER", nil),
				Description: "This is the user id that should be used to make the connection. Can alternatively be sourced from the `KHULNASOFT_USER` environment variable.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("KHULNASOFT_PASSWORD", nil),
				Description: "This is the password that should be used to make the connection. Can alternatively be sourced from the `KHULNASOFT_PASSWORD` environment variable.",
			},
			"khulnasoft_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KHULNASOFT_URL", nil),
				Description: "This is the base URL of your Khulnasoft instance. Can alternatively be sourced from the `KHULNASOFT_URL` environment variable.",
			},
			"verify_tls": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KHULNASOFT_TLS_VERIFY", true),
				Description: "If true, server tls certificates will be verified by the client before making a connection. Defaults to true. Can alternatively be sourced from the `KHULNASOFT_TLS_VERIFY` environment variable.",
			},
			"ca_certificate_path": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KHULNASOFT_CA_CERT_PATH", nil),
				Description: "This is the file path for server CA certificates if they are not available on the host OS. Can alternatively be sourced from the `KHULNASOFT_CA_CERT_PATH` environment variable.",
			},
			"config_path": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KHULNASOFT_CONFIG", "~/.khulnasoft/tf.config"),
				Description: "This is the file path for Khulnasoft provider configuration. The default configuration path is `~/.khulnasoft/tf.config`. Can alternatively be sourced from the `KHULNASOFT_CONFIG` environment variable.",
			},
			"khulnasoft_api_key_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("KHULNASOFT_API_KEY_ID", nil),
				Description: "This is the API key ID that should be used to make the connection. Can alternatively be sourced from the `KHULNASOFT_API_KEY_ID` environment variable.",
			},
			"khulnasoft_api_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("KHULNASOFT_API_SECRET", nil),
				Description: "This is the API secret that should be used to make the connection. Can alternatively be sourced from the `KHULNASOFT_API_SECRET` environment variable.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"khulnasoft_user":                        resourceUser(),
			"khulnasoft_role":                        resourceRole(),
			"khulnasoft_integration_registry":        resourceRegistry(),
			"khulnasoft_firewall_policy":             resourceFirewallPolicy(),
			"khulnasoft_enforcer_groups":             resourceEnforcerGroup(),
			"khulnasoft_service":                     resourceService(),
			"khulnasoft_image":                       resourceImage(),
			"khulnasoft_notification_slack":          resourceNotificationOld(),
			"khulnasoft_container_runtime_policy":    resourceContainerRuntimePolicy(),
			"khulnasoft_function_runtime_policy":     resourceFunctionRuntimePolicy(),
			"khulnasoft_host_runtime_policy":         resourceHostRuntimePolicy(),
			"khulnasoft_host_assurance_policy":       resourceHostAssurancePolicy(),
			"khulnasoft_vmware_assurance_policy":     resourceVMwareAssurancePolicy(),
			"khulnasoft_image_assurance_policy":      resourceImageAssurancePolicy(),
			"khulnasoft_kubernetes_assurance_policy": resourceKubernetesAssurancePolicy(),
			"khulnasoft_function_assurance_policy":   resourceFunctionAssurancePolicy(),
			"khulnasoft_application_scope":           resourceApplicationScope(),
			"khulnasoft_permissions_sets":            resourcePermissionSet(),
			//"khulnasoft_sso":						 resourceSSO(),
			"khulnasoft_role_mapping": resourceRoleMapping(),
			"khulnasoft_khulnasoft_label":   resourceKhulnasoftLabels(),
			"khulnasoft_acknowledge":  resourceAcknowledge(),
			"khulnasoft_notification": resourceSourceNotification(),
			//saas
			"khulnasoft_group":             resourceGroup(),
			"khulnasoft_user_saas":         resourceUserSaas(),
			"khulnasoft_role_mapping_saas": resourceRoleMappingSaas(),
			"khulnasoft_permission_set_saas": resourcePermissionSetSaas(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"khulnasoft_users":                       dataSourceUsers(),
			"khulnasoft_roles":                       dataSourceRoles(),
			"khulnasoft_integration_registries":      dataSourceRegistry(),
			"khulnasoft_firewall_policy":             dataSourceFirewallPolicy(),
			"khulnasoft_enforcer_groups":             dataSourceEnforcerGroup(),
			"khulnasoft_service":                     dataSourceService(),
			"khulnasoft_image":                       dataImage(),
			"khulnasoft_container_runtime_policy":    dataContainerRuntimePolicy(),
			"khulnasoft_function_runtime_policy":     dataFunctionRuntimePolicy(),
			"khulnasoft_host_runtime_policy":         dataHostRuntimePolicy(),
			"khulnasoft_image_assurance_policy":      dataImageAssurancePolicy(),
			"khulnasoft_kubernetes_assurance_policy": dataKubernetesAssurancePolicy(),
			"khulnasoft_host_assurance_policy":       dataHostAssurancePolicy(),
			"khulnasoft_function_assurance_policy":   dataFunctionAssurancePolicy(),
			"khulnasoft_gateways":                    dataSourceGateways(),
			"khulnasoft_application_scope":           dataApplicationScope(),
			"khulnasoft_permissions_sets":            dataSourcePermissionsSets(),
			"khulnasoft_integration_state":           dataIntegrationState(),
			//"khulnasoft_sso":							 	dataSourceSSO(),
			"khulnasoft_roles_mapping": dataSourceRolesMapping(),
			"khulnasoft_khulnasoft_labels":   dataSourceKhulnasoftLabels(),
			"khulnasoft_acknowledges":  dataSourceAcknowledges(),
			"khulnasoft_notifications": dataSourceNotification(),
			//saas:
			"khulnasoft_groups":             dataSourceGroups(),
			"khulnasoft_users_saas":         dataSourceUsersSaas(),
			"khulnasoft_roles_mapping_saas": dataSourceRolesMappingSaas(),
			"khulnasoft_permissions_sets_saas": dataSourcePermissionsSetsSaas(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func getProviderConfigurationFromFile(d *schema.ResourceData) (string, string, string, error) {
	log.Print("[DEBUG] Trying to load configuration from file")
	if configPath, ok := d.GetOk("config_path"); ok && configPath.(string) != "" {
		path, err := homedir.Expand(configPath.(string))
		if err != nil {
			log.Printf("[DEBUG] Failed to expand config file path %s, error %s", configPath, err)
			return "", "", "", nil
		}
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Printf("[DEBUG] Terraform config file %s does not exist, error %s", path, err)
			return "", "", "", nil
		}
		log.Printf("[DEBUG] Terraform configuration file is: %s", path)
		configFile, err := os.Open(path)
		if err != nil {
			log.Printf("[DEBUG] Unable to open Terraform configuration file %s", path)
			return "", "", "", fmt.Errorf("Unable to open terraform configuration file. Error %v", err)
		}
		defer configFile.Close()

		configBytes, _ := io.ReadAll(configFile)
		var config Config
		err = json.Unmarshal(configBytes, &config)
		if err != nil {
			log.Printf("[DEBUG] Failed to parse config file %s", path)
			return "", "", "", fmt.Errorf("Invalid terraform configuration file format. Error %v", err)
		}
		return config.Username, config.Password, config.KhulnasoftURL, nil
	}
	return "", "", "", nil
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	fmt.Println("----------------------------------------")
	var diags diag.Diagnostics
	var err error

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	khulnasoftURL := d.Get("khulnasoft_url").(string)
	verifyTLS := d.Get("verify_tls").(bool)
	caCertPath := d.Get("ca_certificate_path").(string)
	apiKeyID := d.Get("khulnasoft_api_key_id").(string)
	apiSecret := d.Get("khulnasoft_api_secret").(string)

	// Check if using API key authentication
	if apiKeyID != "" && apiSecret != "" {
		// Validate that username/password are not also provided
		if username != "" || password != "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Cannot use both username/password and API key authentication",
				Detail:   "Please provide either username/password OR khulnasoft_api_key_id/khulnasoft_api_secret, not both.",
			})
			return nil, diags
		}
	} else if username == "" && password == "" && khulnasoftURL == "" && apiKeyID == "" && apiSecret == "" {
		// Try to load from config file if all parameters are empty
		username, password, khulnasoftURL, err = getProviderConfigurationFromFile(d)
		if err != nil {
			return nil, diag.FromErr(err)
		}
	}

	// Validate required parameters for username/password auth
	if (username != "" || password != "") && khulnasoftURL == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Initializing provider, khulnasoft_url parameter is missing when using username/password authentication.",
		})
	}

	// Validate required parameters for API key auth
	if (apiKeyID != "" || apiSecret != "") && khulnasoftURL == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Initializing provider, khulnasoft_url parameter is missing when using API key authentication.",
		})
	}

	// Ensure both API key fields are provided together
	if (apiKeyID != "" && apiSecret == "") || (apiKeyID == "" && apiSecret != "") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Both khulnasoft_api_key_id and khulnasoft_api_secret must be provided together.",
		})
	}

	// Validate that we have some form of authentication
	if username == "" && password == "" && apiKeyID == "" && apiSecret == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "No authentication method provided. Please provide either username/password or khulnasoft_api_key_id/khulnasoft_api_secret.",
		})
	}

	if khulnasoftURL == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Initializing provider, khulnasoft_url parameter is missing.",
		})
	}

	var caCertByte []byte
	if caCertPath != "" {
		caCertByte, err = os.ReadFile(caCertPath)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to read CA certificates",
				Detail:   err.Error(),
			})

			return nil, diags
		}
	}

	if diags != nil && len(diags) > 0 {
		return nil, diags
	}

	var khulnasoftClient *client.Client

	// Create client based on authentication method
	if apiKeyID != "" && apiSecret != "" {
		// Use API key authentication
		khulnasoftClient = client.NewClientWithAPIKey(khulnasoftURL, apiKeyID, apiSecret, verifyTLS, caCertByte)
	} else {
		// Use username/password authentication
		khulnasoftClient = client.NewClientWithTokenAuth(khulnasoftURL, username, password, verifyTLS, caCertByte)
	}

	token, tokenPresent := os.LookupEnv("TESTING_AUTH_TOKEN")
	url, urlPresent := os.LookupEnv("TESTING_URL")

	if !tokenPresent || !urlPresent {
		_, _, err = khulnasoftClient.GetAuthToken()

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to fetch token",
				Detail:   err.Error(),
			})

			return nil, diags
		}
	} else {
		khulnasoftClient.SetAuthToken(token)
		khulnasoftClient.SetUrl(url)
	}

	return khulnasoftClient, diags
}