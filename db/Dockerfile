FROM postgres:16.3-bookworm
ENV POSTGRES_PASSWORD=postgres

WORKDIR /db 

COPY ./data ./data

COPY ./stock.sql /docker-entrypoint-initdb.d/

COPY ./seed.py .

ENTRYPOINT [ "docker-entrypoint.sh" ]

CMD ["bash", "-c", "psql -U postgres -f /docker-entrypoint-initdb.d/stock.sql && python3 /db/seed.py"]
