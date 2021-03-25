docker pull yandex/clickhouse-server

docker run -p 9000:9000 -d --name some-clickhouse-server --ulimit nofile=262144:262144 yandex/clickhouse-server
