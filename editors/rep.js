exports.calledName = "";
exports.names=["rep","replace"];
const os = require("os");

exports.parms=[
	{name:"pattern", value:null, defaultValue:null },
	{name:"replacement", value:null, defaultValue:null}
];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "rep - Replaces with a regular expression"+os.EOL+
	"Parameters: "+os.EOL+
	"    pattern is a regex pattern to replace"+os.EOL+
	"    replacement is a replacement string"+os.EOL+
	"  also accepts -migIG flags to modify pattern.  Defaults to -ig"+os.EOL+
	"Syntax: pasty rep \"pattern\" \"replacement\""+os.EOL+os.EOL+
	"Example: echo \"sw33t\" | pasty rep \"\\d\" \"e\""+os.EOL+
	">> sweet";
exports.oneLiner = "replaces with a RegExp";

exports.allowedSwitches = "migIG";

const str = require("../stringHelpers.js");

function enhancedReplacementPattern(){
	let thisRp = exports.parms[1].value;

	const args = Array.prototype.slice.call(arguments);
	const groups = args.splice(0,arguments.length-2);

	thisRp = thisRp.replace(/\\u\$(\d)/i, function(a,b){ return groups[parseInt(b,10)].toUpperCase(); });
	thisRp = thisRp.replace(/\\u\$\{(\d+)\}/i, function(a,b){ return groups[parseInt(b,10)].toUpperCase(); });

	thisRp = thisRp.replace(/\\l\$(\d)/, function(a,b){ return groups[parseInt(b,10)].toLowerCase(); });
	thisRp = thisRp.replace(/\\l\$\{(\d+)\}/i, function(a,b){ return groups[parseInt(b,10)].toLowerCase(); });

	thisRp = thisRp.replace(/\$(\d)/, function(a,b){ return groups[parseInt(b,10)]; });
	thisRp = thisRp.replace(/\$\{(\d+)\}/i, function(a,b){ return groups[parseInt(b,10)]; });

	return thisRp;
}

exports.edit=function(input, switches){
	const pattern = exports.parms[0].value;
	const repl = exports.parms[1].value;
	const regxSwitches = str.getRegexSwitches(switches);
	const rx = new RegExp(pattern, regxSwitches);
	if(repl.match(/(\\[ul]\$(\{\d+\}|\d)|\$0)/i)){
		return input.replace(rx, enhancedReplacementPattern);
	}
	return input.replace(rx, repl);
};

