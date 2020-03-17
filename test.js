var system = require("./editorRunner.js");
require("./stringHelpers.js").keepWindowOpen = function(){};

var chalk = require("chalk");
var errors = 0;

var lastLog = "";
console.log = (function(logFn) {
	return function(msg) {
		lastLog = msg;
		logFn.call(console, msg);
	};
})(console.log);

function test(input, output, args, testName){
	var outcome = system.handleInput(input, args);
	if(output !== outcome){
		console.log("\n");
		console.log(chalk.red(testName+" failed\n"));
		console.log(chalk.red("expected \""+output+"\"\r\n but was \""+outcome+"\""));
		console.log("\n");
		errors++;
		return false;
	}
	console.log(chalk.green(testName+" passed"));
	return true;
}

function testMessage(input, message, args, testName){
	var outcome = system.handleInput(input, args);
	outcome = lastLog;
	if(message !== lastLog){
		console.log("\n");
		console.log(chalk.red(testName+" failed\n"));
		console.log(chalk.red("expected \""+message+"\"\r\n but was \""+outcome+"\""));
		console.log("\n");
		errors++;
		return false;
	}
	console.log(chalk.green(testName+" passed"));
	return true;
}
/*tests go here*/

test("test", "text", ["rep","s","x"],"Standard Rep");
test("FIND this and didn't find this", "found this and didn't find this", ["rep","FIND","found","-I"],"Case Sensitive Rep");
test("the big brown dog barked", "The Big Brown Dog Barked", ["rep","(\\w)(\\w+)","\\u$1$2"],"Replacing with different case");

test("this is upper case","THIS IS UPPER CASE",["cap"],"Capitalize");
test("this is upper case","THIS IS UPPER CASE",["upper"],"Upper case");
test("THIS IS UPPER CASE","this is upper case",["lower"],"Lower case");

test("&lt;root/&gt;","<root/>",["xmldecode"],"Xml Decode");
test("<root/>","&lt;root/&gt;",["xmlencode"],"Xml Encode");

test("<root/>","%3Croot%2F%3E",["urlencode"],"URL Encode");
test("%3Croot%2F%3E","<root/>",["urldecode"],"URL Decode");

test("this is base64","dGhpcyBpcyBiYXNlNjQ=", ["base64encode"],"Base64 Encoding");
test("dGhpcyBpcyBiYXNlNjQ=","this is base64", ["base64decode"],"Base64 Decoding");

test(
	"c1,c2,c3,name,dob\ncol1,column2,column3,test,1/1/01\ncolumn2,c1,c4,success,12/31/2008",
	"c1       c2       c3       name     dob\ncol1     column2  column3  test     1/1/01\ncolumn2  c1       c4       success  12/31/2008",
	["columnAlign","2",","],
	"Column Alignment");

testMessage("123456789🍔","10 characters", ["count"], "Counting with special characters");
testMessage("123456789🍔","11 bytes", ["count","bytes"], "Counting with special characters");
testMessage("123456789🍔","1 lines", ["count","lines"], "Counting with special characters");



/*end of tests*/


if(errors === 0){
	console.log(chalk.green("\n\nAll tests passed"));
}else{
	console.log(chalk.red("\n\n"+errors+" test failed"));
}
