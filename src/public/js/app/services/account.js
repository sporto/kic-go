(function () {

	'use strict';

	angular.module('APP')
		.factory('Account', function(Restangular) {
			return {
				all: function() {
					return Restangular.all('accounts');
				},
				one: function(id) {
					return Restangular.one('accounts', id);
				}
			};
		});

}());