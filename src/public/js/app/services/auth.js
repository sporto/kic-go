(function() {

	'use strict';

	angular.module('APP')
		.factory('authServ', function() {
			var loggedIn = false;

			function canDo(resource, action) {
				return isLoggedIn();
			}
			function canDoFor(model, action) {

			}
			function isLoggedIn() {
				return loggedIn;
			}
			function setLoggedIn(val) {
				loggedIn = val;
			}
			function logOut() {
				loggedIn = false;
			}
			function getCurrentUser(){
				if (loggedIn) {
					return {username: 'Sam'};
				} else {
					return null;
				}
			}
			return {
				canDo:          canDo, //move this to authorization service
				canDoFor:       canDoFor,
				isLoggedIn:     isLoggedIn,
				setLoggedIn:    setLoggedIn,
				logOut:         logOut,
				getCurrentUser: getCurrentUser
			}
		});

}());