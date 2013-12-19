before(function () {
	console.log('starting rethinkdb');
	console.log('starting go');
	global.browser = new Zombie();
});

after(function () {
	console.log('closing rethinkdb');
	console.log('closing go');
	global.browser.close();
});