var os = require("os");
var fs = require("fs");

function getSettings(){
	var settingsPath = getSettingsFile();
	if(fs.existsSync(settingsPath)){
		var settings = fs.readFileSync(settingsPath);
		return JSON.parse(settings);
	} else {
		return {"tabString":"\t"};
	}
	//if file exists, consume and return the object
	//otherwise return the default object
}

function getSettingsFile(){
	return os.homedir()+"/pasty.json";
}

exports.settings = getSettings();
exports.settingsFile = getSettingsFile();
