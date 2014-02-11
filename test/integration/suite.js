before(function () {
	console.log('starting rethinkdb');
	console.log('starting go');
	global.browser = new Zombie();

	// r.connect({host: 'localhost', port: 28015}, function(err, conn) {
	// 	if (err) throw err;
	// 	global.dbConn = conn;
	// });

});

after(function () {
	console.log('closing rethinkdb');
	console.log('closing go');
	global.browser.close();
});