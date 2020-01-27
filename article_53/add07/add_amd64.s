TEXT ·add(SB),$0
    MOVQ    x+0(FP), AX
    ADDQ    y+8(FP), AX
    MOVQ    AX, result+16(FP)
    RET
