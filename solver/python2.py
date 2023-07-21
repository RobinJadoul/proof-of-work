#!/usr/bin/env python3
import hashlib
import sys

def main():
    prefix = sys.argv[1]
    difficulty = int(sys.argv[2])
    bound = 1 << (64 - difficulty)

    i = 0
    while True:
        i += 1
        s = prefix + str(i)
        if int(hashlib.sha256(s.encode()).hexdigest()[:16], 16) < bound:
            print(i)
            exit(0)

if __name__ == "__main__":
    main()
