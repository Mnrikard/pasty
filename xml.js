exports.names=["xml"];
var os = require("os");

exports.parms=[];

exports.helpText = "xml - Pretty Prints xml inside the string"+os.EOL+
	"Needs not be parseable xml, it will do it's best"+os.EOL+
	"Syntax: pasty xml"+os.EOL+os.EOL+
	"Example: echo <root><child/></root> | pasty xml"+os.EOL+
	">> <root>\n\t<child/>\n</root>";

exports.edit=function(input, switches){
	input = initializeSpaces(input);

	var lines = text.split(/\n/);
	var tabcount = 0;
	var inCommentOrCdata = false;

	//todo: process pasty.json settings file
	var tabstr = "\t";//obt.TabString;

	for (var line = 0; line < lines.Length; line++)
	{
		if(inCommentOrCdata)
		{
			if (endOfCommentOrCdata(lines[line]))
			{
				inCommentOrCdata = false;
			}
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

	return lines.join(os.EOL);
};


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
	return !inComment &&
		!lines [line].match(/<\//) &&
		!lines [line].match(/\/>/) &&
		!lines [line].match(/^<\//) &&
		!lines [line].match(/^<\?xml/) &&
		!lines [line].match(/-->$/) &&
		(lines [line].match(/</) || lines [line].match(/>/));
}


function initializeSpaces(text){
	var spaceBetweenTags = new RegExp(/>\s+</);
	var endTags = new RegExp(/([^>\n])</);
	var textAfterClosedTag = new RegExp(/\/>([^\n<])/);
	var emptyTag = new RegExp(/(<([^\s>]+).+)\n(<\/\2[\s>])/);

	var text = text.replace(spaceBetweenTags, "><")
		.replace(/></,">\n<")
		.replace(endTags,"$1\n<")
		.replace(textAfterClosedTag, "/>\n$1")
		.replace(emptyTag, "$1$3");
	return text;
}
