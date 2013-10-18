requirejs.config({
	baseUrl: 'public/js/lib',
	paths: {
		app: '../app'
	}
});

require(["can", "app/controls/accounts"], function(can, AccountsCtrl) {

	'use strict';

	// log.setLevel('info');
	var Control = can.Control({
		defaults: {
			view: 'app-template'
		}
	},{
		init: function (el, options) {
			console.log('dkd')
			console.log(this.options.view)
			var state = can.Map({title: 'Hello'});
			this.element.append(can.view(this.options.view, state));

			new AccountsCtrl($('.accounts', this.element));
		}
	});

	new Control('#app');
	
});