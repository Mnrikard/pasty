exports.calledName = "";
exports.names=["insert"];
const os = require("os");

exports.parms=[
	{name:"table name",value:null,defaultValue:null},
	{name:"delimiter",value:null,defaultValue:"\t"}
];

exports.helpText = "insert - converts a delimited text to SQL insert statement"+os.EOL+
	"Parameters: "+os.EOL+
	"    TableName: the name of the table to insert into"+os.EOL+
	"    [delimiter]=\"\\t\": the original column delimiter"+os.EOL+
	"Syntax: pasty insert mydb.dbo.mytable [delimiter = \\t]"+os.EOL+os.EOL+
	"Example: pasty insert mydb.dbo.mytable"+os.EOL+
	">> insert into [mydb].[dbo].[mytable] (col1, col2) values ('val1','val2');";
exports.oneLiner = "converts a delimited text to SQL insert statement";

const str = require("../stringHelpers.js");

function defineInsertStatement(columns){
	const tableName = exports.parms[0].value.replace(/\./g,"].[").replace(/\[\[/g, "[").replace(/\]\]/g, "]").replace(/\[\]/g,"");
	return "insert into ["+tableName+"] ("+columns.join(", ")+")\nvalues\n ";
}

function getColumns(line)
{
	return line.split(str.escapeRegex(exports.parms[1].value));
}

function writeSingleRow (cols)
{
	let output = "(";
	let j;
	for (j=0; j<cols.length;j++) {
		if (j > 0) {
			output += ", ";
		}

		let apostropheOrNot = str.isNullOrNumber(cols[j]) ? "" : "'";
		output += apostropheOrNot+cols[j].replace(/'/g,"''")+apostropheOrNot;
	}
	output+=")\n";
	return output;
}


exports.edit=function(input, switches){
	let rowOfInsert = 1000;
	let lines = input.split(/\r?\n/g);
	const columnNames = getColumns(lines[0]);
	const insertStatement = defineInsertStatement(columnNames);

	let output = "";
	let i;

	for (i=1; i<lines.length; i++)
	{
		if (rowOfInsert++ >= 1000)
		{
			output += insertStatement;
			rowOfInsert = 0;
		}
		else
		{
			output+=",";
		}

		let cols = getColumns(lines[i]);
		output += writeSingleRow(cols);
	}

	return output;
};

