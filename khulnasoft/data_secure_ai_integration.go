package khulnasoft

import (
	"context"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var secureAIProviderDataSchema = map[string]*schema.Schema{
	"api_key": {
		Type:        schema.TypeString,
		Description: "API key for the Secure AI provider.",
		Computed:    true,
		Sensitive:   true,
	},
	"id": {
		Type:        schema.TypeString,
		Description: "Identifier for the Secure AI provider.",
		Computed:    true,
	},
	"enabled": {
		Type:        schema.TypeBool,
		Description: "Whether this Secure AI provider is enabled.",
		Computed:    true,
	},
	"mode": {
		Type:        schema.TypeString,
		Description: "Mode of the Secure AI provider (external or internal).",
		Computed:    true,
	},
}

func dataSecureAIIntegration() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSecureAIIntegrationRead,
		Schema: map[string]*schema.Schema{
			"selected_ui": {
				Type:        schema.TypeString,
				Description: "The selected UI provider for Secure AI.",
				Computed:    true,
			},
			"akamai": {
				Type:        schema.TypeList,
				Description: "Akamai Secure AI provider configuration.",
				Computed:    true,
				Elem:        &schema.Resource{Schema: secureAIProviderDataSchema},
			},
			"khulnasoft": {
				Type:        schema.TypeList,
				Description: "Khulnasoft Secure AI provider configuration.",
				Computed:    true,
				Elem:        &schema.Resource{Schema: secureAIProviderDataSchema},
			},
		},
	}
}

func dataSecureAIIntegrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ac := m.(*client.Client)
	integration, err := ac.GetSecureAIIntegration()
	if err != nil {
		return diag.FromErr(err)
	}
	if integration == nil {
		d.SetId("")
		return nil
	}
	d.SetId("secure_ai")
	d.Set("selected_ui", integration.SelectedUI)
	d.Set("akamai", flattenSecureAIProvider(integration.Akamai))
	d.Set("khulnasoft", flattenSecureAIProvider(integration.Khulnasoft))
	return nil
}
