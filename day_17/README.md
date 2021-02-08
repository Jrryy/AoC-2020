# Day 17
It's been quite some time. After failing to do day 14 in rust I became pretty unmotivated and stopped coding for fun for a while. Now that I am on vacation, I feel like it's time to continue what I started.

I didn't want to choose a new language so I did this in Go. It's very similar to a challenge from before, except this one is in 3 and 4 dimensions, and that one was in 2 dimensions.

No I didn't write all the different possible moves myself, I used python for this purpose (god bless Guido and the open source community):

```python
from itertools import product
moves = (-1, 0, 1)
for m in product(*[moves]*4):
    print(f'{{{m[0]}, {m[1]}, {m[2]}, {m[3]}}},')
```