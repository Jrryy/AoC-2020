# Day 18
I got python for day 18, and this made me save a LOT of time (it's worth noticing that the challenge was pretty easy too).

For the first part I just made a function that added or multiplied numbers to a total. This function would also be used recursively for expressions inside parentheses.

For the second part I needed to add some priority and I did it in the simplest possible way: append the numbers to a list if they are not going to be added. Then, when we get to the end of the expression, reduce the list multiplying all the numbers.