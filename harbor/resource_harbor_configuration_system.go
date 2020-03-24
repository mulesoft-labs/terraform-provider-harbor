package harbor

import (
	"context"

	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"
	"github.com/sandhose/terraform-provider-harbor/api/client/products"
	apimodels "github.com/sandhose/terraform-provider-harbor/api/models"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceHarborConfigSystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceHarborConfigSystemCreate,
		Read:   resourceHarborConfigSystemRead,
		Update: resourceHarborConfigSystemUpdate,
		Delete: resourceHarborConfigSystemDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"project_creation_restriction": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "adminonly",
				ValidateFunc: validation.StringInSlice([]string{"adminonly", "everyone"}, false),
			},
			"read_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"token_expiration": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			// TODO(burdz): check swagger definition lines up with configuration <12-03-20> //
			// "robot_token_expiration": {
			// 	Type:     schema.TypeInt,
			// 	Optional: true,
			// 	Default:  30,
			// },
			// "webhooks_enabled": {
			// 	Type:     schema.TypeBool,
			// 	Optional: true,
			// 	Default:  false,
			// },
		},
	}
}

func resourceHarborConfigSystemCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	_, err := client.Products.PutConfigurations(
		products.NewPutConfigurationsParamsWithContext(context.TODO()).
			WithConfigurations(harborConfigSystemGet(d)),
		nil,
	)
	if err != nil {
		return err
	}

	return resourceHarborConfigSystemRead(d, meta)
}

func resourceHarborConfigSystemRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	resp, err := client.Products.GetConfigurations(
		products.NewGetConfigurationsParamsWithContext(context.TODO()),
		nil,
	)
	if err != nil {
		return err
	}

	err = harborConfigSystemUpdate(d, resp.Payload)
	if err != nil {
		return err
	}

	return nil
}

func resourceHarborConfigSystemUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	_, err := client.Products.PutConfigurations(
		products.NewPutConfigurationsParamsWithContext(context.TODO()).
			WithConfigurations(harborConfigSystemGet(d)),
		nil,
	)
	if err != nil {
		return err
	}

	return resourceHarborConfigSystemRead(d, meta)
}

func resourceHarborConfigSystemDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func harborConfigSystemGet(d *schema.ResourceData) *apimodels.Configurations {
	return &apimodels.Configurations{
		ProjectCreationRestriction: d.Get("project_creation_restriction").(string),
		ReadOnly:                   d.Get("read_only").(bool),
		TokenExpiration:            d.Get("token_expiration").(int64),
		// RobotTokenExpiration:       d.Get("robot_token_expiration").(int64),
		// WebhooksEnabled:            d.Get("webhooks_enabled").(bool),
	}
}

func harborConfigSystemUpdate(d *schema.ResourceData, c *apimodels.ConfigurationsResponse) error {
	// TODO(burdz): unique ID <10-03-20> //
	d.SetId("harborsystem")

	attributes := map[string]interface{}{
		"project_creation_restriction": c.ProjectCreationRestriction,
		"read_only":                    c.ReadOnly,
		"token_expiration":             c.TokenExpiration,
		// "robot_token_expiration":       c.RobotTokenExpiration,
		// "webhooks_enabled":             c.WebhooksEnabled,
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
