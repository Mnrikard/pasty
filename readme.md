Pasty - The Clipboard Editor
============================

Edit the text directly on your clipboard (paste bin/whatever) with some built in functions.

## Installation

`npm install pasty-clipboard-editor`

## Usage

`pasty-clipboard-editor` will install a global command `pasty` to your bin which will be accessible
via the command prompt (or run box if you're on Windows). You have some text on your clipboard, and
you can call one of the built in functions, or one of your own saved functions and your text is updated.

Built-in commands are:

**rep** - replaces with a regular expression
**cap** - capitalizes or lower cases
**columnAlign** - Aligns delimited data by column
**count** - counts characters or lines
**dedup** - Deduplicates a list
**urlencode** - encode/decode a url/xml/base64
**grep** - you know, GREP...
**help** - gets help on functions
**insert** - converts a delimited text to SQL insert statement
**newid** - generates a new UUID
**rep** - replaces with a RegExp
**setText** - sets the content to the passed in string
**sort** - Sorts a list
**xml** - Pretty prints xml inside the string

You can also define your own functions comprised of a set of the above list in a local ~/pasty.json file.
An example file is given in the repository.

Pasty will also work via a shell pipe, so `echo "test" | pasty rep s x ` will output `text`.  This
is very useful in text editors that work with stdio like VIM.
