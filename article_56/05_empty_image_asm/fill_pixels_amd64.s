TEXT ·fillPixels(SB),7,$0
  	MOVQ pix_data+0(FP), DI	 // adresa
  	MOVQ pix_len+8(FP), CX	 // delka
	SHRQ $2, CX              // delime ctyrmi
	MOVD $0xffffffff, AX     // zapisovana barva pixelu (RGBA)

        CLD                      // smer zapisu
        REP                      // opakovani CX-krat
	STOSL                    // zapis ctyrbajtoveho slova
  	RET			
