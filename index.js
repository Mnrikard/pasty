#!/usr/bin/env node --harmony

var chalk = require("chalk");
var co = require('co');
var prompt = require('co-prompt');

var interactive = true;

var args = process.argv;
args.shift();
args.shift();

if(process.stdin.isTTY){
	var clipboard = require("copy-paste");
	var content = clipboard.paste();
	clipboard.copy(handleInput(content, args));
} else {
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
			co(function *() {
				parms[i].value = yield prompt(parms[i].name);
			});
			continue;
		}
		throw "Parameter:"+parms[i].name+" is not valued";
	}
	return parms;
}

function getSwitches(args){
	var output = "";
	for(var i=0;i<args.length;i++){
		if(args[i].match(/^\-/)){
			output+=args[i].replace(/\-/g,"");
		}
	}
	return output;
}

function handleInput(str, args) {
	var editor = getEditor(args);
	args.shift();
	editor.parms = getParameters(args, editor.parms);
	var switches = getSwitches(args);
	str = editor.edit(str, switches);
	return str;
}

function getEditor(args){
	var editorName = args[0];
	var ed = require("./editors");
	return ed.getEditor(editorName);
}



