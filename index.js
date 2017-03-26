#!/usr/bin/env node

if(process.stdin.isTTY){
	var clipboard = require("copy-paste");
	var content = clipboard.paste();
	handleInput(content);
} else {
	var pipedInput = '';
	var piper = process.stdin;
	piper.on('readable', function() {
		var chunk = this.read();
		if(chunk !== null){
			pipedInput += chunk;
		}
	});
	piper.on('end', function() {
	   handleInput(pipedInput);
	});
}

function handleInput(str) {
	console.log(str);
}


