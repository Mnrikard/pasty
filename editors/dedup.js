
exports.calledName = "";
exports.names=["dedup"];
var os = require("os");

exports.parms=[{
	name:"separator",
	value:null,
	defaultValue:"\n"
}];

exports.helpText = "dedup - Deduplicates a list"+os.EOL+
	"Syntax: pasty dedup [\"separator\"]"+os.EOL+os.EOL+
	"Example: echo '1,1,2,2,3,3' | pasty dedup ','"+os.EOL+
	">> 1,2,3";
exports.oneLiner = "dedup - Deduplicates a list";


function getRegexSwitches(switches){
	var output = "";
	if(switches.indexOf("m")>-1){
		output += "m";
	}
	if(switches.indexOf("I") === -1){
		output += "i";
	}
	if(switches.indexOf("G") === -1){
		output += "g";
	}
	return output;
}

function escapeRegex(actualText){
	return actualText.replace(/\(/g,"\\(")
		.replace(/\)/g,"\\)")
		.replace(/\+/g,"\\+")
		.replace(/\*/g,"\\*")
		.replace(/\-/g,"\\-")
		.replace(/\./g,"\\.");
}

function contains(arr, itm){
	for(var i=0;i<arr.length;i++){
		if(arr[i].trim() == itm.trim()){
			return true;
		}
	}
	return false;
}

exports.edit=function(input, switches){
	var sep = exports.parms[0].value;
	var regxSwitches = getRegexSwitches(switches);
	var rx = new RegExp(escapeRegex(sep), regxSwitches);
	var matches = input.split(rx);
	var output = [];
	matches.forEach(function(el){
		if(!contains(output, el)){
			output.push(el);
		}
	});
	return output.join(sep);
};

