exports.calledName = "";
exports.names=["newid"];
const os = require("os");

exports.parms=[
{ name:"version (v1 or v4)", value:null, defaultValue:"v4" }
];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "newid - generates a new UUID"+os.EOL+
	"Syntax: pasty newid [v1|v4]"+os.EOL+os.EOL+
	"Example: echo abcd | pasty newid v4"+os.EOL+
	">> f00a5cee-3cef-4707-af2f-1fdd08d0cc4b";
exports.oneLiner = "generates a new UUID";

exports.edit=function(input, switches){
	if(exports.parms[0].value.match(/1/)){
		const uuidV1 = require('uuid/v1');
		return uuidV1();
	}
	const uuidV4 = require('uuid/v4');
	return uuidV4();
};

