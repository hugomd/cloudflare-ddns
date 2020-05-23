# Dynamic DNS
[![](https://images.microbadger.com/badges/image/hugomd/cloudflare-ddns.svg)](https://microbadger.com/images/hugomd/cloudflare-ddns "Get your own image badge on microbadger.com") [![](https://images.microbadger.com/badges/version/hugomd/cloudflare-ddns.svg)](https://microbadger.com/images/hugomd/cloudflare-ddns "Get your own version badge on microbadger.com") 

> Updates a given a DNS record with your current IP

Example with the Cloudflare [provider](#supported-providers):
```
docker run \
  -e PROVIDER=cloudflare \
  -e CLOUDFLARE_APITOKEN=YOUR_API_TOKEN \
  -e CLOUDFLARE_ZONEID=YOUR_ZONE_ID \
  -e CLOUDFLARE_HOST=YOUR_DOMAIN \
  hugomd/cloudflare-ddns:2.0.0
```

Example running as a persistant daemon:
```
docker run -d --restart always \
  -e PROVIDER=cloudflare \
  -e CLOUDFLARE_APITOKEN=YOUR_API_TOKEN \
  -e CLOUDFLARE_ZONEID=YOUR_ZONE_ID \
  -e CLOUDFLARE_HOST=YOUR_DOMAIN \
  hugomd/cloudflare-ddns:2.0.0 -duration 2h
```

You can load environment variables through a config file of key/value pairs.

```sh
echo "PROVIDER=YOUR_PROVIDER" > config.env
docker run \
  -v $PWD/config.env:/tmp/config.env \
  hugomd/cloudflare-ddns -config /tmp/config.env
```

**A note about Docker image tags**: `latest` currently points to a now deprecated version of `cloudflare-ddns`, please use versioned tags e.g. `hugomd/cloudflare-ddns:2.0.0`.

# Supported Providers

| Provider                             | Reference (used for `PROVIDER` environment variable) |
|--------------------------------------|------------------------------------------------------|
| [Cloudflare](https://cloudflare.com) | `cloudflare`                                         |

# CLI

| Parameter             | Description                                                                                                                                                                | Example           | Required |
|-----------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------|----------|
| `-duration`           | Runs program perpetually and recheck after specified interval; parses time strings such as `5m`, `15m`, `2h30m5s`. If not specified, or if equal to 0s, run once and exit. | 2h                | `false`  |
| `-config`             | Loads environment variables from a given file. Variables should be specified as lines of `key=value` pairs. No variables will be loaded if a file is not specified.        | `/tmp/config.env` | `false`  |

# Environment Variables

All providers require the following environment variable:

| Environment Variable            | Description                             | Example       | Required |
|---------------------------------|-----------------------------------------|---------------|----------|
| `PROVIDER`                      | The name of the provider you wish to use | `cloudflare` | `true`   |

## Cloudflare

| Environment Variable               | Description                                                                                                                                                | Example                 | Required |
| ---------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------- | -------- |
| `CLOUDFLARE_APITOKEN`              | An [API Token](https://support.cloudflare.com/hc/en-us/articles/200167836-Managing-API-Tokens-and-Keys), with permission to edit DNS records for your zone | `12345`                 | `true`   |
| `CLOUDFLARE_ZONEID`                | The Zone ID of your domain in Cloudflare (you can find this in the "Overview" tab at the bottom of the page)                                               | `dd255baaaaad2e8...`    | `true`   |
| `CLOUDFLARE_HOST`                  | The record you want to update                                                                                                                              | `subdomain.example.com` | `true`   |
| `CLOUDFLARE_ZONE` **DEPRECATED**   | [Cloudflare Zone](https://api.cloudflare.com/#zone-properties)                                                                                             | `example.com`           |          |
| `CLOUDFLARE_EMAIL` **DEPRECATED**  | Email associated with your Cloudflare account                                                                                                              | `john.doe@example.com`  |          |
| `CLOUDFLARE_APIKEY` **DEPRECATED** | [Cloudflare API key](https://support.cloudflare.com/hc/en-us/articles/200167836-Where-do-I-find-my-Cloudflare-API-key-)                                    | `12345`                 |          |

### Deprecated Environment Variables

Cloudflare now [supports API tokens](https://blog.cloudflare.com/api-tokens-general-availability/) as a more secure way of interacting with their API. Instead of using your global API key/email, you should use a token with limited permissions.

When upgrading, you'll need to replace a few existing environment variables.

Instead of providing a `CLOUDFLARE_ZONE` with your domain name, you should specify the Zone ID (`CLOUDFLARE_ZONEID`) of your domain. You can find this in the "Overview" tab for your domain.

Instead of your `CLOUDFLARE_EMAIL` and `CLOUDFLARE_APIKEY`, you should [generate a token](https://support.cloudflare.com/hc/en-us/articles/200167836-Managing-API-Tokens-and-Keys#12345680) (`CLOUDFLARE_APITOKEN`) with permission to edit DNS records for your desired zone.

You don't need to make any changes to your `CLOUDFLARE_HOST`.

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
