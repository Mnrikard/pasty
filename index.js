#!/usr/bin/env node

var args = process.argv;
args.shift();
args.shift();

var editorRunner = require("./editorRunner.js");

debugger;
if(process.stdin.isTTY){
	try{
		var content = "";
		var clipboard = require("clipboardy");
		try{
			content = clipboard.readSync();
		} catch(e) {
			clipboard.writeSync(content);
		}
		var newContent = editorRunner.handleInput(content, args);
		clipboard.writeSync(newContent);
	} catch(err) {
		console.log(err);
		require("./stringHelpers.js").keepWindowOpen();
	}
	process.exit();
} else {
	editorRunner.interactive = false;
	require("./stringHelpers.js").keepWindowOpen = function(){};
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
