
exports.calledName = "";
exports.names=["sort"];
const os = require("os");

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

const str = require("../stringHelpers.js");

const isDate = /^\d+(\-\/)\d+(\-\/)\d+/ig;

exports.allowedSwitches = "ri";

let ignoreCase = false;

function genericSorter(a,b){
	if(a.match(isDate) && b.match(isDate)){
		let dateA = new Date(a);
		let dateB = new Date(b);
		if(dateA !== "Invalid Date" && dateB !== "Invalid Date"){
			if(dateA > dateB) { return 1; }
			if(dateA < dateB) { return -1; }
			return 0;
		}
	}

	let floatA = parseFloat(a);
	let floatB = parseFloat(b);
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

	const sep = exports.parms[0].value;
	const rx = new RegExp(str.escapeRegex(sep), "g");
	const matches = input.trim().split(rx);
	let outlist = matches.sort(genericSorter);

	if(str.isReverse(switches)){
		outlist.reverse();
	}
	return outlist.join(sep);
};

