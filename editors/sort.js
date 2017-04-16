
exports.calledName = "";
exports.names=["sort"];
var os = require("os");

exports.parms=[{
	name:"separator",
	value:null,
	defaultValue:"\n"
}];

exports.helpText = "sort - Sorts a list"+os.EOL+
	"Syntax: pasty sort [\"separator\"]"+os.EOL+os.EOL+
	"Example: echo \"1,3,2,4,6,5\" | pasty sort \",\""+os.EOL+
	">> 1,2,3,4,5,6";

exports.oneLiner = "sort - Sorts a list";

var str = require("../stringHelpers.js");


function dateSorter(a,b){
	if(a > b) { return 1; }
	if(a < b) { return -1; }
	return 0;
}

function genericSorter(a,b){
	var dateA = new Date(a);
	var dateB = new Date(b);

	if(dateA != "Invalid Date" && dateB != "Invalid Date"){
		return dateSorter(dateA, dateB);
	}
	
	if(a > b) { return 1; }
	if(a < b) { return -1; }
	return 0;
}

function genericReverseSorter(a, b){
	return genericSorter(b, a);
}

exports.edit=function(input, switches){
	var sep = exports.parms[0].value;
	var rx = new RegExp(str.escapeRegex(sep), "g");
	var matches = input.split(rx);
	if(str.isReverse(switches)){
		return matches.sort(genericReverseSorter).join(sep);
	}
	return matches.sort(genericSorter).join(sep);
};

