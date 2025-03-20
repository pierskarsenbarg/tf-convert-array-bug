import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const secretData = {
    one: {
        name: "secret_one",
        value: "one",
    },
    two: {
        name: "secret_two",
        value: "two",
    },
    three: {
        name: "secret_three",
        value: "three",
    },
};
const pkSecrets: aws.secretsmanager.Secret[] = [];
for (const range of Object.entries(secretData).map(([k, v]) => ({key: k, value: v}))) {
    pkSecrets.push(new aws.secretsmanager.Secret(`pk_secrets-${range.key}`, {
        name: `mysecret-${range.key}`,
        recoveryWindowInDays: 0,
    }));
}
const pkSecretsVersion: aws.secretsmanager.SecretVersion[] = [];
pkSecrets.apply(rangeBody => {
    for (const range of rangeBody.map((v, k) => ({key: k, value: v}))) {
        pkSecretsVersion.push(new aws.secretsmanager.SecretVersion(`pk_secrets_version-${range.key}`, {
            secretId: range.value.id,
            secretString: JSON.stringify({
                name: secretData[range.key].name,
                value: secretData[range.key].value,
            }),
        }));
    }
});
