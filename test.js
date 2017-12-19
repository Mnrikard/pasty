var system = require("./editorRunner.js");
require("./stringHelpers.js").keepWindowOpen = function(){};

var chalk = require("chalk");

function test(input, output, args, testName){
	var outcome = system.handleInput(input, args);
	if(output !== outcome){
		console.log("\n");
		console.log(chalk.red(testName+" failed\n"));
		console.log(chalk.red("expected \""+output+"\"\r\n but was \""+outcome+"\""));
		console.log("\n");
		return false;
	}
	console.log(chalk.green(testName+" passed"));
	return true;
}

var errors = 0;

/*tests go here*/

test("test", "text", ["rep","s","x"],"Standard Rep");
test("FIND this and didn't find this", "found this and didn't find this", ["rep","FIND","found","-I"],"Case Sensitive Rep");
test("the big brown dog barked", "The Big Brown Dog Barked", ["rep","(\\w)(\\w+)","\\u$1$2"],"Replacing with different case");

test("this is upper case","THIS IS UPPER CASE",["cap"],"Capitalize");
test("this is upper case","THIS IS UPPER CASE",["upper"],"Upper case");
test("THIS IS UPPER CASE","this is upper case",["lower"],"Lower case");

test("&lt;root/&gt;","<root/>",["xmldecode"],"Xml Decode");
test("<root/>","&lt;root/&gt;",["xmlencode"],"Xml Encode");

test("this is base64","dGhpcyBpcyBiYXNlNjQ=", ["base64encode"],"Base64 Encoding");
test("dGhpcyBpcyBiYXNlNjQ=","this is base64", ["base64decode"],"Base64 Decoding");


/*end of tests*/


if(errors === 0){
	console.log(chalk.green("\n\nAll tests passed"));
}else{
	console.log(chalk.red("\n\n"+errors+" test failed"));
}
