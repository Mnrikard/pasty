exports.calledName = "";
exports.names=["math"];
const os = require("os");

exports.parms=[
	{name:"\"answer\" for answer only", value:null, defaultValue:"" }
];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "math - Solves simple math problems"+os.EOL+
	"Syntax: pasty math"+os.EOL+os.EOL+
	"Example: echo \"123+456\" | pasty math"+os.EOL+
	">> 579";
exports.oneLiner = "Solves simple math problems";

const str = require("../stringHelpers.js");

const decimalPattern = "(?!=\\d\\s*)\\-?\\d+(\\.\\d+)?";
const sumUpList = new RegExp("("+decimalPattern+")\\s+("+decimalPattern+")");

exports.edit=function(input, switches){
	input = input.trim();
	const calc = require("../stringMath.js");
	let expression = input.replace(sumUpList, "${1}+${2}");
	expression = expression.replace(sumUpList, "${1}+${2}");

	let answer = calc.evaluateExpression(expression);

	if(str.same(exports.parms[0].value, "answer"))
	{
		return answer;
	}

	const joiner = "=";

	if(input.match(/\n/)) { joiner = "\n"; }
	else if(input.match(/\=/)) { joiner = ""; }

	return [input,answer].join(joiner);
};

