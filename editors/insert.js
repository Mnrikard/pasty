exports.calledName = "";
exports.names=["insert"];
var os = require("os");

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

var str = require("../stringHelpers.js");

function defineInsertStatement(columns){
	var tableName = exports.parms[0].value.replace(/\./g,"].[").replace(/\[\[/g, "[").replace(/\]\]/g, "]").replace(/\[\]/g,"");
	return "insert into ["+tableName+"] ("+columns.join(", ")+")\nvalues\n ";
}

function getColumns(line)
{
	return line.split(str.escapeRegex(exports.parms[1].value));
}

function writeSingleRow (cols)
{
	var output = "(";
	for (var j=0; j<cols.length;j++) {
		if (j > 0) {
			output += ", ";
		}

		var apostropheOrNot = str.isNullOrNumber(cols[j]) ? "" : "'";
		output += apostropheOrNot+cols[j].replace(/'/g,"''")+apostropheOrNot;
	}
	output+=")\n";
	return output;
}


exports.edit=function(input, switches){
	var rowOfInsert = 1000;
	var lines = input.split(/\r?\n/g);
	var columnNames = getColumns(lines[0]);
	var insertStatement = defineInsertStatement(columnNames);

	var output = "";

	for (var i=1; i<lines.length; i++)
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
		
		var cols = getColumns(lines[i]);
		output += writeSingleRow(cols);
	}

	return output;
};

