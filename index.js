#!/usr/bin/env node

var args = process.argv;
args.shift();
args.shift();

var editorRunner = require("./editorRunner.js");

exports.editPipedInput = function(args){
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
		process.stdout.write(editorRunner.handleInput(pipedInput, args));
	});
};

exports.editClipboard = function(args){
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
}

debugger;
if(process.stdin.isTTY){
	exports.editClipboard(args);
	process.exit();
} else {
	exports.editPipedInput(args);
}

