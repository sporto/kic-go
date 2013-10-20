'use strict';

angular.module('APP')
	.controller('TransactionsNewCtrl', function($scope, $routeParams, logger, notifier, Account) {
		$scope.accountId = $routeParams.accountId;
		$scope.kind = $routeParams.kind;

		$scope.account = Account.one($scope.accountId).get();

		

	})