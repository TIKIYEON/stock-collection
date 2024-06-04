import pandas as pd
from sqlalchemy import create_engine

# Create a connection to the database
db_coonection_string = 'postgresql://postgres:postgres@db/stocks'

# Create database engine
engine = create_engine(db_coonection_string)

# Read the CSV file into a DataFrame
csv_file = 'db/combined.csv'
df = pd.read_csv(csv_file)

# Convert columns to appropriate data types
df['date'] = pd.to_datetime(df['date'])
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
df.to_sql('StockElements', engine, if_exists='append', index=False)

print("Data has been inserted into the StockElements table.")