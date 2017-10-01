package cachethq

import (
	"log"
	"strconv"

	"github.com/andygrunwald/cachet"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceComponentGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceComponentGroupCreate,
		Read:   resourceComponentGroupRead,
		Update: resourceComponentGroupUpdate,
		Delete: resourceComponentGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the component group",
				Required:    true,
			},
			"order": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceComponentGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cachet.Client)
	componentGroup := &cachet.ComponentGroup{
		Name:  d.Get("name").(string),
		Order: d.Get("order").(int),
	}

	componentGroup, _, err := client.Components.CreateGroup(componentGroup)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(componentGroup.ID))

	return resourceComponentGroupRead(d, meta)
}

func resourceComponentGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cachet.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	componentGroup, response, err := client.Components.GetGroup(id)
	if err != nil {
		if response.StatusCode == 404 {
			log.Printf("[WARN] removing component group %s from state because it no longer exists in cachethq", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", componentGroup.Name)
	d.Set("order", componentGroup.Order)

	return nil
}

func resourceComponentGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cachet.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	componentGroup, _, err := client.Components.GetGroup(id)
	if err != nil {
		return err
	}

	if d.HasChange("name") {
		componentGroup.Name = d.Get("name").(string)
	}

	if d.HasChange("order") {
		componentGroup.Order = d.Get("order").(int)
	}

	client.Components.UpdateGroup(id, componentGroup)

	return resourceComponentGroupRead(d, meta)
}

func resourceComponentGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cachet.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	_, err = client.Components.DeleteGroup(id)
	if err != nil {
		return err
	}
	return nil
}
