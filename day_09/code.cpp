#include <iostream>
#include <fstream>

using namespace std;

int main() {
    int x;
    long numbers[1000];
    int i = 0;

    ifstream infile("input.txt");

    while (infile >> x) {
        numbers[i] = x;
        i++;
    }

    long n, m;
    __asm__ __volatile__(
        // Declare initial pointer position
            "movq $24, %%r8;"               // pointer = 24
        
        // Outer loop
        "loop:"
        
        // Initialize iterators
            "addq $1, %%r8;"                // pointer++
            "movq %%r8, %%rcx;"             // i = pointer
            "subq $25, %%rcx;"              // i -= 25
            "movq %%rcx, %%rbx;"            // j = i
        
        // Sum the two numbers
        "operations:"
            "movq (%2, %%rcx, 8), %%r9;"    // x = array[i]
            "addq (%2, %%rbx, 8), %%r9;"    // x += array[j]
        
        // Check if they equal array[pointer]
            "cmpq (%2, %%r8, 8), %%r9;"     // compare x to array[pointer]
            "je loop;"                      // if x == array[pointer] jump to loop
        
        // Add 1 to inner iterator and check if it's in the bounds
            "addq $1, %%rbx;"               // j++
            "cmp %%r8, %%rbx;"              // compare j to pointer
            "jne operations;"               // if j != pointer jump to operations
        
        // Add 1 to outer iterator and check if it's in the bounds
            "addq $1, %%rcx;"               // i++
            "movq %%rcx, %%rbx;"            // j = i
            "cmp %%r8, %%rbx;"              // compare j to pointer
            "jne operations;"               // if j != pointer jump to operations
        
        // We've found the sum, put the number in the first returned variable
            "mov (%2, %%r8, 8), %0;"        // return array[pointer]
        
        // Initialize boundaries
            "movq $0, %%rcx;"               // i = 0
            "movq $1, %%rbx;"               // j = 1
        
        // Loop to sum the elements in the boundaries
        "sumlist:"
        
        // Initialize sum, pointer and the max and min numbers
            "movq $0, %%r10;"               // sum = 0
            "movq %%rcx, %%r11;"            // ptr = i
            "movq $0x0fffffff, %%r12;"      // min = maxint
            "movq $0, %%r13;"               // max = 0
        
        // Loop inside the boundaries adding numbers
        "keepadding:"
            "addq (%2, %%r11, 8), %%r10;"   // sum += array[ptr]
            "cmp %0, %%r10;"                // compare n to sum
            "ja toobig;"                    // if sum > n jump to toobig
        
        // Check if the current number is the smallest
            "cmp (%2, %%r11, 8), %%r12;"    // compare array[ptr] to min
            "jbe notmin;"                   // jump if array[ptr] >= min
            "movq (%2, %%r11, 8), %%r12;"   // min = array[ptr]
        "notmin:"
        
        // Check if the current number is the biggest
            "cmp (%2, %%r11, 8), %%r13;"   // compare array[ptr] to max
            "jae notmax;"                   // jump if array[ptr] <= max
            "movq (%2, %%r11, 8), %%r13;"   // max = array[ptr]
        "notmax:"
        
        // Next number and check if it's between the boundaries
            "addq $1, %%r11;"               // ptr++
            "cmp %%r11, %%rbx;"             // compare ptr to j
            "jae keepadding;"               // if ptr <= j jump to keepadding
        
        // Check if the sum equals the number we found
            "cmpq %0, %%r10;"               // compare n to sum
            "je exit;"                      // if n == sum jump to exit
        
        // If not equals, add to the iterators
            "addq $1, %%rbx;"               // j++
            "cmp $1000, %%rbx;"             // compare j to 1000
            "jne sumlist;"                  // if j != 1000 jump to operations
        "toobig:"
            "addq $1, %%rcx;"               // i++
            "movq %%rcx, %%rbx;"            // j = i
            "addq $1, %%rbx;"               // j++
            "cmp $1000, %%rcx;"             // compare i to 1000
            "jmp sumlist;"                  // if i != 1000 jump to operations
        
        // Sum the min and max and exit
        "exit:"
            "movq %%r12, %1;"               // m = min
            "addq %%r13, %1;"               // m += max
        :"=r" (n), "=r" (m)
        :"r" (numbers)
    );
    
    cout << n << endl;
    cout << m << endl;
}
