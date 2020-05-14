package mackerel

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

//Provider returns a terraform.ResourceProvider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MACKEREL_API_KEY", nil),
				Description: "Mackerel API Key",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"mackerel_channel":          resourceMackerelChannel(),
			"mackerel_role":             resourceMackerelRole(),
			"mackerel_role_metadata":    resourceMackerelRoleMetadata(),
			"mackerel_service":          resourceMackerelService(),
			"mackerel_service_metadata": resourceMackerelServiceMetadata(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	config := config{
		APIKey: data.Get("api_key").(string),
	}
	return config.Client()
}