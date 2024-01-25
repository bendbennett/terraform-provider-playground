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

			//// Reveals warning when blocks are mutated in SDKv2
			//"mutual_authentication": {
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	Computed: true,
			//	MaxItems: 1,
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			"ignore_client_certificate_expiry": {
			//				Type:     schema.TypeBool,
			//				Optional: true,
			//				Default:  false,
			//			},
			//			"mode": {
			//				Type:     schema.TypeString,
			//				Required: true,
			//			},
			//			"trust_store_arn": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//			},
			//		},
			//	},
			//},
		},
	}
}

type MutualAuthenticationAttributes struct {
	_ struct{} `type:"structure"`

	// Indicates whether expired client certificates are ignored.
	IgnoreClientCertificateExpiry *bool `type:"boolean"`

	// The client certificate handling method. Options are off, passthrough or verify.
	// The default value is off.
	Mode *string `type:"string"`

	// The Amazon Resource Name (ARN) of the trust store.
	TrustStoreArn *string `type:"string"`
}

func Pointer[T any](in T) *T {
	return &in
}

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("example-id")

	//// Reveals warning when blocks are mutated in SDKv2
	//ma := map[string]interface{}{
	//	"mode":                             "off",
	//	"trust_store_arn":                  "arn:aws:acm-pca:us-east-1:123456789012:truststore/tstore-123456789012",
	//	"ignore_client_certificate_expiry": true,
	//}
	//
	//err := d.Set("mutual_authentication", []interface{}{ma})
	//
	//if err != nil {
	//	return diag.FromErr(err)
	//}

	return nil
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
