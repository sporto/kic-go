define(['logger', 'can'], function(logger, can) {

	'use strict';
	
	logger.info('Router loaded');

	// can.route('accounts/:id/:action', {type: 'accounts'});

	// can.route(':type/:id');

	// can.route.bind('id', function(ev, newVal, oldVal) {
 //    console.log('The hash\'s id changed.');
	// });

	return can.Control({

		// 'route': 'index',
		// 'accounts route': 'accounts',

		// 'accounts/:id route': function (data) {
		// 	logger.info('route accounts/id');
		// 	logger.info(data.id);
		// },

		// index: function () {
		// 	console.log('index');
		// },

		// accounts: function () {
		// 	logger.info('account');
		// },

		// "{can.route} id" : function(route, ev, newVal){
		// 	console.log(route);
		// }
	});

});