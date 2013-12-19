describe('Home Page', function(){
  // var browser;

  // before(function () {
  //    browser = new Zombie();
  // });

  // after(function () {
  //    browser.close();
  // });

  before(function (done) {
  	// console.log(browser)
    browser.visit('http://localhost:9000', function () {
      done();
    });
  });

  it('shows welcome', function () {
    expect(browser.text("h1")).to.eq("Welcome")
  });

   it('shows welcome', function () {
    expect(1).to.eq(1)
  });

});
