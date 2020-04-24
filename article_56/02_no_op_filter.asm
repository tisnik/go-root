TEXT main.copyPixels(SB) /home/ptisnovs/src/go-root/article_56/02_no_op_filter.go
func copyPixels(src []uint8, dst []uint8) {
  0x4b90f0		4883ec18		SUBQ $0x18, SP		
  0x4b90f4		48896c2410		MOVQ BP, 0x10(SP)	
  0x4b90f9		488d6c2410		LEAQ 0x10(SP), BP	
	for i := 0; i < len(src); i++ {
  0x4b90fe		488b542428		MOVQ 0x28(SP), DX	
  0x4b9103		488b5c2438		MOVQ 0x38(SP), BX	
  0x4b9108		488b742440		MOVQ 0x40(SP), SI	
  0x4b910d		488b7c2420		MOVQ 0x20(SP), DI	
  0x4b9112		31c0			XORL AX, AX		
  0x4b9114		eb07			JMP 0x4b911d		
		dst[i] = src[i]
  0x4b9116		44880403		MOVB R8, 0(BX)(AX*1)	
	for i := 0; i < len(src); i++ {
  0x4b911a		48ffc0			INCQ AX			
  0x4b911d		4839d0			CMPQ DX, AX		
  0x4b9120		7d0c			JGE 0x4b912e		
		dst[i] = src[i]
  0x4b9122		440fb60407		MOVZX 0(DI)(AX*1), R8		
  0x4b9127		4839f0			CMPQ SI, AX			
  0x4b912a		72ea			JB 0x4b9116			
  0x4b912c		eb0a			JMP 0x4b9138			
  0x4b912e		488b6c2410		MOVQ 0x10(SP), BP		
  0x4b9133		4883c418		ADDQ $0x18, SP			
  0x4b9137		c3			RET				
  0x4b9138		4889f1			MOVQ SI, CX			
  0x4b913b		e8d0b5f9ff		CALL runtime.panicIndex(SB)	
  0x4b9140		90			NOPL				
