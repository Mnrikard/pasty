exports.names=["rep","replace"];
var os = require("os");

exports.parms=[{
	name:"pattern",
	value:null,
	defaultValue:null
},{
	name:"replacement",
	value:null,
	defaultValue:null
}];

exports.helpText = "rep - Replaces with a regular expression"+os.EOL+
	"Syntax: pasty rep \"pattern\" \"replacement\""+os.EOL+os.EOL+
	"Example: echo \"sw33t\" | pasty rep \"\\d\" \"e\""+os.EOL+
	">> sweet";

exports.edit=function(input, switches){
	var pattern = exports.parms[0].value;
	var repl = exports.parms[1].value;
	var regxSwitches = getRegexSwitches(switches);
	var rx = new RegExp(pattern, regxSwitches);
	return input.replace(rx, repl);
};

function getRegexSwitches(switches){
	var output = "";
	if(switches.indexOf("m")>-1){
		output += "m";
	}
	if(switches.indexOf("I") === -1){
		output += "i";
	}
	if(switches.indexOf("G") === -1){
		output += "g";
	}
	return output;
}
