package nes

import "log"

// type Memory struct {
// 	ChrRom []uint8
// 	Ram    []uint8
// 	Vram   []uint8
// 	// Mapper
// }

// func (m *Memory) ReadChrRom() {
// }

type Input struct {
}

type Apu struct {
}

type Memory struct {
}

type Console struct {
	Cpu       *Cpu
	Apu       *Apu
	Ppu       *Ppu
	Cartridge Cartridge
	Input     Input
}

func (c Console) Read(addr uint16) uint8 {

	switch {
	case addr >= 0x0000 && addr <= 0x1FFF:
		return c.Cpu.Ram[addr]

	case addr >= 0x2000 && addr <= 0x3FFF:

		// TODO: should account for mirroring
		switch addr {
		case 0x2000:
			return c.Ppu.Ctrl
		case 0x2001:
			return c.Ppu.Mask
		case 0x2002:
			return c.Ppu.Status
		case 0x2003:
			return c.Ppu.OamAddr
		case 0x2004:
			return c.Ppu.OamData
		case 0x2005:
			return c.Ppu.Scroll
		case 0x2006:
			return c.Ppu.Addr
		case 0x2007:
			return c.Ppu.Data
		case 0x4014:
			return c.Ppu.OamDma
		}

	// TODO: handle input
	case addr == 0x4016 || addr == 0x4017:
		return 0

	case addr >= 0x8000 && addr <= 0xFFFF:
		return c.Cartridge.Read(addr)

	default:
		log.Fatalf("Console Read unknown addr %x", addr)
	}

	return 0
}

func (c Console) Write(addr uint16, data uint8) {
	switch {
	case addr >= 0x0000 && addr <= 0x1FFF:
		c.Cpu.Ram[addr] = data

	case addr >= 0x2000 && addr <= 0x3FFF:

		// TODO: should account for mirroring
		switch addr {
		case 0x2000:
			c.Ppu.Ctrl = data
		case 0x2001:
			c.Ppu.Mask = data
		case 0x2002:
			c.Ppu.Status = data
		case 0x2003:
			c.Ppu.OamAddr = data
		case 0x2004:
			c.Ppu.OamData = data
		case 0x2005:
			c.Ppu.Scroll = data
		case 0x2006:
			c.Ppu.Addr = data
		case 0x2007:
			c.Ppu.Data = data
		}

	case addr == 0x4014:
		c.Ppu.OamDma = data

	// TODO: handle Apu
	case (addr >= 0x4000 && addr <= 0x4013) || addr == 0x4015 || addr == 0x4017:

	// TODO: handle input
	case addr == 0x4016 || addr == 0x4017:

	case addr >= 0x8000 && addr <= 0xFFFF:
		c.Cartridge.Write(addr, data)

	default:
		log.Fatalf("Console Write unknown addr %x", addr)
	}
}

func (c Console) ReadWord(addr uint16) uint16 {
	low := c.Read(addr)
	high := c.Read(addr + 1)
	return (uint16(high) << 8) | uint16(low)
}

func (c Console) WriteWord(addr uint16, data uint16) {
	low := uint8(data)
	high := uint8(data >> 8)
	c.Write(addr, uint8(low))
	c.Write(addr+1, uint8(high))
}

// func (c Console) ReadPpu(addr uint16) uint8 {

// 	// switch {
// 	// case addr >= 0x0000 && addr <= 0x1FFF:
// 	// 	return c.Cartridge.ChrRom[addr]

// 	// default:
// 	// 	log.Fatal("Console Read unknown addr", addr)
// 	// }

// 	return 0
// }

// func (c Console) WritePpu(addr uint16, data uint8) {

// 	// switch {
// 	// case addr >= 0x0000 && addr <= 0x1FFF:
// 	// 	return c.Cartridge.ChrRom[addr]

// 	// default:
// 	// 	log.Fatal("Console Read unknown addr", addr)
// 	// }
// }

// func WritePpuBUS(mem Memory, addr uint16, data uint8) {
// 	if addr >= 0x0000 && addr <= 0x1FFF {
// 		mem.PpuBUS[addr&0x07FF] = data
// 	}

// 	if addr >= 0x2000 && addr <= 0x3FFF {
// 		mem.PpuBUS[addr&0x0007] = data
// 	}
// }

func (cpu *Cpu) setNegativeZeroFlags(value uint8) {
	if value&Negative == Negative {
		cpu.Status |= Negative
	} else {
		cpu.Status &= ^Negative
	}

	if value == 0 {
		cpu.Status |= Zero
	} else {
		cpu.Status &= ^Zero
	}

}

func adc(cpu *Cpu, operand uint8) {
	result := uint16(cpu.Accumulator) + uint16(operand)

	if cpu.Status&Carry == Carry {
		result += 1
	}

	if result&0x00FF == 0 {
		cpu.Status |= Zero
	} else {
		cpu.Status &= ^Zero
	}

	if result&uint16(Negative) == uint16(Negative) {
		cpu.Status |= Negative
	} else {
		cpu.Status &= ^Negative
	}

	if result&0xFF00 > 0 {
		cpu.Status |= Carry
	} else {
		cpu.Status &= ^Carry
	}

	if ((uint16(operand)^result)&(uint16(cpu.Accumulator)^result))&0x80 != 0 {
		cpu.Status |= Verflow
	} else {
		cpu.Status &= ^Verflow
	}

	cpu.Accumulator = uint8(result)
}

// TODO: handle the NMI by pulling
type Cpu struct {
	ProgramCounter uint16
	StackPointer   uint8
	Accumulator    uint8
	IndirectX      uint8
	IndirectY      uint8
	Status         uint8
	Ram            [2 * 1024]uint8
}

func (cpu *Cpu) Reset(c Console) {
	cpu.ProgramCounter = c.ReadWord(0xFFFC)
	cpu.StackPointer = 0xFF
	cpu.Status = 0x00
}

func (cpu *Cpu) PushStack(data uint8) {
	// TODO: check if needed
	// cpu.Status |= Break2
	cpu.Ram[0x0100+uint16(cpu.StackPointer)] = data
	cpu.StackPointer--
}

func (cpu *Cpu) PopStack() uint8 {
	cpu.StackPointer++
	return cpu.Ram[0x0100+uint16(cpu.StackPointer)]

}

func (cpu *Cpu) Nmi() {
	// push processor status
	// push return addr
	// on stack
}

func (cpu *Cpu) GetOperandAddr(m Console, instruction Instruction) uint16 {

	var operand uint16

	switch instruction.Mode {
	case Immediate:
		operand = cpu.ProgramCounter + 1

	case ZeroPage:
		operand = uint16(m.Read(cpu.ProgramCounter + 1))

	case ZeroPageX:
		zeroPageAddr := m.Read(cpu.ProgramCounter + 1)
		// use the wrapping of the uint8
		zeroPageAddr += cpu.IndirectX
		operand = uint16(zeroPageAddr)

	case ZeroPageY:
		zeroPageAddr := m.Read(cpu.ProgramCounter + 1)
		// use the wrapping of the uint8
		zeroPageAddr += cpu.IndirectY
		operand = uint16(zeroPageAddr)

	case Absolute:
		operand = m.ReadWord(cpu.ProgramCounter + 1)

	case AbsoluteX, AbsoluteX1:
		operand = m.ReadWord(cpu.ProgramCounter + 1)
		operand += uint16(cpu.IndirectX)

	case AbsoluteY, AbsoluteY1:
		operand = m.ReadWord(cpu.ProgramCounter + 1)
		operand += uint16(cpu.IndirectY)

	case Indirect:
		target := m.ReadWord(cpu.ProgramCounter + 1)
		low := m.Read(target)

		// this is replicating an wrapping error on the original 6502
		// https://www.nesdev.org/obelisk-6502-guide/reference.html#JMP
		if target&0x00FF == 0x00FF {
			target &= 0xFF00
		} else {
			target += 1
		}

		high := m.Read(target)
		operand = (uint16(high) << 8) | uint16(low)

	case IndirectX:
		target := m.Read(cpu.ProgramCounter + 1)
		target += cpu.IndirectX
		// TODO: caused a problem because the operand wasn't wrapping
		// target := m.ReadWord(uint16(operand))
		low := m.Read(uint16(target))
		target++
		high := m.Read(uint16(target))
		operand = (uint16(high) << 8) | uint16(low)

	case IndirectY, IndirectY1:
		target := m.Read(cpu.ProgramCounter + 1)
		// TODO: caused a problem because the operand wasn't wrapping
		// target := m.ReadWord(uint16(operand))
		low := m.Read(uint16(target))
		target++
		high := m.Read(uint16(target))
		operand = (uint16(high) << 8) | uint16(low)
		operand += uint16(cpu.IndirectY)
	}

	return operand
}

// TODO: check for negative displacement
func (cpu *Cpu) Branch(m Console, condition bool) {
	if condition {
		displacement := int8(m.Read(cpu.ProgramCounter + 1))
		cpu.ProgramCounter = uint16(int16(cpu.ProgramCounter) + int16(displacement))
	}
}

func (cpu *Cpu) ExecuteInstruction(m Console) Instruction {
	opcode := m.Read(cpu.ProgramCounter)

	instruction := Instructions[opcode]
	operandAddr := cpu.GetOperandAddr(m, instruction)

	switch instruction.Opcode {
	// TODO: will need to write special comments for that
	// TODO: /IRQ /NMI PHP and BRK set B flags bit https://www.nesdev.org/wiki/Status_flags#The_B_flag
	// ts 5 and 4 when reading flags from the stack in the PLP or RTI instruction.
	case BRK_IMP:
		if opcode != uint8(BRK_IMP) {
			log.Fatalf("\nUNKNOW INSTRUCTION\ninstruction%#v\ncpu: %#v\noperandaddr: %#v\noperand: %#v", m.Read(cpu.ProgramCounter), cpu, operandAddr, m.Read(operandAddr))
		}
		return instruction

	case RTI_IMP:
		cpu.Status = cpu.PopStack()
		// TODO: maybe use the function BytesToWord or PopStackWord for ram
		low := cpu.PopStack()
		high := cpu.PopStack()
		cpu.ProgramCounter = (uint16(high) << 8) | uint16(low)
		// TODO: lean why need to do this also in PLP
		cpu.Status &= ^Break
		cpu.Status |= Break2
		return instruction

	// https://www.righto.com/2012/12/the-6502-overflow-flag-explained.html
	case ADC_IMM, ADC_ZER, ADC_ZRX, ADC_ABS, ADC_ABX, ADC_ABY, ADC_IDX, ADC_IDY:
		adc(cpu, m.Read(operandAddr))

	case SBC_IMM, SBC_ZER, SBC_ZRX, SBC_ABS, SBC_ABX, SBC_ABY, SBC_IDX, SBC_IDY:
		value := m.Read(operandAddr)
		// TODO: works but don't know why
		adc(cpu, (^value))

	// TODO: check to simplify
	case AND_IMM, AND_ZER, AND_ZRX, AND_ABS, AND_ABX, AND_ABY, AND_IDX, AND_IDY:
		operand := m.Read(operandAddr)
		operand &= cpu.Accumulator
		cpu.setNegativeZeroFlags(operand)
		cpu.Accumulator = operand

	// TODO: check to simplify
	case ORA_IMM, ORA_ZER, ORA_ZRX, ORA_ABS, ORA_ABX, ORA_ABY, ORA_IDX, ORA_IDY:
		operand := m.Read(operandAddr)
		operand |= cpu.Accumulator
		cpu.setNegativeZeroFlags(operand)
		cpu.Accumulator = operand

	case EOR_IMM, EOR_ZER, EOR_ZRX, EOR_ABS, EOR_ABX, EOR_ABY, EOR_IDX, EOR_IDY:
		operand := m.Read(operandAddr)
		operand ^= cpu.Accumulator
		cpu.setNegativeZeroFlags(operand)
		cpu.Accumulator = operand

		// TODO: maybe need to check again
	case BIT_ZER, BIT_ABS:
		memoryValue := m.Read(operandAddr)
		cpu.Status &= 0b0011_1111
		cpu.Status |= (memoryValue & 0b1100_0000)

		if cpu.Accumulator&memoryValue == 0 {
			cpu.Status |= Zero
		} else {
			cpu.Status &= ^Zero
		}

	case LSR_ACC:
		if cpu.Accumulator&Carry == Carry {
			cpu.Status |= Carry
		} else {
			cpu.Status &= ^Carry
		}

		cpu.Accumulator >>= 1
		cpu.setNegativeZeroFlags(cpu.Accumulator)

		// TODO: maybe can be rewritten to support all type
	case LSR_ZER, LSR_ZRX, LSR_ABS, LSR_ABX:
		operand := m.Read(operandAddr)
		if operand&Carry == Carry {
			cpu.Status |= Carry
		} else {
			cpu.Status &= ^Carry
		}

		operand >>= 1
		m.Write(operandAddr, operand)
		cpu.setNegativeZeroFlags(operand)

	case ASL_ACC:
		if cpu.Accumulator&Negative == Negative {
			cpu.Status |= Carry
		} else {
			cpu.Status &= ^Carry
		}

		cpu.Accumulator <<= 1
		cpu.setNegativeZeroFlags(cpu.Accumulator)

	case ASL_ZER, ASL_ZRX, ASL_ABS, ASL_ABX:
		operand := m.Read(operandAddr)
		if operand&Negative == Negative {
			cpu.Status |= Carry
		} else {
			cpu.Status &= ^Carry
		}

		operand <<= 1
		m.Write(operandAddr, operand)
		cpu.setNegativeZeroFlags(operand)

	case ROR_ACC:
		bit0 := cpu.Accumulator & Carry
		cpu.Accumulator >>= 1

		if cpu.Status&Carry == Carry {
			cpu.Accumulator |= Negative
		} else {
			cpu.Accumulator &= ^Negative
		}

		if bit0 > 0 {
			cpu.Status |= Carry
		} else {
			cpu.Status &= ^Carry
		}

		cpu.setNegativeZeroFlags(cpu.Accumulator)

	case ROR_ZER, ROR_ZRX, ROR_ABS, ROR_ABX:
		operand := m.Read(operandAddr)

		bit0 := operand & Carry
		operand >>= 1

		if cpu.Status&Carry == Carry {
			operand |= Negative
		} else {
			operand &= ^Negative
		}

		if bit0 > 0 {
			cpu.Status |= Carry
		} else {
			cpu.Status &= ^Carry
		}

		m.Write(operandAddr, operand)
		cpu.setNegativeZeroFlags(operand)

	case ROL_ACC:
		bit7 := cpu.Accumulator & Negative
		cpu.Accumulator <<= 1

		if cpu.Status&Carry == Carry {
			cpu.Accumulator |= Carry
		} else {
			cpu.Accumulator &= ^Carry
		}

		if bit7 > 0 {
			cpu.Status |= Carry
		} else {
			cpu.Status &= ^Carry
		}

		cpu.setNegativeZeroFlags(cpu.Accumulator)

	case ROL_ZER, ROL_ZRX, ROL_ABS, ROL_ABX:
		operand := m.Read(operandAddr)

		bit7 := operand & Negative
		operand <<= 1

		if cpu.Status&Carry == Carry {
			operand |= Carry
		} else {
			operand &= ^Carry
		}

		if bit7 > 0 {
			cpu.Status |= Carry
		} else {
			cpu.Status &= ^Carry
		}

		m.Write(operandAddr, operand)
		cpu.setNegativeZeroFlags(operand)

	// TODO: check for negative displacement
	case BCC_REL:
		cpu.Branch(m, cpu.Status&Carry == 0)
	case BCS_REL:
		cpu.Branch(m, cpu.Status&Carry == Carry)
	case BEQ_REL:
		cpu.Branch(m, cpu.Status&Zero == Zero)
	case BMI_REL:
		cpu.Branch(m, cpu.Status&Negative == Negative)
	case BNE_REL:
		cpu.Branch(m, cpu.Status&Zero == 0)
	case BPL_REL:
		cpu.Branch(m, cpu.Status&Negative == 0)
	case BVS_REL:
		cpu.Branch(m, cpu.Status&Verflow == Verflow)
	case BVC_REL:
		cpu.Branch(m, cpu.Status&Verflow == 0)
	case CLD_IMP:
		cpu.Status &= ^Decimal
	case CLV_IMP:
		cpu.Status &= ^Verflow

	// TODO: check if can use setFlags
	case CMP_IMM, CMP_ZER, CMP_ZRX, CMP_ABS, CMP_ABX, CMP_ABY, CMP_IDX, CMP_IDY:
		operand := m.Read(operandAddr)
		// TODO: check maybe wrapping
		result := cpu.Accumulator - operand

		if cpu.Accumulator >= operand {
			cpu.Status |= Carry
		} else {
			cpu.Status &= ^Carry
		}

		if result == 0 {
			cpu.Status |= Zero
		} else {
			cpu.Status &= ^Zero
		}

		cpu.Status &= ^Negative
		cpu.Status |= (result & Negative)

	// TODO: check if can use setFlags
	case CPY_IMM, CPY_ZER, CPY_ABS:
		operand := m.Read(operandAddr)
		result := cpu.IndirectY - operand

		if cpu.IndirectY >= operand {
			cpu.Status |= Carry
		} else {
			cpu.Status &= ^Carry
		}

		if result == 0 {
			cpu.Status |= Zero
		} else {
			cpu.Status &= ^Zero
		}

		cpu.Status &= ^Negative
		cpu.Status |= (result & Negative)

	// TODO: check if can use setFlags
	case CPX_IMM, CPX_ZER, CPX_ABS:
		operand := m.Read(operandAddr)
		result := cpu.IndirectX - operand

		if cpu.IndirectX >= operand {
			cpu.Status |= Carry
		} else {
			cpu.Status &= ^Carry
		}

		if result == 0 {
			cpu.Status |= Zero
		} else {
			cpu.Status &= ^Zero
		}

		cpu.Status &= ^Negative
		cpu.Status |= (result & Negative)

	// TODO: change the format to first get the operand
	case LDA_IMM, LDA_ZER, LDA_ZRX, LDA_ABS, LDA_ABX, LDA_ABY, LDA_IDX, LDA_IDY:
		cpu.Accumulator = m.Read(operandAddr)
		cpu.setNegativeZeroFlags(cpu.Accumulator)

	case TAY_IMP:
		cpu.IndirectY = cpu.Accumulator
		cpu.setNegativeZeroFlags(cpu.IndirectY)

	case TAX_IMP:
		cpu.IndirectX = cpu.Accumulator
		cpu.setNegativeZeroFlags(cpu.IndirectX)

	case TYA_IMP:
		cpu.Accumulator = cpu.IndirectY
		cpu.setNegativeZeroFlags(cpu.Accumulator)

	case TXA_IMP:
		cpu.Accumulator = cpu.IndirectX
		cpu.setNegativeZeroFlags(cpu.Accumulator)

	case TSX_IMP:
		cpu.IndirectX = uint8(cpu.StackPointer)
		cpu.setNegativeZeroFlags(cpu.IndirectX)

	case TXS_IMP:
		cpu.StackPointer = cpu.IndirectX

	// TODO: maybe change the format to get the operand first
	case LDX_IMM, LDX_ZER, LDX_ZRY, LDX_ABS, LDX_ABY:
		cpu.IndirectX = m.Read(operandAddr)
		cpu.setNegativeZeroFlags(cpu.IndirectX)

	case LDY_IMM, LDY_ZER, LDY_ZRX, LDY_ABS, LDY_ABX:
		cpu.IndirectY = m.Read(operandAddr)
		cpu.setNegativeZeroFlags(cpu.IndirectY)

	case INY_IMP:
		cpu.IndirectY++
		cpu.setNegativeZeroFlags(cpu.IndirectY)

	case DEY_IMP:
		cpu.IndirectY--
		cpu.setNegativeZeroFlags(cpu.IndirectY)

	case INX_IMP:
		cpu.IndirectX++
		cpu.setNegativeZeroFlags(cpu.IndirectX)

	case INC_ZER, INC_ZRX, INC_ABS, INC_ABX:
		operand := m.Read(operandAddr)
		operand++
		m.Write(operandAddr, operand)
		cpu.setNegativeZeroFlags(operand)

	case DEX_IMP:
		cpu.IndirectX--
		cpu.setNegativeZeroFlags(cpu.IndirectX)

	case DEC_ZER, DEC_ZRX, DEC_ABS, DEC_ABX:
		operand := m.Read(operandAddr)
		operand--
		m.Write(operandAddr, operand)
		cpu.setNegativeZeroFlags(operand)

	case STX_ZER, STX_ZRY, STX_ABS:
		m.Write(operandAddr, cpu.IndirectX)

	case STY_ZER, STY_ZRX, STY_ABS:
		m.Write(operandAddr, cpu.IndirectY)

	case STA_ZER, STA_ZRX, STA_ABS, STA_ABX, STA_ABY, STA_IDX, STA_IDY:
		m.Write(operandAddr, cpu.Accumulator)

	case JMP_ABS, JMP_IND:
		cpu.ProgramCounter = operandAddr
		return instruction

	case JSR_ABS:
		rts := cpu.ProgramCounter + 2
		cpu.PushStack(uint8(rts >> 8))
		cpu.PushStack(uint8(rts))
		cpu.ProgramCounter = operandAddr
		return instruction

	case RTS_IMP:
		low := cpu.PopStack()
		high := cpu.PopStack()
		addr := (uint16(high) << 8) | uint16(low)
		cpu.ProgramCounter = addr

	case NOP_IMP:
	case CLC_IMP:
		cpu.Status &= ^Carry
	case SEC_IMP:
		cpu.Status |= Carry
	case SEI_IMP:
		cpu.Status |= Interrupt
	case SED_IMP:
		cpu.Status |= Decimal

	case PHA_IMP:
		cpu.PushStack(cpu.Accumulator)
	case PHP_IMP:
		cpu.PushStack(cpu.Status | Break2 | Break)

	// TODO: maybe remove the operand var
	case PLA_IMP:
		operand := cpu.PopStack()
		cpu.Accumulator = operand
		cpu.setNegativeZeroFlags(operand)

	case PLP_IMP:
		cpu.Status = cpu.PopStack()
		cpu.Status &= ^Break
		cpu.Status |= Break2

	default:
		log.Fatalf("unknown upcode %x", instruction.Opcode)
	}

	cpu.ProgramCounter += uint16(instruction.Bytes)
	return instruction
}
