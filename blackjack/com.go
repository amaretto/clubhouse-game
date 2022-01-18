package main

type Com struct {
	strtgy Strategy
}

type Strategy struct {
}

type Aggresive struct {
	Strategy
}

type Guarded struct {
	Strategy
}

type Random struct {
	Strategy
}
