TEXT main.copyPixels(SB) /home/ptisnovs/src/go-root/article_56/03_no_op_filter.go
func copyPixels(src []uint8, dst []uint8) {
  0x4b94e0		64488b0c25f8ffffff	MOVQ FS:0xfffffff8, CX	
  0x4b94e9		483b6110		CMPQ 0x10(CX), SP	
  0x4b94ed		0f86bc000000		JBE 0x4b95af		
  0x4b94f3		4883ec60		SUBQ $0x60, SP		
  0x4b94f7		48896c2458		MOVQ BP, 0x58(SP)	
  0x4b94fc		488d6c2458		LEAQ 0x58(SP), BP	
	copy(dst, src)
  0x4b9501		488b842488000000	MOVQ 0x88(SP), AX		
  0x4b9509		488b8c2480000000	MOVQ 0x80(SP), CX		
  0x4b9511		488b942490000000	MOVQ 0x90(SP), DX		
  0x4b9519		48894c2440		MOVQ CX, 0x40(SP)		
  0x4b951e		4889442448		MOVQ AX, 0x48(SP)		
  0x4b9523		4889542450		MOVQ DX, 0x50(SP)		
  0x4b9528		488b442468		MOVQ 0x68(SP), AX		
  0x4b952d		488b4c2470		MOVQ 0x70(SP), CX		
  0x4b9532		488b542478		MOVQ 0x78(SP), DX		
  0x4b9537		4889442428		MOVQ AX, 0x28(SP)		
  0x4b953c		48894c2430		MOVQ CX, 0x30(SP)		
  0x4b9541		4889542438		MOVQ DX, 0x38(SP)		
  0x4b9546		488b442448		MOVQ 0x48(SP), AX		
  0x4b954b		4889442420		MOVQ AX, 0x20(SP)		
  0x4b9550		4839442430		CMPQ AX, 0x30(SP)		
  0x4b9555		7c02			JL 0x4b9559			
  0x4b9557		eb54			JMP 0x4b95ad			
  0x4b9559		488b442430		MOVQ 0x30(SP), AX		
  0x4b955e		4889442420		MOVQ AX, 0x20(SP)		
  0x4b9563		eb00			JMP 0x4b9565			
  0x4b9565		488b442428		MOVQ 0x28(SP), AX		
  0x4b956a		4839442440		CMPQ AX, 0x40(SP)		
  0x4b956f		7502			JNE 0x4b9573			
  0x4b9571		eb38			JMP 0x4b95ab			
  0x4b9573		488b442420		MOVQ 0x20(SP), AX		
  0x4b9578		4889442418		MOVQ AX, 0x18(SP)		
  0x4b957d		488b442440		MOVQ 0x40(SP), AX		
  0x4b9582		48890424		MOVQ AX, 0(SP)			
  0x4b9586		488b442428		MOVQ 0x28(SP), AX		
  0x4b958b		4889442408		MOVQ AX, 0x8(SP)		
  0x4b9590		488b442418		MOVQ 0x18(SP), AX		
  0x4b9595		4889442410		MOVQ AX, 0x10(SP)		
  0x4b959a		e8a1b9f9ff		CALL runtime.memmove(SB)	
  0x4b959f		eb00			JMP 0x4b95a1			
  0x4b95a1		488b6c2458		MOVQ 0x58(SP), BP		
  0x4b95a6		4883c460		ADDQ $0x60, SP			
  0x4b95aa		c3			RET				
  0x4b95ab		ebf4			JMP 0x4b95a1			
  0x4b95ad		ebb6			JMP 0x4b9565			
func copyPixels(src []uint8, dst []uint8) {
  0x4b95af		e8ec89f9ff		CALL runtime.morestack_noctxt(SB)	
  0x4b95b4		e927ffffff		JMP main.copyPixels(SB)			
