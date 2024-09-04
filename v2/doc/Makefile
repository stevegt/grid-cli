all: ship

ship: README.md
	git add -i
	grok commit | git commit -F -
	git push

.PHONY: README.md

README.md:
	go run README/main.go . > README.md
