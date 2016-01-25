# ASL

ASL stands for Arma Scripting Language, a case sensitive C-style scripting language compiled to SQF.
ASL is intended to simplify Arma 3 mod and mission development and eliminate the pain of SQF's bad syntax.

Main reasons for ASL:

* consistent and clean syntax
* less writing
* easier to read and maintain
* easy to learn and understand
* full replacement of SQF
* compatible with Arma wiki and commands
* comfortable

## Usage

ASL is a command line tool. After you have downloaded it, navigate to the binary and execute it:

```
asl.exe [-v|-r|-pretty|--help] <input directory> <output directory>
```

| Parameter | Optional/Required | Description |
| --------- | ----------------- | ----------- |
| -v | optional | Show ASL version. |
| -r | optional | Read input directory recursively. |
| -pretty | optional | Enable pretty printing to SQF. |
| --help | optional | Show usage. |
| input directory | required | Input directory for ASL files (use ./ for relative paths). |
| output directory | required | Output directory for SQF files. Can be the same as input directory (use ./ for relative paths). |

**Example:**

```
asl.exe ./missions/myMission/myScripts ./missions/myMission/compiledScripts
```

Since 1.2.0 ASL requires a [supportInfo](https://community.bistudio.com/wiki/supportInfo) file, which must be generated, named "types" and placed right next to the binary. The content looks like:

```
...
t:DIARY_RECORD
t:LOCATION
b:ARRAY waypointattachobject SCALAR,OBJECT
b:OBJECT,GROUP enableattack BOOL
...
```

The types file will be delivered with the current release, but not updated when Arma is.

## Syntax

### Comments

Comments are written exactly like in SQF:

```
// single line comment

/* multi
 line
 comment */
```

### Variables

Variables are declared using the keyword *var*. They keep the visibility mechanic used by SQF. Identifiers starting with an underscore are considered private.

```
var publicVariable = "value";
var _privateVariable = "value";

var number = 123;
var floatingPointNumber = 1.23;
var string = "string";
var array = [1, 2, 3];

// accessing array elements:
var one = array[0];

// accessing using a statement:
var zwo = array[33/3-2];

// it is possble to use arrays in expressions:
var emptyArray = one-[1];
```

### Control structures

Controll structure syntax is C-like. Notice the same brackets for all structures and no semicolon at the end, unlike in SQF:

```
if 1 < 2 {
    // ...
} else { // no else if yet
    // ...
}

while 1 < 2 {
    // ...
}

for var _i = 0; _i < 100; _i = _i+1 {
    // ...
}

for _i = 0; _i < 100; _i = _i+1 { // same as above, var is optional before identifier "_i"
    // ...
}

foreach _unit => allUnits { // iterates over all units in this case
    // element is available as "_unit" AND "_x" here ("_x" is used by SQF's foreach)
}

switch x { // there is no "break" in SQF
    case 1:
        // ...
    case 2:
        // ...
    default:
        // ...
}

try { // handles errors in "catch" block
    // errors thrown here...
} catch {
    // ... will be handled here
}
```

### Functions

Functions are declared using the keyword *func*. The parameters will be available by their specified identifier.

```
func add(_a, _b) {
    return _a+_b;
}

// Call it:
var _x = add(1, 2); // result in _x is 3
```

Functions support predefined parameters:

```
func add(_a = 0, _b = 0) {
    return _a+_b;
}

// Call it:
var _x = add(); // result in _x is 0
```

When trying to define a function with a name that exists in SQF's build in function set, you'll get an compile error. So declaring "func hint()..." won't compile.

### Call build in commands

To call SQF build in commands (like hint, getDir, addItem, ...) use the same syntax when using functions. An exception are "binary" functions. These are functions which accept parameters on both sides of the function name. Here is an example for "addItem":

```
addItem(someUnit)("NVGoogles");

// output:
someUnit addItem "NVGoogles";
```

Where the first brackets contain the parameters used in front of SQF command and the second ones the parameters behind the SQF command. If more than one parameter is passed, it will be converted to an array. This syntax can be used for **all** build in functions.

```
foo(x, y, z)(1, 2, 3);

// output:
[x, y, z] foo [1, 2, 3];
```

If the build in function accepts no parameters (null function) or on one side only (unary function), it can be called with a single pair of brackets:

```
hint("your text");
shownWatch();
```

### Special functions

There are some special functions in SQF, which also require special syntax in ASL. The examples presented here shows how they are written in ASL and what the output will look like. Remember that ASL is case sensitive!

**exitwith**

```
exitwith { // NOT exitWith!
    // your code
}

// output:
if (true) exitWith {
    // your code
};
```

**waituntil**

```
waituntil(condition); // NOT waitUntil!
// or
waituntil(expression;condition);

// output:
waitUntil {condition};
// or
waitUntil {expression;condition};
```

**code**

The code function is used to compile inline code. This does **not** replace SQF's compile build in function, but will return the contained ASL code as SQF:

```
var x = code("var y = 5;"); // pass as string

// output:
x = {y = 5;};
```

## Preprocessor

The preprocessor works like the original one, with some limitations.
Please visit the link at the bottom, to read about the preprocessor and how to use it. Generally, preprocessor lines must start with the hash key (#) and must stay in their own line. They are always printed as seperate lines in SQF. These features are not supported:

* replacing parts of words
* multi line preprocessor commands
* EXEC (not used in SQF anyway)

If you use *EXEC*, it will be replaced by a function call to it ([] call __EXEC).
*LINE* and *FILE* can be used like normal identifiers:

```
if __LINE__ == 22 {
    // ...
}

if __FILE__ == "myScript.sqf" {
    // ...
}
```

## List of all keywords

Keywords must not be used as identifiers. Here is a full list of all keywords in ASL. Remember that build in function names must not be used neither, else you'll get an compile error.

| Keyword |
| ------- |
| var |
| if |
| while |
| switch |
| for |
| foreach |
| func |
| true |
| false |
| case |
| default |
| return |
| try |
| catch |
| exitwith |
| waituntil |
| code |

## What's missing?

The following features are not implemented yet, but maybe will be in 1.3.0 or a future version:

* scopes
* else if
* selector in expression

scopes won't be supported, since it's a stupid concept and can be replaced by functions.

Selectors in expressions do not work (yet), so this is not possible:

```
var x = ([1, 2, 3]-[1, 2])[0]; // should result in 3
```

## Contribute

To contribute please create pull requests or explain your ideas in the issue section on GitHub. Report any bugs or incompatible ASL <-> SQF syntax you can find.

## Further information

For further information you can read the SQF tutorial and documentation of scripting commands on the Arma wiki.

* [Arma Wiki](https://community.bistudio.com/wiki/Main_Page)
* [Scripting commands](https://community.bistudio.com/wiki/Category:Scripting_Commands_Arma_3)
* [Scripting preprocessor](https://community.bistudio.com/wiki/PreProcessor_Commands)

Interesting pages to visit:

* [Bohemia forum topic](https://forums.bistudio.com/topic/185649-asl-arma-scripting-language-compiler/)
* [Armaholic page](http://www.armaholic.com/page.php?id=29720)

## License

MIT
