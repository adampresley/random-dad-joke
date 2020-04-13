/*
 * Copyright (c) 2020. Adam Presley All Rights Reserved
 */

export class DateTimeService {
	constructor() {}

	formatDateTime(input) {
		return moment(input).format("YYYY-MM-DD h:mma");
	}

	nowUTC() {
		return moment().utc();
	}

	parse(dateString) {
		return moment(dateString);
	}

	toISO8601(date) {
		return moment(date).format("YYYY-MM-DDTHH:mm:ssZZ");
	}

	toSQLString(date) {
		return moment(date).format("YYYY-MM-DD HH:mm:ss");
	}

	toUSDate(date) {
		return moment(date).format("MM/DD/YYYY");
	}

	toUSDateTime(date) {
		return moment(date).format("MM/DD/YYYY h:mm A");
	}

	toUSTime(date) {
		return moment(date).format("h:mm A");
	}
}

export function DateTimeServiceInstaller(Vue) {
	Vue.prototype.dateTimeService = new DateTimeService();
}
