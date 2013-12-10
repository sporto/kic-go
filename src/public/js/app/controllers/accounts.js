(function() {
	'use strict';

	angular.module('APP')
		.controller('AccountsIndexCtrl', function($scope, Account, logger, notifier) {
			logger.info('Getting accounts');
			$scope.accounts = [];

			Account.all().getList()
				.then(function(accounts) {
						logger.info(accounts);
						$scope.accounts = accounts;
					},
					function errorCallback(err) {
						logger.info('error');
						logger.info(err);
						notifier.error(err);
					});

		})
		.controller('AccountsShowCtrl', function($scope, $routeParams, Account, logger, notifier, authServ) {
			$scope.id = $routeParams.accountId;
			$scope.state = {
				busy: false
			};
			$scope.canDo = authServ.canDo;

			getAccount();
			getTransactions();

			$scope.refreshInterest = refreshInterest;

			function getAccount() {
				logger.info('Getting account', $scope.id);
				Account.one($scope.id).get()
					.then(function (account) {
						$scope.account = account;
					});
			}

			function getTransactions() {
				// get the latest transactions
				Account.one($scope.id).getList('transactions')
					.then(function (transactions) {
						$scope.transactions = transactions;
						drawChart();
					});
			}

			function refreshInterest() {
				$scope.state.busy = true;
				Account.one($scope.id).customPOST({}, 'adjust')
					.then(function () {
						$scope.state.busy = false;
						getTransactions();
					}, function (response) {
						$scope.state.busy = false;
						notifier.error(response.data.e);
					});
			}

			function drawChart() {
				var records = $scope.transactions.slice(0).reverse();
				
				var labels = _.chain(records)
					.pluck('createdAt')
					.map(function (d) {
						return moment(d).format("YYYY-MM-DD");
					})
					.value();

				var values = _.pluck(records, 'balance');

				var data = {
					labels : labels,
					datasets : [
						{
							fillColor : "rgba(151,187,205,0.5)",
							strokeColor : "rgba(151,187,205,1)",
							pointColor : "rgba(151,187,205,1)",
							pointStrokeColor : "#fff",
							data : values
						}
					]
				};
				var ctx = document.getElementById("chart").getContext("2d");
				var myNewChart = new Chart(ctx).Line(data);
			}

		});
}());