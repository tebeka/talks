; calculate factorical 
; inptut n in ax
; return value in ax
proc fact
    push cx

    move cx, ax ; place n-1 in loop register
    dec cx
fact_loop:
    imul ax, cx
    loop fact_loop

    pop cx
    ret
endp fact
