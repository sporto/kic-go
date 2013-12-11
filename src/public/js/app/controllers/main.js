(function() {
	'use strict';

	angular.module('APP')
		.controller('MainCtrl', function($scope, authServ, notifier) {

			$scope.$watch(authServ.isLoggedIn, function (isLoggedIn) {
				// console.log('isLoggedIn', isLoggedIn)
				$scope.currentUser = authServ.getCurrentUser();
			});

			$scope.logout = function () {
				authServ.logOut();
				notifier.success("Logged out");
				$location.path('/');
			};

		});
}());