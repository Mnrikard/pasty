var os = require("os");
var fs = require("fs");

function getSettings(){
	var settingsPath = os.homedir()+"/pasty.json";
	if(fs.existsSync(settingsPath)){
		var settings = fs.readFileSync(settingsPath);
		return JSON.parse(settings);
	} else {
		return {"tabString":"\t"};
	}
	//if file exists, consume and return the object
	//otherwise return the default object
}

exports.settings = getSettings();
