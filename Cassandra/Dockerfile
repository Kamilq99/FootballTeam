FROM cassandra:5.0.4-jammy

RUN sed -i 's/authenticator: AllowAllAuthenticator/authenticator: PasswordAuthenticator/g' /etc/cassandra/cassandra.yaml
RUN sed -i 's/authorizer: AllowAllAuthorizer/authorizer: CassandraAuthorizer/g' /etc/cassandra/cassandra.yaml

COPY docker-entrypoint.sh /docker-entrypoint.sh

RUN chmod +x /docker-entrypoint.sh

COPY init.cql /init.cql

ENTRYPOINT ["/docker-entrypoint.sh"]

EXPOSE 9042