module sigs.k8s.io/aws-iam-authenticator

go 1.13

require (
	github.com/aws/aws-sdk-go-v2 v1.1.0
	github.com/aws/aws-sdk-go-v2/config v1.1.0
	github.com/aws/aws-sdk-go-v2/credentials v1.1.0
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.0.1
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.1.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.1.0
	github.com/aws/smithy-go v1.0.0
	github.com/gofrs/flock v0.7.0
	github.com/prometheus/client_golang v1.1.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	go.hein.dev/go-version v0.1.0
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/api v0.16.8
	k8s.io/apimachinery v0.16.8
	k8s.io/client-go v0.16.8
	k8s.io/code-generator v0.16.8
	k8s.io/component-base v0.16.8
	k8s.io/sample-controller v0.16.8
)
