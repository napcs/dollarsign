mac:
	mkdir -p dist/mac
	env GOOS=darwin GOARCH=386 go build -o dist/mac/$$

ubuntu:
	mkdir -p dist/linux
	env GOOS=linux GOARCH=386 go build -o dist/linux/$$

clean:
	rm -rf dist/
