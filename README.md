![Logo](docs/nesgologo.png)

## Introduction

Pure Go implementation of the NES CPU emulator, capable of running binary CPU instructions designed for the 6502 architecture. This project accurately emulates the 6502's clock cycles, registers and memory layout.

## Features

- **Accurate Emulation:** Implements all registers, including stack pointer, program counter, and index registers (X, Y) and precise clock cycle emulation.
- **Memory Management:** Handles memory layout and mappers, aiming for valid and accurate emulation.
- **Extendable Cartridge Support:** Allows for custom cartridges with the potential to integrate new chipsets and support a wide range of programs.
- **Game Engine Integration:** Includes a client with a custom game engine to play a Snake game, written entirely in 6502 assembly.

## Demo

You can try the webassembly program in my personal website or build the project yourself.

### Installation

1. Clone the repository:

```bash
git clone https://github.com/raphadam/nesgo.git
cd nesgo
```

Run the project:

```bash
go mod tidy
go run .
```

## Roadmap

- Add more sample games.
- Implement support for additional 6502-compatible chipsets.
- Improve documentation with more code examples.
- Integrate a debugger for step-by-step execution analysis.