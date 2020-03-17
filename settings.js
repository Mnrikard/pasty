const os = require("os");
const fs = require("fs");

function getSettings(){
	const settingsPath = getSettingsFile();
	if(fs.existsSync(settingsPath)){
		const settings = fs.readFileSync(settingsPath);
		const output = JSON.parse(settings);
		fullyQualifyHomeDirectory(output);
		return output;
	} else {
		return {"tabString":"\t"};
	}
	//if file exists, consume and return the object
	//otherwise return the default object
}

function fullyQualifyHomeDirectory(settingsObject){
	settingsObject.pluginsDirectory = settingsObject.pluginsDirectory.replace("~", os.homedir());
	settingsObject.localStyleSheet = settingsObject.localStyleSheet.replace("~", os.homedir());
}

function getSettingsFile(){
	return os.homedir()+"/pasty.json";
}

exports.settings = getSettings();
exports.settingsFile = getSettingsFile();
