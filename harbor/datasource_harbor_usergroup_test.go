package harbor

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"
	"github.com/sandhose/terraform-provider-harbor/api/client/products"
)

func TestAccHarborUserGroupDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHarborUserGroupDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHarborUserGroupDataSourceConfig(expectedDataSourceUserGroupName),
				Check: resource.ComposeTestCheckFunc(
					testAccHarborUserGroupDataSource("data.harbor_usergroup.bar"),
				),
			},
		},
	})
}

func testAccHarborUserGroupDataSource(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("Expected to read user group data from Harbor")
		}

		if a["name"] != expectedDataSourceUserGroupName {
			return fmt.Errorf("Expected the user group name to be: %s, but got: %s", expectedDataSourceUserGroupName, a["name"])
		}
		return nil
	}
}

func testAccHarborUserGroupDataSourceConfig(rName string) string {
	return fmt.Sprintf(`
resource "harbor_usergroup" "foo" {
	name = "%[1]s"
	type = 2
}
data "harbor_usergroup" "bar" {
	name = "${harbor_usergroup.foo.name}"
}
`, rName)
}

func testAccCheckHarborUserGroupDataSourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiclient.Harbor)
	for _, r := range s.RootModule().Resources {
		if r.Type != "harbor_usergroup" {
			continue
		}

		id, err := strconv.ParseInt(r.Primary.ID, 10, 32)
		if err != nil {
			return err
		}

		_, err = client.Products.GetUsergroupsGroupID(
			products.NewGetUsergroupsGroupIDParamsWithContext(context.TODO()).
				WithGroupID(int64(id)),
			nil,
		)
		if err == nil {
			return fmt.Errorf("Harbor user group still exists: %s", r.Primary.ID)
		}

	}
	return nil
}
