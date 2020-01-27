TEXT ·add(SB),$0
    MOVQ    $0x0, 0x18(SP)
    MOVQ    0x8(SP), AX
    ADDQ    0x10(SP), AX
    MOVQ    AX, 0x18(SP)
    RET
