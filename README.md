# Pasty CLI

Pasty is a macro tool built for your clipboard. In many editors, you can define macros to handle text manipulation,
but they are limited and editor specific. With **pasty**, you can have powerful macros tied to the text on your
clipboard in any editor you wish.

If you're like me, you have found yourself often needing to do some edits on the text your working on. For instance,
you need to decode this base64 text. So you:

1. Copy the text
2. Find an online base64 decoder
3. Paste your text into that tool
4. Click the button
5. Copy the output text
6. Go back to your editor and paste it in

With Pasty, you simply

1. Copy the text
2. Run "pasty base64decode"
3. Paste the output back to your editor

If you're using an editor that understands stdin/stdout, like vim, you can simply select the text and run:
`:'<,'>!pasty base64decode`

## Installation

### Using go install

```sh
go install github.com/Mnrikard/pasty@latest
```

### Build from source

```sh
git clone https://github.com/Mnrikard/pasty.git
cd pasty
go install ./
```

### Install the tab completion for your terminal (Optional)

Once you have pasty installed, you can run:

#### Powershell
```sh
pasty completion powershell > $HOME/.pastycompletion.ps1
echo ". $HOME/.pastycompletion.ps1" >> $PROFILE
```

#### Posix
```sh
pasty completion [bash|fish|zsh] > ~/.pastycompletion.sh
chmod +x ~/.pastycompletion.sh
echo '. ~/.pastycompletion.sh` >> ~/.zshrc
```
where `~/.zshrc` is either your zshrc, bashrc, or fishrc file

Then, restart your shell for the settings to take effect

## Security

Pasty has zero ability to reach a network. There is no telemetry, no phone-home, nothing to compromise the security
of the text on your clipboard. This is by design. Even the plugins you write for pasty only have basic math/string 
manipulation capabilities.

## Usage

### Clipboard

The primary usage of pasty is to modify the text on your clipboard. So just copy something and run a command to modify
the text on your clipboard. I find it useful to use the quick command interpreters built into most desktop environments

* Windows: Bring up the run box by typing "Win+R" and you can type in your pasty command quickly
* KDE: Alt+F2 brings up a command interpreter
* Gnome: Alt+F2 as well
* MacOs: can finder do this? I don't know, I don't have a mac

Or, if you're running from a terminal, you can take advantage of the tab completion available there.

### StdIn/StdOut

Pasty can pipe input from `stdin` in your terminal. Just run your output through pasty to execute the macros on your
standard input instead of your clipboard

**Example:**
```
cat myfile | pasty sort | pasty dedup > sortedSingles
```

## User defined functions

You can store commonly used macros and chains of macros in your `~/.config/pasty/user_defined.json` file. The syntax
for new UDFs is:

```json
	{
		"name":"nameOfMacro", //define the name of the macro here, you'll call it with `pasty udf nameOfMacro`
		"commands":[//a list of commands to execute in the order defined
			{
				"name":"rep",//The name of the pasty function to run
				"args":["{{replaceWhat}}", "b"] //The list of arguments passed to the function
				//You can create parameters to pass into your UDF with double-handlebars
				//You'll need to define them below
			}
		],
		"parameters":[//a list of parameters to pass to the UDF
			{
				"name":"replaceWhat",//the name should match the handlebar replacement text above
				"defaultValue":null
			}
		]
	}
```

## Plugins

If User Defined Functions don't solve your problem and you need more control over how to manipulate your text, you can
create your own plugin. 

Plugins are Lua files stored in your home directory under `~/.config/pasty/plugins/`. The version of Lua currently
supported is 5.1, and you have access to string and math functions, but no access to spawn processes or break out of
the sandbox.

The minimal structure of these applications is:

```lua
function process(input)
    return "new text"
end
```

Where the input is a string and the output is also a string. You can manipulate that string within your process
function and create cool new text modifications.

## Where is the old nodejs app? Can I still install from npm?

Due to security concerns around how plugins were handled in the old nodejs app and because of the inherent slowness
in interpreted code, it was decided to move to this app written in go and using sandboxed Lua plugins. The old nodejs
app is still available on npm (albeit outdated) and the source code is available on the
[LegacyNodeJs release](https://github.com/Mnrikard/pasty/releases/tag/LegacyNodeJs) on github.
