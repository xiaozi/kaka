#!/usr/bin/env casperjs

var casper = require("casper").create({
	verbose: true,
	logLevel: "info",
	stepTimeout: 1 * 60 * 1000, // 1 min
	pageSettings: {
		userAgent: ""
	}
});

var url = casper.cli.get("url");
var path = casper.cli.get("path");
var useragent = casper.cli.get("useragent");
var width = casper.cli.get("width");
var height = casper.cli.get("height");

if (useragent) {
	casper.userAgent(useragent);
}

casper.start().viewport(width, height).thenOpen(url, function () {
	this.capture(path);
	this.echo(JSON.stringify({
		url: this.getCurrentUrl(),
		path: path
	}));
});

casper.run();
