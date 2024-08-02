package nes_test

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/raphadam/nesgo/nes"
// )

// type sets struct {
// 	addr uint16
// 	data uint8
// }

// func TestDriver(t *testing.T) {
// 	testCases := []struct {
// 		desc     string
// 		prg      []nes.Opcode
// 		memStart nes.Memory
// 		memEnd   nes.Memory
// 		cpuStart nes.Cpu
// 		cpuEnd   nes.Cpu
// 	}{
// 		{
// 			desc: "STA_ZER",
// 			prg:  []nes.Opcode{nes.STA_ZER, 0x36},
// 			cpuStart: nes.Cpu{
// 				Accumulator: 0x78,
// 			},
// 			cpuEnd: nes.Cpu{
// 				Accumulator:    0x78,
// 				ProgramCounter: 0x0602,
// 			},
// 			memEnd: nes.Memory{CpuBus: [65536]uint8{
// 				0x0036: 0x78,
// 			}},
// 		},
// 		{
// 			desc: "STX_ZER",
// 			prg:  []nes.Opcode{nes.STX_ZER, 0x36},
// 			cpuStart: nes.Cpu{
// 				IndirectX: 0x78,
// 			},
// 			cpuEnd: nes.Cpu{
// 				IndirectX:      0x78,
// 				ProgramCounter: 0x0602,
// 			},
// 			memEnd: nes.Memory{CpuBus: [65536]uint8{
// 				0x0036: 0x78,
// 			}},
// 		},
// 		{
// 			desc: "JSR_ABS",
// 			prg:  []nes.Opcode{nes.JSR_ABS, 0x34, 0x12},
// 			cpuStart: nes.Cpu{
// 				StackPointer: 0xFF,
// 			},
// 			cpuEnd: nes.Cpu{
// 				StackPointer:   0xFD,
// 				ProgramCounter: 0x1234,
// 			},
// 			memEnd: nes.Memory{CpuBus: [65536]uint8{
// 				0x01FE: 0x02,
// 				0x01FF: 0x06,
// 			}},
// 		},
// 		{
// 			desc: "RTS_IMP",
// 			prg:  []nes.Opcode{nes.RTS_IMP},
// 			cpuStart: nes.Cpu{
// 				StackPointer: 0xFD,
// 			},
// 			cpuEnd: nes.Cpu{
// 				StackPointer:   0xFF,
// 				ProgramCounter: 0x0603,
// 			},
// 			memStart: nes.Memory{CpuBus: [65536]uint8{
// 				0x01FE: 0x02,
// 				0x01FF: 0x06,
// 			}},
// 			memEnd: nes.Memory{CpuBus: [65536]uint8{
// 				0x01FE: 0x02,
// 				0x01FF: 0x06,
// 			}},
// 		},
// 		{
// 			desc: "PHP_IMP",
// 			prg:  []nes.Opcode{nes.PHP_IMP},
// 			cpuStart: nes.Cpu{
// 				StackPointer: 0xFF,
// 				Status:       nes.Carry | nes.Negative | nes.Decimal,
// 			},
// 			cpuEnd: nes.Cpu{
// 				StackPointer:   0xFE,
// 				ProgramCounter: 0x0601,
// 				Status:         nes.Carry | nes.Negative | nes.Decimal,
// 			},
// 			memEnd: nes.Memory{CpuBus: [65536]uint8{
// 				0x01FF: nes.Carry | nes.Negative | nes.Decimal | nes.Break,
// 			}},
// 		},
// 		{
// 			desc: "PHA_IMP",
// 			prg:  []nes.Opcode{nes.PHA_IMP},
// 			cpuStart: nes.Cpu{
// 				Accumulator:  0x89,
// 				StackPointer: 0xFF,
// 			},
// 			cpuEnd: nes.Cpu{
// 				Accumulator:    0x89,
// 				StackPointer:   0xFE,
// 				ProgramCounter: 0x0601,
// 			},
// 			memEnd: nes.Memory{CpuBus: [65536]uint8{
// 				0x01FF: 0x89,
// 			}},
// 		},
// 		// {
// 		// 	desc: "LDA_IMM",
// 		// 	prg:  []nes.Opcode{nes.LDA_IMM, 0x24},
// 		// 	cpuEnd: nes.Cpu{
// 		// 		Accumulator:    0x24,
// 		// 		ProgramCounter: 0x0602,
// 		// 	},
// 		// },
// 		// {
// 		// 	desc: "LDA_IMM NEGATIVE",
// 		// 	prg:  []nes.Opcode{nes.LDA_IMM, 0xFC},
// 		// 	cpuEnd: nes.Cpu{
// 		// 		Accumulator:    0xFC,
// 		// 		ProgramCounter: 0x0602,
// 		// 		Status:         nes.Negative,
// 		// 	},
// 		// },
// 		// {
// 		// 	desc: "LDA_IMM ZERO",
// 		// 	prg:  []nes.Opcode{nes.LDA_IMM, 0x00},
// 		// 	cpuEnd: nes.Cpu{
// 		// 		Accumulator:    0x00,
// 		// 		ProgramCounter: 0x0602,
// 		// 		Status:         nes.Zero,
// 		// 	},
// 		// },
// 		// {
// 		// 	desc: "LDX_IMM",
// 		// 	prg:  []nes.Opcode{nes.LDX_IMM, 0x24},
// 		// 	cpuEnd: nes.Cpu{
// 		// 		IndirectX:      0x24,
// 		// 		ProgramCounter: 0x0602,
// 		// 	},
// 		// },
// 		// {
// 		// 	desc: "LDX_IMM NEGATIVE",
// 		// 	prg:  []nes.Opcode{nes.LDX_IMM, 0xFC},
// 		// 	cpuEnd: nes.Cpu{
// 		// 		IndirectX:      0xFC,
// 		// 		ProgramCounter: 0x0602,
// 		// 		Status:         nes.Negative,
// 		// 	},
// 		// },
// 		// {
// 		// 	desc: "LDX_IMM ZERO",
// 		// 	prg:  []nes.Opcode{nes.LDX_IMM, 0x00},
// 		// 	cpuEnd: nes.Cpu{
// 		// 		IndirectX:      0x00,
// 		// 		ProgramCounter: 0x0602,
// 		// 		Status:         nes.Zero,
// 		// 	},
// 		// },
// 		// {
// 		// 	desc: "JMP_ABS",
// 		// 	prg:  []nes.Opcode{nes.JMP_ABS, 0xf0, 0x10},
// 		// 	cpuEnd: nes.Cpu{
// 		// 		ProgramCounter: 0x10f0,
// 		// 	},
// 		// },
// 	}
// 	for _, tC := range testCases {

// 		t.Run(tC.desc, func(t *testing.T) {
// 			tC.cpuStart.ProgramCounter = 0x0600

// 			for i, opcode := range tC.prg {
// 				tC.memStart.CpuBus[0x0600+i] = uint8(opcode)
// 				tC.memEnd.CpuBus[0x0600+i] = uint8(opcode)
// 			}

// 			for range tC.prg {
// 				tC.cpuStart.ExecuteInstruction(&tC.memStart)
// 			}

// 			if !reflect.DeepEqual(tC.cpuStart, tC.cpuEnd) {
// 				t.Errorf("\nwant: %#v\ngot : %#v\n", tC.cpuEnd, tC.cpuStart)
// 			}

// 			if !reflect.DeepEqual(tC.memStart, tC.memEnd) {
// 				for i := range tC.memStart.CpuBus {
// 					if tC.memStart.CpuBus[i] != tC.memEnd.CpuBus[i] {
// 						t.Errorf("memory is not equal got [%x]: %x, want [%x]: %x", i, tC.memStart.CpuBus[i], i, tC.memEnd.CpuBus[i])
// 					}
// 				}

// 				// t.Errorf("\nwant: %#v\ngot : %#v\n", tC.memEnd, tC.memStart)
// 			}
// 		})

// 	}
// }

// // for _, set := range tC.sets {
// // 	mem[set.addr] = set.prg
// // }
