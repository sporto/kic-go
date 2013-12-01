desc('This is the default task.');
task('default', function (params) {
	console.log('This is the default task.');
});

namespace('dist', function () {
	task('run', function () {
		jake.Task['dist:lint'].invoke();
		jake.Task['dist:clean'].invoke();
	});
	task('lint', function () {
		console.log('clean');
	});
	task('clean', function () {
		console.log('clean');
	});
	task('copy', function () {
		console.log('clean');
	});
	task('less', function () {
		console.log('clean');
	});
	task('ngmin', function () {
		console.log('clean');
	});
	task('jsmin', function () {
		console.log('clean');
	});
	task('cssmin', function () {
		console.log('clean');
	});
});