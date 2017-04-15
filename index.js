#!/usr/bin/env node

var args = process.argv;
args.shift();
args.shift();

var editorRunner = require("./editorRunner.js");

debugger;
if(process.stdin.isTTY){
	var content = "";
	var clipboard = require("copy-paste");
	try{
		content = clipboard.paste();
	} catch(e) {
		clipboard.copy(content);
		//console.log("error in getting clipboard content:"+e);
	}
	clipboard.copy(editorRunner.handleInput(content, args));
	debugger;
} else {
	editorRunner.interactive = false;
	var pipedInput = '';
	process.stdin.on('readable', function() {
		var chunk = this.read();
		if(chunk !== null){
			pipedInput += chunk;
		}
	});
	process.stdin.on('end', function() {
	   console.log(editorRunner.handleInput(pipedInput, args));
	});
}
