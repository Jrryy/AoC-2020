# Day 19

I gave up using the randomizer and from now on I'm only using Go.
 
This one challenge wasn't especially difficult to understand or code, but I struggled a little with its second part.
I'm using a `[]int` as a stack, but for some reason, using the end of this list as the top part of the stack was giving me a lot of trouble. Instead, I used the beginning as the top and solved the problem immediately.

Also, I figured it could be a good time to use some parallelism, using a goroutine to check each string. It works perfectly fine, but unfortunately this doesn't improve this code's execution time as it's already fast enough without parallelism and synchronization. Still, I used it. Good enough for me.