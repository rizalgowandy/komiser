package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	. "github.com/mlabouardy/komiser/models/aws"
)

func (awsClient AWS) DescribeInstances(cfg awsConfig.Config) (map[string]interface{}, error) {
	outputInstancesPerRegion := make(map[string]int, 0)
	outputInstancesPerState := make(map[string]int, 0)
	outputInstancesPerFamily := make(map[string]int, 0)
	totalPublicInstances := 0
	totalPrivateInstances := 0
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return map[string]interface{}{}, err
	}
	for _, region := range regions {
		instances, err := awsClient.getInstances(cfg, region.Name)
		if err != nil {
			return map[string]interface{}{}, err
		}
		for _, instance := range instances {
			outputInstancesPerState[instance.State]++
			outputInstancesPerFamily[instance.InstanceType]++
			if instance.Public {
				totalPublicInstances++
			} else {
				totalPrivateInstances++
			}
		}
		outputInstancesPerRegion[region.Name] = len(instances)
	}
	return map[string]interface{}{
		"region":  outputInstancesPerRegion,
		"state":   outputInstancesPerState,
		"family":  outputInstancesPerFamily,
		"public":  totalPublicInstances,
		"private": totalPrivateInstances,
	}, nil
}

func (awsClient AWS) getInstances(cfg awsConfig.Config, region string) ([]EC2, error) {
	cfg.Region = region
	ec2Svc := ec2.NewFromConfig(cfg)
	params := &ec2.DescribeInstancesInput{}
	result, err := ec2Svc.DescribeInstances(context.Background(), params)
	if err != nil {
		return []EC2{}, err
	}
	listOfInstances := make([]EC2, 0)
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			instanceTags := make([]string, 0)
			for _, tag := range instance.Tags {
				instanceTags = append(instanceTags, *tag.Value)
			}
			isPublic := false
			if instance.PublicIpAddress != nil {
				isPublic = true
			}
			listOfInstances = append(listOfInstances, EC2{
				ID:           *instance.InstanceId,
				InstanceType: string(instance.InstanceType),
				LaunchTime:   *instance.LaunchTime,
				Tags:         instanceTags,
				State:        string(instance.State.Name),
				Public:       isPublic,
			})
		}
	}
	return listOfInstances, nil
}

func (awsClient AWS) DescribeScheduledInstances(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := ec2.NewFromConfig(cfg)
		res, _ := svc.DescribeScheduledInstances(context.Background(), &ec2.DescribeScheduledInstancesInput{})

		if res != nil {
			for _, set := range res.ScheduledInstanceSet {
				sum += int64(*set.InstanceCount)
			}
		}
	}
	return sum, nil
}

func (awsClient AWS) DescribeReservedInstances(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := ec2.NewFromConfig(cfg)
		res, err := svc.DescribeReservedInstances(context.Background(), &ec2.DescribeReservedInstancesInput{})
		if err != nil {
			return sum, err
		}

		for _, reservation := range res.ReservedInstances {
			sum += int64(*reservation.InstanceCount)
		}
	}
	return sum, nil
}

func (awsClient AWS) DescribeSpotInstances(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := ec2.NewFromConfig(cfg)
		res, err := svc.DescribeSpotFleetRequests(context.Background(), &ec2.DescribeSpotFleetRequestsInput{})
		if err != nil {
			return sum, err
		}

		for _, request := range res.SpotFleetRequestConfigs {
			res2, err := svc.DescribeSpotFleetInstances(context.Background(), &ec2.DescribeSpotFleetInstancesInput{
				SpotFleetRequestId: request.SpotFleetRequestId,
			})
			if err != nil {
				return sum, err
			}

			sum += int64(len(res2.ActiveInstances))
		}
	}
	return sum, nil
}
