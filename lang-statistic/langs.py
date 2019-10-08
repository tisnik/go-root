from collections import Counter


def add_langs(langs, items):
    replacement = {"Js":            "JavaScript",
                   "Javascript":    "JavaScript",
                   "Ts":            "TypeScript",
                   "Typescript":    "TypeScript",
                   "Shell":         "Bash",
                   "Ruby on rails": "Ruby",
                   "Php":           "PHP"}

    items = [replacement.get(item.capitalize(), item.capitalize()) for item in items]

    langs.extend(items)


with open("langs.txt") as fin:
    langs = []
    for line in fin.readlines():
        items = [item.strip() for item in line.strip().replace("/", ",").split(",")]
        add_langs(langs, items)
    for lang, cnt in Counter(langs).most_common():
        print(lang, cnt)
