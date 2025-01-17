package tags

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
)

func IamResourceHasTag(iamTags []*iam.Tag, tagKey string, tagValue string) bool {
	for _, tag := range iamTags {
		if aws.StringValue(tag.Key) == tagKey && aws.StringValue(tag.Value) == tagValue {
			return true
		}
	}

	return false
}

func Ec2ResourceHasTag(ec2Tags []*ec2.Tag, tagKey string, tagValue string) bool {
	for _, tag := range ec2Tags {
		if aws.StringValue(tag.Key) == tagKey && aws.StringValue(tag.Value) == tagValue {
			return true
		}
	}

	return false
}
