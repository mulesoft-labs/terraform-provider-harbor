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
	d.Set("admiral_endpoint", i.AdmiralEndpoint)
	d.Set("auth_mode", i.AuthMode)
	d.Set("external_url", i.ExternalURL)
	d.Set("harbor_version", i.HarborVersion)
	d.Set("has_ca_root", i.HasCaRoot)
	d.Set("next_scan_all", i.NextScanAll)
	d.Set("project_creation_restriction", i.ProjectCreationRestriction)
	d.Set("registry_url", i.RegistryURL)
	d.Set("self_registration", i.SelfRegistration)
	d.Set("with_admiral", i.WithAdmiral)
	d.Set("with_clair", i.WithClair)
	d.Set("with_notary", i.WithNotary)

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
