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
		.replace(/\\\n/g, "\\n");
}

function setParameters(args, parms) {
	var i;
	if(!parms) { parms = []; }
	for(i=0;i<parms.length;i++){
		if(parms[i].value && parms[i].value !== null){
			// do nothing
		} else if(args && args.length > i){
			parms[i].value = getArg(args[i]);
		} else if(parms[i].defaultValue !== null){
			parms[i].value = getArg(parms[i].defaultValue);
		} else if(exports.interactive){
			parms[i].value = getArg(getUserInput(parms[i].name));
		} else {
			throw "Parameter:"+parms[i].name+" is not valued";
		}
	}
	return parms;
}

function isASwitch(arg){
	return arg.match(/^\-[grimGIL]{1,5}$/);
}

function getSwitches(args){
	var output = "",i;
	for(i=0;i<args.length;i++){
		if(isASwitch(args[i])){
			output+=args[i].replace(/\-/g,"");
		}
	}
	return output;
}

function clearParams(editor){
	var i;
	for(i=0;i<editor.parms.length;i++){
		editor.parms[i].value = null;
	}
}

function getEditor(editorName){
	var ed = require("./editors");
	var output = ed.getEditor(editorName);
	if(output !== null){
		clearParams(output);
	}
	return output;
}

Array.prototype.pluck = function(index){
	return this.splice(index,1);
};

function removeSwitches(args){
	var i;
	for(i=0;i<args.length;i++){
		if(isASwitch(args[i])){
			debugger;
			args.pluck(i);
		}
	}
}


exports.handleInput = function(str, args) {
	var editor = args[0];
	args.shift();
	return exports.runNamedEditor(str, editor, args);
};

exports.runNamedEditor = function(input, name, args){

	input = input.replace(/\r/g,"");

	var editor = getEditor(name);
	if(editor === null){
		args.unshift(name);
		editor = getEditor("help");
	}

	editor.calledName = name;

	var parms = getParameters(editor) || [];

	var switches = getSwitches(args);
	debugger;
	removeSwitches(args);

	editor.parms = setParameters(args, parms);
	var output = editor.edit(input, switches);
	output = output.replace(/\r/g,"");
	debugger;
	return output.replace(/\n/g, os.EOL);
};


