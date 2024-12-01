package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func Databoxes(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	retryOptions := policy.RetryOptions{
		MaxRetries:    6,
		RetryDelay:    2 * time.Second,
		MaxRetryDelay: 120 * time.Second,
	}

	clientOptions := &arm.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Retry: retryOptions,
		},
	}

	svc, err := armdatabox.NewJobsClient(client.AzureClient.SubscriptionId, client.AzureClient.Credentials, clientOptions)
	if err != nil {
		return resources, err
	}

	pager := svc.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return resources, err
		}

		for _, val := range page.Value {
			tags := make([]models.Tag, 0)

			for key, value := range val.Tags {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: *value,
				})
			}

			resources = append(resources, models.Resource{
				Provider:   "Azure",
				Account:    client.Name,
				Service:    "Databox",
				Region:     *val.Location,
				ResourceId: *val.ID,
				Cost:       0,
				Name:       *val.Name,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https://portal.azure.com/#resource%s", *val.ID),
			})
		}

	}

	log.WithFields(log.Fields{
		"provider":  "Azure",
		"account":   client.Name,
		"service":   "Databox",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
