package nes

const (
	ResetVector     uint16 = 0xFFFC
	StackBeginning  uint16 = 0x1FCC
	StackEnd        uint16 = 0x1000
	PrgRomLowerBank uint16 = 0x8000
	PrgRomUpperBank uint16 = 0xC000
)

type Instruction struct {
	Opcode Opcode
	Mode   Mode
	Bytes  int
	Cycles int
	Name   string
}

type Mode int
type Opcode uint8

const (
	Implied Mode = iota
	Immediate
	ZeroPage
	ZeroPageX
	ZeroPageY
	Absolute
	AbsoluteX
	AbsoluteX1
	AbsoluteY
	AbsoluteY1
	Indirect
	IndirectX
	IndirectY
	IndirectY1
	Accumulator
	Relative
)

const (
	Carry     uint8 = (1 << 0)
	Zero      uint8 = (1 << 1)
	Interrupt uint8 = (1 << 2)
	Decimal   uint8 = (1 << 3)
	Break     uint8 = (1 << 4)
	Break2    uint8 = (1 << 5)
	Verflow   uint8 = (1 << 6)
	Negative  uint8 = (1 << 7)
)

const (
	LDA_IMM Opcode = 0xA9
	LDA_ZER Opcode = 0xA5
	LDA_ZRX Opcode = 0xB5
	LDA_ABS Opcode = 0xAD
	LDA_ABX Opcode = 0xBD
	LDA_ABY Opcode = 0xB9
	LDA_IDX Opcode = 0xA1
	LDA_IDY Opcode = 0xB1

	LDX_IMM Opcode = 0xA2
	LDX_ZER Opcode = 0xA6
	LDX_ZRY Opcode = 0xB6
	LDX_ABS Opcode = 0xAE
	LDX_ABY Opcode = 0xBE

	LDY_IMM Opcode = 0xA0
	LDY_ZER Opcode = 0xA4
	LDY_ZRX Opcode = 0xB4
	LDY_ABS Opcode = 0xAC
	LDY_ABX Opcode = 0xBC

	STA_ZER Opcode = 0x85
	STA_ZRX Opcode = 0x95
	STA_ABS Opcode = 0x8D
	STA_ABX Opcode = 0x9D
	STA_ABY Opcode = 0x99
	STA_IDX Opcode = 0x81
	STA_IDY Opcode = 0x91

	STX_ZER Opcode = 0x86
	STX_ZRY Opcode = 0x96
	STX_ABS Opcode = 0x8E

	STY_ZER Opcode = 0x84
	STY_ZRX Opcode = 0x94
	STY_ABS Opcode = 0x8C

	TAX_IMP Opcode = 0xAA
	TAY_IMP Opcode = 0xA8
	TXA_IMP Opcode = 0x8A
	TYA_IMP Opcode = 0x98

	TSX_IMP Opcode = 0xBA
	TXS_IMP Opcode = 0x9A
	PHA_IMP Opcode = 0x48
	PHP_IMP Opcode = 0x08
	PLA_IMP Opcode = 0x68
	PLP_IMP Opcode = 0x28

	AND_IMM Opcode = 0x29
	AND_ZER Opcode = 0x25
	AND_ZRX Opcode = 0x35
	AND_ABS Opcode = 0x2D
	AND_ABX Opcode = 0x3D
	AND_ABY Opcode = 0x39
	AND_IDX Opcode = 0x21
	AND_IDY Opcode = 0x31

	EOR_IMM Opcode = 0x49
	EOR_ZER Opcode = 0x45
	EOR_ZRX Opcode = 0x55
	EOR_ABS Opcode = 0x4D
	EOR_ABX Opcode = 0x5D
	EOR_ABY Opcode = 0x59
	EOR_IDX Opcode = 0x41
	EOR_IDY Opcode = 0x51

	ORA_IMM Opcode = 0x09
	ORA_ZER Opcode = 0x05
	ORA_ZRX Opcode = 0x15
	ORA_ABS Opcode = 0x0D
	ORA_ABX Opcode = 0x1D
	ORA_ABY Opcode = 0x19
	ORA_IDX Opcode = 0x01
	ORA_IDY Opcode = 0x11

	BIT_ZER Opcode = 0x24
	BIT_ABS Opcode = 0x2C

	ADC_IMM Opcode = 0x69
	ADC_ZER Opcode = 0x65
	ADC_ZRX Opcode = 0x75
	ADC_ABS Opcode = 0x6D
	ADC_ABX Opcode = 0x7D
	ADC_ABY Opcode = 0x79
	ADC_IDX Opcode = 0x61
	ADC_IDY Opcode = 0x71

	SBC_IMM Opcode = 0xE9
	SBC_ZER Opcode = 0xE5
	SBC_ZRX Opcode = 0xF5
	SBC_ABS Opcode = 0xED
	SBC_ABX Opcode = 0xFD
	SBC_ABY Opcode = 0xF9
	SBC_IDX Opcode = 0xE1
	SBC_IDY Opcode = 0xF1

	CMP_IMM Opcode = 0xC9
	CMP_ZER Opcode = 0xC5
	CMP_ZRX Opcode = 0xD5
	CMP_ABS Opcode = 0xCD
	CMP_ABX Opcode = 0xDD
	CMP_ABY Opcode = 0xD9
	CMP_IDX Opcode = 0xC1
	CMP_IDY Opcode = 0xD1

	CPX_IMM Opcode = 0xE0
	CPX_ZER Opcode = 0xE4
	CPX_ABS Opcode = 0xEC

	CPY_IMM Opcode = 0xC0
	CPY_ZER Opcode = 0xC4
	CPY_ABS Opcode = 0xCC

	INC_ZER Opcode = 0xE6
	INC_ZRX Opcode = 0xF6
	INC_ABS Opcode = 0xEE
	INC_ABX Opcode = 0xFE
	INX_IMP Opcode = 0xE8
	INY_IMP Opcode = 0xC8

	DEC_ZER Opcode = 0xC6
	DEC_ZRX Opcode = 0xD6
	DEC_ABS Opcode = 0xCE
	DEC_ABX Opcode = 0xDE
	DEX_IMP Opcode = 0xCA
	DEY_IMP Opcode = 0x88

	ASL_ACC Opcode = 0x0A
	ASL_ZER Opcode = 0x06
	ASL_ZRX Opcode = 0x16
	ASL_ABS Opcode = 0x0E
	ASL_ABX Opcode = 0x1E

	LSR_ACC Opcode = 0x4A
	LSR_ZER Opcode = 0x46
	LSR_ZRX Opcode = 0x56
	LSR_ABS Opcode = 0x4E
	LSR_ABX Opcode = 0x5E

	ROL_ACC Opcode = 0x2A
	ROL_ZER Opcode = 0x26
	ROL_ZRX Opcode = 0x36
	ROL_ABS Opcode = 0x2E
	ROL_ABX Opcode = 0x3E

	ROR_ACC Opcode = 0x6A
	ROR_ZER Opcode = 0x66
	ROR_ZRX Opcode = 0x76
	ROR_ABS Opcode = 0x6E
	ROR_ABX Opcode = 0x7E

	JMP_ABS Opcode = 0x4C
	JMP_IND Opcode = 0x6C

	JSR_ABS Opcode = 0x20
	RTS_IMP Opcode = 0x60

	BCC_REL Opcode = 0x90
	BCS_REL Opcode = 0xB0
	BEQ_REL Opcode = 0xF0
	BMI_REL Opcode = 0x30
	BNE_REL Opcode = 0xD0
	BPL_REL Opcode = 0x10
	BVC_REL Opcode = 0x50
	BVS_REL Opcode = 0x70

	CLC_IMP Opcode = 0x18
	CLD_IMP Opcode = 0xD8
	CLI_IMP Opcode = 0x58
	CLV_IMP Opcode = 0xB8

	SEC_IMP Opcode = 0x38
	SED_IMP Opcode = 0xF8
	SEI_IMP Opcode = 0x78

	BRK_IMP Opcode = 0x00
	NOP_IMP Opcode = 0xEA
	RTI_IMP Opcode = 0x40

	LAST_OP Opcode = 0xFF
)

var Instructions = [LAST_OP]Instruction{

	// Load Operations
	LDA_IMM: {Name: "LDA_IMM", Opcode: LDA_IMM, Bytes: 2, Cycles: 2, Mode: Immediate},
	LDA_ZER: {Name: "LDA_ZER", Opcode: LDA_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	LDA_ZRX: {Name: "LDA_ZRX", Opcode: LDA_ZRX, Bytes: 2, Cycles: 4, Mode: ZeroPageX},
	LDA_ABS: {Name: "LDA_ABS", Opcode: LDA_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},
	LDA_ABX: {Name: "LDA_ABX", Opcode: LDA_ABX, Bytes: 3, Cycles: 4, Mode: AbsoluteX1},
	LDA_ABY: {Name: "LDA_ABY", Opcode: LDA_ABY, Bytes: 3, Cycles: 4, Mode: AbsoluteY1},
	LDA_IDX: {Name: "LDA_IDX", Opcode: LDA_IDX, Bytes: 2, Cycles: 6, Mode: IndirectX},
	LDA_IDY: {Name: "LDA_IDY", Opcode: LDA_IDY, Bytes: 2, Cycles: 5, Mode: IndirectY1},

	LDX_IMM: {Name: "LDX_IMM", Opcode: LDX_IMM, Bytes: 2, Cycles: 2, Mode: Immediate},
	LDX_ZER: {Name: "LDX_ZER", Opcode: LDX_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	LDX_ZRY: {Name: "LDX_ZRY", Opcode: LDX_ZRY, Bytes: 2, Cycles: 4, Mode: ZeroPageY},
	LDX_ABS: {Name: "LDX_ABS", Opcode: LDX_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},
	LDX_ABY: {Name: "LDX_ABY", Opcode: LDX_ABY, Bytes: 3, Cycles: 4, Mode: AbsoluteY1},

	LDY_IMM: {Name: "LDY_IMM", Opcode: LDY_IMM, Bytes: 2, Cycles: 2, Mode: Immediate},
	LDY_ZER: {Name: "LDY_ZER", Opcode: LDY_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	LDY_ZRX: {Name: "LDY_ZRX", Opcode: LDY_ZRX, Bytes: 2, Cycles: 4, Mode: ZeroPageX},
	LDY_ABS: {Name: "LDY_ABS", Opcode: LDY_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},
	LDY_ABX: {Name: "LDY_ABX", Opcode: LDY_ABX, Bytes: 3, Cycles: 4, Mode: AbsoluteX1},

	// Store Operations
	STA_ZER: {Name: "STA_ZER", Opcode: STA_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	STA_ZRX: {Name: "STA_ZRX", Opcode: STA_ZRX, Bytes: 2, Cycles: 4, Mode: ZeroPageX},
	STA_ABS: {Name: "STA_ABS", Opcode: STA_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},
	STA_ABX: {Name: "STA_ABX", Opcode: STA_ABX, Bytes: 3, Cycles: 5, Mode: AbsoluteX},
	STA_ABY: {Name: "STA_ABY", Opcode: STA_ABY, Bytes: 3, Cycles: 5, Mode: AbsoluteY},
	STA_IDX: {Name: "STA_IDX", Opcode: STA_IDX, Bytes: 2, Cycles: 6, Mode: IndirectX},
	STA_IDY: {Name: "STA_IDY", Opcode: STA_IDY, Bytes: 2, Cycles: 6, Mode: IndirectY},

	STX_ZER: {Name: "STX_ZER", Opcode: STX_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	STX_ZRY: {Name: "STX_ZRY", Opcode: STX_ZRY, Bytes: 2, Cycles: 4, Mode: ZeroPageY},
	STX_ABS: {Name: "STX_ABS", Opcode: STX_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},

	STY_ZER: {Name: "STY_ZER", Opcode: STY_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	STY_ZRX: {Name: "STY_ZRX", Opcode: STY_ZRX, Bytes: 2, Cycles: 4, Mode: ZeroPageX},
	STY_ABS: {Name: "STY_ABS", Opcode: STY_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},

	// Register Transfers
	TAX_IMP: {Name: "TAX_IMP", Opcode: TAX_IMP, Bytes: 1, Cycles: 2, Mode: Implied},
	TAY_IMP: {Name: "TAY_IMP", Opcode: TAY_IMP, Bytes: 1, Cycles: 2, Mode: Implied},
	TXA_IMP: {Name: "TXA_IMP", Opcode: TXA_IMP, Bytes: 1, Cycles: 2, Mode: Implied},
	TYA_IMP: {Name: "TYA_IMP", Opcode: TYA_IMP, Bytes: 1, Cycles: 2, Mode: Implied},

	// Stack
	TSX_IMP: {Name: "TSX_IMP", Opcode: TSX_IMP, Bytes: 1, Cycles: 2, Mode: Implied},
	TXS_IMP: {Name: "TXS_IMP", Opcode: TXS_IMP, Bytes: 1, Cycles: 2, Mode: Implied},
	PHA_IMP: {Name: "PHA_IMP", Opcode: PHA_IMP, Bytes: 1, Cycles: 3, Mode: Implied},
	PHP_IMP: {Name: "PHP_IMP", Opcode: PHP_IMP, Bytes: 1, Cycles: 3, Mode: Implied},
	PLA_IMP: {Name: "PLA_IMP", Opcode: PLA_IMP, Bytes: 1, Cycles: 4, Mode: Implied},
	PLP_IMP: {Name: "PLP_IMP", Opcode: PLP_IMP, Bytes: 1, Cycles: 4, Mode: Implied},

	// Logical
	AND_IMM: {Name: "AND_IMM", Opcode: AND_IMM, Bytes: 2, Cycles: 2, Mode: Immediate},
	AND_ZER: {Name: "AND_ZER", Opcode: AND_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	AND_ZRX: {Name: "AND_ZRX", Opcode: AND_ZRX, Bytes: 2, Cycles: 4, Mode: ZeroPageX},
	AND_ABS: {Name: "AND_ABS", Opcode: AND_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},
	AND_ABX: {Name: "AND_ABX", Opcode: AND_ABX, Bytes: 3, Cycles: 4, Mode: AbsoluteX1},
	AND_ABY: {Name: "AND_ABY", Opcode: AND_ABY, Bytes: 3, Cycles: 4, Mode: AbsoluteY1},
	AND_IDX: {Name: "AND_IDX", Opcode: AND_IDX, Bytes: 2, Cycles: 6, Mode: IndirectX},
	AND_IDY: {Name: "AND_IDY", Opcode: AND_IDY, Bytes: 2, Cycles: 5, Mode: IndirectY1},

	EOR_IMM: {Name: "EOR_IMM", Opcode: EOR_IMM, Bytes: 2, Cycles: 2, Mode: Immediate},
	EOR_ZER: {Name: "EOR_ZER", Opcode: EOR_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	EOR_ZRX: {Name: "EOR_ZRX", Opcode: EOR_ZRX, Bytes: 2, Cycles: 4, Mode: ZeroPageX},
	EOR_ABS: {Name: "EOR_ABS", Opcode: EOR_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},
	EOR_ABX: {Name: "EOR_ABX", Opcode: EOR_ABX, Bytes: 3, Cycles: 4, Mode: AbsoluteX1},
	EOR_ABY: {Name: "EOR_ABY", Opcode: EOR_ABY, Bytes: 3, Cycles: 4, Mode: AbsoluteY1},
	EOR_IDX: {Name: "EOR_IDX", Opcode: EOR_IDX, Bytes: 2, Cycles: 6, Mode: IndirectX},
	EOR_IDY: {Name: "EOR_IDY", Opcode: EOR_IDY, Bytes: 2, Cycles: 5, Mode: IndirectY1},

	ORA_IMM: {Name: "ORA_IMM", Opcode: ORA_IMM, Bytes: 2, Cycles: 2, Mode: Immediate},
	ORA_ZER: {Name: "ORA_ZER", Opcode: ORA_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	ORA_ZRX: {Name: "ORA_ZRX", Opcode: ORA_ZRX, Bytes: 2, Cycles: 4, Mode: ZeroPageX},
	ORA_ABS: {Name: "ORA_ABS", Opcode: ORA_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},
	ORA_ABX: {Name: "ORA_ABX", Opcode: ORA_ABX, Bytes: 3, Cycles: 4, Mode: AbsoluteX1},
	ORA_ABY: {Name: "ORA_ABY", Opcode: ORA_ABY, Bytes: 3, Cycles: 4, Mode: AbsoluteY1},
	ORA_IDX: {Name: "ORA_IDX", Opcode: ORA_IDX, Bytes: 2, Cycles: 6, Mode: IndirectX},
	ORA_IDY: {Name: "ORA_IDY", Opcode: ORA_IDY, Bytes: 2, Cycles: 5, Mode: IndirectY1},

	BIT_ZER: {Name: "BIT_ZER", Opcode: BIT_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	BIT_ABS: {Name: "BIT_ABS", Opcode: BIT_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},

	// Arithmetic
	ADC_IMM: {Name: "ADC_IMM", Opcode: ADC_IMM, Bytes: 2, Cycles: 2, Mode: Immediate},
	ADC_ZER: {Name: "ADC_ZER", Opcode: ADC_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	ADC_ZRX: {Name: "ADC_ZRX", Opcode: ADC_ZRX, Bytes: 2, Cycles: 4, Mode: ZeroPageX},
	ADC_ABS: {Name: "ADC_ABS", Opcode: ADC_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},
	ADC_ABX: {Name: "ADC_ABX", Opcode: ADC_ABX, Bytes: 3, Cycles: 4, Mode: AbsoluteX1},
	ADC_ABY: {Name: "ADC_ABY", Opcode: ADC_ABY, Bytes: 3, Cycles: 4, Mode: AbsoluteY1},
	ADC_IDX: {Name: "ADC_IDX", Opcode: ADC_IDX, Bytes: 2, Cycles: 6, Mode: IndirectX},
	ADC_IDY: {Name: "ADC_IDY", Opcode: ADC_IDY, Bytes: 2, Cycles: 5, Mode: IndirectY1},

	SBC_IMM: {Name: "SBC_IMM", Opcode: SBC_IMM, Bytes: 2, Cycles: 2, Mode: Immediate},
	SBC_ZER: {Name: "SBC_ZER", Opcode: SBC_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	SBC_ZRX: {Name: "SBC_ZRX", Opcode: SBC_ZRX, Bytes: 2, Cycles: 4, Mode: ZeroPageX},
	SBC_ABS: {Name: "SBC_ABS", Opcode: SBC_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},
	SBC_ABX: {Name: "SBC_ABX", Opcode: SBC_ABX, Bytes: 3, Cycles: 4, Mode: AbsoluteX1},
	SBC_ABY: {Name: "SBC_ABY", Opcode: SBC_ABY, Bytes: 3, Cycles: 4, Mode: AbsoluteY1},
	SBC_IDX: {Name: "SBC_IDX", Opcode: SBC_IDX, Bytes: 2, Cycles: 6, Mode: IndirectX},
	SBC_IDY: {Name: "SBC_IDY", Opcode: SBC_IDY, Bytes: 2, Cycles: 5, Mode: IndirectY1},

	CMP_IMM: {Name: "CMP_IMM", Opcode: CMP_IMM, Bytes: 2, Cycles: 2, Mode: Immediate},
	CMP_ZER: {Name: "CMP_ZER", Opcode: CMP_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	CMP_ZRX: {Name: "CMP_ZRX", Opcode: CMP_ZRX, Bytes: 2, Cycles: 4, Mode: ZeroPageX},
	CMP_ABS: {Name: "CMP_ABS", Opcode: CMP_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},
	CMP_ABX: {Name: "CMP_ABX", Opcode: CMP_ABX, Bytes: 3, Cycles: 4, Mode: AbsoluteX1},
	CMP_ABY: {Name: "CMP_ABY", Opcode: CMP_ABY, Bytes: 3, Cycles: 4, Mode: AbsoluteY1},
	CMP_IDX: {Name: "CMP_IDX", Opcode: CMP_IDX, Bytes: 2, Cycles: 6, Mode: IndirectX},
	CMP_IDY: {Name: "CMP_IDY", Opcode: CMP_IDY, Bytes: 2, Cycles: 5, Mode: IndirectY1},

	CPX_IMM: {Name: "CPX_IMM", Opcode: CPX_IMM, Bytes: 2, Cycles: 2, Mode: Immediate},
	CPX_ZER: {Name: "CPX_ZER", Opcode: CPX_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	CPX_ABS: {Name: "CPX_ABS", Opcode: CPX_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},

	CPY_IMM: {Name: "CPY_IMM", Opcode: CPY_IMM, Bytes: 2, Cycles: 2, Mode: Immediate},
	CPY_ZER: {Name: "CPY_ZER", Opcode: CPY_ZER, Bytes: 2, Cycles: 3, Mode: ZeroPage},
	CPY_ABS: {Name: "CPY_ABS", Opcode: CPY_ABS, Bytes: 3, Cycles: 4, Mode: Absolute},

	// Increments
	INC_ZER: {Name: "INC_ZER", Opcode: INC_ZER, Bytes: 2, Cycles: 5, Mode: ZeroPage},
	INC_ZRX: {Name: "INC_ZRX", Opcode: INC_ZRX, Bytes: 2, Cycles: 6, Mode: ZeroPageX},
	INC_ABS: {Name: "INC_ABS", Opcode: INC_ABS, Bytes: 3, Cycles: 6, Mode: Absolute},
	INC_ABX: {Name: "INC_ABX", Opcode: INC_ABX, Bytes: 3, Cycles: 7, Mode: AbsoluteX},
	INX_IMP: {Name: "INX_IMP", Opcode: INX_IMP, Bytes: 1, Cycles: 2, Mode: Implied},
	INY_IMP: {Name: "INY_IMP", Opcode: INY_IMP, Bytes: 1, Cycles: 2, Mode: Implied},

	// Decrements
	DEC_ZER: {Name: "DEC_ZER", Opcode: DEC_ZER, Bytes: 2, Cycles: 5, Mode: ZeroPage},
	DEC_ZRX: {Name: "DEC_ZRX", Opcode: DEC_ZRX, Bytes: 2, Cycles: 6, Mode: ZeroPageX},
	DEC_ABS: {Name: "DEC_ABS", Opcode: DEC_ABS, Bytes: 3, Cycles: 6, Mode: Absolute},
	DEC_ABX: {Name: "DEC_ABX", Opcode: DEC_ABX, Bytes: 3, Cycles: 7, Mode: AbsoluteX},
	DEX_IMP: {Name: "DEX_IMP", Opcode: DEX_IMP, Bytes: 1, Cycles: 2, Mode: Implied},
	DEY_IMP: {Name: "DEY_IMP", Opcode: DEY_IMP, Bytes: 1, Cycles: 2, Mode: Implied},

	// Shifts
	ASL_ACC: {Name: "ASL_ACC", Opcode: ASL_ACC, Bytes: 1, Cycles: 2, Mode: Accumulator},
	ASL_ZER: {Name: "ASL_ZER", Opcode: ASL_ZER, Bytes: 2, Cycles: 5, Mode: ZeroPage},
	ASL_ZRX: {Name: "ASL_ZRX", Opcode: ASL_ZRX, Bytes: 2, Cycles: 6, Mode: ZeroPageX},
	ASL_ABS: {Name: "ASL_ABS", Opcode: ASL_ABS, Bytes: 3, Cycles: 6, Mode: Absolute},
	ASL_ABX: {Name: "ASL_ABX", Opcode: ASL_ABX, Bytes: 3, Cycles: 7, Mode: AbsoluteX},

	LSR_ACC: {Name: "LSR_ACC", Opcode: LSR_ACC, Bytes: 1, Cycles: 2, Mode: Accumulator},
	LSR_ZER: {Name: "LSR_ZER", Opcode: LSR_ZER, Bytes: 2, Cycles: 5, Mode: ZeroPage},
	LSR_ZRX: {Name: "LSR_ZRX", Opcode: LSR_ZRX, Bytes: 2, Cycles: 6, Mode: ZeroPageX},
	LSR_ABS: {Name: "LSR_ABS", Opcode: LSR_ABS, Bytes: 3, Cycles: 6, Mode: Absolute},
	LSR_ABX: {Name: "LSR_ABX", Opcode: LSR_ABX, Bytes: 3, Cycles: 7, Mode: AbsoluteX},

	ROL_ACC: {Name: "ROL_ACC", Opcode: ROL_ACC, Bytes: 1, Cycles: 2, Mode: Accumulator},
	ROL_ZER: {Name: "ROL_ZER", Opcode: ROL_ZER, Bytes: 2, Cycles: 5, Mode: ZeroPage},
	ROL_ZRX: {Name: "ROL_ZRX", Opcode: ROL_ZRX, Bytes: 2, Cycles: 6, Mode: ZeroPageX},
	ROL_ABS: {Name: "ROL_ABS", Opcode: ROL_ABS, Bytes: 3, Cycles: 6, Mode: Absolute},
	ROL_ABX: {Name: "ROL_ABX", Opcode: ROL_ABX, Bytes: 3, Cycles: 7, Mode: AbsoluteX},

	ROR_ACC: {Name: "ROR_ACC", Opcode: ROR_ACC, Bytes: 1, Cycles: 2, Mode: Accumulator},
	ROR_ZER: {Name: "ROR_ZER", Opcode: ROR_ZER, Bytes: 2, Cycles: 5, Mode: ZeroPage},
	ROR_ZRX: {Name: "ROR_ZRX", Opcode: ROR_ZRX, Bytes: 2, Cycles: 6, Mode: ZeroPageX},
	ROR_ABS: {Name: "ROR_ABS", Opcode: ROR_ABS, Bytes: 3, Cycles: 6, Mode: Absolute},
	ROR_ABX: {Name: "ROR_ABX", Opcode: ROR_ABX, Bytes: 3, Cycles: 7, Mode: AbsoluteX},

	// Jumps
	/* byte size set to 0 because to not change the prg counter after jump */
	JMP_ABS: {Name: "JMP_ABS", Opcode: JMP_ABS, Bytes: 3, Cycles: 3, Mode: Absolute},
	JMP_IND: {Name: "JMP_IND", Opcode: JMP_IND, Bytes: 3, Cycles: 5, Mode: Indirect},

	JSR_ABS: {Name: "JSR_ABS", Opcode: JSR_ABS, Bytes: 3, Cycles: 6, Mode: Absolute},
	RTS_IMP: {Name: "RTS_IMP", Opcode: RTS_IMP, Bytes: 1, Cycles: 6, Mode: Implied},

	// Branching
	BCC_REL: {Name: "BCC_REL", Opcode: BCC_REL, Bytes: 2, Cycles: 2 /*to+2*/, Mode: Relative},
	BCS_REL: {Name: "BCS_REL", Opcode: BCS_REL, Bytes: 2, Cycles: 2 /*to+2*/, Mode: Relative},
	BEQ_REL: {Name: "BEQ_REL", Opcode: BEQ_REL, Bytes: 2, Cycles: 2 /*to+2*/, Mode: Relative},
	BMI_REL: {Name: "BMI_REL", Opcode: BMI_REL, Bytes: 2, Cycles: 2 /*to+2*/, Mode: Relative},
	BNE_REL: {Name: "BNE_REL", Opcode: BNE_REL, Bytes: 2, Cycles: 2 /*to+2*/, Mode: Relative},
	BPL_REL: {Name: "BPL_REL", Opcode: BPL_REL, Bytes: 2, Cycles: 2 /*to+2*/, Mode: Relative},
	BVC_REL: {Name: "BVC_REL", Opcode: BVC_REL, Bytes: 2, Cycles: 2 /*to+2*/, Mode: Relative},
	BVS_REL: {Name: "BVS_REL", Opcode: BVS_REL, Bytes: 2, Cycles: 2 /*to+2*/, Mode: Relative},

	// Status Flag Changes
	CLC_IMP: {Name: "CLC_IMP", Opcode: CLC_IMP, Bytes: 1, Cycles: 2, Mode: Implied},
	CLD_IMP: {Name: "CLD_IMP", Opcode: CLD_IMP, Bytes: 1, Cycles: 2, Mode: Implied},
	CLI_IMP: {Name: "CLI_IMP", Opcode: CLI_IMP, Bytes: 1, Cycles: 2, Mode: Implied},
	CLV_IMP: {Name: "CLV_IMP", Opcode: CLV_IMP, Bytes: 1, Cycles: 2, Mode: Implied},

	SEC_IMP: {Name: "SEC_IMP", Opcode: SEC_IMP, Bytes: 1, Cycles: 2, Mode: Implied},
	SED_IMP: {Name: "SED_IMP", Opcode: SED_IMP, Bytes: 1, Cycles: 2, Mode: Implied},
	SEI_IMP: {Name: "SEI_IMP", Opcode: SEI_IMP, Bytes: 1, Cycles: 2, Mode: Implied},

	// System Functions
	BRK_IMP: {Name: "BRK_IMP", Opcode: BRK_IMP, Bytes: 1, Cycles: 7, Mode: Implied},
	NOP_IMP: {Name: "NOP_IMP", Opcode: NOP_IMP, Bytes: 1, Cycles: 2, Mode: Implied},
	RTI_IMP: {Name: "RTI_IMP", Opcode: RTI_IMP, Bytes: 1, Cycles: 6, Mode: Implied},
}
