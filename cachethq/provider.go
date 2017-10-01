package cachethq

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider is the cachethq provider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CACHETHQ_TOKEN", nil),
				Description: "API Token to read and update CachetHQ",
			},
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CACHETHQ_BASE_URL", nil),
				Description: "Cachethq base api url",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cachethq_component_group": resourceComponentGroup(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:   d.Get("token").(string),
		BaseURL: d.Get("base_url").(string),
	}

	return config.Client()
}
