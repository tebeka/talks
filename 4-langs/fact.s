.section .data
.section .text
.globl _start

_start:
    movl $5, %eax   # Send value to compute in eax
    jmp fact
end:
    movl %eax, %ebx # Set result as return value
    movl $1, %eax   # System exit
    int $0x80

fact:
    movl %eax, %ecx  # Loop n -1 times
    decl %ecx
fact_loop:
    imul %ecx, %eax  # n = n * n-1
    loop fact_loop
    jmp end
