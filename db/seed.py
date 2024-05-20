import psycopg2
import os
import csv
import glob

count = 0

# db connection
conn = psycopg2.connect(
        dbname="stocks",
        user="postgres",
        password=os.getenv("POSTGRES_PASSWORD"),
        host="localhost",
        port="5432"
    )

cur = conn.cursor()

# Dir of the csv files
data_dir = "./data"

# Import the csv files
def seed_data(file, sid, stock_id):
    with open(file, 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            cur.execute("""
            INSERT INTO StockElements (sid, date, adj_close, close, open, low, high, stock_id)
            VALUES (%s, %s, %s, %s, %s, %s, %s)
            """, (
                    sid,
                    row['Date'],
                    row['Adj_close'],
                    row['Close'],
                    row['Open'],
                    row['Low'],
                    row['High'],
                    stock_id
            ))
    conn.commit()

# For every file create a new stock and for each row in the file insert the datainto stockelements
for file in glob.glob(f"{data_dir}/*.csv"):
    stock_id = os.path.basename(file).split(".")[0]
    cur.execute("""
    INSERT INTO Stocks (stock_id)
    VALUES (%s)
    """, (stock_id))
    seed_data(file, count, stock_id)
    count += 1

# Close the connection
cur.close()
conn.close()
