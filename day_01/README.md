# Day 1
I got PHP :) welp, better now than halfway through.

We meet again old friend. Unfortunately I never used it for scripting nor coding algorithms so I had to read through some documentation before developing my solution. It was kinda easy although I forgot how verbose some languages can be.

## Part 1
I tried to save some time by first sorting the array of numbers, then I check the sum of the first one in the list and the last one:
* If it's 2020, I print their product.
* If it's < 2020 I use the one before the last one.
* If it's > 2020, I start again with the second one in the list.

And so on. This way, the algorithm is not O(n^2) anymore but rather... O(nlogn) I guess? Assuming that PHP's built-in sorting algorithms are efficient enough.

## Part 2
I could have made this a little more efficient, but... what's the point anyways? Let's just add another iterator and move on.
