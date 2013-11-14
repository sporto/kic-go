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
		$scope.id = $routeParams.accountId;

		logger.info('Getting account', $scope.id);
		Account.one($scope.id).get()
			.then(function (account) {
				$scope.account = account;
			});

		// get the latest transactions
		Account.one($scope.id).getList('transactions')
			.then(function (transactions) {
				$scope.transactions = transactions;
			});
	});