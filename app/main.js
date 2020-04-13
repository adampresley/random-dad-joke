/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

import router from "/app/router.js";
import { AlertServiceInstaller } from "/app/modules/ui/services/AlertService.js";
import { DateTimeServiceInstaller } from "/app/modules/datetime/services/DateTimeService.js";
import { InstallAppHttpInterceptors, InstallGlobalHttpInterceptors } from "/app/HttpInterceptors.js";
import { VersionServiceInstaller } from "/app/modules/version/services/VersionService.js";
import { JokeServiceInstaller } from "/app/modules/jokes/services/JokeService.js";

/*
 * Core plugins
 */
Vue.use(VueRouter);
Vue.use(VueResource);
Vue.use(VueLoading);

/*
 * Services
 */
Vue.use(AlertServiceInstaller);
Vue.use(DateTimeServiceInstaller);
Vue.use(VersionServiceInstaller);
Vue.use(JokeServiceInstaller);

Vue.component("loading", VueLoading);

InstallGlobalHttpInterceptors(Vue);

new Vue({
	el: "#app",
	router,

	async beforeCreate() {
		InstallAppHttpInterceptors(Vue, this);
	},

	template: `
	<div style="width: 100%">
		<header></header>

		<main role="main" class="flex-shrink-0 main-body">
			<div class="container-fluid">
				<router-view></router-view>
			</div>
		</main>
	</div>
	`
});


