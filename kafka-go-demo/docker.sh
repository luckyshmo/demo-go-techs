# docker pull bitnami/kafka

# curl -sSL https://raw.githubusercontent.com/bitnami/bitnami-docker-kafka/master/docker-compose.yml > docker-compose.yml
docker pull wurstmeister/kafka
docker pull wurstmeister/zookeeper
docker pull yandex/clickhouse-server

docker-compose up -d