/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

import { Joke } from "/app/modules/jokes/model/Joke.js";

export default {
	data() {
		return {
			version: "",
			joke: new Joke(),
		};
	},

	async created() {
		this.version = await this.versionService.GetVersion();
		this.joke = await this.jokeService.getRandomJoke();
	},

	template: `
		<div>
			<div class="row mt-4 justify-content-center">
				<div class="col-12">
					<p class="joke-text">{{joke.joke}}</p>
				</div>
			</div>

			<nav class="navbar fixed-bottom navbar-light bg-light">
				<p>Server version: <strong>{{version}}</strong></p>
			</nav>
		</div>
	`
};

