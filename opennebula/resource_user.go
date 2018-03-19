package opennebula

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the user",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password of the user",
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	// Create base object
	client.Call(
		"one.user.allocate",
		d.Get("username").(string),d.Get("password").(string))

	return nil;
}
