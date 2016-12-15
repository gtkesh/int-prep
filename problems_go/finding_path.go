package main

type Ticket struct {
	from string
	to   string
}

type Path struct {
	start *node
}

type node struct {
	name string
	prev *node
	next *node
}

func findPath(tickets []Ticket) *node {
	return nil
}
