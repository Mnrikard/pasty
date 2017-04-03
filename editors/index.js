var eds = [];

var path = __dirname;

require('fs').readdirSync(path).forEach(function(file){
	if(file != "index.js" && file.match(/\.js$/i)){
		eds.push(require("./"+file));
	}
});

exports.editors = eds;

exports.getEditor = function(name, returnHelp){
	var searchTerm = new RegExp("^"+name+"$","gi");
	for(var ed=0;ed < eds.length;ed++){
		for(var n=0;n<eds[ed].names.length;n++){
			if(eds[ed].names[n].match(searchTerm)){
				return eds[ed];
			}
		}
	}
	if(returnHelp == "undefined"){ returnHelp = true; }
	if(returnHelp){
		return exports.getEditor("help");
	} else {
		return null;
	}
};

exports.getEditorNames = function(){
	var output = [];
	for(var ed=0;ed < eds.length;ed++){
		var name = eds[ed].names[0];
		var addEd = {"name":name,"aliases":[],"description":eds[ed].oneLiner};
		for(var n=1;n<eds[ed].names.length;n++){
			addEd["aliases"].push(eds[ed].names[n]);
		}
		output.push(addEd);
	}
	return output;
};
