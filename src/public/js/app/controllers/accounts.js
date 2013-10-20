'use strict';

angular.module('APP')
	.controller('AccountsIndexCtrl', function($scope, Account, logger, notifier) {
		logger.info('Getting accounts');
		// notifier.success('Getting accounts');
		$scope.accounts = [];

		Account.all().getList()
			.then(function(accounts) {
					logger.info(accounts)
					$scope.accounts = accounts;
					// notifier.success(accounts);
				},
				function errorCallback(err) {
					logger.info('error');
					logger.info(err);
					notifier.error(err);
				});

		

	})
	.controller('AccountsShowCtrl', function($scope, $routeParams, Account, logger, notifier) {
		var id = $routeParams.id;

		$scope.account = Account.one($routeParams.id).get();
		// get the latest transactions
		// $scope.transactions = Account.one(id).all('transactions').getList();
	});