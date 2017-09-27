var complete = require('complete');

var completionObj = {
  program: 'pasty',
  // Commands
  commands: { },
  // Position-independent options.
  // These will attempted to be
  // matched if `commands` fails
  // to match.
  options: {
    '-g': {},
    '-r': {},
    '-i': {},
    '-m': {},
    '-p': {}
  }
};

var nameonlyComplete = function(words,prev,cur){
	complete.output(cur, []);
};

var names = require("./editors").getEditorNames();
var i,a;
for(i=0;i<names.length;i++){
	completionObj.commands[names[i].name] = nameonlyComplete;
	for(a=0;a<names[i].aliases.length;a++){
		completionObj.commands[names[i].aliases[a]] = nameonlyComplete;
	}
}

complete(completionObj);
