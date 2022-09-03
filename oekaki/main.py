import mybezier
from timethis import *

a = {}


# @timethis
def main1():
    return mybezier.genWeights(4)

# @timethis
def main2():
    if a.get((4,100)) is  None:
        a[(4,100)] = mybezier.genWeights(4)
    return a[(4,100)]

@timethis
def main():
    for i in range(10000):
        b = main1()
if __name__ == '__main__':
    main()