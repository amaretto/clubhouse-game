package main

type Player struct {
	name   string
	hand   []Card
	result int
	chip   int
	bet    int
	com    bool
	status int //joined:1, skipped:2, surrendered:3
}
