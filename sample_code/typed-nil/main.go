package main

func main() {
	var p *int = nil
	var a any = p
	if a == nil {
		return
	}
}
