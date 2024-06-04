import pandas as pd
from sqlalchemy import create_engine, text
import os

# Create a connection to the database
#db_coonection_string = 'postgresql://postgres:postgres@db/stocks'
db_connection_string = f"postgresql://{os.getenv('DB_USER')}:{os.getenv('DB_PASSWORD')}@{os.getenv('DB_HOST')}:{os.getenv('DB_PORT')}/stocks"

# Create database engine
engine = create_engine(db_connection_string)


# Function to populate the Stocks table with 77 entries
def populate_stocks_table(engine):
    with engine.connect() as connection:
        # SQL command to insert 77 rows with default values
        insert_command = text("""
        DO $$
        BEGIN
            FOR i IN 1..77 LOOP
                INSERT INTO Stocks DEFAULT VALUES;
            END LOOP;
        END $$;
        """)
        connection.execute(insert_command)
        print("77 entries have been inserted into the Stocks table.")

# Populate the Stocks table
populate_stocks_table(engine)

# Read the CSV file into a DataFrame
csv_file = 'combined.csv'
df = pd.read_csv(csv_file)

# Convert columns to appropriate data types
df['date'] = pd.to_datetime(df['Date'])
df['open'] = df['Open'].astype(float).round(16)
df['high'] = df['High'].astype(float).round(16)
df['low'] = df['Low'].astype(float).round(16)
df['close'] = df['Close'].astype(float).round(16)
df['adj_close'] = df['Adj Close'].astype(float).round(16)
df['volume'] = df['Volume'].astype(int)
df['stock_id'] = df['sid'].astype(int)

# Select only the necessary columns
df = df[['date', 'open', 'high', 'low', 'close', 'adj_close', 'volume', 'stock_id']]

# Insert data into the SQL table
df.to_sql('stockelements', engine, if_exists='append', index=False)

print("Data has been inserted into the StockElements table.")