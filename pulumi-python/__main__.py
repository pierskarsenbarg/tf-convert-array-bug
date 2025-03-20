import pulumi
import json
import pulumi_aws as aws

secret_data = {
    "one": {
        "name": "secret_one",
        "value": "one",
    },
    "two": {
        "name": "secret_two",
        "value": "two",
    },
    "three": {
        "name": "secret_three",
        "value": "three",
    },
}
pk_secrets = []
for range in [{"key": k, "value": v} for [k, v] in enumerate(secret_data)]:
    pk_secrets.append(aws.secretsmanager.Secret(f"pk_secrets-{range['key']}",
        name=f"mysecret-{range['key']}",
        recovery_window_in_days=0))
pk_secrets_version = []
def create_pk_secrets_version(range_body):
    for range in [{"key": k, "value": v} for [k, v] in enumerate(range_body)]:
        pk_secrets_version.append(aws.secretsmanager.SecretVersion(f"pk_secrets_version-{range['key']}",
            secret_id=range["value"],
            secret_string=json.dumps({
                "name": secret_data[range["key"]]["name"],
                "value": secret_data[range["key"]]["value"],
            })))

pk_secrets.apply(create_pk_secrets_version)
