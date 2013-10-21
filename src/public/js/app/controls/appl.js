define([
	'logger', 
	'can', 
	'js/app/controls/accounts/index',
	'js/app/controls/accounts/show',
	], 
	function(
		logger, 
		can, 
		AccountsIndexCtrl,
		AccountsShowCtrl
	) {

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

			this.$view = $('.view', this.element);

			can.route.ready(false);
		},

		'accounts route': 'accounts',
		'accounts/:id route': 'account',

		accounts: function () {
			logger.info('accounts');
			new AccountsIndexCtrl(this.$view);
		},

		account: function (data) {
			logger.info('route accounts/id');
			logger.info(data.id);
			new AccountsShowCtrl(this.$view);
		},

	});

});