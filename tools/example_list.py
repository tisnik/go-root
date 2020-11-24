#!/usr/bin/python3
# vim: set fileencoding=utf-8

from pathlib import Path
from sys import exit
import subprocess


def main():
    files = list(Path(".").rglob("*.go"))
    files.sort()

    for filename in files:
        git_cmd = "git log --pretty=oneline --abbrev-commit " + str(filename)
        p = subprocess.Popen(git_cmd,
                             stdout=subprocess.PIPE,
                             shell=True)

        (output, err) = p.communicate()
        p_status = p.wait()

        if p_status == 0:
            lines = output.decode('utf-8').split("\n")
            commit = lines[-2]
            space = commit.find(" ")
            commit = commit[space:].strip()

            colon = commit.find(": ")
            if colon >= 0:
                commit = commit[colon+1:].strip()

            commit = commit[0].upper() + commit[1:]
            print("1. [{}]({})".format(commit, filename))


if __name__ == "__main__":
    main()
