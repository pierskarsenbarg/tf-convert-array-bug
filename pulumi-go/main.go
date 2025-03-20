package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		secretData := map[string]interface{}{
			"one": map[string]interface{}{
				"name":  "secret_one",
				"value": "one",
			},
			"two": map[string]interface{}{
				"name":  "secret_two",
				"value": "two",
			},
			"three": map[string]interface{}{
				"name":  "secret_three",
				"value": "three",
			},
		}
		var pkSecrets []*secretsmanager.Secret
		for key0, _ := range secretData {
			__res, err := secretsmanager.NewSecret(ctx, fmt.Sprintf("pk_secrets-%v", key0), &secretsmanager.SecretArgs{
				Name:                 pulumi.Sprintf("mysecret-%v", key0),
				RecoveryWindowInDays: pulumi.Int(0),
			})
			if err != nil {
				return err
			}
			pkSecrets = append(pkSecrets, __res)
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"name":  secretData[key0].Name,
			"value": secretData[key0].Value,
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		var pkSecretsVersion []*secretsmanager.SecretVersion
		for key0, val0 := range pkSecrets {
			__res, err := secretsmanager.NewSecretVersion(ctx, fmt.Sprintf("pk_secrets_version-%v", key0), &secretsmanager.SecretVersionArgs{
				SecretId:     pulumi.String(val0),
				SecretString: pulumi.String(json0),
			})
			if err != nil {
				return err
			}
			pkSecretsVersion = append(pkSecretsVersion, __res)
		}
		return nil
	})
}
