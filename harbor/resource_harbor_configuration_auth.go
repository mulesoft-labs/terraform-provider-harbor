package harbor

import (
	"context"
	"fmt"

	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"
	"github.com/sandhose/terraform-provider-harbor/api/client/products"
	apimodels "github.com/sandhose/terraform-provider-harbor/api/models"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceHarborConfigAuth() *schema.Resource {
	return &schema.Resource{
		Create: resourceHarborConfigAuthCreate,
		Read:   resourceHarborConfigAuthRead,
		Update: resourceHarborConfigAuthUpdate,
		Delete: resourceHarborConfigAuthDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"auth_mode": {
				Type:     schema.TypeString,
				Required: true,
			},
			"oidc_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oidc_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oidc_client_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oidc_client_secret": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			// "oidc_groups_claim": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			"oidc_scope": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oidc_verify_cert": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceHarborConfigAuthCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)
	authMode := d.Get("auth_mode").(string)

	_, err := client.Products.PutConfigurations(
		products.NewPutConfigurationsParamsWithContext(context.TODO()).
			WithConfigurations(harborConfigAuthGet(d)),
		nil,
	)
	if err != nil {
		return err
	}

	resp, err := client.Products.GetConfigurations(
		products.NewGetConfigurationsParamsWithContext(context.TODO()),
		nil,
	)
	if err != nil {
		return err
	}

	if resp.Payload.AuthMode.Value != authMode {
		return fmt.Errorf("auth_mode not set correctly")
	}

	err = harborConfigAuthUpdate(d, resp.Payload)
	if err != nil {
		return err
	}

	return nil
}

func resourceHarborConfigAuthRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	resp, err := client.Products.GetConfigurations(
		products.NewGetConfigurationsParamsWithContext(context.TODO()),
		nil,
	)
	if err != nil {
		return err
	}

	err = harborConfigAuthUpdate(d, resp.Payload)
	if err != nil {
		return err
	}

	return nil
}

func resourceHarborConfigAuthUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)
	authMode := d.Get("auth_mode").(string)

	_, err := client.Products.PutConfigurations(
		products.NewPutConfigurationsParamsWithContext(context.TODO()).
			WithConfigurations(harborConfigAuthGet(d)),
		nil,
	)
	if err != nil {
		return err
	}

	resp, err := client.Products.GetConfigurations(
		products.NewGetConfigurationsParamsWithContext(context.TODO()),
		nil,
	)
	if err != nil {
		return err
	}

	if resp.Payload.AuthMode.Value != authMode {
		return fmt.Errorf("auth_mode not set correctly")
	}

	err = harborConfigAuthUpdate(d, resp.Payload)
	if err != nil {
		return err
	}

	return nil
}

func resourceHarborConfigAuthDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func harborConfigAuthGet(d *schema.ResourceData) *apimodels.Configurations {
	return &apimodels.Configurations{
		AuthMode:         d.Get("auth_mode").(string),
		OidcName:         d.Get("oidc_name").(string),
		OidcEndpoint:     d.Get("oidc_endpoint").(string),
		OidcClientID:     d.Get("oidc_client_id").(string),
		OidcClientSecret: d.Get("oidc_client_secret").(string),
		// OidcGroupsClaim:  d.Set("oidc_groups_claim").(string),
		OidcScope:      d.Get("oidc_scope").(string),
		OidcVerifyCert: d.Get("oidc_verify_cert").(bool),
	}
}

func harborConfigAuthUpdate(d *schema.ResourceData, c *apimodels.ConfigurationsResponse) error {
	d.SetId(fmt.Sprint(c.AuthMode))

	attributes := map[string]interface{}{
		"auth_mode":      c.AuthMode,
		"oidc_name":      c.OidcName,
		"oidc_endpoint":  c.OidcEndpoint,
		"oidc_client_id": c.OidcClientID,
		// "oidc_groups_claim": c.OidcClientID,
		"oidc_scope":       c.OidcScope,
		"oidc_verify_cert": c.OidcVerifyCert,
	}

	var err error
	for key, val := range attributes {
		err = d.Set(key, val)
		if err != nil {
			return err
		}
	}
	return nil
}
