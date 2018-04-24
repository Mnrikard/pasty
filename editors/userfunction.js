var settings = require("../settings.js").settings;
var str = require("../stringHelpers.js");

function getNames(){
	var output = ["userfunc"];
	if(settings && settings.savedCommands){
		for(var i=0;i<settings.savedCommands.length;i++){
			output.push(settings.savedCommands[i].name);
		}
	}
	return output;
}

function getSavedCommand(name){
	if(settings && settings.savedCommands){
		for(var i=0;i<settings.savedCommands.length;i++){
			if(str.same(settings.savedCommands[i].name, name)){
				return settings.savedCommands[i];
			}
		}
	}
}

exports.getParms = function(){
	//{ name:"pattern", value:null, defaultValue:null }
	var cmd = getSavedCommand(exports.calledName);
	var output = cmd.parameters;
	if(output === null){
		return [];
	}
	exports.parms = output;
	return output;
};

exports.updateHelpText = function(){
	var savedCmd = getSavedCommand(exports.calledName);
	if(!savedCmd){
		return "no saved command with the name \""+exports.calledName+"\" exists: see ~/pasty.json";
	}
	exports.oneLiner = savedCmd.description;
	exports.helpText = savedCmd.description;
	exports.parms = savedCmd.parameters;
};

exports.calledName = "";
exports.names=getNames();
exports.parms=[];
exports.helpText = "executes user defined functions from ~/pasty.json";
exports.oneLiner = exports.helpText;

function getReplacedArgs(args){
	var output = [];
	args.forEach(function(el){
		exports.parms.forEach(function(replacement){
			var replaceRx = new RegExp(str.escapeRegex("{{"+replacement.name+"}}"),"g");
			el = el.replace(replaceRx, replacement.value);
		});
		output.push(el);
	});
	return output;
}

exports.edit=function(input, switches){
	debugger;
	var savedCmd = getSavedCommand(exports.calledName);
	if(!savedCmd){
		return "no saved command with the name \""+exports.calledName+"\" exists: see ~/pasty.json";
	}

	var runr = require("../editorRunner.js");
	for(var i=0;i<savedCmd.commands.length;i++){
		debugger;
		var replacedArgs = getReplacedArgs(savedCmd.commands[i].args);
		input = runr.runNamedEditor(input, savedCmd.commands[i].name, replacedArgs);
	}
	return input;
};

