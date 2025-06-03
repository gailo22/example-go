package main

type Item struct{}

func main() {
	sem := make(chan struct{}, 10)
	bigList := make([]Item, 10000)
	for _, item := range bigList {
		sem <- struct{}{}
		go func(i Item) {
			defer func() { <-sem }()
			process(i)
		}(item)
	}
}

func process(i Item) {}
