exports.calledName = "";
exports.names=["xmlencode","xmldecode"];
var os = require("os");

exports.parms=[
];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "xmlencode - encode/decode an xml string"+os.EOL+
	"Syntax: pasty <xmlencode|xmldecode> [-r]"+os.EOL+os.EOL+
	"Example: echo \"this&that\" | pasty xmlencode"+os.EOL+
	">> this&amp;that";
exports.oneLiner = "xmlencode - encode/decode a xml";

var str = require("../stringHelpers.js");

function xmlEncode(input){
	return input.replace(/[<>&'"]/g, function (c) {
		switch (c) {
			case '<': return '&lt;';
			case '>': return '&gt;';
			case '&': return '&amp;';
			case '\'': return '&apos;';
			case '"': return '&quot;';
		}
	});
}

function xmlDecode(input){
	return input.replace(/&(amp|lt|gt|apos|quot|#\\d+);/g, function (c,g1) {
		switch (g1) {
			case 'lt': return '<';
			case 'gt': return '>';
			case 'amp': return '&';
			case 'apos': return "'";
			case 'quot': return '"';
			default: return String.fromCharCode(g1);
		}
	});
}
exports.edit=function(input, switches){
	if(exports.calledName.match(/encode/i)){
		if(str.isReverse(switches)){
			return xmlDecode(input);
		}
		return xmlEncode(input);
	}
	if(str.isReverse(switches)){
		return xmlEncode(input);
	}
	return xmlDecode(input);
};



