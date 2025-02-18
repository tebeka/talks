package main

func loop(n int) {
	loop(n + 1)
}

func main() {
	loop(0)
}
