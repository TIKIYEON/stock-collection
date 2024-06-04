import os
import pandas as pd

# Directory where the data files are stored
DATA_DIR = 'data/'

# Initialize the starting sid
sid = 1

# Iterate through all files in the directory
for filename in os.listdir(DATA_DIR):
    if filename.endswith('.csv'):
        file_path = os.path.join(DATA_DIR, filename)
        
        # Read the CSV file into a DataFrame
        df = pd.read_csv(file_path)
        
        # Add the 'sid' column with the current sid value
        df['sid'] = sid
        
        # Save the modified DataFrame back to the CSV file
        df.to_csv(file_path, index=False)
        
        print(f"Processed {filename} with sid {sid}")
        
        # Increment the sid for the next file
        sid += 1

print("All CSV files have been updated.")