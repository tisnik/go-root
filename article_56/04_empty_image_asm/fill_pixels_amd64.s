TEXT ·fillPixels(SB),7,$0
  	MOVQ pix_data+0(FP), CX	 // adresa
  	MOVQ pix_len+8(FP), AX	 // delka
  	XORL DX, DX		 // offset
	MOVD $0xffffffff, BX     // zapisovana barva pixelu (RGBA)

LOOP:
	MOVD BX, 0(CX)(DX*1)     // zapis bajtu
	ADDQ $4, DX              // zvyseni hodnoty offsetu
	SUBQ $4, AX              // zmenseni pocitadla
	JNZ LOOP                 // pocitadlo vetsi nez 0? ok, skok
  	RET			
