# Cloudflare Dynamic DNS
[![](https://images.microbadger.com/badges/image/hugomd/cloudflare-ddns.svg)](https://microbadger.com/images/hugomd/cloudflare-ddns "Get your own image badge on microbadger.com") [![](https://images.microbadger.com/badges/version/hugomd/cloudflare-ddns.svg)](https://microbadger.com/images/hugomd/cloudflare-ddns "Get your own version badge on microbadger.com") 

> Updates a given Cloudflare DNS record with your current IP

```
docker run \
  -e APIKEY=YOUR_API_KEY \
  -e ZONE=YOUR_ZONE \
  -e HOST=YOUR_DOMAIN \
  -e EMAIL=YOUR_CLOUDFLARE_EMAIL \
  hugomd/cloudflare-ddns
```

# Environment Variables

| Environment Variable | Description                                                                                                             | Example                 | Required |
|----------------------|-------------------------------------------------------------------------------------------------------------------------|-------------------------|----------|
| `APIKEY`             | [Cloudflare API key](https://support.cloudflare.com/hc/en-us/articles/200167836-Where-do-I-find-my-Cloudflare-API-key-) | `12345`                 | `true`   |
| `ZONE`               | [Cloudflare Zone](https://api.cloudflare.com/#zone-properties)                                                          | `example.com`           | `true`   |
| `HOST`               | The record you want to update                                                                                  | `subdomain.example.com` | `true`   |
| `EMAIL`              | Email associated with your Cloudflare account                                                                           | `john.doe@example.com`  | `true`   |

# Contributing
* Fork this repository 🍴
* Make your changes
* Open a pull request and ask for review
* 🎉

# License
MIT, see [LICENSE](./LICENSE).
