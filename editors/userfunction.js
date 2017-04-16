var settings = require("../settings.js").settings;

function getNames(){
	var output = [];
	if(settings && settings.savedCommands){
		for(var i=0;i<settings.savedCommands.length;i++){
			output.push(settings.savedCommands[i].name);
		}
	}
	return output;
}

exports.calledName = "";
exports.names=getNames();
exports.parms=[];
exports.helpText = "executes user defined functions from ~/pasty.json";
exports.oneLiner = exports.helpText;

function same(text1, text2){
	return text1.toLowerCase().trim() == text2.toLowerCase().trim();
}

function getSavedCommand(name){
	if(settings && settings.savedCommands){
		for(var i=0;i<settings.savedCommands.length;i++){
			if(same(settings.savedCommands[i].name, name)){
				return settings.savedCommands[i];
			}
		}
	}
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
		input = runr.runNamedEditor(input, savedCmd.commands[i].name, savedCmd.commands[i].args);
	}
	return input;
};

