
exports.calledName = "";
exports.names=["sort"];
var os = require("os");

exports.parms=[{
	name:"separator",
	value:null,
	defaultValue:"\n"
}];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "sort - Sorts a list"+os.EOL+
	"Syntax: pasty sort [\"separator\"]"+os.EOL+os.EOL+
	"Example: echo \"1,3,2,4,6,5\" | pasty sort \",\""+os.EOL+
	">> 1,2,3,4,5,6";

exports.oneLiner = "Sorts a list";

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

exports.edit=function(input, switches){
	var sep = str.escapeRegex(exports.parms[0].value);
	var rx = new RegExp(sep, "g");
	var matches = input.split(rx);
	var outlist = matches.sort(genericSorter);

	if(str.isReverse(switches)){
		outlist.reverse();
	}
	return outlist.join(sep);
};

