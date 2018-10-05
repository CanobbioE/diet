let index = {
	init: function() {
		// Init
		asticode.loader.init();
		asticode.modaler.init();
		asticode.notifier.init();

		document.addEventListener('astilectron-ready', function() {
			index.listen();
			index.showmain();
		})
	},
	addMeal: function() {
		// TODO add meal
	},
	newFood: function() {
		// TODO new food item
	},
	listen: function() {
		astilectron.onMessage(function(message) {
			switch (message.name) {
				case "event.name":
					// TODO
					break;
			}
		});
	}
};
