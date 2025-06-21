#!/bin/bash
#entrypoint.sh

/docker-entrypoint.sh cassandra -R &

until cqlsh -u cassandra -p test1234 -e "DESC KEYSPACE football_team" > /dev/null 2>&1; do
    echo "Cassandra is unavailable - sleeping"
    sleep 10
done

cqlsh -f /init.cql

fg