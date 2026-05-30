import os

directories = sorted(os.listdir())

files = sorted(os.listdir())

for file in files:
    if file.endswith(".go") and file != "00_index.go":
        print("// " + file + ":")
        title = False
        title_written = False
        with open(file, "r") as fin:
            for line in fin:
                line = line.strip()
                if title_written:
                    if line == "//":
                        break
                if title:
                    print(line)
                    title_written = True
                if line.startswith("// Demonstrační příklad"):
                    title = True
        print("//")

for entry in os.scandir("."):
    if entry.is_dir():
        directory=entry
        files = sorted(os.listdir(directory))
        for file in files:
            if file.endswith(".go") and file != "00_index.go":
                path=str(directory.path)+"/"+file
                print("// " + path[2:] + ":")
                title = False
                title_written = False
                with open(path, "r") as fin:
                    for line in fin:
                        line = line.strip()
                        if title_written:
                            if line == "//":
                                break
                        if title:
                            print(line)
                            title_written = True
                        if line.startswith("// Demonstrační příklad"):
                            title = True
                print("//")
