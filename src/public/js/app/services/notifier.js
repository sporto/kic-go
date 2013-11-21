(function() {

	'use strict';

	angular.module('APP')
		.factory('notifier', function() {
			return toastr;
		});

}());