# Day 16
I finally used Java! I was waiting for this day to come. God is it verbose. But I have to admit it's pretty good, using it again after many years makes you see the beauty it it along with its latest updates that make it better than ever.

## Part 1
To begin with, I made an ArrayList of the ranges of accepted values for each field. Then, for each ticket, ignoring my own, I went number by number checking if it was in any of them, and if it wasn't I'd add it to the sum.

## Part 2
The code I made here was quite inefficient. I'm pretty sure I could have optimized a pair of things but, why would I? I'm already 4 days away from the current challenge, not gonna spend more time on this one. I used a matrix, in which one dimension was each one of the ranges I created in part 1, and the other dimension was the field of the ticket that would be correct in that range. For example, if we have two ranges:
- departure 1-6 or 9-11
- arrival 1-3 or 5-13

and a ticket like `4,10`, the matrix we would have is `[[0, 1], [1]]`, with the 0 missing in the second position because the first field of the tickets wouldn't be correct in that range.

With this, I could create a 20 x 20 matrix and initialize it with all the fields for all the ranges. Then, whenever I find a valid ticket while doing the process of part 1, I would check all the numbers to see which numbers and ranges are not available together, and remove them from the matrix.

After this, I made the bold (but correct, apparently) assumption that there would be at least one range in which only one of the fields would be correct for it. If the range is also one of the first 6, I will use it for the product. Then I remove that field from all other ranges, because this one is not a possibility anymore, and thus at least one other range will have now only one available field, and I start the process again, and so on until the matrix is empty.

Again, all of this could have been improved, but it works fine and that's honestly all that matters to me right now.