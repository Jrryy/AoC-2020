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

    long n;
    __asm__ __volatile__(
        "movq $24, %%r8;"               // pointer = 24
        "loop:"
        "addq $1, %%r8;"                // pointer++
        "movq %%r8, %%rcx;"             // i = pointer
        "subq $25, %%rcx;"              // i -= 25
        "movq %%rcx, %%rbx;"            // j = i
        "operations:"
        "movq (%1, %%rcx, 8), %%r9;"    // x = array[i]
        "addq (%1, %%rbx, 8), %%r9;"    // x += array[j]
        "cmpq (%1, %%r8, 8), %%r9;"     // compare x to array[pointer]
        "je loop;"                      // if x == array[pointer] jump to loop
        "addq $1, %%rbx;"               // j++
        "cmp %%r8, %%rbx;"              // compare j to pointer
        "jne operations;"               // if j != pointer jump to operations
        "addq $1, %%rcx;"               // i++
        "movq %%rcx, %%rbx;"            // j = i
        "cmp %%r8, %%rbx;"              // compare j to pointer
        "jne operations;"               // if j != pointer jump to operations
        "mov (%1, %%r8, 8), %0;"        // return array[pointer]
        :"=r" (n)
        :"r" (numbers)
    );
    cout << n << endl;
}
