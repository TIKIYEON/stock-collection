# Use the official Python image from the Docker Hub
FROM python:3.9-slim

# Set the working directory
WORKDIR /seeder

# Copy the requirements file
COPY requirements.txt ./

# Install the required packages
RUN pip install --no-cache-dir -r requirements.txt

# Copy the Python script and CSV file into the container
COPY ./stockselementsseed.py ./stockselementsseed.py
COPY ./combined.csv ./combined.csv

# Run the Python script
CMD ["python", "./stockselementsseed.py"]
