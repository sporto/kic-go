'use strict';

log.setLevel('info');

var app = angular.module('APP', ['restangular']);

app.constant('apiBase', '//localhost:5000');

app.config(function ($routeProvider) {
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
});

app.run(function(Restangular, apiBase) {
  Restangular.setBaseUrl(apiBase);
});
