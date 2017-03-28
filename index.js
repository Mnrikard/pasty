#!/usr/bin/env node

var args = process.argv;
args.shift();
args.shift();

if(process.stdin.isTTY){
	var clipboard = require("copy-paste");
	var content = clipboard.paste();
	clipboard.copy(handleInput(content, args));
} else {
	var pipedInput = '';
	process.stdin.on('readable', function() {
		var chunk = this.read();
		if(chunk !== null){
			pipedInput += chunk;
		}
	});
	process.stdin.on('end', function() {
	   console.log(handleInput(pipedInput, args));
	});
}

function handleInput(str, args) {
	var editor = getEditor(args);
	return str;
}

function getEditor(args){
	if(args.length === 0 || args[0].toLower() === "help") {
		return new Help(args);
	}
}

function Help(args){
	
}


