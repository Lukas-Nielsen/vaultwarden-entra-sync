# go-vaultwarden-admin

a docker container to to automaticaly invite users to vaultwarden

## azure app registration

- create a single tenant app
- needed permission (microsoft graph > application permission)
  - User.Read.All
  - GroupMember.Read.All

## example

```yaml
services:
  vaultwarden-entra-sync:
    image: ghcr.io/lukas-nielsen/vaultwarden-entra-sync
    container_name: vaultwarden-entra-sync
    restart: unless-stopped
	environment:
		# azure app registration
		- CLIENT_ID=00000000-0000-0000-0000-000000000000
		- TENANT_ID=00000000-0000-0000-0000-000000000000
		- CLIENT_SECRET=topsecret
		# entra group ids, which contains the users to sync
		- GROUPS=00000000-0000-0000-0000-000000000000,00000000-0000-0000-0000-000000000000
		# vaultwarden base url (/admin is appended automaticaly)
		- BASE_URL=https://vw.example.com
		# plaintext vaultwarden admin token
		- ADMIN_TOKEN=supersecret
		# whitelist users that are not synced, eg. manualy created admin user
		- WHITELIST=admin@example.com,test@example.com
		# optional default value
		- CRON=15/* * * * *
```
