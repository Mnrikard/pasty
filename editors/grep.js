
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

exports.helpText = "grep - Gets a regex and prints"+os.EOL+
	"Syntax: pasty grep \"pattern\""+os.EOL+os.EOL+
	"Example: echo \"sw33t\" | pasty grep \"\\d\""+os.EOL+
	">> 3"+os.EOL+
	"3";
exports.oneLiner = "grep - you know, GREP...";


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

function enhancedReplacementPattern(){
	
}

exports.edit=function(input, switches){
	var pattern = exports.parms[0].value;
	var sep = exports.parms[1].value;
	var regxSwitches = getRegexSwitches(switches);
	var rx = new RegExp(pattern, regxSwitches);
	var matches = input.match(rx);
	return matches.join(sep);
};

