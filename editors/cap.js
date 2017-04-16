exports.calledName = "";
exports.names=["cap","lower","upper"];
var os = require("os");

exports.parms=[];

exports.helpText = "cap - capitalizes or lower cases"+os.EOL+
	"Syntax: pasty <cap|lower|upper>"+os.EOL+os.EOL+
	"Example: echo abcd | pasty cap"+os.EOL+
	">> ABCD";
exports.oneLiner = "cap - capitalizes or lower cases";

var str = require("../stringHelpers.js");

var _columnLengths = [];

exports.edit=function(input, switches){
	if(exports.calledName.toLowerCase() == "cap" || exports.calledName.toLowerCase() == "upper"){
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

