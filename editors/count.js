exports.calledName = "";
exports.names=["count","len"];
const os = require("os");

exports.parms=[{name:"chars or lines",value:null,defaultValue:"chars"}];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "count - counts characters or lines"+os.EOL+
	"Parameters: [char/line]=\"char\" - whether to count characters or lines"+os.EOL+
	"Syntax: pasty <count|len> [lines]"+os.EOL+os.EOL+
	"Example: echo abcd | pasty len"+os.EOL+
	">> 4 characters";
exports.oneLiner = "counts characters or lines";

const str = require("../stringHelpers.js");

exports.edit=function(input, switches){
	if(exports.parms[0].value.trim().toLowerCase().indexOf("line")>-1){
		console.log(input.split(/\n/g).length + " lines");
	} else if(exports.parms[0].value.trim().toLowerCase().indexOf("byte")>-1) {
		console.log(input.length + " bytes");
	} else {
		console.log([...input].length + " characters");
	}
	str.keepWindowOpen();
	return input;
};

