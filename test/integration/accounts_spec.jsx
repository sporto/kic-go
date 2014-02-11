var r = require('rethinkdb');

describe('Accounts', function() {

	var account;

	// before all
	before(function (done) {

		// add one account
		

		browser.visit('http://localhost:9000', function () {

			browser.clickLink('Accounts')
				.then(function () {
					done();
				});

		});
	});

	it('shows the account index page', function (done) {
		expect(browser.text("h1")).to.eq("Accounts");
	});

	it('shows the account link', function (done) {
		// expect(browser.)
	});

	// it('shows the account show page')

	// it('shows the account name')

	// it('shows the current balance')

	// it('shows the transactions')
});
