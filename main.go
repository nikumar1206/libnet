package main

func main() {
	c := initFlags()

	conn := NewConnection(*c)
	repl := NewREPL(conn)

	repl.Run()
}
