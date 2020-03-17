
exports.calledName = "";
exports.names=["dedup"];
const os = require("os");

exports.parms=[{
	name:"separator",
	value:null,
	defaultValue:"\n"
}];

exports.getParms = function(){
	return exports.parms;
};

exports.allowedSwitches = "migIG";

exports.helpText = "dedup - Deduplicates a list"+os.EOL+
	"Parameters: [separator]=new-line - how to separate the list"+os.EOL+
	"Syntax: pasty dedup [\"separator\"]"+os.EOL+os.EOL+
	"Example: echo '1,1,2,2,3,3' | pasty dedup ','"+os.EOL+
	">> 1,2,3";
exports.oneLiner = "Deduplicates a list";

const str = require("../stringHelpers.js");

function contains(arr, itm){
	let i;
	for(i=0;i<arr.length;i++){
		if(arr[i].trim() == itm.trim()){
			return true;
		}
	}
	return false;
}

exports.edit=function(input, switches){
	const sep = exports.parms[0].value;
	const regxSwitches = str.getRegexSwitches(switches);
	const rx = new RegExp(str.escapeRegex(sep), regxSwitches);
	const matches = input.split(rx);
	let output = [];
	matches.forEach(function(el){
		if(!contains(output, el)){
			output.push(el);
		}
	});
	return output.join(sep);
};

