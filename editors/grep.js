
exports.calledName = "";
exports.names=["grep","grab"];
var os = require("os");

exports.parms=[{
	name:"pattern",
	value:null,
	defaultValue:null
},{
	name:"separator",
	value:null,
	defaultValue:"\n"
}];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "grep - Gets a regex and prints"+os.EOL+
	"Parameters: "+os.EOL+
	"    Pattern: a regex pattern to GREP"+os.EOL+
	"    [separator]=new-line - how to separate the matches when printing"+os.EOL+
	"Syntax: pasty grep \"pattern\""+os.EOL+os.EOL+
	"Example: echo \"sw33t\" | pasty grep \"\\d\""+os.EOL+
	">> 3"+os.EOL+
	">> 3";
exports.oneLiner = "you know, GREP...";

var str = require("../stringHelpers.js");

function enhancedReplacementPattern(){
	
}

exports.edit=function(input, switches){
	var pattern = exports.parms[0].value;
	var sep = exports.parms[1].value;
	var regxSwitches = str.getRegexSwitches(switches);
	var rx = new RegExp(pattern, regxSwitches);
	var matches = input.match(rx);
	return matches.join(sep);
};

