package provider_sdk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func exampleSdkResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: create,
		ReadContext:   read,
		UpdateContext: update,
		DeleteContext: schema.NoopContext,
		CustomizeDiff: func(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error {
			tflog.Info(ctx, "calling CustomizeDiff")
			return nil
		},

		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				return nil, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"attr_settings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"id": {
				Computed: true,
				Type:     schema.TypeString,
			},
		},
	}
}

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("example-id")

	return nil
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
