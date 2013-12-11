(function() {
	'use strict';

	angular.module('APP')
		.controller('SessionsNewCtrl', function (authServ, $scope, notifier, $location) {

			$scope.user = {
				password: ''
			};

			$scope.login = function () {
				if ($scope.user.password === '123') {
					notifier.success("Logged in");
					authServ.setLoggedIn(true);
					$location.path('/');
				} else {
					notifier.error('Wrong');
				}
			};
		});

}());