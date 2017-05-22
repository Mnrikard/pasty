var os = require("os");
exports.interactive = true;

function getUserInput(parmName){
	var readlineSync = require('readline-sync');
	var output = readlineSync.question(parmName+": ");
	return output;
}

function getParameters(editor){
	if(editor.getParms){
		return editor.getParms();
	}
	return editor.parms;
}

function getArg(arg){
	return arg.replace(/\\t/g, "\t")
		.replace(/\\q/g, "\"")
		.replace(/\\n/g, "\n")
		.replace(/\\p/g, "|")
		.replace(/\\\n/g, "\\n");
}

function setParameters(args, parms) {
	if(!parms) { parms = []; }
	for(var i=0;i<parms.length;i++){
		if(parms[i].value && parms[i].value !== null){
			continue;
		}
		if(args && args.length > i){
			parms[i].value = getArg(args[i]);
			continue;
		}
		if(parms[i].defaultValue !== null){
			parms[i].value = getArg(parms[i].defaultValue);
			continue;
		}
		if(exports.interactive){
			parms[i].value = getArg(getUserInput(parms[i].name));
			continue;
		}
		throw "Parameter:"+parms[i].name+" is not valued";
	}
	return parms;
}

function isASwitch(arg){
	return arg.match(/^\-[grimL]{1,5}$/);
}

function getSwitches(args){
	var output = "";
	for(var i=0;i<args.length;i++){
		if(isASwitch(args[i])){
			output+=args[i].replace(/\-/g,"");
		}
	}
	return output;
}

function clearParams(editor){
	for(var i=0;i<editor.parms.length;i++){
		editor.parms[i].value = null;
	}
}

function getEditor(editorName){
	var ed = require("./editors");
	var output = ed.getEditor(editorName);
	if(output != null){
		clearParams(output);
	}
	return output;
}

exports.handleInput = function(str, args) {
	var editor = args[0];
	args.shift();
	return exports.runNamedEditor(str, editor, args).replace(/\r?\n/g,os.EOL);
};

exports.runNamedEditor = function(input, name, args){
	input = input.replace(/\r/g,"");
	var editor = getEditor(name);
	if(editor == null){
		console.log("No editor named:"+name+" found, exiting");
		return input;
	}
	editor.calledName = name;
	debugger;
	var parms = getParameters(editor);
	if(parms === null){
		parms = [];
	}
	var switches = getSwitches(args);
	for(var i=0;i<args.length;i++){
		if(isASwitch(args[i])){
			args.splice(i,1);
		}
	}
	editor.parms = setParameters(args, parms);
	var output = editor.edit(input, switches);
	return output.replace(/\r*\n/g, os.EOL);
}


