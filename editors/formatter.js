
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

var tabstr = "tab";


function tabRightC(code) {
	var tabcount = 0;
	var keywords = ["abstract","as","base","break","case","catch","checked","continue","default","delegate","do","else","event","explicit","extern","false","finally","fixed","for","foreach","goto","if","implicit","in","interface","internal","is","lock","namespace","new","null","object","operator","out","override","params","private","protected","public","readonly","ref","return","sealed","sizeof","stackalloc","switch","this","throw","true","try","typeof","unchecked","unsafe","using","virtual","while"];
	var types = [ "bool", "byte", "char", "class", "const", "decimal", "double", "enum", "float", "int", "long", "sbyte", "short", "static", "string", "struct", "uint", "ulong", "ushort", "void" ];

	var inString = false;
	var inComment = false;
	var blockString = false;
	var blockComment = false;
	var prevChar = "";
	var setPrvChr = "";

	var sr = code.split(/\r*\n/g);
	var output = [];
	var i;

	sr.forEach(function(line){
		var subtracttab = 0;
		if(!blockString && !blockComment){
			line = line.trim();
		}

		prevChar = os.EOL;
		var c;
		var concurrentQuot;
		var concurrentSlash;
		for (i = 0; i < line.length; i++ ){
			c = line[i];
			setPrvChr = c;

			//strings
			if (c === '"' || blockString){
				inString = true;
				if (prevChar === "@"){
					blockString = true;
					concurrentQuot = 0;
					while (++i < line.length){
						if (line[i] === '"'){
							concurrentQuot++;
						} else if (concurrentQuot > 0) {
							if (concurrentQuot % 2 === 1) {
								i--;
								inString = false;
								blockString = false;
								break;
							}
							concurrentQuot = 0;
						}
					}
				} else {//simple string
					concurrentSlash = 0;
					
					while (++i < line.length) {
						if (line[i] === "\\"){
							concurrentSlash++;
						} else {
							if(line[i] === '"') {
								if (concurrentSlash % 2 === 0) {
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
			else if ((prevChar==="/" && c === '*') || blockComment) {
				if(!blockComment) {
					i++;
					prevChar = "*";
				}
				inComment = true;
				blockComment = true;
				while (++i < line.length) {
					if (line[i] === '/' && prevChar === "*") {
						inComment = false;
						blockComment = false;
						break;
					}
					prevChar = line[i].ToString();
				}
			}

			//comments
			else if (prevChar==="/" && c === '/') {
				while (++i < line.length) {
					buildLine += line[i];
				}
			}

			else if (c === '{' && !inComment && !inString) {
				buildLine += c;
				tabcount++;
			} else if (c === '}' && !inComment && !inString) {
				tabcount--;
				buildLine += `${str.makeString(tabstr, tabcount)}${c}`;
			} else {
				buildLine += c;
			}

			prevChar = setPrvChr;
		}

		if (blockString || blockComment) {
			subtracttab = tabcount * -1;
		}

		output.push(`${str.makeString(tabstr,tabcount+subtracttab)}${buildLine}`);
	});

	return output.join("\n");
}

RegExp.prototype.IsMatch = function(teststr){
	this.lastIndex = 0;
	return this.test(teststr);
};

function tabRightVb(code){
	code = code.replace(/_\s*\n\s*/g," ");
	code = code.replace(/(else)\s*('.+)/gi, "$1\n$2");
	code = code.replace(/(then)\s*('.+)/gi, "$1\n$2");

	var tplus = new RegExp("^(while|for|if|elseif|class|function|sub|select\\s+case|do|(private\\s+|public\\s+|friend\\s+|protected\\s+)?(shared\\s+|mustinherit\\s+|sealed\\s+)?(function|class|sub|property|module))[\\s]", "gi");
	var singleLineIf = new RegExp("then$", "gi");
	var tabBecauseIf = new RegExp("^if", "gi");
	var tminus = new RegExp("^(wend|until|loop|next|elseif|end\\s+(if|function|sub|class|select|property))[\\s]?", "gi");
	var isLabel = new RegExp("[^\\s]:$","g");
	var elif = new RegExp("^elseif.+then$", "gi");

	//special instructions for "select case" statements
	var casebound = new RegExp("^(select case|end select)", "gi");
	var caseitem = new RegExp("^case\s", "gi");

	var incase = false;

	//get rid of multiline rows
	code = code.replace(/_[ \t]+\n[ \t]*/i, " ");

	var rows = code.split(/\n/g);
	var ii = 0;
	var tabcount=0;

	for (ii = 0; ii < rows.length; ii++) {
		var currTabc = 0;

		rows[ii] = rows[ii].trim();

		if (casebound.IsMatch(rows[ii])) {
			if (incase) {
				tabcount--;
				incase = false;
			} else {
				incase = true;
				tabcount++;
				//I know I'm doubling the tab count, I want to
			}
		}

		if (tminus.IsMatch(rows[ii]) && !isLabel.IsMatch(rows[ii])) {
			tabcount--;
		}

		if (tabcount < 0) {
			tabcount = 0;
		}
		currTabc = tabcount;

		var currentLine = "";
		if (isLabel.IsMatch(rows[ii])) {
			currentLine = rows[ii];
		} else if (incase && caseitem.IsMatch(rows[ii]) && tabcount > 0) {
			currentLine = str.makeString(tabstr, tabcount-1) + rows[ii];
		} else if (rows[ii].toLowerCase().trim() == "else" && tabcount > 0) {
			currentLine = str.makeString(tabstr, tabcount - 1) + rows[ii];
		} else {
			currentLine = str.makeString(tabstr, currTabc) + rows[ii];
		}

		if (tplus.IsMatch(rows[ii]) && !isLabel.IsMatch(rows[ii])) {
			console.log("tplus");
			if (tabBecauseIf.IsMatch(rows[ii])) {
				if (singleLineIf.IsMatch(rows[ii])) {
					tabcount++;
				}
			} else {
				tabcount++;
			}
		}
		currTabc = tabcount;
		if (currTabc < 0) { currTabc = 0; }
		rows[ii] = currentLine;
	}
	return rows.join("\n").replace(/(\t|    )(Select\s+Case)/gi,"$2");
}

exports.edit=function(input, switches){
	var obt = require("../settings.js").settings;
	tabstr = obt.tabString;

	if(exports.parms[0].value.trim().toLowerCase() === "c"){
		return tabRightC(input);
	}

	if(exports.parms[0].value.trim().toLowerCase() === "vb"){
		return tabRightVb(input);
	}
};

