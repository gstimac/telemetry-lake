# Telemetry lake

## Getting started

### Start and configure Minio
```
docker compose -f minio-compose.yml up -d 
```
- open http://localhost:9090
- log in with the username and password you set (or the default ones)
- create a "github" bucket
- create an access-key-id and secret-access-key and set them in config/development.yaml as well

### Installing Clickhouse
Follow the clickhouse quickstart to download, install and start the server
```
sudo ./clickhouse install
sudo clickhouse start
cp ./clickhouse/etc/clickhouse-server/config.xml /etc/clickhouse-server/config.xml
clickhouse-client --port 19000

docker run -d -p 18123:8123 -p19000:9000 \
    --name some-clickhouse-server \
    --ulimit nofile=262144:262144 \
    -v $(realpath ./ch_data):/var/lib/clickhouse/ \
    -v $(realpath ./ch_logs):/var/log/clickhouse-server/ \
    -v $(realpath ./clickhouse/etc/clickhouse-server/config.xml):/etc/clickhouse-server/config.xml \
    clickhouse/clickhouse-server:24.6.2.17
echo 'SELECT version()' | curl 'http://localhost:18123/' --data-binary @-

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