var eds = [];

var path = __dirname;

require('fs').readdirSync(path).forEach(function(file){
	if(file != "index.js" && file.match(/\.js$/i)){
		eds.push(require("./"+file));
	}
});

var settings = require("../settings.js").settings;

if(settings["pluginsDirectory"]){
	var pluginPath = settings["pluginsDirectory"];
	if(!/\/$/.test(pluginPath)){
		pluginPath += "/";
	}
	require('fs').readdirSync(pluginPath).forEach(function(file){
		if(file.match(/\.js$/i)){
			eds.push(require(pluginPath+file));
		}
	});
}

exports.editors = eds;

exports.getEditor = function(name, returnHelp){
	var searchTerm = new RegExp("^"+name+"$","gi");
	var ed,n;
	for(ed=0;ed < eds.length;ed++){
		for(n=0;n<eds[ed].names.length;n++){
			if(eds[ed].names[n].match(searchTerm)){
				return eds[ed];
			}
		}
	}
	if(returnHelp === "undefined"){ returnHelp = true; }
	if(returnHelp){
		return exports.getEditor("help");
	}
	return null;
};

exports.getAllFuncNames = function(){
	var output = [];
	var ed,n;
	for(ed=0;ed < eds.length;ed++){
		for(n=0;n<eds[ed].names.length;n++){
			if(eds[ed].names[n] === "userfunc"){
				continue;
			}
			output.push(eds[ed].names[n]);
		}
	}
	return output.sort();
};

exports.getEditorNames = function(){
	var output = [];
	var ed,name,addEd,n;
	for(ed=0;ed < eds.length;ed++){
		name = eds[ed].names[0];
		addEd = {"name":name,"aliases":[],"description":eds[ed].oneLiner};
		for(n=1;n<eds[ed].names.length;n++){
			addEd.aliases.push(eds[ed].names[n]);
		}
		output.push(addEd);
	}
	return output;
};
