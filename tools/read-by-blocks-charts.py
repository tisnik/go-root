#!/usr/bin/env python3

import sys
import csv
import matplotlib.pyplot as plt


def read_csv(filename):
    # Try to open the CSV file specified.
    with open(filename) as csv_input:
        # And open this file as CSV
        csv_reader = csv.reader(csv_input)

        # Skip header
        next(csv_reader, None)

        # Read all rows from the provided CSV file
        durations = [(int(row[0]), int(row[1])) for row in csv_reader]

    # Create new graph
    x = [i[0] for i in durations]
    y = [i[1] for i in durations]

    return x, y


# Check if command line argument is specified (it is mandatory).
if len(sys.argv) < 3:
    print("Usage:")
    print("  read-by-blocks-charts.py input_file.csv input_file.csv")
    print("Example:")
    print("  read-by-blocks-charts.py durations-1.csv durations-100.csv")
    sys.exit(1)

# First command line argument should contain name of input CSV.
input_csv_1 = sys.argv[1]
input_csv_2 = sys.argv[2]

x1, y1 = read_csv(input_csv_1)
x2, y2 = read_csv(input_csv_2)

plt.plot(x1, y1, "b", label="1 reader goroutine")
plt.plot(x2, y2, "r", label="100 reader goroutines")

# Title of a graph
plt.title("Reading by block with size N")

# Add a label to x-axis
plt.xlabel("Block size")

# Add a label to y-axis
plt.ylabel("Duration time [ns]")

# Add a legend
plt.legend()

# Set the plot layout
plt.tight_layout()

# And save the plot into raster format and vector format as well
plt.savefig("read-by-block-time.png")
plt.savefig("read-by-block-time.svg")

# Try to show the plot on screen
plt.show()
