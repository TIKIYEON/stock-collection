#!/bin/bash

# Data directory
DATA_DIR="/db/data"

# Go through all the files in the data directory
for file in "$DATA_DIR"/*.csv; do
    # Extraxt sid from the file name
    stock_id=$(basename "$file" .csv)

    # Insert the stock ids into the Stocks table
    psql -U postgres -d stocks -c "INSERT INTO Stocks (stock_id) VALUES ('$stock_id') ON CONFLICT DO NOTHING;"

    # Load the data from the csv files into the StockElements table
    psql -U postgres -d stocks -c "\copy StockElements(date, open, high, low, close, adj_close, volume) FROM '$file' DELIMITER ',' CSV HEADER;"
done
