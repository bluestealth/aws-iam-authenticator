package ec2provider

import (
	"context"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const (
	DescribeDelay = 100
)

type mockEc2Client struct {
	EC2API
	Reservations []*ec2Types.Reservation
}

func (c mockEc2Client) DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	// simulate the time it takes for aws to return
	time.Sleep(DescribeDelay * time.Millisecond)
	var reservations []ec2Types.Reservation
	for _, res := range c.Reservations {
		var reservation ec2Types.Reservation
		for _, inst := range res.Instances {
			for _, id := range params.InstanceIds {
				if id == aws.ToString(inst.InstanceId) {
					reservation.Instances = append(reservation.Instances, inst)
				}
			}
		}
		if len(reservation.Instances) > 0 {
			reservations = append(reservations, reservation)
		}
	}
	return &ec2.DescribeInstancesOutput{
		Reservations: reservations,
	}, nil
}

func newMockedEC2ProviderImpl() *ec2ProviderImpl {
	dnsCache := ec2PrivateDNSCache{
		cache: make(map[string]string),
		lock:  sync.RWMutex{},
	}
	ec2Requests := ec2Requests{
		set:  make(map[string]bool),
		lock: sync.RWMutex{},
	}
	return &ec2ProviderImpl{
		ec2:                &mockEc2Client{},
		privateDNSCache:    dnsCache,
		ec2Requests:        ec2Requests,
		instanceIdsChannel: make(chan string, maxChannelSize),
	}
}

func TestGetPrivateDNSName(t *testing.T) {
	ec2Provider := newMockedEC2ProviderImpl()
	ec2Provider.ec2 = &mockEc2Client{Reservations: prepareSingleInstanceOutput()}
	go ec2Provider.StartEc2DescribeBatchProcessing()
	dns_name, err := ec2Provider.GetPrivateDNSName("ec2-1")
	if err != nil {
		t.Error("There is an error which is not expected when calling ec2 API with setting up mocks")
	}
	if dns_name != "ec2-dns-1" {
		t.Errorf("want: %v, got: %v", "ec2-dns-1", dns_name)
	}
}

func prepareSingleInstanceOutput() []*ec2Types.Reservation {
	reservations := []*ec2Types.Reservation{
		{
			Groups: nil,
			Instances: []ec2Types.Instance{
				{
					InstanceId:     aws.String("ec2-1"),
					PrivateDnsName: aws.String("ec2-dns-1"),
				},
			},
			OwnerId:       nil,
			RequesterId:   nil,
			ReservationId: nil,
		},
	}
	return reservations
}

func TestGetPrivateDNSNameWithBatching(t *testing.T) {
	ec2Provider := newMockedEC2ProviderImpl()
	reservations := prepare100InstanceOutput()
	ec2Provider.ec2 = &mockEc2Client{Reservations: reservations}
	go ec2Provider.StartEc2DescribeBatchProcessing()
	var wg sync.WaitGroup
	for i := 1; i < 101; i++ {
		instanceString := "ec2-" + strconv.Itoa(i)
		dnsString := "ec2-dns-" + strconv.Itoa(i)
		wg.Add(1)
		// This code helps test the batch functionality twice
		if i == 50 {
			time.Sleep(200 * time.Millisecond)
		}
		go getPrivateDNSName(ec2Provider, instanceString, dnsString, t, &wg)
	}
	wg.Wait()
}

func getPrivateDNSName(ec2provider *ec2ProviderImpl, instanceString string, dnsString string, t *testing.T, wg *sync.WaitGroup) {
	defer wg.Done()
	dnsName, err := ec2provider.GetPrivateDNSName(instanceString)
	if err != nil {
		t.Error("There is an error which is not expected when calling ec2 API with setting up mocks")
	}
	if dnsName != dnsString {
		t.Errorf("want: %v, got: %v", dnsString, dnsName)
	}
}

func prepare100InstanceOutput() []*ec2Types.Reservation {
	var reservations []*ec2Types.Reservation

	for i := 1; i < 101; i++ {
		instanceString := "ec2-" + strconv.Itoa(i)
		dnsString := "ec2-dns-" + strconv.Itoa(i)
		instance := ec2Types.Instance{
			InstanceId:     aws.String(instanceString),
			PrivateDnsName: aws.String(dnsString),
		}
		var instances []ec2Types.Instance
		instances = append(instances, instance)
		res1 := &ec2Types.Reservation{
			Groups:        nil,
			Instances:     instances,
			OwnerId:       nil,
			RequesterId:   nil,
			ReservationId: nil,
		}
		reservations = append(reservations, res1)
	}
	return reservations
}
