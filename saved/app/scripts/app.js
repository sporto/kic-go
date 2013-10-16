'use strict';

log.setLevel('info');

var app = angular.module('APP', ['restangular']);

app.constant('apiBase', '//localhost:5000');

app.config(function ($routeProvider, $httpProvider, RestangularProvider) {

  $routeProvider
    .when('/', {
      templateUrl: 'views/main.html',
      controller: 'MainCtrl'
    })
    .when('/accounts', {
      templateUrl: 'views/accounts/index.html',
      controller: 'AccountsIndexCtrl'
    })
    .otherwise({
      redirectTo: '/'
    });

  // Deal with CORS issues
  // $httpProvider.defaults.useXDomain = true;
  // delete $httpProvider.defaults.headers.common['X-Requested-With'];

  RestangularProvider.setResponseExtractor(function(response, operation, what, url) {
    // This is a get for a list
    var newResponse;
    if (operation === "getList") {
      // Here we're returning an Array which has one special property metadata with our extra information
      newResponse = response.d;
      // newResponse.metadata = response.data.meta;
    } else {
      // This is an element
      newResponse = response.d;
    }
    return newResponse;
  });

});

app.run(function(Restangular, apiBase) {
  Restangular.setBaseUrl(apiBase);
});