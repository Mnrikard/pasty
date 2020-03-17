#!/usr/bin/env node

let args = process.argv;
args.shift();
args.shift();

const editorRunner = require("./editorRunner.js");

exports.editPipedInput = function(args){
	editorRunner.interactive = false;
	require("./stringHelpers.js").keepWindowOpen = function(){};
	let pipedInput = '';
	process.stdin.on('readable', function() {
		let chunk = this.read();
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
		let content = "";
		const clipboard = require("clipboardy");
		try{
			content = clipboard.readSync();
		} catch(e) {
			clipboard.writeSync(content);
		}
		let newContent = editorRunner.handleInput(content, args);
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

