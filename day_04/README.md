# Day 4
My language for today was Ruby. I have to say I was pretty happy because I thought it was a language pretty similar to python thus it was going to be easy to learn, but boy was this an experience.

It's easier to use than purely compiled languages, that's for sure, it has some pretty cool features, it has similarities with other scripting languages... but as a python engineer I can clearly see the many things my main language does better, or lets you do in a way easier way.

## Part 1
Pretty easy. I built a Hash (Ruby's name for a hashmap, or a dictionary) with the fields I cared about initialized as false. Then I just go line by line, setting as true the ones that I find. Whenever I find a blank line, or the EOF, I check if all the values in the Hash are true. If so, I add 1 to a counter.

## Part 2
Here it got a little complicated, I had to change my code a little bit and search up for how to use methods and regex. The Hash's entries were now going to be initialized as empty strings, and for the counter of the part 1 I would just check that the value was not an empty string. If this test didn't pass, I wouldn't have to check the values for the part 2; if it did, I executed the method that checked if the value was correctly formatted. That method is basically a Ruby's switch in which, depending on the key, I would try and match the value against a regex, and for some of them also check if they belonged to a certain range of integers (you can do `10.between?(5, 20)`, how cool is that?), and that was pretty much it.
