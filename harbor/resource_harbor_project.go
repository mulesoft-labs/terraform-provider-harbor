package harbor

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"
	"github.com/sandhose/terraform-provider-harbor/api/client/products"
	apimodels "github.com/sandhose/terraform-provider-harbor/api/models"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// TODO(burdz): update schema to include current_user_role_ids, cve_whitelist <10-03-20> //
func resourceHarborProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceHarborProjectCreate,
		Read:   resourceHarborProjectRead,
		Update: resourceHarborProjectUpdate,
		Delete: resourceHarborProjectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"public": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"auto_scan": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"content_trust": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"prevent_vulnerability": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"severity": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "low",
			},
			"reuse_sys_cve_whitelist": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			// computed
			"project_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"owner_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"owner_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deleted": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"repo_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"chart_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceHarborProjectCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)
	name := d.Get("name").(string)

	projectParams := &apimodels.ProjectReq{
		ProjectName: name,
		Metadata: &apimodels.ProjectMetadata{
			Public:               strconv.FormatBool(d.Get("public").(bool)),
			AutoScan:             strconv.FormatBool(d.Get("auto_scan").(bool)),
			PreventVul:           strconv.FormatBool(d.Get("prevent_vulnerability").(bool)),
			Severity:             d.Get("severity").(string),
			EnableContentTrust:   strconv.FormatBool(d.Get("content_trust").(bool)),
			ReuseSysCveWhitelist: strconv.FormatBool(d.Get("content_trust").(bool)),
		},
	}

	_, err := client.Products.PostProjects(
		products.NewPostProjectsParamsWithContext(context.TODO()).
			WithProject(projectParams),
		nil,
	)

	if err != nil {
		return err
	}

	resp, err := client.Products.GetProjects(
		products.NewGetProjectsParamsWithContext(context.TODO()).
			WithName(&name),
		nil,
	)

	if err != nil {
		return err
	}

	var project *apimodels.Project
	for _, p := range resp.Payload {
		if strings.EqualFold(p.Name, name) {
			project = p
			break
		}
	}

	if project == nil {
		return fmt.Errorf("Project not found")
	}

	harborProjectUpdate(d, project)

	return nil
}

func resourceHarborProjectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)

	if err != nil {
		return err
	}

	resp, err := client.Products.GetProjectsProjectID(
		products.NewGetProjectsProjectIDParamsWithContext(context.TODO()).
			WithProjectID(projectID),
		nil,
	)

	if err != nil {
		return err
	}

	harborProjectUpdate(d, resp.Payload)

	return nil
}

func resourceHarborProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	projectParams := &apimodels.ProjectReq{
		Metadata: &apimodels.ProjectMetadata{
			Public:               strconv.FormatBool(d.Get("public").(bool)),
			AutoScan:             strconv.FormatBool(d.Get("auto_scan").(bool)),
			PreventVul:           strconv.FormatBool(d.Get("prevent_vulnerability").(bool)),
			Severity:             d.Get("severity").(string),
			EnableContentTrust:   strconv.FormatBool(d.Get("content_trust").(bool)),
			ReuseSysCveWhitelist: strconv.FormatBool(d.Get("content_trust").(bool)),
		},
	}

	_, err = client.Products.PutProjectsProjectID(
		products.NewPutProjectsProjectIDParamsWithContext(context.TODO()).
			WithProjectID(projectID).
			WithProject(projectParams),
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func resourceHarborProjectDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	_, err = client.Products.DeleteProjectsProjectID(
		products.NewDeleteProjectsProjectIDParamsWithContext(context.TODO()).
			WithProjectID(projectID),
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

func harborProjectUpdate(d *schema.ResourceData, p *apimodels.Project) {
	d.SetId(fmt.Sprint(p.ProjectID))

	attributes := map[string]interface{}{
		"project_id":              p.ProjectID,
		"name":                    p.Name,
		"public":                  p.Metadata.Public,
		"auto_scan":               p.Metadata.AutoScan,
		"prevent_vulnerability":   p.Metadata.PreventVul,
		"severity":                p.Metadata.Severity,
		"content_trust":           p.Metadata.EnableContentTrust,
		"reuse_sys_cve_whitelist": p.Metadata.ReuseSysCveWhitelist,
		"owner_id":                p.OwnerID,
		"owner_name":              p.OwnerName,
		"creation_time":           p.CreationTime,
		"update_time":             p.UpdateTime,
		"deleted":                 p.Deleted,
		"repo_count":              p.RepoCount,
		"chart_count":             p.ChartCount,
	}

	for key, val := range attributes {
		d.Set(key, val)
		// TODO(burdz): return error <12-03-20> //
		// err = d.Set(key, val)
		// if err != nil {
		// 	return err
		// }
	}
}
