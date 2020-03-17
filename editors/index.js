let eds = [];

let path = __dirname;

require('fs').readdirSync(path).forEach(function(file){
	if(file != "index.js" && file.match(/\.js$/i)){
		eds.push(require("./"+file));
	}
});

const settings = require("../settings.js").settings;

if(settings["pluginsDirectory"]){
	let pluginPath = settings["pluginsDirectory"];
	if(!/\/$/.test(pluginPath)){
		pluginPath += "/";
	}
	const fs = require('fs');
	if(fs.existsSync(pluginPath)){
		fs.readdirSync(pluginPath).forEach(function(file){
			if(file.match(/\.js$/i)){
				eds.push(require(pluginPath+file));
			}
		});
	}
}

exports.editors = eds;

exports.getEditor = function(name, returnHelp){
	const searchTerm = new RegExp("^"+name+"$","gi");
	let ed,n;
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
	let output = [];
	let ed,n;
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
	let output = [];
	let ed,name,addEd,n;
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
