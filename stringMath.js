//var str = require("./stringHelpers.js");

var decimalPattern = "(?!=\\d\\s*)\\-?\\d+(\\.\\d+)?";
var insideParens = new RegExp("\\(([^\\(\\)]+)\\)");


function getEquationArguments(regx, equation){
	var expr = regx.exec(equation);
	if(expr === null) { return null; }
	return {left: parseFloat(expr[1]), right: parseFloat(expr[4]), operator: expr[3], expression:expr[0]};
}

function evaluatePattern(math, mathFunction)
{
	var mathArgs = getEquationArguments(mathFunction.equationPattern, math);
	var answer;
	while(mathArgs)
	{
		answer = mathFunction.evaluate(mathArgs);
		//fix floating point errors (sort of)
		answer = parseFloat(parseFloat(answer).toFixed(10));
		math = math.replace(mathArgs.expression, answer);
		mathArgs = getEquationArguments(mathFunction.equationPattern, math);
	}

	return math;
}

var AddSubtract = {
	equationPattern:new RegExp("("+decimalPattern+")\\s*([\\+\\-])\\s*("+decimalPattern+")","gi")
};
AddSubtract.evaluate=function(expr){
	if(expr.operator.match(/\+/)){
		return (expr.left+expr.right).toString();
	}

	return (expr.left-expr.right).toString();
};

var MultiplyDivide = {
	equationPattern:new RegExp("("+decimalPattern+")\\s*([\\*\\/])\\s*("+decimalPattern+")")
};
MultiplyDivide.evaluate=function(expr) {
	if(expr.operator.match(/\*/)){
		return (expr.left * expr.right).toString();
	}

	return (expr.left / expr.right).toString();
};

var Exponent = {
	equationPattern:new RegExp("("+decimalPattern+")\\s*(\\^)\\s*("+decimalPattern+")")
};
Exponent.evaluate=function(expr) {
	return Math.pow(expr.left,expr.right).toString();
};

var Modulo = {
	equationPattern : new RegExp("("+decimalPattern+")\\s*(\\%)\\s*("+decimalPattern+")")
};

Modulo.evaluate = function(expr){
	return (expr.left % expr.right).toString();
};

var DivideWithRemainder = {
	equationPattern : new RegExp("("+decimalPattern+")\\s*(\\/\\%)\\s*("+decimalPattern+")")
};

DivideWithRemainder.evaluate = function(expr) {
	return Math.floor(expr.left/expr.right).toString()+","+(expr.left % expr.right).toString();
};

function allAreNumbers(numArr){
	var i;
	for(i=0;i<numArr.length;i++){
		if(isNaN(numArr[i])){
			return false;
		}
	}
	return true;
}


exports.evaluateExpression = function(math){
	var originalMath = math.trim().replace(/\=$/,"");
	math = originalMath;

	if(allAreNumbers(originalMath.split(/,/g))){
		return math;
	}

	var paren = math.match(insideParens);
	var insideMath;
	while(paren)
	{
		insideMath = paren[1];
		math = math.replace(paren[0], exports.evaluateExpression(insideMath));
		paren = math.match(insideParens);
	}

	math = evaluatePattern(math, Exponent);
	math = evaluatePattern(math, MultiplyDivide);
	math = evaluatePattern(math, Modulo);
	math = evaluatePattern(math, DivideWithRemainder);
	math = evaluatePattern(math, AddSubtract);

	if(math === originalMath)
	{
		throw "Could not parse "+math;
	}

	return exports.evaluateExpression(math);

};
