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

exports.helpText = "why are you looking for help on help?  What did you expect to find?";
exports.oneLiner = "gets help on functions";

function prettyWriteHelp(helpText){
	var output = "";
	var lines = helpText.split(/\r?\n/g);
	for(var i=0;i<lines.length;i++){
		if(i==0){
			var edname = lines[i].match(/[\w]+ - /);
			debugger;
			output += chalk.blue.bold(edname)+lines[i].replace(edname,"") + os.EOL;
		} else {
			if(lines[i].match(/syntax:/i)){
				output += chalk.red("Syntax:") + chalk.green(lines[i].replace(/syntax:/i,"")) + os.EOL;
			} else if(lines[i].match(/example:/i)){
				output += chalk.red("Example:") + chalk.green(lines[i].replace(/example:/i,"")) + os.EOL;
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
		console.log(chalk.red("No editor found matching: "+searchEd));
		listEditors();
		str.keepWindowOpen();
		return input;
	}
	prettyWriteHelp(editor.helpText);
	str.keepWindowOpen();
	return input;
};

function listEditors(){
	var names = require("./index.js").getEditorNames();
	for(var i=0;i<names.length;i++){
		console.log(chalk.blue.bold(names[i].name+getAliases(names[i].aliases))+" "+names[i].description);
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
