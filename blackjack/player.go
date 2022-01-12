package main

type Player struct {
	name   string
	hand   []int
	result int
	chip   int
	com    bool
	status int //joined:1, skipped:2, surrendered:3
}
