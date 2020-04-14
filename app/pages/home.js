/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

import { Joke } from "/app/modules/jokes/model/Joke.js";

export default {
	data() {
		return {
			version: "",
			joke: new Joke(),
			jokeImage: "",
		};
	},

	async created() {
		this.version = await this.versionService.GetVersion();
		this.joke = await this.jokeService.getRandomJoke();
	},

	methods: {
		getRandomJokeImage() {
			const images = [
				"haha.jpeg",
				"troll-face.jpg",
			];

			const imageIndex = Math.floor(Math.random() * images.length);

			return `/app/assets/random-dad-joke/images/${images[imageIndex]}`;
		},
	},

	template: `
		<div>
			<div class="row mt-4 justify-content-center">
				<div class="col-12">
					<h4>Random Dad Joke</h4>
				</div>
			</div>

			<div class="row mt-4 justify-content-center">
				<div class="col-12">
					<p class="joke-text">{{joke.joke}}</p>
					<p class="mt-4 joke-image"><img :src="getRandomJokeImage()" width="30%" /></p>
				</div>
			</div>

			<nav class="navbar fixed-bottom navbar-light bg-light">
				<p>Server version: <strong>{{version}}</strong></p>
			</nav>
		</div>
	`
};

