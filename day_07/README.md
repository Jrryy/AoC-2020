# Day 7
I got go again. Sweet, I sure love using it!

For this problem I first thought that we were supposed to build an N-ary tree. However, some roots could have the same children so that was not possible, but the idea of building a directed graph still looked appealing to me.
 
So I did that, I created a map to represent it and mapped in each bag all the bags that it contained.

## Part 1
I went recursively through all the bags and the ones inside them and counted how many times I found the shiny gold one.
I didn't bother to optimize this at all in the first version, but it could be very easily done by parallelizing this process and dividing the graph, or simply obtaining the "root" bag, which is the one that is not inside any other, assuming that there is one and only one that follows this rule, and counting how many bags inside of it we need to go through until finding the shiny gold.

## Part 2
This part was very fast with the graph. We can recursively calculate how many bags are inside the shiny gold one because we already have the amount of other bags mapped in the bag structure.