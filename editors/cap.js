exports.calledName = "";
exports.names=["cap","lower","upper"];
const os = require("os");

exports.parms=[];
exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "cap - capitalizes or lower cases"+os.EOL+
	"Parameters: one of cap/upper/lower to define how to \"cap\""+os.EOL+
	"Syntax: pasty <cap|lower|upper>"+os.EOL+os.EOL+
	"Example: echo abcd | pasty cap"+os.EOL+
	">> ABCD";
exports.oneLiner = "capitalizes or lower cases";
exports.allowedSwitches = "r";

const str = require("../stringHelpers.js");

let _columnLengths = [];

exports.edit=function(input, switches){
	if(exports.calledName.toLowerCase() === "cap" || exports.calledName.toLowerCase() === "upper"){
		if(str.isReverse(switches)){
			return input.toLowerCase();
		}
		return input.toUpperCase();
	}
	if(str.isReverse(switches)){
		return input.toUpperCase();
	}
	return input.toLowerCase();
};

