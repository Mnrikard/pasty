
exports.calledName = "";
exports.names=["sort"];
var os = require("os");

exports.parms=[
	{ name:"separator", value:null, defaultValue:"\n" }
];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "sort - Sorts a list"+os.EOL+
	"Parameters: [separator], defaults to new-line."+os.EOL+
	"  accepts the -r flag to reverse the sort"+os.EOL+
	"Syntax: pasty sort [\"separator\"]"+os.EOL+os.EOL+
	"Example: echo \"1,3,2,4,6,5\" | pasty sort \",\""+os.EOL+
	">> 1,2,3,4,5,6";

exports.oneLiner = "Sorts a list";

var str = require("../stringHelpers.js");

var isDate = /^\d+(\-\/)\d+(\-\/)\d+/ig;

exports.allowedSwitches = "ri";

var ignoreCase = false;

function genericSorter(a,b){
	if(a.match(isDate) && b.match(isDate)){
		var dateA = new Date(a);
		var dateB = new Date(b);
		if(dateA !== "Invalid Date" && dateB !== "Invalid Date"){
			if(dateA > dateB) { return 1; }
			if(dateA < dateB) { return -1; }
			return 0;
		}
	}

	var floatA = parseFloat(a);
	var floatB = parseFloat(b);
	if(!isNaN(floatA) && !isNaN(floatB)){
		if(floatA > floatB){ return 1; }
		if(floatA < floatB){ return -1; }
		return 0;
	}

	if(ignoreCase){
		a = a.toLowerCase();
		b = b.toLowerCase();
	}

	if(a > b) { return 1; }
	if(a < b) { return -1; }
	return 0;
}

exports.edit=function(input, switches){
	ignoreCase = (switches.indexOf("i") > -1);

	var sep = exports.parms[0].value;
	var rx = new RegExp(str.escapeRegex(sep), "g");
	var matches = input.trim().split(rx);
	var outlist = matches.sort(genericSorter);

	if(str.isReverse(switches)){
		outlist.reverse();
	}
	return outlist.join(sep);
};

