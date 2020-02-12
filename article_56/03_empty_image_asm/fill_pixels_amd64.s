TEXT ·fillPixels(SB),7,$0
  	MOVQ pix_data+0(FP), CX	 // adresa
  	MOVQ pix_len+8(FP), AX	 // delka
  	XORL DX, DX		 // offset

LOOP:
	MOVB $0xff, 0(CX)(DX*1)  // zapis bajtu
	INCQ DX                  // zvyseni hodnoty offsetu
	DECQ AX                  // zmenseni pocitadla
	JNZ LOOP                 // pocitadlo vetsi nez 0? ok, skok
  	RET			
