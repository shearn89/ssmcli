package internal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	log "github.com/sirupsen/logrus"
)

var (
	ec2Svc *ec2.EC2
)

func init() {
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	// Load session from shared config, sess declared in api.go
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}

func describeInstances() (*ec2.DescribeInstancesOutput, error) {
	// Create new EC2 client
	ec2Svc = ec2.New(sess)

	// Call to get detailed information on each instance
	result, err := ec2Svc.DescribeInstances(nil)
	return result, err
}

func getNameTag(data []*ec2.Tag) string {
	for _, tag := range data {
		if aws.StringValue(tag.Key) == "Name" {
			return aws.StringValue(tag.Value)
		}
	}
	return ""
}

func parseInstances(data *ec2.DescribeInstancesOutput) map[string]string {
	output := make(map[string]string)
	for _, reservation := range data.Reservations {
		for _, instance := range reservation.Instances {
			name := getNameTag(instance.Tags)
			id := aws.StringValue(instance.InstanceId)
			log.WithFields(log.Fields{"id": id, "name": name}).Debug("instance found")
			output[name] = id
		}
	}
	return output
}

func GetInstances() map[string]string {
	data, err := describeInstances()
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr == credentials.ErrNoValidProvidersFoundInChain {
				log.Error("no credentials provider found, do you need to set AWS_* variables?")
			} else {
				log.WithFields(log.Fields{
					"code":    awsErr.Code(),
					"message": awsErr.Message(),
				}).Error("aws error")
			}
		} else {
			log.WithError(err).Error("unable to describe instances")
		}
		return nil
	}
	return parseInstances(data)
}
