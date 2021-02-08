/*
Copyright 2017-2021 by the contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package partitions

var partitionNames = []string{"aws", "aws-cn", "aws-us-gov", "aws-iso", "aws-iso-b"}

var partitions = map[string]interface{}{
	"aws": map[string]interface{}{
		"id": "aws",
		"regions": []string{
			"aws-global",
			"af-south-1",
			"ap-east-1", "ap-northeast-1", "ap-northeast-2", "ap-south-1", "ap-southeast-1", "ap-southeast-2",
			"ca-central-1",
			"eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3",
			"me-south-1",
			"sa-east-1",
			"us-east-1", "us-east-2", "us-west-1", "us-west-2",
			"us-east-1-fips", "us-east-2-fips", "us-west-1-fips", "us-west-2-fips",
		}},
	"aws-cn": map[string]interface{}{
		"id": "aws-cn",
		"regions": []string{
			"cn-north-1", "cn-northwest-1",
		}},
	"aws-us-gov": map[string]interface{}{
		"id": "aws-us-gov",
		"regions": []string{
			"us-gov-east-1", "us-gov-west-1",
			"us-gov-east-1-fips", "us-gov-west-1-fips",
		}},
	"aws-iso": map[string]interface{}{
		"id": "aws-iso",
		"regions": []string{
			"us-iso-east-1",
		}},
	"aws-iso-b": map[string]interface{}{
		"id": "aws-isb-b",
		"regions": []string{
			"us-isob-east-1",
		}},
}

func GetDefaultPartitionsNames() []string {
	return partitionNames
}

func GetDefaultPartitions() map[string]interface{} {
	return partitions
}

func GetRegions(id string) []string {
	if value, ok := partitions[id]; ok {
		return (value.(map[string]interface{}))["regions"].([]string)
	}

	return nil
}

func ValidPartition(id string) bool {
	_, ok := partitions[id]
	return ok
}

func GetDefaultPartitionId() string {
	return "aws"
}
