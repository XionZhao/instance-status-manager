package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func getamiid(key, val string) []string {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("cn-north-1"),
	}))
	svc := ec2.New(sess)

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag-key"),
				Values: []*string{
					aws.String(key),
				},
			},
			{
				Name: aws.String("tag-value"),
				Values: []*string{
					aws.String(val),
				},
			},
		},
	}
	results, err := svc.DescribeInstances(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}

	id := make([]string, 0)
	for i := 0; i < len(results.Reservations); i++ {
		val := *(results.Reservations[i].Instances[0].InstanceId)
		id = append(id, val)
	}
	return id
}

func startinstance(key, val string) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("cn-north-1"),
	}))
	svc := ec2.New(sess)

	id := getamiid(key, val)

	input := &ec2.StartInstancesInput{
		InstanceIds: aws.StringSlice(id),
	}
	results, err := svc.StartInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	for i := 0; i < len(results.StartingInstances); i++ {
		if *(results.StartingInstances[i].PreviousState.Name) == "stopped" {
			fmt.Printf("%v The running\n", *(results.StartingInstances[i].InstanceId))
		} else {
			fmt.Printf("%v Status is %v\n", *(results.StartingInstances[i].InstanceId), *(results.StartingInstances[i].PreviousState.Name))
		}
	}
}

func stopinstance(key, val string) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("cn-north-1"),
	}))
	svc := ec2.New(sess)

	id := getamiid(key, val)

	input := &ec2.StopInstancesInput{
		InstanceIds: aws.StringSlice(id),
	}
	results, err := svc.StopInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	for i := 0; i < len(results.StoppingInstances); i++ {
		if *(results.StoppingInstances[i].PreviousState.Name) == "stopped" {
			fmt.Printf("%v Is stopped\n", *(results.StoppingInstances[i].InstanceId))
		} else {
			fmt.Printf("%v Status is %v\n", *(results.StoppingInstances[i].InstanceId), *(results.StoppingInstances[i].PreviousState.Name))
		}
	}
}

func main() {

	status := flag.String("status", "", "Select start or stop")
	key := flag.String("tagkey", "", "Choose to stop or start ec2 tag key")
	value := flag.String("tagvalue", "", "Choose to stop or start ec2 tag value")

	flag.Parse()

	if *status == "start" {
		startinstance(*key, *value)
	} else if *status == "stop" {
		stopinstance(*key, *value)
	} else {
		fmt.Println("Please Input [start|stop]")
	}
}

