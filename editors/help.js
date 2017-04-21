exports.names = ["help"];
const chalk = require("chalk");
exports.parms = [{
	name:"function",
	value:null,
	defaultValue:"."
}];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "help [functionName]";
exports.oneLiner = "gets help on functions";

exports.edit = function(input, switches){
	var ed = require("./");
	var searchEd = exports.parms[0].value;
	var editor = ed.getEditor(searchEd, false);
	if(editor === null){
		console.log(chalk.red("No editor found matching: "+searchEd));
		listEditors();
		return input;
	}
	console.log(editor.helpText);
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
