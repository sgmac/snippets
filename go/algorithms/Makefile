

all: bubble

bubble: bubble_sort.go
	go build -o $@ $^

benchmark:
	go test -bench=.
clean:
	rm bubble
