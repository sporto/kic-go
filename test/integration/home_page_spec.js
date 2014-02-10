describe('Home Page', function(){
  // var browser;

  // before(function () {
  //    browser = new Zombie();
  // });

  // after(function () {
  //    browser.close();
  // });

	before(function (done) {
		browser.visit('http://localhost:9000', function () {
			done();
			expect(browser.statusCode).to.eq(200)
		});
	});

	it('shows welcome', function () {
		expect(browser.text("h1")).to.eq("Welcome")
	});

});
