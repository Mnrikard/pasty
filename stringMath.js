//const str = require("./stringHelpers.js");

const decimalPattern = "(?!=\\d\\s*)\\-?\\d+(\\.\\d+)?";
const insideParens = new RegExp("\\(([^\\(\\)]+)\\)");


function getEquationArguments(regx, equation){
	const expr = regx.exec(equation);
	if(expr === null) { return null; }
	return {left: parseFloat(expr[1]), right: parseFloat(expr[4]), operator: expr[3], expression:expr[0]};
}

function evaluatePattern(math, mathFunction)
{
	let mathArgs = getEquationArguments(mathFunction.equationPattern, math);
	let answer;
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

const AddSubtract = {
	equationPattern:new RegExp("("+decimalPattern+")\\s*([\\+\\-])\\s*("+decimalPattern+")","gi")
};
AddSubtract.evaluate=function(expr){
	if(expr.operator.match(/\+/)){
		return (expr.left+expr.right).toString();
	}

	return (expr.left-expr.right).toString();
};

const MultiplyDivide = {
	equationPattern:new RegExp("("+decimalPattern+")\\s*([\\*\\/])\\s*("+decimalPattern+")")
};
MultiplyDivide.evaluate=function(expr) {
	if(expr.operator.match(/\*/)){
		return (expr.left * expr.right).toString();
	}

	return (expr.left / expr.right).toString();
};

const Exponent = {
	equationPattern:new RegExp("("+decimalPattern+")\\s*(\\^)\\s*("+decimalPattern+")")
};
Exponent.evaluate=function(expr) {
	return Math.pow(expr.left,expr.right).toString();
};

const Modulo = {
	equationPattern : new RegExp("("+decimalPattern+")\\s*(\\%)\\s*("+decimalPattern+")")
};

Modulo.evaluate = function(expr){
	return (expr.left % expr.right).toString();
};

const DivideWithRemainder = {
	equationPattern : new RegExp("("+decimalPattern+")\\s*(\\/\\%)\\s*("+decimalPattern+")")
};

DivideWithRemainder.evaluate = function(expr) {
	return Math.floor(expr.left/expr.right).toString()+","+(expr.left % expr.right).toString();
};

function allAreNumbers(numArr){
	let i;
	for(i=0;i<numArr.length;i++){
		if(isNaN(numArr[i])){
			return false;
		}
	}
	return true;
}


exports.evaluateExpression = function(math){
	const originalMath = math.trim().replace(/\=$/,"");
	math = originalMath;

	if(allAreNumbers(originalMath.split(/,/g))){
		return math;
	}

	let paren = math.match(insideParens);
	let insideMath;
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
