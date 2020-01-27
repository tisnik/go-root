TEXT ·add(SB),$0
    MOVW 0x8(R13), R0
    MOVW 0x4(R13), R1
    ADD R0, R1, R0
    MOVW R0, 0xc(R13)
    ADD $0, R14, R15
