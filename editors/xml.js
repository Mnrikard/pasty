exports.calledName = "";
exports.names=["xml"];
const os = require("os");

exports.parms=[];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "xml - Pretty Prints xml inside the string"+os.EOL+
	"Needs not be parseable xml, it will do it's best"+os.EOL+
	"Syntax: pasty xml"+os.EOL+os.EOL+
	"Example: echo <root><child/></root> | pasty xml"+os.EOL+
	">> <root>"+os.EOL+
	">> \t<child/>"+os.EOL+
	">> </root>";
exports.oneLiner = "Pretty prints xml inside the string";

exports.edit=function(input, switches){
	input = initializeSpaces(input);

	let lines = input.split(/\n/g);
	let tabcount = 0;
	let inCommentOrCdata = false;
	let line;

	for (line = 0; line < lines.length; line++)
	{
		if(inCommentOrCdata)
		{
			inCommentOrCdata = !endOfCommentOrCdata(lines[line]);
		}
		else
		{
			inCommentOrCdata = startsCommentOrCdata(lines[line]);
		}

		if (lines[line].match(/^<\//) && !inCommentOrCdata)
		{
			tabcount--;
			if (tabcount < 0) { tabcount = 0; }
		}

		lines[line] = makeTabs(tabcount) + lines[line];

		if (shouldIncrementTabCount(lines, inCommentOrCdata, line))
		{
			tabcount++;
		}

	}

	return lines.join("\n");
};

function makeTabs(tabcount){
	const settings = require("../settings.js").settings;
	const tabstr = settings.tabString;
	let output = "",i;
	for(i=0;i<tabcount;i++){
		output += tabstr;
	}
	return output;
}


function endOfCommentOrCdata(line)
{
	return line.trim().match(/\-\->$/) || line.trim().match(/\]\]>$/);
}

function startsCommentOrCdata(line)
{
	line = line.trim();
	return (line.match(/^<!\-\-/) && !line.match(/\-\->/)) ||
		(line.match(/^<!\[CDATA\[/) && !line.match(/\]\]>/));
}

function shouldIncrementTabCount(lines, inComment, line)
{
	let output = !inComment &&
		!lines[line].match(/<\//) &&
		!lines[line].match(/\/>/) &&
		!lines[line].match(/^<\//) &&
		!lines[line].match(/^<\?xml/) &&
		!lines[line].match(/-->$/) &&
		(lines[line].match(/</) || lines[line].match(/>/));
	return output;
}


function initializeSpaces(text){
	const spaceBetweenTags = new RegExp(/>\s+</g);
	const endTags = new RegExp(/([^>\n])</g);
	const textAfterClosedTag = new RegExp(/\/>([^\n<])/g);
	const emptyTag = new RegExp(/(<([^\s>]+).+)\n(<\/\2[\s>])/g);

	text = text.replace(spaceBetweenTags, "><")
		.replace(/></g,">\n<")
		.replace(endTags,"$1\n<")
		.replace(textAfterClosedTag, "/>\n$1")
		.replace(emptyTag, "$1$3");
	return text;
}
