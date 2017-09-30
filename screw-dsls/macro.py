def SWAP(r1, r2):
    # Swaps two registers (uses r9)
    MOV(r9, r1)
    MOV(r1, r2)
    MOV(r2, r9)
