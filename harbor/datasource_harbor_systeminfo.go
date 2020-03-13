package harbor

import (
	"context"

	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"
	"github.com/sandhose/terraform-provider-harbor/api/client/products"
	"github.com/sandhose/terraform-provider-harbor/api/models"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHarborSystemInfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHarborSystemInfoRead,
		Schema: map[string]*schema.Schema{
			"with_notary": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"with_clair": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"with_admiral": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"admiral_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"registry_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_creation_restriction": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"self_registration": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"has_ca_root": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"next_scan_all": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"harbor_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"clair_vulnerability_status": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"overall_last_update": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						// "details": {
						// 	Type:     schema.TypeInt,
						// 	Computed: true,
						// },
					},
				},
			},
		},
	}
}

func dataSourceHarborSystemInfoRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)
	resp, err := client.Products.GetSysteminfo(products.NewGetSysteminfoParamsWithContext(context.TODO()), nil)
	if err != nil {
		return err
	}

	i := resp.Payload

	// TODO(burdz): unique ID <10-03-20> //
	d.SetId("systeminfo")
	attributes := map[string]interface{}{
		"admiral_endpoint":             i.AdmiralEndpoint,
		"auth_mode":                    i.AuthMode,
		"external_url":                 i.ExternalURL,
		"harbor_version":               i.HarborVersion,
		"has_ca_root":                  i.HasCaRoot,
		"next_scan_all":                i.NextScanAll,
		"project_creation_restriction": i.ProjectCreationRestriction,
		"registry_url":                 i.RegistryURL,
		"self_registration":            i.SelfRegistration,
		"with_admiral":                 i.WithAdmiral,
		"with_clair":                   i.WithClair,
		"with_notary":                  i.WithNotary,
	}

	for key, val := range attributes {
		err = d.Set(key, val)
		if err != nil {
			return err
		}
	}

	clairVulnerabilityStatus, err := flattenClairVulnerabilityStatus(i.ClairVulnerabilityStatus, d)
	if err != nil {
		return err
	}

	if clairVulnerabilityStatus != nil {
		if err := d.Set("clair_vulnerability_status", clairVulnerabilityStatus); err != nil {
			return err
		}
	}
	return nil
}

func flattenClairVulnerabilityStatus(i *models.GeneralInfoClairVulnerabilityStatus, d *schema.ResourceData) (map[string]interface{}, error) {
	if i == nil {
		return nil, nil
	}

	clairStatus := make(map[string]interface{})

	clairStatus["overall_last_update"] = i.OverallLastUpdate
	return clairStatus, nil
}
