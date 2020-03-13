package harbor

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HARBOR_HOST", nil),
				Description: descriptions["host"],
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HARBOR_USERNAME", nil),
				Description: descriptions["username"],
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HARBOR_PASSWORD", nil),
				Description: descriptions["password"],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"harbor_project":                    resourceHarborProject(),
			"harbor_project_robot_account":      resourceHarborProjectRobotAccount(),
			"harbor_usergroup":                  resourceHarborUserGroup(),
			"harbor_system_configuration_auth":  resourceHarborConfigAuth(),
			"harbor_system_configuration_email": resourceHarborConfigEmail(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"harbor_system_info": dataSourceHarborSystemInfo(),
			"harbor_project":     dataSourceHarborProject(),
			"harbor_usergroup":   dataSourceHarborUserGroup(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}

	return provider
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	cfg := &Config{
		Host:             d.Get("host").(string),
		Username:         d.Get("username").(string),
		Password:         d.Get("password").(string),
		terraformVersion: terraformVersion,
	}
	return cfg.Config()
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"host":     "The hostname (in form of URI) of the Harbor server.",
		"username": "The username to use for HTTP basic authentication when accessing the Harbor server.",
		"password": "The password to use for HTTP basic authentication when accessing the Harbor server.",
	}
}
