exports.calledName = "";
exports.names=["count","len"];
var os = require("os");

exports.parms=[{"name":"chars or lines","value":null,"defaultValue":"chars"}];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "count - counts characters or lines"+os.EOL+
	"Syntax: pasty <count|len> [lines]"+os.EOL+os.EOL+
	"Example: echo abcd | pasty len"+os.EOL+
	">> 4 characters";
exports.oneLiner = "count - counts characters or lines";

var str = require("../stringHelpers.js");

exports.edit=function(input, switches){
	if(exports.parms[0].value.trim().toLowerCase().indexOf("line")>-1){
		console.log(input.split(/\n/g).length + " lines");
	} else {
		console.log(input.length + " characters");
	}
	return input;
};

