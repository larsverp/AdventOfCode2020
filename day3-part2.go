package main

import (
	"fmt"
)

var field = []string{
	"..#.......#..##...#...#..#.#...",
	"..##..#..#.....#.........#....#",
	"...#.##..#.#......#.#....#.....",
	"...#.....#......#...#..........",
	".......#.#..#..#....##....##...",
	".#......#......#.#..#....#.#...",
	".#..........#.....###.##..#.#..",
	"....#...##...........#.........",
	"##......#.#...#...#....##.#...#",
	".....#.....#.#..#....###...#..#",
	".#....#.........#...#.......#.#",
	".##......#.#.........#....#.#..",
	"...#........###..#......##....#",
	".#.....###..........#...#....##",
	"............#......#...#...#..#",
	".....#.#....#.#.#...........#.#",
	"....#.....#...##.##.....###..#.",
	".....#.##......#.##.#...#.#.#..",
	"##.....#.##.##....#..##......#.",
	".....#.....#........#........##",
	"#..#...#.#.#..##....#....##..#.",
	".#......................##..#..",
	"#.#.#................#.....#...",
	"..#....#.#.#..##..#..#.....##..",
	".....###..#.............#....##",
	"..#....#.#...#......###..#.....",
	".......####.....##......#......",
	"..##...#..#...#.#..#......#....",
	"..##..#..#.....#......##.##....",
	".#......#................#..#.#",
	"...............#....#.......#..",
	"#...........#.#.#........#.#...",
	"...#...#..#.....##..#.##...#.#.",
	".#.#..#.#.....#...#..##....#.##",
	".........#...#.#.#....#.......#",
	"............##.#..#.#........#.",
	"#..#..#........#......####.....",
	"..#.#...#...#...#.#...#..#..#..",
	"#....................#.#.#....#",
	".......#.....#....#....##..#...",
	".......#.............##.....#..",
	"..##..#.......#..#.........#..#",
	"..##.#........#...#...#..#..#..",
	"..##.#...#................#..#.",
	"..#.....##...##......###.....#.",
	".....#.....#......#......#.....",
	"....##.#.#....#.....#..........",
	"...#...#.#.....#.###.....#....#",
	"......##....#..##...##.#.......",
	"......................##......#",
	".##......#...#...#...##.#.##.#.",
	"#.......#...#.....#........#.#.",
	"...............#......#.......#",
	".#..##...#....#.....###....##..",
	"...#..###...#..#....#.#..#.....",
	".#...#....#.................#..",
	".......##....#..##......#.#....",
	"....#.....#.....###...#....#...",
	"..##..#..##........#...........",
	"....#..#.#............#........",
	"####.....##.........#....#.....",
	"........#.....#......#......##.",
	"#..........#........#....#...##",
	"#..#...###.....##.....#.##...#.",
	"......#....##.............#..#.",
	"..#......###...#...#..##.....#.",
	"#.##...#......##.###....#....#.",
	"...............#....#.....###.#",
	"#......#........#.#..##.#.....#",
	"#..............#..##.#....#....",
	".##....###...#...#.#....#....#.",
	".......#...#.......#....#..##..",
	"..##.#.....#..#...............#",
	".##..............#.......#...#.",
	".....#...#.#..#.........#..#...",
	"........#.#.###......#..#......",
	"..#.......###..#...#...........",
	"............#.#.....#....#...#.",
	"#...#..#......#..#......#......",
	"..##..#......#..#.......#.....#",
	"............#.##.....#.#...#...",
	"....#.#...#.#...#........##....",
	"........##...#...##.....#.#..#.",
	".#..........#.##...............",
	"###.#..#...###..###..#........#",
	"....#..#............#...#.#.#..",
	".#....###........#.......#...#.",
	"..........#..#.....#......#....",
	"..##..#.#....#..#.....##....#..",
	"...#...............#......#....",
	"......#.#.#...###.....##.#.....",
	".#...#.#.#.#....#.....#..#..#.#",
	"..#.....#..##....#......#.#.#..",
	"##.#....#.......#.#.#.......#..",
	".#.#.#........##.#....#........",
	"...#..#...#.#.........#..#....#",
	"#.#......#.#.#..#.#............",
	"......#.....#.....#.......#..#.",
	".........#..#.##..#..###....#.#",
	".......#..........#..#.........",
	"......#.#.##...#.#...##........",
	"...........##..##.##....#......",
	"..#..#...#.###...##.....#..#..#",
	".#...##.......#.#..........#..#",
	"##......#...........#.#........",
	"..#..#.....##..#.#.......#...#.",
	".#....#..#.....###.......#...#.",
	"...#..#...#...###.#.###..#.....",
	".......#...#.##......#.#.#.#...",
	".#.......#...#...#...##........",
	"...#........#....#..#.#...#.#..",
	"..#............#.....#.#....#..",
	"#.....#.#..#.#....#...#.#.#....",
	"#......#......##.#...........#.",
	".#..#.........##.#........#....",
	".#....#.#...#........#......#..",
	"....#..#..#....#..#.#.....#..#.",
	".##.#.....#..#.#...#.....#..###",
	"#..#......#..##....##..#.......",
	".......#..##....#.###..........",
	".#......#..........#........#..",
	".........#.....#......#.#......",
	".....#..#.......#...#.#....#...",
	".#......#...#..#......#.....#.#",
	"#.#.#..#.........#.......##....",
	"...#..#.....#.#..#......#...#..",
	".##...#...#.#....#.....#.#...#.",
	"..#......##......#....#.#..#...",
	"....#.#.....#..#...........#...",
	".#........#....#....#...#......",
	"..#.#.....#.......#.#.#.##..##.",
	"...#..#.##.......#...#.........",
	"....##.#.#..#.#..#.#.#..#.#.#.#",
	"......#.......#.....#...##.#...",
	"..#.##.....#....#...#...##.##.#",
	"..##.........##.........#..#...",
	".....###......####..#.#...#....",
	"...#....#....##.............#.#",
	".#.........#.......##..#.#.....",
	"...#..#........#...............",
	"........#........##......#.#...",
	"##...#......#.....#.##........#",
	".............#.#........#......",
	".......##.........#.......#....",
	"....#.......#......#..###....#.",
	".#.##........#.....#......#....",
	"#...........##...#..##.........",
	".....#.#........#........##...#",
	"#.........#..##.....#...#....#.",
	".........#...####..#....###....",
	"###.#..##.......#....#....#.##.",
	"...........#.....#.#...#..#....",
	".##......###.#.#.#....#........",
	"....#..................#.###.#.",
	"#..##...#......#...#......###..",
	".#.....#..##......#.#.#.#......",
	"..##...#..#.......##.#.......#.",
	"...#..#..##..##..##.#..####....",
	"......#..#..............#.#####",
	"....#....#..#...#...#.#........",
	".###...#.......#..#........#...",
	"........#.#......##...#........",
	".#..#.......##.......#.....##..",
	"...##..........##...........#..",
	"......................#..#.#.#.",
	"#..#.....#......#.....#....#.##",
	"..#......#.....#....#...#.##.#.",
	"............#...#...#......#.##",
	"........##.......#......##.....",
	"..#.##.....#....#..##...#......",
	"........#.#...##.#..#...#....#.",
	"...##............##.....#..#...",
	"...###.##....#....#.#.#.#..#...",
	"............#..#....#..#.....##",
	"...#......##.......#......##..#",
	".......#......#........#.....#.",
	".#....#.##..........#..........",
	"..###.........#..#...##.....#..",
	"##....##..#.......###....#.#.#.",
	"#.......#.#.##.................",
	"..#..........#................#",
	"....#.#....#...###...#......#..",
	".#..........#..#..##.....#...##",
	".#.....###....#.#...#.#........",
	".........##.........#..#.#.....",
	".#.....##....#......##.#....#..",
	"###..#...#..#.........#......#.",
	"..##.....#....#............#...",
	".....##.##....#.....#..#..#....",
	".......###...#..........#......",
	"...#........#....#.##.#........",
	"........###....##........##.#..",
	"....#....#........#.....##.....",
	".#........#.#........#...#..#..",
	"#..#..#......#....##.#..###....",
	"...........#...#...#....#.#...#",
	".#..#.....#.##..#......##......",
	"..#.#...............##..##.###.",
	"...#.#...#............##.#..#..",
	".#.#....#....#................#",
	"...#..#.#.##.#.#....#......#...",
	".........#..#.......##.##.#....",
	"..............#..##.#.....#....",
	"......#.........#...#...##.#..#",
	"....##.....#.#...#.#.####.#....",
	"#..#.#....#.##.......#....#....",
	"...##....#...................##",
	"##.#.......#....#.#.........#.#",
	".##.#..###...#.#.........#.....",
	"...#.#.#..#...#...#.##....#..#.",
	"....#.....###...#....##........",
	".............#....#....#....#..",
	"#...#.....#.#...#.#............",
	"#.#....#...........##.......#..",
	"..#..#.........#....###.......#",
	"....###..................##...#",
	".#........#.....##...#......#..",
	"#..#......#..#.....#.##..#...#.",
	"....#........#.##.#......#.....",
	"#.#...#...............####...##",
	"#.#.....##..#.####.............",
	"##...##..#.##........#.#...#...",
	"...#...........#............#..",
	"...#..#..#........##...........",
	"..#.##..#.#...#...#..#......#..",
	".....##......#...##.....##.....",
	".......##.....##....#..........",
	"..........#.#...............#.#",
	"#.#..........#..#..##..#...#..#",
	".#.#..#.###................#...",
	"#...#...#....#...#....#.##.##..",
	".#................###.....##...",
	".....##.#.....##.#.....#.....##",
	".......#.......#......#.#......",
	"..#....#......#.....#.....#..#.",
	"#......#..##..####.....#.#..#..",
	".......#...###..#...#.....#....",
	"#.#.#......#.............#..###",
	".#.#.......#..#.#..#..#...#.#..",
	"..#.#......#......#.#....#.....",
	"..#..###..#.#.....#....#.......",
	"..........#........#......#..#.",
	"##..........##....##..###.##...",
	"...#....#.......#..##.......#..",
	"##...#............##...#.#.#.#.",
	"...........#...#.#....#.....##.",
	".#................#.......#...#",
	"....##.#..#.#.........###.###.#",
	"#..#...###..#...#......#..##...",
	"..##........#.#..##.#..........",
	"...#..#...#...............#.#..",
	"##.##....#...#........#...#....",
	"#..#......#.##.................",
	".....#..##.##.......#..........",
	"...#.....#........#......#.....",
	".....#..#......#.....##...#....",
	"#......#...###....#....###.....",
	"................###..#..#..#.##",
	"......#.......#..........#.#..#",
	"###..#..#.##.............#.#...",
	"....##.#.......#...#..##.......",
	"..#.....##..#..#..#.#..........",
	".#.......#.#..#........##......",
	"..............#.........#......",
	"..#........##....#.#.#......##.",
	".#.#.........#.#...#.#.........",
	"..........#..#.##.#....#...#.#.",
	"............##....#.....###...#",
	"#....#..#...#.#...#.....#....#.",
	".#...##.....#......#..#........",
	".#..#...###.#..##........#...#.",
	"#..#...##.####.......#.....#...",
	"#.##..#...#......#.#.......#.#.",
	"#.#.....##...#.#...............",
	"#...........##.##.#.........#..",
	"...#...........#.#.#.#....#..#.",
	"..#......#.#.#....##..#.#.....#",
	".........#.#.##...##...#.#.#...",
	"...........#..#.####.#..#.#.###",
	"#........#.#......#..#...#.....",
	"...#....#......##...#.#........",
	"......#.#....#.##....#....#..##",
	".......###......#.#.....#......",
	"#..........##..................",
	".###.......##..#..##...##......",
	"##.#..............#....#....##.",
	"#....#.....#.##.....#..#....#.#",
	".......#.......#..#..#..##..#..",
	"......#...........#..##....#...",
	".....#.......#..#......#..#..##",
	".##...#......#........#....#...",
	".....#..#....#...............##",
	"..#.....#....#..#.##....#.#....",
	"#.#......####...#..#.........#.",
	"#.#........#..#........#...#...",
	"..#..............#.......###.#.",
	"...#.#....##.#......#........#.",
	"....###.......##.#.....##.....#",
	".........#......#.#.......##.##",
	".......#.#....#.#...#...#....#.",
	"....#....#....#.#.......##.....",
	"#...#.....#..#.....#...........",
	"#...#..#...#.#..#.............#",
	"........###..#........#........",
	".............#....#..###..#.#.#",
	"#...............#..#.#.........",
	"##.....###..#......#...#....#..",
	".#...................##.#.##...",
	"#..#........#.#...#..#...#.....",
	"................#.#.........#..",
	"#.....##.#.#...#..#.........##.",
	"...#.....#....#..####.#........",
	"....#.#...........#............",
	".....#........##..........#....",
	"..#.......#.#.#.####..#......#.",
	"#.......#...##.#.#..#.#.#......",
	"##........#.#...###............",
	"..##........#........#...#.#.#.",
	"#.#..#.#.......#....#........#.",
	"..............#....#.........#.",
	"#....#.#........#.............#",
	"..##...#..........#........#...",
	"..#..#..#....#....#............",
	"",
}

func main() {
	ch := make(chan int)

	go calculateTrees(1, 1, ch)
	go calculateTrees(3, 1, ch)
	go calculateTrees(5, 1, ch)
	go calculateTrees(7, 1, ch)
	go calculateTrees(1, 2, ch)

	total := <-ch * <-ch * <-ch * <-ch * <-ch
	fmt.Printf("Encountered trees is %d \n", total)
}

func calculateTrees(rigthStep int, downStep int, ch chan int) {
	encounteredTrees := 0
	currentRow := 0
	currentCol := 0
	for currentRow < len(field)-(downStep+1) {
		currentRow += downStep
		if currentCol > len(field[currentRow])-(rigthStep+1) {
			currentCol = rigthStep - (len(field[currentRow]) - currentCol)
		} else {
			currentCol += rigthStep
		}
		if string(field[currentRow][currentCol]) == "#" {
			encounteredTrees++
		}
	}
	ch <- encounteredTrees
}
