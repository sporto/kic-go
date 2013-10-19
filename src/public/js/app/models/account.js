define(['can'], function (can) {

	"use strict";

	return can.Model({
		models: 'd',
		findAll: 'GET /api/accounts'
	}, {});

});