package main

type DanceInstructor struct {
	Position Position
	Name     string
}

type Position struct {
	x         int
	y         int
	Direction Direction
}

type Direction struct {
	Direction string
}

type Dancefloorsize struct {
	x int //0 to 5
	y int
}

type DanceType struct {
	//to be implemented later
	Genera           string //eg freestyle
	Qualifiedteacher bool
	Rating           int //ie rate the level the dancer can go at
}
