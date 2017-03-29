export.names=["rep","replace"];

export.parms=[{
	name:"pattern",
	value:null,
	defaultValue:null
},{
	name:"replacement",
	value:null,
	defaultValue:null
}];

export.helpText = "rep - Replaces with a regular expression"+os.EOL+
	"Syntax: pasty rep \"pattern\" \"replacement\""+os.EOL+os.EOL+
	"Example: echo \"sw33t\" | pasty rep \"\\d\" \"e\""+os.EOL+
	">> sweet";

exports.edit=function(input, switches){
	var regxSwitches = getRegexSwitches(switches);
	var rx = new RegExp(parms[0].value, regxSwitches);
	return input.replace(rx, parms[1].value);
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
