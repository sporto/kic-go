define(['logger', 'can'], function (logger, can) {

	"use strict";

	logger.info('Account model loaded');

	return can.Model({
		models: 'd',
		findAll: 'GET /api/accounts'
	}, {});

});