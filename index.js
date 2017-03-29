#!/usr/bin/env node

var chalk = require("chalk");
var editors = [
	require("rep")
];
var interactive = true;

var args = process.argv;
args.shift();
args.shift();

if(process.stdin.isTTY){
	console.log("is tty");
	var clipboard = require("copy-paste");
	var content = clipboard.paste();
	clipboard.copy(handleInput(content, args));
} else {
	console.log("is tty...not");
	interactive = false;
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

function getParameters(args, parms) {
	for(var i=0;i<parms.length;i++){
		if(parms[i].value !== null){
			continue;
		}
		if(args && args.length > i){
			parms[i].value = args[i];
			continue;
		}
		if(parms[i].defaultValue !== null){
			parms[i].value = parms[i].defaultValue;
			continue;
		}
		if(interactive){
			parms[i].value = yield prompt(parms[i].name);
			continue;
		}
		throw "Parameter:"+parms[i].name+" is not valued";
	}
	return parms;
}

function handleInput(str, args) {
	var editor = getEditor(args);
	editor.parms = getParameters(args, editor.parms);
	str = editor.edit(str);
	return str;
}

function getEditor(args){
	if(args.length === 0 || args[0].toLower() === "help") {
		return new Help(args);
	}

	var searchTerm = new RegExp(args[1],"gi");
	for(var ed=0;ed < editors.length;ed++){
		for(var n=0;n<editors[ed].names.length;n++){
			if(editors[ed].names[n].match(searchTerm)){
				return editors[ed];
			}
		}
	}
	throw "No editor found matching: "+args[0];
}

function Help(args){
	var searchTerm = new RegExp(args[1],"gi");
	for(var ed=0;ed < editors.length;ed++){
		for(var n=0;n<editors[ed].names.length;n++){
			if(editors[ed].names[n].match(searchTerm)){
				console.log(editors[ed].helpText);
			}
		}
	}
}


