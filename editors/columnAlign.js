exports.calledName = "";
exports.names=["columnAlign","align"];
const os = require("os");

exports.parms=[{
	name:"number of spaces",
	value:null,
	defaultValue:"2"
},{
	name:"delimiter",
	value:null,
	defaultValue:"\t"
}];

exports.getParms = function(){
	return exports.parms;
};

exports.helpText = "columnAlign - Aligns delimited data by column"+os.EOL+
	"Parameters: "+os.EOL+
	"    [number of spaces]=2, how many spaces between columns"+os.EOL+
	"    [delimiter]=\"\\t\", original column delimiter"+os.EOL+
	"Syntax: pasty columnAlign [\"number\"] [\"delimiter\"]"+os.EOL+os.EOL+
	"Example: cat tabDelimited.file | pasty align"+os.EOL+
	">> col1   col2             col3"+os.EOL+
	">> names  some other data  1234";
exports.oneLiner = "Aligns delimited data by column";

const str = require("../stringHelpers.js");

let _columnLengths = [];

function splitLine(line){
	const splitter = new RegExp(str.escapeRegex(exports.parms[1].value), "ig");
	return line.split(splitter);
}

function rebuildRows (rows) {
	const columnSeparator = str.makeString(" ", exports.parms[0].value);
	let r,c;

	for(r=0; r<rows.length; r++){
		const cols = splitLine(rows[r]);
		for (c = 0; c < cols.length; c++)
		{
			cols[c] = str.padRight(cols[c], ' ', _columnLengths[c]);
		}
		rows[r] = cols.join(columnSeparator).trim();
	}
}

function defineColumnLengths(rows){
	_columnLengths = [];
	let r,c;
	for (r=0; r<rows.length;r++){
		const cols = splitLine(rows[r]);
		for (c=0; c<cols.length; c++){
			if (c >= _columnLengths.length)
			{
				_columnLengths.push(cols[c].length);
			}
			_columnLengths[c] = str.max(_columnLengths[c], cols[c].length);
		}
	}
}

exports.edit=function(input, switches){
	let rows = input.split(/\r?\n/g);
	defineColumnLengths(rows);
	rebuildRows(rows);
	return rows.join("\n");
};

