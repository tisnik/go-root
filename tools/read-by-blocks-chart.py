#!/usr/bin/env python3

import sys
import csv
import matplotlib.pyplot as plt

# Check if command line argument is specified (it is mandatory).
if len(sys.argv) < 2:
    print("Usage:")
    print("  read-by-blocks-chart.py input_file.csv")
    print("Example:")
    print("  read-by-blocks-chart.py durations.csv")
    sys.exit(1)

# First command line argument should contain name of input CSV.
input_csv = sys.argv[1]

# Try to open the CSV file specified.
with open(input_csv) as csv_input:
    # And open this file as CSV
    csv_reader = csv.reader(csv_input)

    # Skip header
    next(csv_reader, None)

    # Read all rows from the provided CSV file
    durations = [(int(row[0]), int(row[1])) for row in csv_reader]

# Create new graph
x = [i[0] for i in durations]
y = [i[1] for i in durations]

plt.plot(x, y, "b")

# Title of a graph
plt.title("Reading by block with size N")

# Add a label to x-axis
plt.xlabel("Block size")

# Add a label to y-axis
plt.ylabel("Duration time [ns]")

# Set the plot layout
plt.tight_layout()

# And save the plot into raster format and vector format as well
plt.savefig("read-by-block-time.png")
plt.savefig("read-by-block-time.svg")

# Try to show the plot on screen
plt.show()
