.PHONY: build-css
build-css: tailwindcss
	./tailwindcss -i tailwind.css -o static/css/tailwind.css --minify

tailwindcss:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-x64
	chmod +x tailwindcss-macos-x64
	mkdir -p bin
	mv tailwindcss-macos-x64 bin/tailwindcss

.PHONY: watch-css
watch-css: tailwindcss
	./bin/tailwindcss -i tailwind.css -o public/styles/tailwind.css --watch