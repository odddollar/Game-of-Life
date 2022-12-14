# Game of Life

A simple implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) in Go.

The game creates a grid of cells that are randomly initialised to either **Alive** or **Dead**. It then steps through subsequent generations where a cell's state is determined by simple rules:

1. If an **Alive** cell has fewer than 2 or more than 3 living neighbours it becomes **Dead**
2. If a **Dead** cell has exactly 3 living neighbours it becomes **Alive**
3. If an **Alive** cell has 2 or 3 living neighbours it continues living in the next generation

Can be built to an executable with:
```
git clone https://github.com/odddollar/Game-of-Life.git
cd Game-of-Life
go build Game-of-Life
```

