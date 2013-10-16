'use strict';

angular.module('APP')
  .factory('Account', function (Restangular) {
    return {
      all: Restangular.all('accounts')
    }
  });
