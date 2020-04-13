/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

const router = new VueRouter({
	routes: [
		{
			path: "/",
			name: "home",
			component: () => import("/app/pages/home.js"),
			meta: {
				title: "Home",
			},
		},
	],

	scrollBehavior: () => {
		return {
			x: 0,
			y: 0,
		};
	},
});

router.beforeEach(async (to, from, next) => {
	/*
	 * Here is where you can do things like check to see if the user has a valid
	 * session, setup global Vue objects on successful session validation, etc...
	 */

	document.title = `${to.meta.title} | Random Dad Joke`;
	return next();
});

export default router;
