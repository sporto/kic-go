'use strict';

angular.module('APP')
  .controller('AccountsIndexCtrl', function ($scope, Account, logger, notifier) {
    logger.info('Getting accounts');
    // notifier.success('Getting accounts');
    $scope.accounts = [];

    Account.all()
      .then(function (accounts) {
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
  .controller('AccountsShowCtrl', function ($scope, $routeParams, Account, logger, notifier) {
    // $scope.params = $routeParams;
    $scope.account = Account.one($routeParams.id);
  });