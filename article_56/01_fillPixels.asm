TEXT main.fillPixels(SB) /home/ptisnovs/src/go-root/article_56/01_no_op_filter.go
	for i := 0; i < len(pixels); i++ {
  0x4b87b0		488b442410		MOVQ 0x10(SP), AX	
  0x4b87b5		488b4c2408		MOVQ 0x8(SP), CX	
  0x4b87ba		31d2			XORL DX, DX		
  0x4b87bc		eb07			JMP 0x4b87c5		
		pixels[i] = 255
  0x4b87be		c60411ff		MOVB $0xff, 0(CX)(DX*1)	
	for i := 0; i < len(pixels); i++ {
  0x4b87c2		48ffc2			INCQ DX			
  0x4b87c5		4839c2			CMPQ AX, DX		
  0x4b87c8		7cf4			JL 0x4b87be		
  0x4b87ca		c3			RET			
