# Dynamic DNS
[![](https://images.microbadger.com/badges/image/hugomd/cloudflare-ddns.svg)](https://microbadger.com/images/hugomd/cloudflare-ddns "Get your own image badge on microbadger.com") [![](https://images.microbadger.com/badges/version/hugomd/cloudflare-ddns.svg)](https://microbadger.com/images/hugomd/cloudflare-ddns "Get your own version badge on microbadger.com") 

> Updates a given a DNS record with your current IP

Example with the Cloudflare [provider](#supported-providers):
```
docker run \
  -e PROVIDER=cloudflare \
  -e CLOUDFLARE_APIKEY=YOUR_API_KEY \
  -e CLOUDFLARE_ZONE=YOUR_ZONE \
  -e CLOUDFLARE_HOST=YOUR_DOMAIN \
  -e CLOUDFLARE_EMAIL=YOUR_CLOUDFLARE_EMAIL \
  hugomd/cloudflare-ddns
```

# Supported Providers

| Provider                             | Reference (used for `PROVIDER` environment variable) |
|--------------------------------------|------------------------------------------------------|
| [Cloudflare](https://cloudflare.com) | `cloudflare`                                         |

# Environment Variables

All providers require the following environment variable:

| Environment Variable            | Description                             | Example       | Required |
|---------------------------------|-----------------------------------------|---------------|----------|
| `PROVIDER`                      | The name of the provider you wish to use | `cloudflare` | `true`   |

## Cloudflare

| Environment Variable            | Description                                                                                                             | Example                 | Required |
|---------------------------------|-------------------------------------------------------------------------------------------------------------------------|-------------------------|----------|
| `CLOUDFLARE_APIKEY`             | [Cloudflare API key](https://support.cloudflare.com/hc/en-us/articles/200167836-Where-do-I-find-my-Cloudflare-API-key-) | `12345`                 | `true`   |
| `CLOUDFLARE_ZONE`               | [Cloudflare Zone](https://api.cloudflare.com/#zone-properties)                                                          | `example.com`           | `true`   |
| `CLOUDFLARE_HOST`               | The record you want to update                                                                                           | `subdomain.example.com` | `true`   |
| `CLOUDFLARE_EMAIL`              | Email associated with your Cloudflare account                                                                           | `john.doe@example.com`  | `true`   |

# Contributing

## Adding a new provider

To add a new provider:
1. Create a new folder in [`lib/providers`](https://github.com/hugomd/cloudflare-ddns/tree/master/lib/providers), called `your_provider`
2. Create a package for your provider in the previously created folder, `your_provider`.
3. Ensure your provider implements the `Provider` interface
4. Import your provider in [`lib/providers/_all/all.go`](https://github.com/hugomd/cloudflare-ddns/blob/master/lib/providers/_all/all.go)
5. Open a PR 🎉

# License

MIT, see [LICENSE](./LICENSE).
