package nes

import (
	"errors"
	"log"
	"os"

	"golang.org/x/exp/slices"
)

var (
	INES_HEADER = [4]uint8{0x4E, 0x45, 0x53, 0x1A}
)

const (
	INES_HEADER_SIZE   = 16
	INES_PRG_BANK_SIZE = 16 * 1024
	INES_CHR_BANK_SIZE = 8 * 1024
	INES_TRAINER_SIZE  = 512
)

type MapperId uint8

const (
	NROM MapperId = 0x0000
)

type Mirroring uint8

const (
	Vertical Mirroring = iota
	Horizontal
	FourScreen
)

type Cartridge struct {
	PrgBanks  int // 16KB
	ChrBanks  int // 8KB
	PrgRom    []uint8
	ChrRom    []uint8
	Mirroring Mirroring
	MapperId  MapperId
}

func (c Cartridge) mapper(addr uint16, chr func(mapAddr uint16), prg func(mapAddr uint16)) {
	switch c.MapperId {
	case NROM:
		// TODO: PRG_RAM Family basic
		// TODO: Nametable mirroring

		if addr >= 0x0000 && addr <= 0x1FFF {
			chr(addr)
			return
		}

		if addr >= 0x8000 && addr <= 0xBFFF {
			prg(addr - 0x8000)
			return
		}

		if addr >= 0xC000 && c.PrgBanks == 1 {
			prg(addr - 0xC000)
			return
		}

		if addr >= 0xC000 && c.PrgBanks == 2 {
			prg(addr - 0x8000)
			return
		}

		log.Fatal("mapper cartridge unmapped addr", addr)

	default:
		panic("unknown mapper")
	}
}

// TODO: change this to avoid func param
func (c Cartridge) Read(addr uint16) uint8 {
	var data uint8

	chr := func(mapAddr uint16) {
		data = c.ChrRom[mapAddr]
	}

	prg := func(mapAddr uint16) {
		data = c.PrgRom[mapAddr]
	}

	c.mapper(addr, chr, prg)
	return data
}

func (c Cartridge) Write(addr uint16, data uint8) {
	chr := func(mapAddr uint16) {
		c.ChrRom[mapAddr] = data
	}

	prg := func(mapAddr uint16) {
		c.PrgRom[mapAddr] = data
	}

	c.mapper(addr, chr, prg)
}

func LoadRom(path string) (Cartridge, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Cartridge{}, err
	}

	// TODO: support other file type
	// TODO: detect ines vs nes 2.0 format
	if !slices.Equal(INES_HEADER[:], data[:4]) {
		return Cartridge{}, errors.New("must be an ines file")
	}

	cartidge := Cartridge{}
	cartidge.PrgBanks = int(data[4])
	cartidge.ChrBanks = int(data[5])
	cartidge.MapperId = MapperId((data[6] >> 4) | ((data[7] >> 4) << 4))

	// TODO: check for battery backed PRG RAM
	if data[6]&0b0000_00010 == 0b0000_00010 {
		return Cartridge{}, errors.New("PrgRam not supported")
	}

	// TODO: support trainer
	if data[6]&0b0000_00100 == 0b0000_00100 {
		return Cartridge{}, errors.New("Trainer not supported")
	}

	if data[6]&0b0000_0001 == 0b0000_0001 {
		cartidge.Mirroring = Vertical
	} else {
		cartidge.Mirroring = Horizontal
	}
	if data[6]&0b0000_1000 == 0b0000_1000 {
		cartidge.Mirroring = FourScreen
	}

	totalPRGBankSize := INES_PRG_BANK_SIZE * cartidge.PrgBanks
	cartidge.PrgRom = make([]uint8, totalPRGBankSize)
	cartidge.PrgRom = slices.Replace(cartidge.PrgRom, 0, totalPRGBankSize, data[INES_HEADER_SIZE:INES_HEADER_SIZE+totalPRGBankSize]...)

	totalCHRBankSize := INES_CHR_BANK_SIZE * cartidge.ChrBanks
	cartidge.ChrRom = make([]uint8, totalCHRBankSize)
	cartidge.ChrRom = slices.Replace(cartidge.ChrRom, 0, totalCHRBankSize, data[INES_HEADER_SIZE+totalPRGBankSize:INES_HEADER_SIZE+totalPRGBankSize+totalCHRBankSize]...)

	return cartidge, nil
}
