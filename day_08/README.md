# Day 8
Finally, the "assembly" problem that we all were waiting for, and I have to solve it in JavaScript! Nothing to say about it, I've used it before and it's okay I guess, there's no other choice to develop web front-end so far so I, currently a back-end web developer, must know at least a little bit of JavaScript.

## Part 1
Easy, just execute the operations. We've all done it a hundred times before.

## Part 2
So I have to find which `nop` or `jmp` instruction I have to swap so that the program correctly ends in the last instruction. Well, first of all, in the first part, I will be saving the state of the pointer each time I execute one of those, and in case it's a `nop`, only if the quantity in the instruction doesn't equal to 0 and the addition of that quantity to the pointer doesn't point to an instruction that has already been visited (I feel like this saves some time, not much, but it does). Then for each one of those instructions I swap the corresponding `nop` or `jmp`, perform the execution again, and test if the program does indeed reach the last instruction in the input. 

I'm pretty sure there's a way to optimize this but, once again, the input is not big enough for me to bother, it already takes only 45ms to run, why would I want less.