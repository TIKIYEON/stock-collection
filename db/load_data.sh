#!/bin/bash

# Data directory
DATA_DIR="/db/data"

# Check if the data directory exists
if [ ! -d "$DATA_DIR" ]; then
    echo "Data directory $DATA_DIR does not exist."
    exit 1
fi

# Check if there are any CSV files in the data directory
if [ -z "$(ls -A $DATA_DIR/*.csv 2>/dev/null)" ]; then
    echo "No CSV files found in $DATA_DIR."
    exit 1
fi

# Go through all the files in the data directory
for file in "$DATA_DIR"/*.csv; do
    # Extract stock identifier from the file name
    stock_identifier=$(basename "$file" .csv)

    # Insert the stock identifier into the Stocks table
    psql -U postgres -d stocks -c "INSERT INTO Stocks (sid) VALUES (DEFAULT) RETURNING sid;" > sid.txt

    # Get the sid for the stock identifier
    sid=$(cat sid.txt | xargs)


    # Load the data from the csv files into the StockElements table
    psql -U postgres -d stocks -c "\copy StockElements(date, open, high, low, close, adj_close, volume) FROM '$file' DELIMITER ',' CSV HEADER;"
    # Update the stock_id in the StockElements table
    psql -U postgres -d stocks -c "UPDATE StockElements SET stock_id = $sid WHERE stock_id IS NULL;"
done

# psql -U postgres -d stocks -c "TRUNCATE TABLE StockElements RESTART IDENTITY CASCADE;"