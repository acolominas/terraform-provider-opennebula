package opennebula

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type User struct {
	XMLName  xml.Name `xml:"USER"`
	ID       int      `xml:"ID"`
	GID      int      `xml:"GID"`
	Name     string   `xml:"NAME"`
	GroupIDs []int    `xml:"GROUPS>ID"`
	//Password   string   `xml:"PASSWORD"`
	AuthDriver string `xml:"AUTH_DRIVER"`
}

type UserPool struct {
	XMLName xml.Name `xml:"USER_POOL"`
	Users   []User   `xml:"USER"`
}


func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Exists: resourceUserExists,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
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
			"auth_driver": {
				Type: 				schema.TypeString,
				Optional:			true,
				Description:	"Authentication driver for the user",
				Default:			"",
			},
			"groups": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Groups assigned to user",
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	var groups = []int{}
	if v, ok := d.GetOk("groups"); ok {
		groupItf := v.([]interface{})
		for _, gp := range groupItf {
			groups = append(groups, gp.(int))
		}
	}

	id, err := client.Call(
			"one.user.allocate",
			d.Get("username").(string),d.Get("password").(string),d.Get("auth_driver"),groups,
	)
	if err != nil {
		return err
	}
	d.SetId(id)
	return nil;
}

func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceUserExists(d *schema.ResourceData, meta interface{})  (bool, error) {
	return false, nil
}

func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
