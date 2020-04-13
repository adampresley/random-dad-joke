/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

import { Joke } from "/app/modules/jokes/model/Joke.js";

export class JokeService {
	constructor($http) {
		this.$http = $http;
	}

	/**
	 * Returns a random dad-joke
	 * @returns {Promise<Joke>}
	 */
	async getRandomJoke() {
		let response = await this.$http.get(`/api/joke/random`);
		return this.mapJSONtoJoke(response.body);
	}

	mapJSONtoJoke(j) {
		return new Joke({
			id: j.id,
			joke: j.joke,
			status: j.status,
		});
	}
}

export function JokeServiceInstaller(Vue) {
	Vue.prototype.jokeService = new JokeService(Vue.http);

}
