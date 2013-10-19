define(['logger', 'can', 'js/app/controls/router', 'js/app/controls/accounts'], function(logger, can, Router, AccountsCtrl) {

	'use strict';
	
	logger.info('ApplControl loaded');

	return can.Control({
		defaults: {
			view: 'app-template'
		}
	},{
		init: function (el, options) {
			var state = can.Map({title: 'Hello'});
			this.element.append(can.view(this.options.view, state));

			new Router(document);
			new AccountsCtrl($('.accounts', this.element));
		},

	});

});