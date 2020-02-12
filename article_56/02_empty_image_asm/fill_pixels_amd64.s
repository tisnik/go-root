TEXT ·fillPixels(SB),7,$0
  	MOVQ pix_data+0(FP), CX	 // adresa
  	MOVQ pix_len+8(FP), AX	 // delka
  	XORL DX, DX		 // pocitadlo
	JMP  NEXT                // reseni problemu len(pixels)==0

LOOP:
	MOVB $0xff, 0(CX)(DX*1)  // zapis bajtu
	INCQ DX                  // zvyseni hodnoty pocitadla
NEXT:	CMPQ DX, AX              // porovnani s delkou rezu
	JL LOOP                  // pocitadlo mensi? ok, skok
  	RET			
