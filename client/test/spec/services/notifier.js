'use strict';

describe('Service: notifier', function () {

  // load the service's module
  beforeEach(module('kicClientApp'));

  // instantiate service
  var notifier;
  beforeEach(inject(function (_notifier_) {
    notifier = _notifier_;
  }));

  it('should do something', function () {
    expect(!!notifier).toBe(true);
  });

});
