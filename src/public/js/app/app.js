(function() {
	'use strict';

	// log.setLevel('info');

	var app = angular.module('APP', ['ngRoute', 'restangular']);

	app.constant('apiBase', '/api');

	app.config(function($routeProvider, $httpProvider, RestangularProvider) {

		$routeProvider
			.when('/', {
				templateUrl: 'public/views/main.html',
				controller: 'MainCtrl'
			})
			.when('/accounts', {
				templateUrl: 'public/views/accounts/index.html',
				controller: 'AccountsIndexCtrl'
			})
			.when('/accounts/:accountId', {
				templateUrl: 'public/views/accounts/show.html',
				controller: 'AccountsShowCtrl'
			})
			.when('/accounts/:accountId/transactions/new', {
				templateUrl: 'public/views/transactions/new.html',
				controller: 'TransactionsNewCtrl'
			})
			.when('/login', {
				templateUrl: 'public/views/sessions/new.html',
				controller: 'SessionsNewCtrl'
			})
			.otherwise({
				redirectTo: '/'
			});

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

}());