get-tailwindcss:
ifeq (,$(wildcard ./bin/tailwindcss))
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-x64
	chmod +x tailwindcss-macos-x64
	mkdir -p bin
	mv tailwindcss-macos-x64 bin/tailwindcss
endif

get-air:
ifeq (,$(wildcard ./bin/air))
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
endif

.PHONY: build-css
build-css: get-tailwindcss
	./bin/tailwindcss -i tailwind.css -o public/styles/tailwind.css --minify

.PHONY: watch-css
watch-css: get-tailwindcss
	./bin/tailwindcss -i tailwind.css -o public/styles/tailwind.css --watch

watch-go: get-air
	./bin/air

.PHONY: dev
dev: 
	$(MAKE) -j 2 watch-go watch-css