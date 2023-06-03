package compute

import (
	"context"
	"fmt"
	"time"

	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

const createdLayout = "2006-01-02 15:04:05 +0000 +0000"

func Servers(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	instanceSvc := instance.NewAPI(client.ScalewayClient)

	regions := []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}

	for _, region := range regions {
		for _, zone := range region.GetZones() {
			serversTypes, err := instanceSvc.ListServersTypes(&instance.ListServersTypesRequest{
				Zone: zone,
			})
			if err != nil {
				return resources, err
			}

			output, err := instanceSvc.ListServers(&instance.ListServersRequest{
				Zone: zone,
			})
			if err != nil {
				return resources, err
			}

			for _, inst := range output.Servers {
				serverType := serversTypes.Servers[inst.CommercialType]
				hourlyPrice := serverType.HourlyPrice
				currentTime := time.Now()
				currentMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)
				created, err := time.Parse(createdLayout, inst.CreationDate.String())
				if err != nil {
					return nil, err
				}

				var duration time.Duration
				if created.Before(currentMonth) {
					duration = currentTime.Sub(currentMonth)
				} else {
					duration = currentTime.Sub(created)
				}

				monthlyCost := hourlyPrice * float32(duration.Hours())

				resources = append(resources, models.Resource{
					Provider:   "Scaleway",
					Account:    client.Name,
					Service:    "Server",
					Region:     inst.Zone.String(),
					ResourceId: inst.ID,
					Cost:       float64(monthlyCost),
					// inst.Tags is a slice of strings so, we can't extract them in key and value pair.
					//Tags:       inst.Tags,
					Name:      inst.Name,
					FetchedAt: time.Now(),
					Link:      fmt.Sprintf("https://console.scaleway.com/instance/servers/%s/%s", inst.Zone.String(), inst.ID),
				})
			}
		}
	}
	log.WithFields(log.Fields{
		"provider":  "Scaleway",
		"account":   client.Name,
		"service":   "Server",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
