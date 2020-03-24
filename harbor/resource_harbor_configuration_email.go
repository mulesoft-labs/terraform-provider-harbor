package harbor

import (
	"context"

	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"
	"github.com/sandhose/terraform-provider-harbor/api/client/products"
	apimodels "github.com/sandhose/terraform-provider-harbor/api/models"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceHarborConfigEmail() *schema.Resource {
	return &schema.Resource{
		Create: resourceHarborConfigEmailCreate,
		Read:   resourceHarborConfigEmailRead,
		Update: resourceHarborConfigEmailUpdate,
		Delete: resourceHarborConfigEmailDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"email_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email_port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"email_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_identity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_from": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email_ssl": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"email_verify_cert": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceHarborConfigEmailCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	_, err := client.Products.PutConfigurations(
		products.NewPutConfigurationsParamsWithContext(context.TODO()).
			WithConfigurations(harborConfigEmailGet(d)),
		nil,
	)
	if err != nil {
		return err
	}

	return resourceHarborConfigEmailRead(d, meta)
}

func resourceHarborConfigEmailRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	resp, err := client.Products.GetConfigurations(
		products.NewGetConfigurationsParamsWithContext(context.TODO()),
		nil,
	)
	if err != nil {
		return err
	}

	err = harborConfigEmailUpdate(d, resp.Payload)
	if err != nil {
		return err
	}

	return nil
}

func resourceHarborConfigEmailUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	_, err := client.Products.PutConfigurations(
		products.NewPutConfigurationsParamsWithContext(context.TODO()).
			WithConfigurations(harborConfigEmailGet(d)),
		nil,
	)
	if err != nil {
		return err
	}

	return resourceHarborConfigEmailRead(d, meta)
}

func resourceHarborConfigEmailDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func harborConfigEmailGet(d *schema.ResourceData) *apimodels.Configurations {
	return &apimodels.Configurations{
		EmailHost:     d.Get("email_host").(string),
		EmailPort:     d.Get("email_port").(int64),
		EmailUsername: d.Get("email_username").(string),
		EmailIdentity: d.Get("email_identity").(string),
		EmailFrom:     d.Get("email_from").(string),
		EmailSsl:      d.Get("email_ssl").(bool),
		EmailInsecure: d.Get("email_verify_cert").(bool),
	}
}

func harborConfigEmailUpdate(d *schema.ResourceData, c *apimodels.ConfigurationsResponse) error {
	// TODO(burdz): unique ID <10-03-20> //
	d.SetId("harboremail")

	attributes := map[string]interface{}{
		"email_host":        c.EmailHost,
		"email_port":        c.EmailPort,
		"email_username":    c.EmailUsername,
		"email_identity":    c.EmailIdentity,
		"email_from":        c.EmailFrom,
		"email_ssl":         c.EmailSsl,
		"email_verify_cert": c.EmailInsecure,
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
