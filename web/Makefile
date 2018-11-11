.PHONY: install
install:
	@npm i
	@npm i -g browserify

.PHONY: build
build:
	@browserify main.js > bundle.js

.PHONY: serve
serve:
	@python -m SimpleHTTPServer 3000
