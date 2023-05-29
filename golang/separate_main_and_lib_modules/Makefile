
.PHONY: all clean run

all:
	mkdir -p genfiles
	$(MAKE) run

clean:
	rm -rf genfiles

run: genfiles/prog
	@echo 'Running....'
	@echo '--------------------------------------------------------------------------------'
	@genfiles/prog
	@echo '--------------------------------------------------------------------------------'

genfiles/prog: $(shell find main) $(shell find lib)
	go build -o $@ example.com/skeleton/main

