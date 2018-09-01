package provider

import (
	"github.com/alexashley/terraform-provider-kong/kong/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceKongPlugin() *schema.Resource {
	return &schema.Resource{
		Create: resourceKongPluginCreate,
		Read:   resourceKongPluginRead,
		Update: resourceKongPluginUpdate,
		Delete: resourceKongPluginDelete,
		Schema: map[string]*schema.Schema{
			"service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"route_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"consumer_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"config": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceKongPluginCreate(data *schema.ResourceData, meta interface{}) error {
	kongClient := meta.(*client.KongClient)
	plugin, err := kongClient.CreatePlugin(client.KongPlugin{
		ServiceId:  data.Get("service_id").(string),
		RouteId:    data.Get("route_id").(string),
		ConsumerId: data.Get("consumer_id").(string),
		Name:       data.Get("name").(string),
		Config:     data.Get("config").(map[string]interface{}),
		Enabled:    data.Get("enabled").(bool),
	})

	if err != nil {
		return err
	}

	data.SetId(plugin.Id)
	data.Set("created_at", plugin.CreatedAt)

	return nil
}

func resourceKongPluginRead(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceKongPluginUpdate(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceKongPluginDelete(data *schema.ResourceData, meta interface{}) error {
	return nil
}