define(['can', 'app/models/account'], function (can, Account) {

	"use strict";

	return can.Control({
		init: function (ele, options) {
			var self = this;

			this.accounts = new Account.List([]);
			this.state = can.Map({

			});

			this.element.html(can.view('accounts-template', {accounts: this.accounts, state: this.state}));

			console.log('getting accounts');

			Account.findAll({}, function (accounts) {
				console.log('Got accounts')
				console.log(accounts);
				self.accounts.replace(accounts);
			}, function (xhr) {
				console.log('errorp')
				console.log(xhr);
			});
		}
	});

});