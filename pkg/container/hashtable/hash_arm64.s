// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#include "textflag.h"

// func crc32Int64BatchHash(data *uint64, hashes *uint64, length int)
// Requires: CRC32
TEXT ·crc32Int64BatchHash(SB), NOSPLIT, $0-24
	MOVD data+0(FP), R0
	MOVD hashes+8(FP), R1
	MOVD length+16(FP), R2

loop:
	SUBS $8, R2
	BLT  tail

	MOVD $-1, R3
	MOVD $-1, R4
	MOVD $-1, R5
	MOVD $-1, R6
	MOVD $-1, R7
	MOVD $-1, R8
	MOVD $-1, R9
	MOVD $-1, R10

	LDP.P 16(R0), (R11, R12)
	LDP.P 16(R0), (R13, R14)
	LDP.P 16(R0), (R15, R16)
	LDP.P 16(R0), (R17, R19)

	CRC32CX R11, R3
	CRC32CX R12, R4
	CRC32CX R13, R5
	CRC32CX R14, R6
	CRC32CX R15, R7
	CRC32CX R16, R8
	CRC32CX R17, R9
	CRC32CX R19, R10

	STP.P (R3, R4), 16(R1)
	STP.P (R5, R6), 16(R1)
	STP.P (R7, R8), 16(R1)
	STP.P (R9, R10), 16(R1)

	JMP loop

tail:
	ADDS $8, R2
	BEQ  done

tailLoop:
	MOVD    $-1, R3
	MOVD.P  8(R0), R4
	CRC32CX R4, R3
	MOVD.P  R3, 8(R1)

	SUBS $1, R2
	BNE  tailLoop

done:
	RET

// func crc32Int64CellBatchHash(data *uint64, hashes *uint64, length int)
// Requires: CRC32
TEXT ·crc32Int64CellBatchHash(SB), NOSPLIT, $0-24
	MOVD data+0(FP), R0
	MOVD hashes+8(FP), R1
	MOVD length+16(FP), R2

loop:
	SUBS $8, R2
	BLT  tail

	MOVD $-1, R3
	MOVD $-1, R4
	MOVD $-1, R5
	MOVD $-1, R6
	MOVD $-1, R7
	MOVD $-1, R8
	MOVD $-1, R9
	MOVD $-1, R10

	MOVD.P 16(R0), R11
	MOVD.P 16(R0), R12
	MOVD.P 16(R0), R13
	MOVD.P 16(R0), R14
	MOVD.P 16(R0), R15
	MOVD.P 16(R0), R16
	MOVD.P 16(R0), R17
	MOVD.P 16(R0), R19

	CRC32CX R11, R3
	CRC32CX R12, R4
	CRC32CX R13, R5
	CRC32CX R14, R6
	CRC32CX R15, R7
	CRC32CX R16, R8
	CRC32CX R17, R9
	CRC32CX R19, R10

	STP.P (R3, R4), 16(R1)
	STP.P (R5, R6), 16(R1)
	STP.P (R7, R8), 16(R1)
	STP.P (R9, R10), 16(R1)

	JMP loop

tail:
	ADDS $8, R2
	BEQ  done

tailLoop:
	MOVD    $-1, R4
	MOVD.P  16(R0), R3
	CRC32CX R3, R4
	MOVD.P  R4, 8(R1)

	SUBS $1, R2
	BNE  tailLoop

done:
	RET

////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////

DATA Pi<>+0x00(SB)/8, $0x3243f6a8885a308d
DATA Pi<>+0x08(SB)/8, $0x313198a2e0370734
DATA Pi<>+0x10(SB)/8, $0x4a4093822299f31d
DATA Pi<>+0x18(SB)/8, $0x0082efa98ec4e6c8
DATA Pi<>+0x20(SB)/8, $0x9452821e638d0137
DATA Pi<>+0x28(SB)/8, $0x7be5466cf34e90c6
DATA Pi<>+0x30(SB)/8, $0xcc0ac29b7c97c50d
DATA Pi<>+0x38(SB)/8, $0xd3f84d5b5b547091
DATA Pi<>+0x40(SB)/8, $0x79216d5d98979fb1
DATA Pi<>+0x48(SB)/8, $0xbd1310ba698dfb5a
DATA Pi<>+0x50(SB)/8, $0xc2ffd72dbd01adfb
DATA Pi<>+0x58(SB)/8, $0x7b8e1afed6a267e9
DATA Pi<>+0x60(SB)/8, $0x6ba7c9045f12c7f9
DATA Pi<>+0x68(SB)/8, $0x924a19947b3916cf
DATA Pi<>+0x70(SB)/8, $0x70801f2e2858efc1
DATA Pi<>+0x78(SB)/8, $0x6636920d871574e6
GLOBL Pi<>(SB), (NOPTR+RODATA), $0x80

DATA CryptedPi<>+0x00(SB)/8, $0x822233b93c11087c
DATA CryptedPi<>+0x08(SB)/8, $0xd2b32f4adde873da
DATA CryptedPi<>+0x10(SB)/8, $0xae9c2fc7dd17bcdb
DATA CryptedPi<>+0x18(SB)/8, $0x859110441a1569fc
DATA CryptedPi<>+0x20(SB)/8, $0x47087d794fffb5c9
DATA CryptedPi<>+0x28(SB)/8, $0xb7b6c8f565414445
DATA CryptedPi<>+0x30(SB)/8, $0xfd260edabb308f8d
DATA CryptedPi<>+0x38(SB)/8, $0x3ddefc67bc565a13
DATA CryptedPi<>+0x40(SB)/8, $0xe4c1d50223544f10
DATA CryptedPi<>+0x48(SB)/8, $0xaf40e05725c3192b
DATA CryptedPi<>+0x50(SB)/8, $0x281d8ab9a16382e9
DATA CryptedPi<>+0x58(SB)/8, $0xddc10c903b63a6cf
DATA CryptedPi<>+0x60(SB)/8, $0x852d3ad603e8df72
DATA CryptedPi<>+0x68(SB)/8, $0xa6642b57d1011deb
DATA CryptedPi<>+0x70(SB)/8, $0x5063d25a1cb7b6b9
DATA CryptedPi<>+0x78(SB)/8, $0xb2623e6241e8e46e
GLOBL CryptedPi<>(SB), (NOPTR+RODATA), $0x80

// func aesBytesBatchGenHashStates(data *[]byte, states *[3]uint64, length int)
// Requires: AES
TEXT ·aesBytesBatchGenHashStates(SB), NOSPLIT, $0-24
	MOVD data+0(FP), R0
	MOVD states+8(FP), R1
	MOVD length+16(FP), R2

	MOVD $CryptedPi<>(SB), R3
	VLD1.P 64(R3), [V0.B16, V1.B16, V2.B16, V3.B16]
	VLD1 (R3), [V4.B16, V5.B16, V6.B16, V7.B16]
	VEOR V31.B16, V31.B16, V31.B16

loop:
	LDP.P 24(R0), (R4, R5)
	MOVD  R5, R6

	ADD R4, R5
	SUB $0x40, R5

	VMOV V0.B16, V8.B16
	VMOV V1.B16, V9.B16
	VMOV V2.B16, V10.B16
	VMOV V3.B16, V11.B16
	VMOV V4.B16, V12.B16
	VMOV V5.B16, V13.B16
	VMOV V6.B16, V14.B16
	VMOV V7.B16, V15.B16

innerLoop:
	CMP R4, R5
	BLE tail

	VLD1.P 0x40(R4), [V16.B16, V17.B16, V18.B16, V19.B16]

	AESE  V31.B16, V8.B16
	AESMC V8.B16, V8.B16
	VEOR  V16.B16, V8.B16, V8.B16

	AESE  V31.B16, V12.B16
	AESMC V12.B16, V12.B16
	VEOR  V16.B16, V12.B16, V12.B16

	AESE  V31.B16, V9.B16
	AESMC V9.B16, V9.B16
	VEOR  V17.B16, V9.B16, V9.B16

	AESE  V31.B16, V13.B16
	AESMC V13.B16, V13.B16
	VEOR  V17.B16, V13.B16, V13.B16

	AESE  V31.B16, V10.B16
	AESMC V10.B16, V10.B16
	VEOR  V18.B16, V10.B16, V10.B16

	AESE  V31.B16, V14.B16
	AESMC V14.B16, V14.B16
	VEOR  V18.B16, V14.B16, V14.B16

	AESE  V31.B16, V11.B16
	AESMC V11.B16, V11.B16
	VEOR  V19.B16, V11.B16, V11.B16

	AESE  V31.B16, V15.B16
	AESMC V15.B16, V15.B16
	VEOR  V19.B16, V15.B16, V15.B16

	JMP innerLoop

tail:
	ADD $0x30, R5
	CMP R4, R5
	BLE done

	VLD1.P 0x10(R4), [V16.B16]

	AESE  V31.B16, V8.B16
	AESMC V8.B16, V8.B16
	VEOR  V16.B16, V8.B16, V8.B16

	AESE  V31.B16, V12.B16
	AESMC V12.B16, V12.B16
	VEOR  V16.B16, V12.B16, V12.B16

	CMP R4, R5
	BLE done

	VLD1.P 0x10(R4), [V17.B16]

	AESE  V31.B16, V9.B16
	AESMC V9.B16, V9.B16
	VEOR  V17.B16, V9.B16, V9.B16

	AESE  V31.B16, V13.B16
	AESMC V13.B16, V13.B16
	VEOR  V17.B16, V13.B16, V13.B16

	CMP R4, R5
	BLE done

	VLD1 (R4), [V18.B16]

	AESE  V31.B16, V10.B16
	AESMC V10.B16, V10.B16
	VEOR  V18.B16, V10.B16, V10.B16

	AESE  V31.B16, V14.B16
	AESMC V14.B16, V14.B16
	VEOR  V18.B16, V14.B16, V14.B16

done:
	VLD1  (R5), [V19.B16]

	AESE  V31.B16, V11.B16
	AESMC V11.B16, V11.B16
	VEOR  V19.B16, V11.B16, V11.B16

	AESE  V31.B16, V15.B16
	AESMC V15.B16, V15.B16
	VEOR  V19.B16, V15.B16, V15.B16

	AESE  V31.B16, V8.B16
	AESMC V8.B16, V8.B16
	VEOR  V9.B16, V8.B16, V8.B16

	AESE  V31.B16, V11.B16
	AESMC V11.B16, V11.B16

	AESE  V10.B16, V11.B16
	AESMC V11.B16, V11.B16
	VEOR  V8.B16, V11.B16, V9.B16

	AESE  V8.B16, V11.B16
	AESMC V11.B16, V11.B16
	VEOR  V9.B16, V11.B16, V10.B16

	AESE  V9.B16, V11.B16
	AESMC V11.B16, V11.B16
	VEOR  V10.B16, V11.B16, V8.B16

	AESE  V10.B16, V11.B16
	AESMC V11.B16, V11.B16
	VEOR  V8.B16, V11.B16, V11.B16

	AESE  V31.B16, V12.B16
	AESMC V12.B16, V12.B16

	AESE  V31.B16, V13.B16
	AESMC V13.B16, V13.B16
	VEOR  V14.B16, V13.B16, V13.B16

	AESE  V15.B16, V12.B16
	AESMC V12.B16, V12.B16
	VEOR  V13.B16, V12.B16, V12.B16

	VMOV V11.D[0], R7
	VMOV V11.D[1], R8
	EOR  R8, R7
	EOR  R6, R7

	MOVD.P R7, 8(R1)
	VST1.P [V12.B16], 16(R1)

	SUBS $1, R2
	BNE  loop

	RET

// func aesInt192BatchGenHashStates(data *[3]uint64, states *[3]uint64, length int)
// Requires: AES
TEXT ·aesInt192BatchGenHashStates(SB), NOSPLIT, $0-24
	MOVD data+0(FP), R0
	MOVD states+8(FP), R1
	MOVD length+16(FP), R2

	MOVD $CryptedPi<>(SB), R3
	VLD1.P 64(R3), [V0.B16, V1.B16, V2.B16, V3.B16]
	VLD1 (R3), [V4.B16, V5.B16, V6.B16, V7.B16]
	VEOR V31.B16, V31.B16, V31.B16

	VMOV V0.B16, V30.B16

	AESE  V31.B16, V0.B16
	AESMC V0.B16, V0.B16

	AESE  V31.B16, V1.B16
	AESMC V1.B16, V1.B16

	AESE  V31.B16, V3.B16
	AESMC V3.B16, V3.B16
	VEOR  V2.B16, V3.B16, V3.B16

	AESE  V31.B16, V4.B16
	AESMC V4.B16, V4.B16

	AESE  V31.B16, V5.B16
	AESMC V5.B16, V5.B16

	AESE  V31.B16, V6.B16
	AESMC V6.B16, V6.B16
	VEOR  V7.B16, V6.B16, V6.B16

loop:
	VLD1   (R0), [V8.B16]
	ADD    $0x08, R0
	VLD1.P 0x10(R0), [V9.B16]

	VEOR V0.B16, V8.B16, V10.B16
	VEOR V5.B16, V9.B16, V11.B16

	AESE  V1.B16, V9.B16
	AESMC V9.B16, V9.B16

	AESE  V10.B16, V9.B16
	AESMC V9.B16, V9.B16
	VEOR  V3.B16, V9.B16, V10.B16

	AESE  V3.B16, V9.B16
	AESMC V9.B16, V9.B16
	VEOR  V10.B16, V9.B16, V12.B16

	AESE  V10.B16, V9.B16
	AESMC V9.B16, V9.B16
	VEOR  V12.B16, V9.B16, V9.B16

	VMOV  V9.D[0], R4
	VMOV  V9.D[1], R5
	EOR   R5, R4

	AESE  V4.B16, V8.B16
	AESMC V8.B16, V8.B16

	AESE  V11.B16, V8.B16
	AESMC V8.B16, V8.B16
	VEOR  V6.B16, V8.B16, V8.B16

	MOVD.P R4, 0x08(R1)
	VST1.P [V8.B16], 0x10(R1)

	SUBS $1, R2
	BNE  loop

done:
	RET

// func aesInt256BatchGenHashStates(data *[4]uint64, states *[3]uint64, length int)
// Requires: AES
TEXT ·aesInt256BatchGenHashStates(SB), NOSPLIT, $0-24
	MOVD data+0(FP), R0
	MOVD states+8(FP), R1
	MOVD length+16(FP), R2

	MOVD $CryptedPi<>(SB), R3
	VLD1.P 64(R3), [V0.B16, V1.B16, V2.B16, V3.B16]
	VLD1 (R3), [V4.B16, V5.B16, V6.B16, V7.B16]
	VEOR V31.B16, V31.B16, V31.B16

	VMOV V0.B16, V30.B16

	AESE  V31.B16, V0.B16
	AESMC V0.B16, V0.B16

	AESE  V31.B16, V1.B16
	AESMC V1.B16, V1.B16

	AESE  V31.B16, V3.B16
	AESMC V3.B16, V3.B16
	VEOR  V2.B16, V3.B16, V3.B16

	AESE  V31.B16, V4.B16
	AESMC V4.B16, V4.B16

	AESE  V31.B16, V5.B16
	AESMC V5.B16, V5.B16

	AESE  V31.B16, V6.B16
	AESMC V6.B16, V6.B16
	VEOR  V7.B16, V6.B16, V6.B16

loop:
	VLD1.P 0x20(R0), [V8.B16, V9.B16]

	VEOR V0.B16, V8.B16, V10.B16
	VEOR V5.B16, V9.B16, V11.B16

	AESE  V1.B16, V9.B16
	AESMC V9.B16, V9.B16

	AESE  V10.B16, V9.B16
	AESMC V9.B16, V9.B16
	VEOR  V3.B16, V9.B16, V10.B16

	AESE  V3.B16, V9.B16
	AESMC V9.B16, V9.B16
	VEOR  V10.B16, V9.B16, V12.B16

	AESE  V10.B16, V9.B16
	AESMC V9.B16, V9.B16
	VEOR  V12.B16, V9.B16, V9.B16

	VMOV  V9.D[0], R4
	VMOV  V9.D[1], R5
	EOR   R5, R4

	AESE  V4.B16, V8.B16
	AESMC V8.B16, V8.B16

	AESE  V11.B16, V8.B16
	AESMC V8.B16, V8.B16
	VEOR  V6.B16, V8.B16, V8.B16

	MOVD.P R4, 0x08(R1)
	VST1.P [V8.B16], 0x10(R1)

	SUBS $1, R2
	BNE  loop

done:
	RET

// func aesInt320BatchGenHashStates(data *[5]uint64, states *[3]uint64, length int)
// Requires: AES
TEXT ·aesInt320BatchGenHashStates(SB), NOSPLIT, $0-24
	MOVD data+0(FP), R0
	MOVD states+8(FP), R1
	MOVD length+16(FP), R2

	MOVD $CryptedPi<>(SB), R3
	VLD1.P 64(R3), [V0.B16, V1.B16, V2.B16, V3.B16]
	VLD1 (R3), [V4.B16, V5.B16, V6.B16, V7.B16]
	VEOR V31.B16, V31.B16, V31.B16

	AESE  V31.B16, V0.B16
	AESMC V0.B16, V0.B16

	AESE  V31.B16, V1.B16
	AESMC V1.B16, V1.B16

	AESE  V31.B16, V3.B16
	AESMC V3.B16, V3.B16

	AESE  V31.B16, V4.B16
	AESMC V4.B16, V4.B16

	AESE  V31.B16, V5.B16
	AESMC V5.B16, V5.B16

	AESE  V31.B16, V6.B16
	AESMC V6.B16, V6.B16

loop:
	VLD1 (R0), [V8.B16, V9.B16]
	ADD  $0x18, R0
	VLD1.P 0x10(R0), [V10.B16]

	VEOR V4.B16, V8.B16, V11.B16
	VEOR V5.B16, V9.B16, V12.B16

	VEOR V3.B16, V10.B16, V13.B16

	AESE  V0.B16, V8.B16
	AESMC V8.B16, V8.B16

	AESE  V1.B16, V9.B16
	AESMC V9.B16, V9.B16
	VEOR  V2.B16, V9.B16, V9.B16

	AESE  V13.B16, V8.B16
	AESMC V8.B16, V8.B16
	VEOR  V9.B16, V8.B16, V13.B16

	AESE  V9.B16, V8.B16
	AESMC V8.B16, V8.B16
	VEOR  V13.B16, V8.B16, V9.B16

	AESE  V13.B16, V8.B16
	AESMC V8.B16, V8.B16
	VEOR  V9.B16, V8.B16, V8.B16

	VMOV  V8.D[0], R4
	VMOV  V8.D[1], R5
	EOR   R5, R4

	AESE  V31.B16, V11.B16
	AESMC V11.B16, V11.B16

	AESE  V6.B16, V10.B16
	AESMC V10.B16, V10.B16
	VEOR  V7.B16, V10.B16, V10.B16

	AESE  V12.B16, V11.B16
	AESMC V11.B16, V11.B16
	VEOR  V10.B16, V11.B16, V11.B16

	MOVD.P R4, 0x08(R1)
	VST1.P [V11.B16], 0x10(R1)

	SUBS $1, R2
	BNE  loop

done:
	RET
