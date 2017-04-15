var co = require('co');
var prompt = require('co-prompt');
exports.interactive = true;

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
		if(exports.interactive){
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

function getEditor(editorName){
	var ed = require("./editors");
	return ed.getEditor(editorName);
}

exports.handleInput = function(str, args) {
	var editor = args[0];
	args.shift();
	return exports.runNamedEditor(str, editor, args);
};

exports.runNamedEditor = function(input, name, args){
	var editor = getEditor(name);
	editor.calledName = name;
	editor.parms = getParameters(args, editor.parms);
	var switches = getSwitches(args);
	input = editor.edit(input, switches);
	return input;
}


