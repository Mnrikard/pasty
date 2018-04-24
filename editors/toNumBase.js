
exports.calledName = "";
exports.names=["tobase"];
var os = require("os");

exports.parms=[
	{ name:"base", value:null, defaultValue:"16" }
];

exports.getParms = function(){
	return exports.parms;
};

exports.allowedSwitches = "r";

exports.helpText = "toBase - converts a number to/from decimal from/to a base."+os.EOL+
	"Parameters: [base], defaults to 16"+os.EOL+
	"  accepts the -r flag to reverse the conversion"+os.EOL+
	"Syntax: pasty tobase [integer base]"+os.EOL+os.EOL+
	"Example: echo \"16\" | pasty tobase 16"+os.EOL+
	">> F";

exports.oneLiner = "converts a number to/from decimal from/to a base";

var str = require("../stringHelpers.js");

exports.edit=function(input, switches){
	var base = parseInt(exports.parms[0].value);

	if(exports.parms[0].value == "0x"){ base=16; }
	if(base == 0){ base=8; }
	debugger;

	if(str.isReverse(switches)){
		return parseInt(input, base).toString();
	}
	return parseInt(input).toString(base).toUpperCase();
};

