package common

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/require"
)

func TestIamPolicy(t *testing.T, ctx types.TestContext) {
	iamClient := iam.NewFromConfig(GetAWSConfig(t))

	t.Run("TestDoesSinglePolicyExist", func(t *testing.T) {
		ctx.EnabledOnlyForTests(t, "single_policy")

		policyArn := terraform.Output(t, ctx.TerratestTerraformOptions(), "policy_arn")
		policyName := terraform.Output(t, ctx.TerratestTerraformOptions(), "policy_name")

		output := GetPolicy(t, iamClient, policyArn)
		require.Equal(t, policyName, *output.Policy.PolicyName, "Expected name to be %s, got %s", policyName, *output.Policy.PolicyName)
	})

	t.Run("TestDoMultiplePoliciesExist", func(t *testing.T) {
		ctx.EnabledOnlyForTests(t, "multiple_policies")

		policyArns := terraform.OutputList(t, ctx.TerratestTerraformOptions(), "policy_arns")
		policyNames := terraform.OutputList(t, ctx.TerratestTerraformOptions(), "policy_names")

		for i, policyArn := range policyArns {
			output := GetPolicy(t, iamClient, policyArn)
			require.Equal(t, policyNames[i], *output.Policy.PolicyName, "Expected name to be %s, got %s", policyNames[i], *output.Policy.PolicyName)
		}
	})
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}

func GetPolicy(t *testing.T, c *iam.Client, policyArn string) *iam.GetPolicyOutput {
	output, err := c.GetPolicy(context.TODO(), &iam.GetPolicyInput{
		PolicyArn: &policyArn,
	})
	require.NoErrorf(t, err, "unable to get policy, %v", err)
	return output
}
