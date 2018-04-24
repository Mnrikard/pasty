exports.names = ["help"];
const chalk = require("chalk");
const os = require("os");
exports.parms = [{
	name:"function",
	value:null,
	defaultValue:"."
}];

var str = require("../stringHelpers.js");

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "why are you looking for help on help? What did you expect to find?";
exports.oneLiner = "gets help on functions";

function prettyWriteHelp(searchedName, helpText){
	var output = "";
	if(!helpText || helpText === ""){
		helpText = "No help available for this command";
	}

	var lines = helpText.split(/\r?\n/g);
	for(var i=0;i<lines.length;i++){
		if(i==0){
			var edname = lines[i].match(/[\w]+ - /);
			if(!edname){
				edname = searchedName + " - ";
			}

			output += chalk.cyan.bold(edname)+lines[i].replace(edname,"") + os.EOL;
		} else {
			if(lines[i].match(/syntax:/i)){
				output += chalk.red.bold("Syntax:") + chalk.green(lines[i].replace(/syntax:/i,"")) + os.EOL;
			} else if(lines[i].match(/example:/i)){
				output += chalk.red.bold("Example:") + chalk.green(lines[i].replace(/example:/i,"")) + os.EOL;
			} else if (lines[i].match(/^>>/)){
				output += chalk.yellow(lines[i])+os.EOL;
			} else{
				output += lines[i]+os.EOL;
			}
		}
	}
	console.log(output);
}

exports.edit = function(input, switches){
	var ed = require("./");
	var searchEd = exports.parms[0].value;
	var editor = ed.getEditor(searchEd, false);
	if(editor === null){
		console.log(chalk.red.bold("No editor found matching: "+searchEd));
		listEditors();
		findSimilarEditors(searchEd);
		str.keepWindowOpen();
		return input;
	}
	if(editor.updateHelpText){
		editor.calledName = searchEd;
		editor.updateHelpText();
	}
	prettyWriteHelp(searchEd, editor.helpText);
	str.keepWindowOpen();
	return input;
};

function listEditors(){
	var names = require("./index.js").getEditorNames();
	for(var i=0;i<names.length;i++){
		console.log(chalk.cyan.bold(names[i].name+getAliases(names[i].aliases))+" "+names[i].description);
	}
}

function getCamelNames(editorName){
	var output = [];

	for(var c=0;c<editorName.length;c++){
		var chr = editorName[c].charCodeAt(0);
		if(c ===0 || chr >= 65 && chr <= 90){
			output.push(editorName[c]);
		} else {
			output[output.length-1] += editorName[c];
		}
	}

	return output;
}

function findSimilarEditors(editorName){
	var editors = require("./index.js").getEditorNames();
	var matches = [];
	var camelNames = getCamelNames(editorName);
	var pattern = new RegExp("("+camelNames.join("|")+")","i");

	console.log(chalk.red.bold("Attempting to find functions matching the words:"));
	console.log(chalk.red.bold(camelNames.join(", ")));
	console.log();
	var i,j;
	for(i=0;i<editors.length;i++){
		if(editors[i].name.match(pattern)){
			matches.push(editors[i].name);
		}

		for(j=0;j<editors[i].aliases.length;j++){
			if(editors[i].aliases[j].match(pattern)){
				matches.push(editors[i].aliases[j]);
			}
		}
	}
	if(matches.length === 0){
		console.log(chalk.cyan("No functions found"));
	} else {
		for(i=0;i<matches.length;i++){
			console.log(chalk.cyan(matches[i]));
		}
	}
}

function getAliases(names){
	if(names.length > 0){
		return " (alias:"+names.join()+")";
	}
	return "";
}

function listEditorsInBlock(){
	var names = require("./index.js").getEditorNames();
	var maxwidth = process.stdout.columns;

}
