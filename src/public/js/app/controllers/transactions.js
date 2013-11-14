'use strict';

angular.module('APP')
	.controller('TransactionsNewCtrl', function($scope, $routeParams, $location, logger, notifier, Account) {
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

			if ($scope.transaction.kind === 'deposit') {
				$scope.transaction.credit = $scope.transaction.amount;
				$scope.transaction.debit = 0;
			} else {
				$scope.transaction.debit = $scope.transaction.amount;
				$scope.transaction.credit = 0;
			}

			$scope.account.post('transactions', $scope.transaction)
				.then(function () {
					notifier.success('Saved');
					$location.path('/accounts/' + $scope.account.id);
				}, function (response) {
					notifier.error(response.data.e);
					$scope.state.busy = false;
				});
		}
	})