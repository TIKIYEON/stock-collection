FROM postgres:16.3-bookworm
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_USER=postgres
ENV POSTGRES_DB=stocks

# Change ownership of the data directory
RUN chown -R postgres:postgres /var/lib/postgresql/data

# Set the user to postgres
USER postgres

WORKDIR /db 

COPY ./data ./data

COPY stock.sql /docker-entrypoint-initdb.d
COPY load_data.sh /docker-entrypoint-initdb.d
