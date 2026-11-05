package khulnasoft

import (
	"context"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var secureAIProviderSchema = map[string]*schema.Schema{
	"api_key": {
		Type:        schema.TypeString,
		Description: "API key for the Secure AI provider.",
		Optional:    true,
		Sensitive:   true,
	},
	"id": {
		Type:        schema.TypeString,
		Description: "Identifier for the Secure AI provider.",
		Optional:    true,
	},
	"enabled": {
		Type:        schema.TypeBool,
		Description: "Whether this Secure AI provider is enabled.",
		Optional:    true,
	},
	"mode": {
		Type:        schema.TypeString,
		Description: "Mode of the Secure AI provider (external or internal).",
		Optional:    true,
	},
}

func resourceSecureAIIntegration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSecureAIIntegrationCreate,
		ReadContext:   resourceSecureAIIntegrationRead,
		UpdateContext: resourceSecureAIIntegrationUpdate,
		DeleteContext: resourceSecureAIIntegrationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"selected_ui": {
				Type:        schema.TypeString,
				Description: "The selected UI provider for Secure AI (e.g. akamai, khulnasoft).",
				Required:    true,
			},
			"akamai": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Akamai Secure AI provider configuration.",
				Optional:    true,
				Elem:        &schema.Resource{Schema: secureAIProviderSchema},
			},
			"khulnasoft": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Khulnasoft Secure AI provider configuration.",
				Optional:    true,
				Elem:        &schema.Resource{Schema: secureAIProviderSchema},
			},
		},
	}
}

func expandSecureAIProvider(d []interface{}) *client.SecureAIProvider {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	v := d[0].(map[string]interface{})
	return &client.SecureAIProvider{
		ApiKey:  v["api_key"].(string),
		ID:      v["id"].(string),
		Enabled: v["enabled"].(bool),
		Mode:    v["mode"].(string),
	}
}

func expandSecureAIIntegration(d *schema.ResourceData) client.SecureAIIntegration {
	return client.SecureAIIntegration{
		SelectedUI: d.Get("selected_ui").(string),
		Akamai:     expandSecureAIProvider(d.Get("akamai").([]interface{})),
		Khulnasoft:       expandSecureAIProvider(d.Get("khulnasoft").([]interface{})),
	}
}

func flattenSecureAIProvider(p *client.SecureAIProvider) []map[string]interface{} {
	if p == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"api_key": p.ApiKey,
			"id":      p.ID,
			"enabled": p.Enabled,
			"mode":    p.Mode,
		},
	}
}

func resourceSecureAIIntegrationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(*client.Client)
	integration := expandSecureAIIntegration(d)
	err := ac.SaveSecureAIIntegration(integration)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("secure_ai")
	return resourceSecureAIIntegrationRead(ctx, d, m)
}

func resourceSecureAIIntegrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(*client.Client)
	integration, err := ac.GetSecureAIIntegration()
	if err != nil {
		return diag.FromErr(err)
	}
	if integration == nil {
		d.SetId("")
		return nil
	}
	d.Set("selected_ui", integration.SelectedUI)

	// Server clears sensitive fields (api_key, id) on GET via ClearSensitiveFields().
	// Preserve user-provided values from state to avoid perpetual diff.
	if integration.Akamai != nil {
		if prev, ok := d.GetOk("akamai"); ok {
			prevList := prev.([]interface{})
			if len(prevList) > 0 && prevList[0] != nil {
				prevMap := prevList[0].(map[string]interface{})
				if integration.Akamai.ApiKey == "" {
					integration.Akamai.ApiKey = prevMap["api_key"].(string)
				}
				if integration.Akamai.ID == "" {
					integration.Akamai.ID = prevMap["id"].(string)
				}
			}
		}
	}
	if integration.Khulnasoft != nil {
		if prev, ok := d.GetOk("khulnasoft"); ok {
			prevList := prev.([]interface{})
			if len(prevList) > 0 && prevList[0] != nil {
				prevMap := prevList[0].(map[string]interface{})
				if integration.Khulnasoft.ApiKey == "" {
					integration.Khulnasoft.ApiKey = prevMap["api_key"].(string)
				}
				if integration.Khulnasoft.ID == "" {
					integration.Khulnasoft.ID = prevMap["id"].(string)
				}
			}
		}
	}

	d.Set("akamai", flattenSecureAIProvider(integration.Akamai))
	d.Set("khulnasoft", flattenSecureAIProvider(integration.Khulnasoft))
	return nil
}

func resourceSecureAIIntegrationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(*client.Client)
	integration := expandSecureAIIntegration(d)
	err := ac.SaveSecureAIIntegration(integration)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceSecureAIIntegrationRead(ctx, d, m)
}

func resourceSecureAIIntegrationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
