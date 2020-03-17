exports.calledName = "";
exports.names=["encode","urlencode","urldecode","xmlencode","xmldecode","base64encode","base64decode","base64"];
const os = require("os");

exports.parms=[
];

exports.getParms = function(){
	return exports.parms;
};

exports.allowedSwitches = "r";

exports.helpText = "encode - encode/decode a url/xml/base64"+os.EOL+
	"Parameters: different names define different encodings and whether to encode or decode"+os.EOL+
	"Syntax: pasty <url|xml|base64<encode|decode>> [-r]"+os.EOL+os.EOL+
	"Example: echo \"this&that\" | pasty urlencode | pasty xmlencode | pasty base64"+os.EOL+
	">> this%26that";
exports.oneLiner = "encode/decode a url/xml/base64";

const str = require("../stringHelpers.js");

const encoder = {
	"encode":function(input){return input},
	"decode":function(input){return input},
};

const xmlEncode = function(input){
	return input.replace(/[<>&'"]/g, function (c) {
		switch (c) {
			case '<': return '&lt;';
			case '>': return '&gt;';
			case '&': return '&amp;';
			case '\'': return '&apos;';
			case '"': return '&quot;';
		}
	});
};

const xmlDecode = function(input){
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
};

const b64EncodeUnicode = function(input) {
	const btoa = require("btoa");
	return btoa(encodeURIComponent(input).replace(/%([0-9A-F]{2})/g, function(match, p1) {
		return String.fromCharCode('0x' + p1);
	}));
};

const b64DecodeUnicode = function(input) {
	const atob = require("atob");
	return decodeURIComponent(atob(input).split('').map(function(c) {
		return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
	}).join(''));
};

exports.edit=function(input, switches){
	if(exports.calledName.match(/base64/i)){
		encoder.encode = b64EncodeUnicode;
		encoder.decode = b64DecodeUnicode;
	} else if (exports.calledName.match(/url/i)){
		encoder.encode = encodeURIComponent;
		encoder.decode = decodeURIComponent;
	} else if (exports.calledName.match(/xml/i)){
		encoder.encode = xmlEncode;
		encoder.decode = xmlDecode;
	}

	if(exports.calledName.match(/decode/i)){
		if(str.isReverse(switches)){
			return encoder.encode(input);
		}
		return encoder.decode(input);
	}
	if(str.isReverse(switches)){
		return encoder.decode(input);
	}
	return encoder.encode(input);
};


