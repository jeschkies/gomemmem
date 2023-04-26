// Code generated by command: go run asm.go -out bytes.s -stubs bytes.go. DO NOT EDIT.

#include "textflag.h"

// func Mask(first byte, haystack []byte) int32
// Requires: AVX, AVX2
TEXT ·Mask(SB), NOSPLIT, $0-36
	MOVQ         haystack_base+8(FP), AX
	VPBROADCASTB first+0(FP), Y0
	VMOVDQU      (AX), Y1
	VPCMPEQB     Y0, Y1, Y0
	VPMOVMSKB    Y0, AX
	MOVL         AX, ret+32(FP)
	RET

// func Search(haystack []byte, needle []byte) bool
// Requires: AVX, AVX2, BMI
TEXT ·Search(SB), NOSPLIT, $0-49
	MOVQ needle_base+24(FP), AX
	MOVQ needle_len+32(FP), CX
	MOVQ haystack_base+0(FP), DX
	MOVQ DX, BX
	ADDQ haystack_len+8(FP), BX
	SUBQ $0x20, BX

	// create vector filled with first and last character
	VPBROADCASTB needle+0(FP), Y0
	MOVQ         needle+0(FP), SI
	ADDQ         CX, SI
	DECQ         SI
	VPBROADCASTB (SI), Y1

chunk_loop:
	CMPQ DX, BX
	JG   chunk_loop_end
	MOVQ DX, SI
	ADDQ CX, SI
	DECQ SI

	// compare blocks against first and last character
	VMOVDQU  (DX), Y2
	VMOVDQU  (SI), Y3
	VPCMPEQB Y0, Y2, Y2
	VPCMPEQB Y1, Y3, Y3

	// create mask and determine position
	VPAND     Y2, Y3, Y2
	VPMOVMSKB Y2, SI

mask_loop:
	CMPL   SI, $0x00
	JE     mask_loop_done
	TZCNTL SI, DI
	MOVQ   DX, R8
	ADDQ   DI, R8

	// compare two slices
	MOVQ CX, DI
	MOVQ AX, R9

memcmp_loop:
	// the loop is done; the chunks must be equal
	CMPQ DI, $0x00
	JE   memcmp_equal
	MOVB (R9), R10
	CMPB (R8), R10
	JNE  memcmp_not_equal
	ADDQ $0x01, R8
	ADDQ $0x01, R9
	DECQ DI
	JMP  memcmp_loop

memcmp_equal:
	MOVB $0x01, ret+48(FP)
	RET

memcmp_not_equal:
	MOVL SI, DI
	DECL DI
	ANDL DI, SI
	JMP  mask_loop

mask_loop_done:
	ADDQ $0x20, DX
	JMP  chunk_loop

chunk_loop_end:
	MOVB $0x00, ret+48(FP)
	RET
