package khulnasoft

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/go-homedir"
)

// Config - godoc
type Config struct {
	Username         string   `json:"tenant"`
	Password         string   `json:"token"`
	KhulnasoftURL          string   `json:"khulnasoft_url"`
	APIKey           string   `json:"khulnasoft_api_key"`
	SecretKey        string   `json:"khulnasoft_api_secret"`
	Validity         int      `json:"validity"`
	AllowedEndpoints []string `json:"allowed_endpoints"`
	CSPRoles         []string `json:"csp_roles"`
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
			"khulnasoft_api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("KHULNASOFT_API_KEY", nil),
				Description: "API key for authentication. If set, API key mode is used instead of token-based auth.",
			},
			"khulnasoft_api_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("KHULNASOFT_API_SECRET", nil),
				Description: "Shared secret for API key HMAC signing.",
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
			"validate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Skip provider credential validation when set to false.",
			},
			"validity": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     240,
				Description: "Lifetime of the token, in minutes. Set between 1 and 1500. Once the token expires, need to generate a new one",
			},
			"allowed_endpoints": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "API methods the token has access to",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"csp_roles": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"khulnasoft_application_scope_saas":      resourceApplicationScopeSaas(),
			"khulnasoft_permissions_sets":            resourcePermissionSet(),
			//"khulnasoft_sso":						 resourceSSO(),
			"khulnasoft_role_mapping": resourceRoleMapping(),
			"khulnasoft_khulnasoft_label":   resourceKhulnasoftLabels(),
			"khulnasoft_acknowledge":  resourceAcknowledge(),
			"khulnasoft_notification": resourceSourceNotification(),
			//saas
			"khulnasoft_group":                   resourceGroup(),
			"khulnasoft_user_saas":               resourceUserSaas(),
			"khulnasoft_role_mapping_saas":       resourceRoleMappingSaas(),
			"khulnasoft_permission_set_saas":     resourcePermissionSetSaas(),
			"khulnasoft_assurance_custom_script": resourceAssuranceScript(),
			"khulnasoft_khulnasoft_api_key":            resourceAPIKey(),
			"khulnasoft_scanner_group":           resourceScannerGroup(),
			"khulnasoft_log_management":          resourceLogManagement(),
			"khulnasoft_serverless_application":  resourceServerlessApplication(),
			"khulnasoft_monitoring_system":       resourceMonitoringSystem(),
			"khulnasoft_suppression_rule":        resourceSuppressionRule(),
			"khulnasoft_response_policy":         resourceResponsePolicy(),
			"khulnasoft_secure_ai_integration":   resourceSecureAIIntegration(),
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
			"khulnasoft_application_scope_saas":      dataApplicationScopeSaas(),
			"khulnasoft_permissions_sets":            dataSourcePermissionsSets(),
			"khulnasoft_integration_state":           dataIntegrationState(),
			//"khulnasoft_sso":							 	dataSourceSSO(),
			"khulnasoft_roles_mapping": dataSourceRolesMapping(),
			"khulnasoft_khulnasoft_labels":   dataSourceKhulnasoftLabels(),
			"khulnasoft_acknowledges":  dataSourceAcknowledges(),
			"khulnasoft_notifications": dataSourceNotification(),
			//saas:
			"khulnasoft_groups":                  dataSourceGroups(),
			"khulnasoft_users_saas":              dataSourceUsersSaas(),
			"khulnasoft_roles_mapping_saas":      dataSourceRolesMappingSaas(),
			"khulnasoft_permissions_sets_saas":   dataSourcePermissionsSetsSaas(),
			"khulnasoft_assurance_custom_script": dataSourceAssuranceScript(),
			"khulnasoft_khulnasoft_api_keys":           dataSourceAPIKeys(),
			"khulnasoft_scanner_group":           dataSourceScannerGroup(),
			"khulnasoft_vmware_assurance_policy": dataVmwareAssurancePolicy(),
			"khulnasoft_log_managements":         dataLogManagement(),
			"khulnasoft_serverless_applications": dataSourceServerlessApplication(),
			"khulnasoft_monitoring_systems":      dataSourceMonitoringSystem(),
			"khulnasoft_suppression_rules":       dataSourceSuppressionRule(),
			"khulnasoft_response_policies":       dataResponsePolicy(),
			"khulnasoft_response_policy_config":  dataResponsePolicyConfig(),
			"khulnasoft_secure_ai_integration":   dataSecureAIIntegration(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func getProviderConfigurationFromFile(d *schema.ResourceData) (string, string, string, string, string, error) {
	log.Print("[DEBUG] Trying to load configuration from file")
	if configPath, ok := d.GetOk("config_path"); ok && configPath.(string) != "" {
		path, err := homedir.Expand(configPath.(string))
		if err != nil {
			log.Printf("[DEBUG] Failed to expand config file path %s, error %s", configPath, err)
			return "", "", "", "", "", nil
		}
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Printf("[DEBUG] Terraform config file %s does not exist, error %s", path, err)
			return "", "", "", "", "", nil
		}
		log.Printf("[DEBUG] Terraform configuration file is: %s", path)
		configFile, err := os.Open(path)
		if err != nil {
			log.Printf("[DEBUG] Unable to open Terraform configuration file %s", path)
			return "", "", "", "", "", fmt.Errorf("Unable to open terraform configuration file. Error %v", err)
		}
		defer configFile.Close()

		configBytes, _ := io.ReadAll(configFile)
		var config Config
		err = json.Unmarshal(configBytes, &config)
		if err != nil {
			log.Printf("[DEBUG] Failed to parse config file %s", path)
			return "", "", "", "", "", fmt.Errorf("Invalid terraform configuration file format. Error %v", err)
		}
		return config.Username, config.Password, config.KhulnasoftURL, config.APIKey, config.SecretKey, nil
	}
	return "", "", "", "", "", nil
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	fmt.Println("----------------------------------------")
	var diags diag.Diagnostics
	var err error

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	khulnasoftURL := d.Get("khulnasoft_url").(string)
	apiKey := d.Get("khulnasoft_api_key").(string)
	secretkey := d.Get("khulnasoft_api_secret").(string)
	verifyTLS := d.Get("verify_tls").(bool)
	caCertPath := d.Get("ca_certificate_path").(string)
	validate := d.Get("validate").(bool)

	if username == "" || password == "" || apiKey == "" || secretkey == "" {
		uF, pF, fileURL, akF, skF, ferr := getProviderConfigurationFromFile(d)
		if ferr != nil && validate {
			return nil, diag.FromErr(ferr)
		}
		if username == "" {
			username = uF
		}
		if password == "" {
			password = pF
		}
		if khulnasoftURL == "" && fileURL != "" {
			khulnasoftURL = fileURL
		}
		if apiKey == "" {
			apiKey = akF
		}
		if secretkey == "" {
			secretkey = skF
		}
	}

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	apiKey = strings.TrimSpace(apiKey)
	secretkey = strings.TrimSpace(secretkey)
	khulnasoftURL = strings.TrimSpace(khulnasoftURL)

	if validate {
		if khulnasoftURL == "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Initializing provider, khulnasoft_url parameter is missing.",
			})
		}
		apiPairOK := apiKey != "" && secretkey != ""
		upOK := username != "" && password != ""
		if !apiPairOK && !upOK {
			if username != "" && password == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Initializing provider, password parameter is missing.",
				})
			}

			if password != "" && username == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Initializing provider, username parameter is missing.",
				})
			}

			if apiKey != "" && secretkey == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Initializing provider, khulnasoft_api_secret parameter is missing.",
				})
			}

			if secretkey != "" && apiKey == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Initializing provider, khulnasoft_api_key parameter is missing.",
				})
			}

			if username == "" && password == "" && apiKey == "" && secretkey == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Initializing provider, credentials are missing.",
					Detail:   "Provide either username+password or khulnasoft_api_key+khulnasoft_api_secret.",
				})
			}
		}
	}

	var caCertByte []byte
	if caCertPath != "" {
		caCertByte, err = os.ReadFile(caCertPath)
		if err != nil && validate {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to read CA certificates",
				Detail:   err.Error(),
			})

			return nil, diags
		}
	}

	if len(diags) > 0 {
		return nil, diags
	}

	var khulnasoftClient *client.Client
	if apiKey != "" {
		khulnasoftClient, err = client.NewClientWithAPIKey(khulnasoftURL, apiKey, secretkey, verifyTLS, caCertByte)
		if err != nil {
			return nil, diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Error creating Khulnasoft client with API key",
				Detail:   err.Error(),
			}}
		}
		if v, ok := d.GetOk("validity"); ok {
			khulnasoftClient.Validity = v.(int)
		}
		if v, ok := d.GetOk("allowed_endpoints"); ok {
			khulnasoftClient.AllowedEndpoints = convertStringArr(v.([]interface{}))
		}
		if v, ok := d.GetOk("csp_roles"); ok {
			khulnasoftClient.CSPRoles = convertStringArr(v.([]interface{}))
		}
	} else if username != "" && password != "" {
		khulnasoftClient, err = client.NewClientWithTokenAuth(khulnasoftURL, username, password, verifyTLS, caCertByte)
		if err != nil {
			return nil, diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Error creating Khulnasoft client with token authentication",
				Detail:   err.Error(),
			}}
		}
	} else {
		return nil, diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing credentials",
			Detail:   "Provide username+password or khulnasoft_api_key+khulnasoft_api_secret.",
		}}
	}

	if validate {
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
	}

	return khulnasoftClient, diags
}
