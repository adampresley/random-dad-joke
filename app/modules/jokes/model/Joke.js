/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

export class Joke {
	constructor(config = {
		id: "",
		joke: "",
		status: 0,
	}) {
		this.id = config.id || "";
		this.joke = config.joke || "";
		this.status = config.status || 0;
	}
}
