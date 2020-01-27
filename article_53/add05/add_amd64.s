TEXT ·add(SB),$0
    MOVQ    x+8(SP), AX
    ADDQ    y+16(SP), AX
    MOVQ    AX, 24(SP)
    RET
