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
		.replace(/\\q/g, "\"");
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


function getSwitches(args){
	var output = "";
	for(var i=0;i<args.length;i++){
		if(args[i].match(/^\-/)){
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
	clearParams(output);
	return output;
}

exports.handleInput = function(str, args) {
	var editor = args[0];
	args.shift();
	return exports.runNamedEditor(str, editor, args).replace(/\r?\n/g,os.EOL);
};

exports.runNamedEditor = function(input, name, args){
	var editor = getEditor(name);
	editor.calledName = name;
	debugger;
	var parms = getParameters(editor);
	if(parms === null){
		parms = [];
	}
	debugger;
	editor.parms = setParameters(args, parms);
	var switches = getSwitches(args);
	var output = editor.edit(input, switches);
	return output;
}


