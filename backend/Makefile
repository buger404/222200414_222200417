.PHONY: run
run:
	go build main.go router.go router_gen.go
	rm -f main.pid
	./main & echo $$! > main.pid

.PHONY: stop
stop:
	kill `cat main.pid` && rm main.pid