exports.names = ["help"];
exports.parms = [{
	name:"function",
	value:null,
	defaultValue:"."
}];


exports.helpText = "help [functionName]";
exports.edit = function(input, switches){
	var searchTerm = new RegExp(exports.parms[0].value,"gi");
	var editors = require("./editorlist.js").editors;
	for(var ed=0;ed < editors.length;ed++){
		for(var n=0;n<editors[ed].names.length;n++){
			if(editors[ed].names[n].match(searchTerm)){
				console.log(editors[ed].helpText);
			}
		}
	}
	return input;
};
