import random
from time import sleep


def floordiv(a: int, b: int) -> int:
    return a // b


if __name__ == "__main__":
    while True:
        a, b = 100, random.randint(1, 10)
        res = floordiv(a, b)
        print(f"tracee: floordiv({a},{b}) = {res}")
        sleep(0.5)
