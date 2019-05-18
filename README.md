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

* Fork this repository 🍴
* Make your changes
* Open a pull request and ask for review
* 🎉

# License

MIT, see [LICENSE](./LICENSE).
