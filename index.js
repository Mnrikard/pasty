#!/usr/bin/env node

var args = process.argv;
args.shift();
args.shift();

var editorRunner = require("./editorRunner.js");

debugger;
if(process.stdin.isTTY){
	var content = "";
	var clipboard = require("clipboardy");
	try{
		content = clipboard.readSync();
	} catch(e) {
		clipboard.writeSync(content);
		console.log(e);
		require("./stringHelpers.js").keepWindowOpen();
	}
	var newContent = editorRunner.handleInput(content, args);
	clipboard.writeSync(newContent);
	process.exit();
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
