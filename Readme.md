# Telemetry lake

## Getting started

### Start and configure Minio
```
docker compose -f docker-compose.yml up -d 
```
- open http://localhost:9090
- log in with the username and password you set (or the default ones)
- create a "github" bucket
- create an access-key-id and secret-access-key and set them in config/development.yaml as well

### Clickhouse Setup 
Follow the clickhouse quickstart to download, install and start the server
```
https://clickhouse.com/docs/en/install#setup-the-debian-repository
clickhouse-client --port 19000
```
### Configuring Github webhooks
To ingest event from Github, configure an organization or repository webhook.

https://docs.github.com/en/webhooks/using-webhooks/creating-webhooks

Configure the following settings:
- Payload URL: set the ngrok public https url that forwards to your localhost:8080
- Content Type: application/json
- Secret: a 21+ characters secret used for HMAC validation. Set the same value in your config/development.yaml config on the application side
- SSL verification: Enable SSL verification
- Which events: all or select the specific ones
- Active: true

### Start the application
```
go run main.go
ngrok http http://localhost:8080
```

### Querying the events from Clickhouse
```
SELECT *
FROM s3(
             'http://localhost:9000/clickhouse/*.jsonl',
             '${access-key-id_value}', '${secret_access_key_value}',
             'JSONEachRow'
     )
LIMIT 100;
```