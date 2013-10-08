'use strict';

angular.module('APP')
  .factory('Account', function (Restangular, apiBase) {
    return Restangular.all('accounts');
  });
