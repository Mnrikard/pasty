exports.makeString = function(text, count){
	var output = "";
	for(var i=0;i<count;i++){
		output += text;
	}
	return output;
};

exports.makeTabs = function(tabcount){
	var settings = require("../settings.js").settings;
	var tabstr = settings.tabString;
	return exports.makeString(tabstr, tabcount);
};

exports.getRegexSwitches = function(switches){
	var output = "";
	if(switches.indexOf("m")>-1){
		output += "m";
	}
	if(switches.indexOf("I") === -1){
		output += "i";
	}
	if(switches.indexOf("G") === -1){
		output += "g";
	}
	return output;
};

exports.isReverse = function(switches){
	return (switches.indexOf("r")>-1);
};

exports.escapeRegex = function(actualText){
	return actualText.replace(/\(/g,"\\(")
		.replace(/\)/g,"\\)")
		.replace(/\+/g,"\\+")
		.replace(/\*/g,"\\*")
		.replace(/\-/g,"\\-")
		.replace(/\./g,"\\.")
		.replace(/\|/g,"\\|");
};

exports.max = function(a, b){
	if(a>b) { return a; }
	return b;
};

exports.same = function(text1, text2){
	return text1.toLowerCase().trim() == text2.toLowerCase().trim();
};

exports.padRight = function(text, withstr, count){
	var whatsleft = count - text.length;
	if(whatsleft > 0){
		return text+exports.makeString(' ', whatsleft);
	}
	return text;
};

exports.isNullOrNumber = function(text){
	if(text.match(/^0\d/)){
		return false;
	}

	return (text === "NULL" || !isNaN(text));
}
