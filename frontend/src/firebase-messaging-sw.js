'use strict';
 
importScripts("https://www.gstatic.com/firebasejs/8.2.4/firebase-app.js");
importScripts("https://www.gstatic.com/firebasejs/8.2.4/firebase-messaging.js");
 
const FIREBASE_CONFIG = {
	apiKey: "",
	authDomain: "",
	projectId: "",
	storageBucket: "",
	messagingSenderId: "",
	appId: "",
	measurementId: ""
};
 
// Initialize the firebase in the service worker.
firebase.initializeApp(FIREBASE_CONFIG);
 
self.addEventListener('push', function (event) {
	var data = event.data.json();
 
	const title = data.Title;
	data.Data.actions = data.Actions;
	const options = {
		body: data.Message,
		data: data.Data
	};
	event.waitUntil(self.registration.showNotification(title, options));
});
 
self.addEventListener('notificationclick', function (event) {});
 
self.addEventListener('notificationclose', function (event) {});