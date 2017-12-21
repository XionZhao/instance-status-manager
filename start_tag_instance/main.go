package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	tagkey   string = "Env"
	tagvalue string = "SystemTest"
)

func getamiid() []string {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("cn-north-1"),
	}))
	svc := ec2.New(sess)

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag-key"),
				Values: []*string{
					aws.String(tagkey),
				},
			},
			{
				Name: aws.String("tag-value"),
				Values: []*string{
					aws.String(tagvalue),
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

func startinstance() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("cn-north-1"),
	}))
	svc := ec2.New(sess)

	id := getamiid()

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
	fmt.Println(results.StartingInstances)
}

func main() {
	startinstance()
}
