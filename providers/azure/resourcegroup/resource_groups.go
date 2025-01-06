package resourcegroup

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	log "github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
)

func ResourceGroups(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
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

	resourceGroupClient, err := armresources.NewResourceGroupsClient(
		client.AzureClient.SubscriptionId,
		client.AzureClient.Credentials,
		clientOptions,
	)

	if err != nil {
		return resources, err
	}

	pager := resourceGroupClient.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return resources, err
		}

		for _, resourceGroup := range page.ResourceGroupListResult.Value {

			tags := make([]models.Tag, 0)

			for key, value := range resourceGroup.Tags {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: *value,
				})
			}

			resources = append(resources, models.Resource{
				Provider:   "Azure",
				Account:    client.Name,
				Service:    "Resource Group",
				Region:     *resourceGroup.Location,
				ResourceId: *resourceGroup.ID,
				Cost:       0,
				Name:       *resourceGroup.Name,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https://portal.azure.com/#resource%s", *resourceGroup.ID),
			})
		}
	}

	log.WithFields(log.Fields{
		"provider":  "Azure",
		"account":   client.Name,
		"service":   "Resource Group",
		"resources": len(resources),
	}).Info("Fetched resource groups")

	return resources, nil
}
