'use strict';

angular.module('APP')
	.controller('TransactionsNewCtrl', function($scope, $routeParams, logger, notifier, Account) {
		logger.info('TransactionsNewCtrl');
		var id = $routeParams.accountId;

		$scope.account = {};
		$scope.state = {
			busy: true
		}

		Account.one(id).get()
			.then(function (account) {
				$scope.account = account;
				$scope.state.busy = false;
			});

		$scope.transaction = {
			accountId: $scope.account.id,
			amount: 0,
			kind: $routeParams.kind,
		}

		$scope.submit = function () {
			$scope.state.busy = true;
			$scope.account.post('transactions', $scope.transaction)
				.then(function () {
					notifier.success('Saved');
					$scope.state.busy = false;
				});
		}
	})