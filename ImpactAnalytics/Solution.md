Run Python Code:

``` bash
python3 main.py --n 5 --m 4 
```


In a university, your attendance determines whether you will be
allowed to attend your graduation ceremony.
You are not allowed to miss classes for four or more consecutive days.
Your graduation ceremony is on the last day of the academic year,
which is the Nth day.

  Your task is to determine the following:

1. The number of ways to attend classes over N days.
2. The probability that you will miss your graduation ceremony.

Represent the solution in the string format as "Answer of (2) / Answer
of (1)", don't actually divide or reduce the fraction to decimal

```diff
- Question is wrong
```

1) No. of ways of being absent at the graduation ceremony
2) No. of ways to attend classes over N days
3) probability = (1)/(2)

which is the 2nd point in the task that have been asked.

WHY???
Let say
```
n = 5, m = 4 
```

2) `29 ways` to attend classes over 5 days which are as below

```python3
AAAPA,AAAPP,AAPAA,AAPAP,AAPPA,AAPPP,APAAA,APAAP,APAPA,APAPP,APPAA,APPAP,APPPA,APPPP,PAAAP,PAAPA,PAAPP,PAPAA,PAPAP,PAPPA,PAPPP,PPAAA,PPAAP,PPAPA,PPAPP,PPPAA,PPPAP,PPPPA,PPPPP
```

Now we need to pick only those from the above where I am missing from the graduation day.
`14 ways` are there as follows:

```python3
AAPA,AAPP,APAA,APAP,APPA,APPP,PAAA,PAAP,PAPA,PAPP,PPAA,PPAP,PPPA,PPPP 
```
(remaining 4 days which belongs to above set)

We are marking ourself absent on that day so `m -> m-1` (take atleast 3 consecutive days now as we don't want to include the invalid cases)
// last day has been marked as "A"

```diff
- APPROACH
```

**RECURSION:**

In order to find the 1) and 2) 

We will have two choices:

- choice 1 : miss the class (only if possible i.e if not having >=4 consecutive days of absent)
- choice 2 : don't miss the class

if we are able to generate the pattern return 1 for all the cases and add them up to give the total

so time complexity is O(2^n) as having two choices for n days

```
Time complexity: O(2^N)
Space complexity: O(N)
```

for `1st part` we will take `n` days and allow limit to be `m`

for 2nd part we will take `n-1` days and allow limit to be `m-1`

as we want to avoid failure cases + we are absent on graduation day.

**RECURSION + MEMOIZATION**

```
Time complexity: O(N^2)
Space complexity: O(N^2)
```






