exports.calledName = "";
exports.names=["rep","replace"];
var os = require("os");

exports.parms=[
	{name:"pattern", value:null, defaultValue:null },
	{name:"replacement", value:null, defaultValue:null}
];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "rep - Replaces with a regular expression"+os.EOL+
	"Syntax: pasty rep \"pattern\" \"replacement\""+os.EOL+os.EOL+
	"Example: echo \"sw33t\" | pasty rep \"\\d\" \"e\""+os.EOL+
	">> sweet";
exports.oneLiner = "replaces with a RegExp";

var str = require("../stringHelpers.js");

function enhancedReplacementPattern(){
	
}

exports.edit=function(input, switches){
	var pattern = exports.parms[0].value;
	var repl = exports.parms[1].value;
	var regxSwitches = str.getRegexSwitches(switches);
	var rx = new RegExp(pattern, regxSwitches);
	return input.replace(rx, repl);
};

