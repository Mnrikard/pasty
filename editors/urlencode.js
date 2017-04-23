exports.calledName = "";
exports.names=["urlencode","urldecode"];
var os = require("os");

exports.parms=[
];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "urlencode - encode/decode a url"+os.EOL+
	"Syntax: pasty <urlencode|urldecode> [-r]"+os.EOL+os.EOL+
	"Example: echo \"this&that\" | pasty urlencode"+os.EOL+
	">> this%26that";
exports.oneLiner = "urlencode - encode/decode a url";

var str = require("../stringHelpers.js");

exports.edit=function(input, switches){
	if(exports.calledName.match(/encode/i)){
		if(str.isReverse(switches)){
			return decodeURIComponent(input);
		}
		return encodeURIComponent(input);
	}
	if(str.isReverse(switches)){
		return encodeURIComponent(input);
	}
	return decodeURIComponent(input);
};


