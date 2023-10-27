package scaleway

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/providers/ovh/alerting"
	"github.com/tailwarden/komiser/providers/ovh/image"
	"github.com/tailwarden/komiser/providers/ovh/instance"
	"github.com/tailwarden/komiser/utils"

	"github.com/tailwarden/komiser/providers"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []providers.FetchDataFunction {
	return []providers.FetchDataFunction{
		alerting.Alerts,
		image.Images,
		instance.Instances,
	}
}

func FetchResources(ctx context.Context, client providers.ProviderClient, db *bun.DB, telemetry bool, analytics utils.Analytics) {
	for _, fetchResources := range listOfSupportedServices() {
		resources, err := fetchResources(ctx, client)
		if err != nil {
			log.Printf("[%s][OVH] %s", client.Name, err)
		} else {
			for _, resource := range resources {
				_, err := db.NewInsert().Model(&resource).On("CONFLICT (resource_id) DO UPDATE").Set("cost = EXCLUDED.cost").Exec(context.Background())
				if err != nil {
					logrus.WithError(err).Errorf("db trigger failed")
				}
			}
			if telemetry {
				analytics.TrackEvent("discovered_resources", map[string]interface{}{
					"provider":  "OVH",
					"resources": len(resources),
				})
			}
		}
	}
}
