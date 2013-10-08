'use strict';

angular.module('APP')
  .controller('AccountsIndexCtrl', function ($scope, Account, logger, notifier) {
    logger.info('Getting accounts');
    notifier.success('Getting accounts');

    Account.getList()
      .then(function (accounts) {
        logger.info('success');
        logger.info(accounts);
        notifier.success(accounts);
      }, 
      function errorCallback(err) {
        logger.info('error');
        logger.info(err);
        notifier.error(err);
      });

  });
