package migration

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func SubscriptionPricingV0ToV1() schema.StateUpgrader {
	return schema.StateUpgrader{
		Type:    subscriptionPricingSchemaForV0().CoreConfigSchema().ImpliedType(),
		Upgrade: subscriptionPricingUpgradeV0ToV1,
		Version: 0,
	}
}

func subscriptionPricingSchemaForV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tier": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func subscriptionPricingUpgradeV0ToV1(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating ResourceType from v0 to v1 format")
	oldId := rawState["id"].(string)
	newId := strings.Replace(oldId, "/default", "/VirtualMachines", 1)

	log.Printf("[DEBUG] Updating ID from %q to %q", oldId, newId)

	rawState["id"] = newId

	return rawState, nil
}
