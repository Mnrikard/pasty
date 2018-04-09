
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

exports.allowedSwitches = "migIGL";

exports.helpText = "grep - Gets a regex and prints"+os.EOL+
	"Parameters: "+os.EOL+
	"    Pattern: a regex pattern to GREP"+os.EOL+
	"    [separator]=new-line - how to separate the matches when printing"+os.EOL+
	"This command accepts -migIG switches as well as -r to negate the pattern"+os.EOL+
	"  and it accepts the -L switch to grep not on a line"+os.EOL+
	"Syntax: pasty grep \"pattern\""+os.EOL+os.EOL+
	"Example: echo \"sw33t\" | pasty grep \"\\d\""+os.EOL+
	">> 3"+os.EOL+
	">> 3";
exports.oneLiner = "you know, GREP...";

var str = require("../stringHelpers.js");


exports.edit=function(input, switches){
	var pattern = exports.parms[0].value;
	var sep = exports.parms[1].value;
	var regxSwitches = str.getRegexSwitches(switches);
	var rx = new RegExp(pattern, regxSwitches);
	var reverse = str.isReverse(switches);

	//non-linear matching
	if(switches.indexOf("L") > -1){
		if(reverse){
			return input.replace(rx,"");
		}
		var matches = input.match(rx);
		return matches.join(sep);
	}

	//standard GREP
	var sepRx = new RegExp(str.escapeRegex(sep),"gi");
	var lines = input.split(sepRx);
	var output = [];

	var ln;
	for(ln=0;ln<lines.length;ln++){
		rx.lastIndex = 0;
		if(rx.test(lines[ln]) !== reverse){
			output.push(lines[ln]);
		}
	}
	return output.join(sep);
};

