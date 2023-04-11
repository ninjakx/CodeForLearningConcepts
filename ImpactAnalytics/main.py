import argparse
from implementation import Implementation

def main():
    parser = argparse.ArgumentParser(
            description='Takes the no of days and consecutive days allowed to miss the classes')
    parser.add_argument(
        "--n",
        help="Number of academic days)",
        default=10,
        type=int
    )
    parser.add_argument(
        "--m",
        help="No of the consecutive days you can miss the classes)",
        default=10,
        type=int
    )
    args = parser.parse_args()
    n = args.n
    m = args.m
    impl = Implementation(n, m) 
    print(impl.result())

if __name__ == "__main__":
    main()

