'use strict';

angular.module('APP', [])
  .config(function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl'
      })
      .when('/accounts', {
        templateUrl: 'views/accounts.html',
        controller: 'AccountsIndexCtrl'
      })
      .otherwise({
        redirectTo: '/'
      });
  });
