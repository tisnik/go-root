import os

files = sorted(os.listdir())

for file in files:
    if file.endswith(".go") and file != "00_index.go":
        print("// " + file + ":")
        title = False
        with open(file, "r") as fin:
            for line in fin:
                line = line.strip()
                if title:
                    print(line)
                    break
                if line.startswith("// Demonstrační příklad "):
                    title = True
        print("//")
