
exports.calledName = "";
exports.names=["format","tabright"];
var os = require("os");

exports.parms=[{
	name:"language",
	value:null,
	defaultValue:null
}];

exports.helpText = "format - indents your code"+os.EOL+
	"Parameters: "+os.EOL+
	"    Language: one of c|vb|json"+os.EOL+
	"Syntax: pasty format <c|vb|json>"+os.EOL+os.EOL+
	"Example: echo \"{name:value}\" | pasty format json"+os.EOL+
	">> {"+os.EOL+
	">> 	name:value"+os.EOL+
	">> }";
exports.oneLiner = "indents your code";

var str = require("../stringHelpers.js");


function tabRightC(code)
{
	var tabcount = 0;
	var keywords = ["abstract","as","base","break","case","catch","checked","continue","default","delegate","do","else","event","explicit","extern","false","finally","fixed","for","foreach","goto","if","implicit","in","interface","internal","is","lock","namespace","new","null","object","operator","out","override","params","private","protected","public","readonly","ref","return","sealed","sizeof","stackalloc","switch","this","throw","true","try","typeof","unchecked","unsafe","using","virtual","while"];
	var types = [ "bool", "byte", "char", "class", "const", "decimal", "double", "enum", "float", "int", "long", "sbyte", "short", "static", "string", "struct", "uint", "ulong", "ushort", "void" ];

	var inString = false;
	var inComment = false;
	var blockString = false;
	var blockComment = false;
	var prevChar = "";
	var setPrvChr = "";

	var obt = require("../settings.js").settings;
	var tabstr = obt.tabString;

	var sr = code.split(/\r*\n/g);
	var output = "";
	var i;

	sr.forEach(function(line){
		var subtracttab = 0;
		var buildLine = "";
		if(!blockString && !blockComment){
			line = line.trim();
		}
		prevChar = os.EOL;
		for (i = 0; i < line.length; i++ ){
			var c = line[i];
			setPrvChr = c;

			//strings
			if (c === '"' || blockString){
				buildLine += c;
				inString = true;
				if (prevChar === "@"){
					blockString = true;
					var concurrentQuot = 0;
					while (++i < line.length){
						if (line[i] === '"'){
							concurrentQuot++;
							buildLine += line[i];
						} else if (concurrentQuot > 0) {
							if (concurrentQuot % 2 == 1) {
								i--;
								inString = false;
								blockString = false;
								break;
							} else {
								buildLine += line[i];
							}
							concurrentQuot = 0;
						} else {
							buildLine += line[i];
						}
					}
				} else {//simple string
					var concurrentSlash = 0;
					
					while (++i < line.length) {
						buildLine += line[i];

						if (line[i] === "\\"){
							concurrentSlash++;
						} else {
							if(line[i] == '"') {
								if (concurrentSlash % 2 == 0) {
									inString = false;
									break;
								}
							}
							concurrentSlash = 0;
						}
					}
				}
			}

			//Block comments
			else if ((prevChar=="/" && c == '*') || blockComment) {
				buildLine += c;
				if(!blockComment) {
					i++;
					prevChar = "*";
				}
				inComment = true;
				blockComment = true;
				while (++i < line.length) {
					buildLine += line[i];

					if (line[i] == '/' && prevChar == "*") {
						inComment = false;
						blockComment = false;
						break;
					}
					prevChar = line[i].ToString();
				}
			}

			//comments

			else if (prevChar=="/" && c == '/') {
				buildLine += c;
				while (++i < line.Length) {
					buildLine += line[i];
				}
			}

			else if (c == '{' && !inComment && !inString) {
				buildLine += c + os.EOL;
				tabcount++;
			} else if (c == '}' && !inComment && !inString) {
				tabcount--;
				buildLine += `${os.EOL}${str.makeString(tabstr, tabcount)}${c}${os.EOL}`;
			} else if (c == ';' && !inComment && !inString) {
				buildLine += `${c}${os.EOL}`;
			} else {
				buildLine += c;
			}

			prevChar = setPrvChr;
		}

		if (blockString || blockComment) {
			subtracttab = tabcount * -1;
		}

		output += `${str.makeString(tabstr,tabcount+subtracttab)}${buildLine}${os.EOL}`;
	});

	//output = Regex.Replace(output, "\n+", "\n");

	return output;
}

exports.edit=function(input, switches){
	return tabRightC(input);
};

