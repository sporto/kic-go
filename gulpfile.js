var gulp      = require('gulp');
var jshint    = require('gulp-jshint');
var less      = require('gulp-less');
var ngmin     = require('gulp-ngmin');
var concat    = require('gulp-concat');
var clean     = require('gulp-clean');
var uglify    = require('gulp-uglify');
var path      = require('path');
// var uglify = require('gulp-uglify');
// var imagemin = require('gulp-imagemin');

gulp.task('lint', function() {
	gulp.src('./src/public/js/app/**/*.js')
		.pipe(jshint())
		.pipe(jshint.reporter(require('jshint-stylish')));
});

// compile less files
gulp.task('less', function () {
	gulp.src('./src/public/css/app/main.less')
		.pipe(less({
			paths: [ path.join(__dirname, 'less', 'includes') ]
		}))
		.pipe(gulp.dest('./src/public/css/app'));
});

gulp.task('dist:clean', function () {
	gulp.src('./dist2', {read: false})
		.pipe(clean({force: true}));
});

// copy css to dist
gulp.task('dist:css', function () {
	//less and copy
	gulp.run('less');
	gulp
		.src('./src/public/css/app/main.css')
		.pipe(gulp.dest('./dist2/public/css/app'));

	// concat lib css and copy
	gulp
		.src('./src/public/css/lib/**/*.css')
		.pipe(concat('lib.css'))
		.pipe(gulp.dest('./dist2/public/css/lib'));
});

gulp.task('dist:js', function () {
	gulp.src('./src/public/js/app/**/*.js')
		.pipe(concat('app.js'))
		.pipe(ngmin())
		.pipe(uglify())
		.pipe(gulp.dest('./dist2/public/js'));

	gulp.src('./src/public/js/lib/**/*.js')
		.pipe(concat('lib.js'))
		.pipe(uglify())
		.pipe(gulp.dest('./dist2/public/js'));
});

gulp.task('dist:views', function () {
	gulp.src('./src/public/views/**/*.html')
		.pipe(gulp.dest('./dist2/public/views'))
});

gulp.task('dist:index', function () {
	gulp.src('./src/index.html')
		.pipe(gulp.dest('./dist2'));
});

gulp.task('dist', function () {
	gulp.run('lint');
	gulp.run('dist:clean');
	gulp.run('dist:css');
	gulp.run('dist:js');
	gulp.run('dist:views');
	gulp.run('dist:index');
});