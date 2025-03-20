using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var secretData = 
    {
        { "one", 
        {
            { "name", "secret_one" },
            { "value", "one" },
        } },
        { "two", 
        {
            { "name", "secret_two" },
            { "value", "two" },
        } },
        { "three", 
        {
            { "name", "secret_three" },
            { "value", "three" },
        } },
    };

    var pkSecrets = new List<Aws.SecretsManager.Secret>();
    foreach (var range in secretData.Select(pair => new { pair.Key, pair.Value }))
    {
        pkSecrets.Add(new Aws.SecretsManager.Secret($"pk_secrets-{range.Key}", new()
        {
            Name = $"mysecret-{range.Key}",
            RecoveryWindowInDays = 0,
        }));
    }
    var pkSecretsVersion = new List<Aws.SecretsManager.SecretVersion>();
    foreach (var range in pkSecrets.Select((v, k) => new { Key = k, Value = v }))
    {
        pkSecretsVersion.Add(new Aws.SecretsManager.SecretVersion($"pk_secrets_version-{range.Key}", new()
        {
            SecretId = range.Value.Id,
            SecretString = JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                ["name"] = secretData[range.Key].Name,
                ["value"] = secretData[range.Key].Value,
            }),
        }));
    }
});

