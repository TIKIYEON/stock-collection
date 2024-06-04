import os
import pandas as pd

# Directory where the data files are stored
DATA_DIR = 'data/'
# output file
OUTPUT_FILE = 'combined.csv'

# list to store the dataframes
dataframes = []

# loop through the files in the data directory
for file in os.listdir(DATA_DIR):
    # check if the file is a csv file
    if file.endswith('.csv'):
        filepath = os.path.join(DATA_DIR, file)
        # read the csv file into a dataframe
        df = pd.read_csv(filepath)
        # append the dataframe to the list
        dataframes.append(df)

        print(f'Read {file}')

# combine the dataframes into a single dataframe
combined_df = pd.concat(dataframes, ignore_index=True)

# save the combined dataframe to a csv file
combined_df.to_csv(os.path.join(DATA_DIR, OUTPUT_FILE), index=False)

print(f'Combined {len(dataframes)} files into {OUTPUT_FILE}')