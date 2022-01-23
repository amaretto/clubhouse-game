package main

import "github.com/rivo/tview"

type UI struct {
	*tview.Flex
}

type DealerUI struct {
	*tview.Flex
	hand HandView
	chip ChipView
}

type PlayerUI struct {
	*tview.Flex
	hand HandView
	chip ChipView
}

type HandView struct {
	*tview.TextView
}

type ChipView struct {
	*tview.TextView
}
