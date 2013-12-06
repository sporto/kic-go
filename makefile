default:
	go test -v ./...

dist: dist-clean dist-copy dist-less dist-jsmin dist-cssmin

dist-clean:
	-rm -r ./.tmp
	-rm -r ./dist

dist-copy:
	-mkdir ./dist
	-cp ./src/index.html ./dist/index.html

dist-less:
	-lessc ./src/public/css/app/main.less ./src/public/css/app/main.css

dist-jsmin:
	-uglifyjs ./src/public/js/app/**/*.js -o ./dist/js/app.js -m
	-uglifyjs ./src/public/js/lib/**/*.js -o ./dist/js/lib.js -m

dist-cssmin: