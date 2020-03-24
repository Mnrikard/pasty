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

//cap
test("this is upper case","THIS IS UPPER CASE",["cap"],"Capitalize");
test("this is upper case","THIS IS UPPER CASE",["upper"],"Upper case");
test("THIS IS UPPER CASE","this is upper case",["lower"],"Lower case");

//columnAlign
test(
	"c1,c2,c3,name,dob\ncol1,column2,column3,test,1/1/01\ncolumn2,c1,c4,success,12/31/2008",
	"c1       c2       c3       name     dob\ncol1     column2  column3  test     1/1/01\ncolumn2  c1       c4       success  12/31/2008",
	["columnAlign","2",","],
	"Column Alignment");

//count
testMessage("123456789🍔","10 characters", ["count"], "Counting with special characters");
testMessage("123456789🍔","11 bytes", ["count","bytes"], "Counting with special characters");
testMessage("123456789🍔","1 lines", ["count","lines"], "Counting with special characters");

//dedup
test("one, two, three, two, seven, three","one, two, three, seven",["dedup", ", "], "Dedup")
test("one\ntwo\nthree\ntwo\nseven\nthree","one\ntwo\nthree\nseven",["dedup"], "Dedup")

//encode
test("&lt;root/&gt;","<root/>",["xmldecode"],"Xml Decode");
test("<root/>","&lt;root/&gt;",["xmlencode"],"Xml Encode");

test("<root/>","%3Croot%2F%3E",["urlencode"],"URL Encode");
test("%3Croot%2F%3E","<root/>",["urldecode"],"URL Decode");

test("this is base64","dGhpcyBpcyBiYXNlNjQ=", ["base64encode"],"Base64 Encoding");
test("dGhpcyBpcyBiYXNlNjQ=","this is base64", ["base64decode"],"Base64 Decoding");

//grep
test("test\ntest2\nnot found\nnot 2 found\n","test2\nnot 2 found", ["grep","\\d"], "Grep for pattern")
test("test\ntest2\nnot found\nnot 2 found\n","2\n2", ["grep","\\d","-L"], "Grep for pattern, just pattern")

//insert
test(
	"c1\tc2\tc3\tname\tdob\ncol1\tcolumn2\tcolumn3\ttest\t1/1/01\ncolumn2\tc1\tc4\tsuccess\t12/31/2008",
	"insert into [db].[schema].[table] (c1, c2, c3, name, dob)\nvalues\n ('col1', 'column2', 'column3', 'test', '1/1/01')\n,('column2', 'c1', 'c4', 'success', '12/31/2008')\n",
	["insert","db.schema.table"],
	"Standard Insert");

test(
	"c1\tc2\tc3\tname\tdob\ncol1\t123.45\tcolumn3\ttest\t1/1/01\ncolumn2\t-23\tc4\tnull\t12/31/2008",
	"insert into [db].[schema].[table] (c1, c2, c3, name, dob)\nvalues\n ('col1', 123.45, 'column3', 'test', '1/1/01')\n,('column2', -23, 'c4', null, '12/31/2008')\n",
	["insert","db.schema.table"],
	"Weird Insert");

//math (experimental)
test("1+2", "1+2=3",["math"],"Simple Math")
test("(1+2)*3", "9",["math","answer"],"Pemdas")

//rep
test("test", "text", ["rep","s","x"],"Standard Rep");
test("FIND this and didn't find this", "found this and didn't find this", ["rep","FIND","found","-I"],"Case Sensitive Rep");
test("the big brown dog barked", "The Big Brown Dog Barked", ["rep","(\\w)(\\w+)","\\u$1$2"],"Replacing with different case");
test("the big brown dog barked", "the big brown dog barked 7 times", ["rep",".+","$0 7 times"],"Special zero group");

//setText
test("", "This is the text I set", ["settext","This is the text I set"],"Set Text")

//sort
test("1\n2019-1-1\n7\nAardvark\nPickles\nMayonnaise\n2020-2-27\n23\14","1\n7\n23\n2019-1-1\n2020-2-27\nAardvark\nMayonnaise\nPickles",["sort"],"Sort Normal")
test("1\n2019-1-1\n7\nAardvark\nPickles\nMayonnaise\n2020-2-27\n23\14","Pickles\nMayonnaise\nAardvark\n2020-2-27\n2019-1-1\n23\n7\n1",["sort","-r"],"Sort Reverse")

//toNumBase
test("15","F",["ToBase","16"],"10 to 16")
test("16","10",["ToBase","16"],"10 to 16")
test("253","FD",["ToBase","16"],"10 to 16 again")

/*end of tests*/


if(errors === 0){
	console.log(chalk.green("\n\nAll tests passed"));
}else{
	console.log(chalk.red("\n\n"+errors+" test failed"));
}
