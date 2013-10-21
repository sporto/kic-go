define(['logger', 'can', '../../models/account'], function (logger, can, Account) {

	"use strict";

	logger.info('AccountsShowCtrl loaded');

	return can.Control({
		init: function (ele, options) {
			var self = this;

			this.element.html(can.view('account-show', {account: this.account}));
		}
	});

});