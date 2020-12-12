# Day 10
Swift again. This was a very fun problem to be honest. It's easy to know how to get to the solution, but the difficulty lies on the time you are willing to wait for your code to return something.

## Part 1
Very easy, just sort the input and count the differences. Could be the problem for the first day, if there wasn't a...

## Part 2
Yes, the important part of this problem. We all thought that you could solve this by just counting all the different possible sequences, but the problem said it clearly: `there must be more than a trillion valid ways`, which in computer scientist language translates to `forget brute forcing this`. Thus, I took a notepad and started to write down some small sequences. 

The first thing I noticed was that you didn't care about numbers with a difference of 3, there was only one way to arrange them, so what you really care about is the numbers with a difference of 1. This way, if you take the amount of ways to arrange the different sequences of consecutive numbers individually and then multiply them together, you'd get the result.

I started looking at the different possible sequences with 1, 2, 3 and 4 consecutive numbers and started to see a pattern: `1, 1, 2, 4...`. "Got it!", I thought. The ways to arrange consecutive numbers are multiplied by 2 every number you added! Wrong answer.

Then I added one more number (as I should have done before) and the sequence changed: `1, 1, 2, 4, 7...`. Ah man... Then I noticed... If we consider the middle numbers of a sequence as bits, we can see that those bits translate to the correct amount of ways to arrange them. For example, if there are 5 numbers, that means there are 3 in the middle, which is `111`, which is `7` in binary. The numbers matched! Wrong, that was not it, but the reasoning was pretty close actually. I was not considering the fact that if I had 6 numbers, there could not be a jump of 4 numbers, which would make 15 ways to arrange them impossible.

Then I finally saw it: the tribonacci numbers. `1 + 1 + 2 = 4`, `1 + 2 + 4 = 7`, `2 + 4 + 7 = 13`... That had to be it. Not just that, but I noticed that the next number in the sequence was always going to be the current number\*2 - the number 3 positions before it, thus I could find the solution extremely efficiently. I coded it and... Yes, it worked.

This was a very fun problem and this second part was really challenging, and I'm pretty happy I could find the most efficient solution in such a short amount of time.
