exports.calledName = "";
exports.names=["setText"];
const os = require("os");

exports.parms=[
{ name:"text", value:null, defaultValue:null }
];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "setText - sets the content to the passed in string"+os.EOL+
	"Parameters: \"text\": the text to set"+os.EOL+
	"Syntax: pasty setText <anyValue>"+os.EOL+os.EOL+
	"Example: echo abcd | pasty setText \"something else\""+os.EOL+
	">> something else";
exports.oneLiner = "sets the content to the passed in string";

exports.edit=function(input, switches){
	return exports.parms[0].value;
};

