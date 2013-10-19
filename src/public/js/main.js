require(['logger', 'js/app/controls/appl'], function(logger, ApplControl) {

	'use strict';

	logger.info('APP');
	
	new ApplControl('#app');
	
});