define(['logger', 'can', 'js/app/controls/router', 'js/app/controls/accounts'], function(logger, can, Router, AccountsCtrl) {

	'use strict';

	logger.info('ApplControl loaded');

	return can.Control({
		defaults: {
			view: 'app-template'
		}
	}, {
		init: function(el, options) {
			var state = can.Map({
				title: 'Hello'
			});
			this.element.append(can.view(this.options.view, state));

			new AccountsCtrl($('.accounts', this.element));

			can.route.ready(false);

		},

		'accounts route': 'accounts',
		'accounts/:id route': 'account',

		accounts: function () {
			logger.info('accounts');
		},

		account: function (data) {
			logger.info('route accounts/id');
			logger.info(data.id);
		},

	});

});