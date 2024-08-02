package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/raphadam/nesgo/nes"
)

// NESTEST

type StateTest struct {
	ProgramCounter uint16
	Accumulator    uint8
	IndirectX      uint8
	IndirectY      uint8
	Cycles         int
	StackPointer   uint8
	Status         uint8
}

// func RenderTests() {
// 	cartidge, err := nes.LoadRom("nestest.nes")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	cpu := nes.Cpu{}
// 	ppu := nes.Ppu{}
// 	console := nes.Console{
// 		Cpu:       &cpu,
// 		Ppu:       &ppu,
// 		Cartridge: cartidge,
// 	}
// 	cpu.Reset(console)

// 	ebiten.SetWindowSize(640, 480)
// 	ebiten.SetWindowTitle("RenderTests")
// 	ebiten.SetVsyncEnabled(true)

// 	if err := ebiten.RunGame(&RenderTest{cpu: &cpu, ppu: &ppu, car: &cartidge, console: console}); err != nil {
// 		log.Fatal("ERROR IS", err)
// 	}
// }

func NesTests() {

	cartridge, err := nes.LoadRom("nestest.nes")
	if err != nil {
		log.Fatal("unable to load rom", err)
	}

	// TEST LOG
	logFile, err := os.Open("nestest.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	states := []StateTest{}

	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {
		line := scanner.Text()

		pc, err := strconv.ParseUint(line[:4], 16, 16)
		if err != nil {
			log.Fatal("unable to parse")
		}

		a, err := strconv.ParseUint(line[50:52], 16, 8)
		if err != nil {
			log.Fatal("unable to parse")
		}

		x, err := strconv.ParseUint(line[55:57], 16, 8)
		if err != nil {
			log.Fatal("unable to parse")
		}

		y, err := strconv.ParseUint(line[60:62], 16, 8)
		if err != nil {
			log.Fatal("unable to parse")
		}

		cycles, err := strconv.ParseUint(line[90:], 10, 32)
		if err != nil {
			log.Fatal("unable to parse")
		}

		sp, err := strconv.ParseUint(line[71:73], 16, 16)
		if err != nil {
			log.Fatal("unable to parse")
		}

		status, err := strconv.ParseUint(line[65:67], 16, 8)
		if err != nil {
			log.Fatal("unable to parse")
		}

		inst := StateTest{
			ProgramCounter: uint16(pc),
			Accumulator:    uint8(a),
			IndirectX:      uint8(x),
			IndirectY:      uint8(y),
			Cycles:         int(cycles),
			StackPointer:   uint8(sp),
			Status:         uint8(status),
		}
		states = append(states, inst)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("scanner error", err)
	}

	cpu := nes.Cpu{
		ProgramCounter: 0xC000,
		StackPointer:   0xFD,
		Status:         0x24,
	}

	m := nes.Console{
		Cpu:       &cpu,
		Cartridge: cartridge,
	}

	for i, state := range states {
		if cpu.ProgramCounter != state.ProgramCounter {
			log.Fatalf("\ni: %d\nProgramCounter\ntrue: %#02v\ncurr: %#02v", i, state.ProgramCounter, cpu.ProgramCounter)
		}

		if cpu.Accumulator != state.Accumulator {
			log.Fatalf("\ni: %d\nAccumulator\ntrue: %#02v\ncurr: %#02v", i, state.Accumulator, cpu.Accumulator)
		}

		if cpu.IndirectX != state.IndirectX {
			log.Fatalf("\ni: %d\nIndirectX\ntrue: %#02v\ncurr: %#02v", i, state.IndirectX, cpu.IndirectX)
		}

		if cpu.IndirectY != state.IndirectY {
			log.Fatalf("\ni: %d\nIndirectY\ntrue: %#02v\ncurr: %#02v", i, state.IndirectY, cpu.IndirectY)
		}

		if cpu.StackPointer != state.StackPointer {
			log.Fatalf("\ni: %d\nStackPointer\ntrue: %#02v\ncurr: %#02v", i, state.StackPointer, cpu.StackPointer)
		}

		if cpu.Status != state.Status {
			log.Fatalf("\ni: %d\nStatus\ntrue: %#02v\ncurr: %#02v", i, state.Status, cpu.Status)
		}

		// TODO: check on memory insertion
		// TODO: check cycles
		cpu.ExecuteInstruction(m)
	}
}
