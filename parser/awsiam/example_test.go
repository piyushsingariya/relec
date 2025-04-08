package policy

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Example_newBucketPolicy() {
	p := Policy{
		Version: VersionLatest,
		Id:      "CloudTrailBucketPolicy",
		Statements: NewStatements([]Statement{
			{
				Sid:       "AWSCloudTrailWrite20150319",
				Effect:    EffectAllow,
				Principal: NewServicePrincipal("cloudtrail.amazonaws.com"),
				Action:    NewStringOrSlice(false, "s3:PutObject"),
				Resource:  NewStringOrSlice(false, "arn:aws:s3:::examplebucket/AWSLogs/123456789012/*"),
				Condition: map[string]map[string]*ConditionValue{
					"StringEquals": {
						"s3:x-amz-acl": NewConditionValueString(true, "bucket-owner-full-control"),
					},
				},
			},
			{
				Sid:       "AWSCloudTrailAclCheck20150319",
				Effect:    EffectAllow,
				Principal: NewServicePrincipal("cloudtrail.amazonaws.com"),
				Action:    NewStringOrSlice(true, "s3:GetBucketAcl"),
				Resource:  NewStringOrSlice(true, "arn:aws:s3:::examplebucket"),
			},
		}...),
	}
	out, _ := json.MarshalIndent(p, "", "\t")
	fmt.Println(string(out))
	// Output:
	//{
	//	"Id": "CloudTrailBucketPolicy",
	//	"Statement": [
	//		{
	//			"Action": [
	//				"s3:PutObject"
	//			],
	//			"Condition": {
	//				"StringEquals": {
	//					"s3:x-amz-acl": "bucket-owner-full-control"
	//				}
	//			},
	//			"Effect": "Allow",
	//			"Principal": {
	//				"Service": "cloudtrail.amazonaws.com"
	//			},
	//			"Resource": [
	//				"arn:aws:s3:::examplebucket/AWSLogs/123456789012/*"
	//			],
	//			"Sid": "AWSCloudTrailWrite20150319"
	//		},
	//		{
	//			"Action": "s3:GetBucketAcl",
	//			"Effect": "Allow",
	//			"Principal": {
	//				"Service": "cloudtrail.amazonaws.com"
	//			},
	//			"Resource": "arn:aws:s3:::examplebucket",
	//			"Sid": "AWSCloudTrailAclCheck20150319"
	//		}
	//	],
	//	"Version": "2012-10-17"
	//}
}

/*
This is an example of how to use the strict JSON decoding functionality to
ensure that the JSON document is valid according to this package.

The strict decoding functionality is enabled by default on Statements
because it implements a custom UnmarshalJSON method. If you are modifying an
existing policy document, or if you are using a policy document that you did
not create, you can enable strict decoding by setting the DisallowUnknownFields
field on a custom json.Decoder.

This example will fail because the JSON document has a hypothetical new field
"Foo" that is not part of the IAM policy statement grammar.
*/
func Example_strictJSONDecoding() {

	invalidPolicyJSON := []byte(`{
		"Id": "CloudTrailBucketPolicy",
		"Foo": "hypothetical new field",
		"Statement": [
			{
				"Sid": "AWSCloudTrailWrite20150319",
				"Effect": "Allow",
				"Principal": {
					"Service": "cloudtrail.amazonaws.com"
				},
				"Action": "s3:PutObject",
				"Resource": "arn:aws:s3:::examplebucket/AWSLogs/123456789012/*"
			}
		]
	}`)
	var p Policy
	decoder := json.NewDecoder(bytes.NewBuffer(invalidPolicyJSON))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&p)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// json: unknown field "Foo"
}
