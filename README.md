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

ASL is a command line tool. After you have downloaded it, navigation to the binary and execute it:

```
asl [-v|-r|-pretty|--help] <input directory> <output directory>
```

| Parameter | Optional/Required | Meaning |
| --------- | ----------------- | ------- |
| -v | optional | Shows ASL version. |
| -r | optional | Read input directory recursively. |
| -pretty | optional | Enable pretty printing to SQF. |
| --help | optional | Show usage. |
| input directory | required | Directory to read ASL files from (use ./ for relative paths). |
| output directory | required | Directory for SQF output. Can be the same as input directory (use ./ for relative paths). |

**Example:**

```
asl ./missions/myMission/myScripts ./missions/myMission/compiledScripts
```

Since 1.2.0 ASL requires a [supportInfo](https://community.bistudio.com/wiki/supportInfo) file, which must be generated, named "types" and placed right next to the binary. So if you have the asl.exe there must also be a file called types right next to it. The content looks like:

```
...
t:DIARY_RECORD
t:LOCATION
b:ARRAY waypointattachobject SCALAR,OBJECT
b:OBJECT,GROUP enableattack BOOL
...
```

A current types file will be delivered with the current release, but not updated when Arma is.

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
var emptyArray = one-[0];
```

### Control structures

Controll structure syntax is C-like. Notice they are all using the same brackets and do not require to set a semicolon at the end, unlike in SQF.

```
if 1 < 2 {
    // ...
} else { // no else if yet
    // ...
}

while 1 < 2 {
    // ...
}

for var _i = 0; _i < 100; _i = _i+1 { // var before identifier is optional
    // ...
}

foreach unit => allUnits { // foreach, iterates over all units in this case
    // element is available as "unit" here
    // _x is still available due to how SQF works!
}

switch x { // there is no "break" in SQF
    case 1:
        // ...
    case 2:
        // ...
    default:
        // ...
}

try {
    // ...
} catch {
    // ...
}
```

### Functions

Functions are declared using the keyword *func*. The parameters will be available by their specified identifier.

```
func add(_a, _b) {
    return _a+_b;
}

// Call it:

var _x = add(1, 2);
// result in _x is 3
```

Functions support predefined parameters:

```
func add(_a = 0, _b = 0) {
    return _a+_b;
}

// Call it:

var _x = add();
// result in _x is 0
```

When trying to define a function with a name that exists in the build in function set, you'll get an compile error.

### Call build in commands

To call SQF build in commands (like hint, getDir, addItem, ...) we have to use a different syntax.

```
addItem(someUnit)("NVGoogles");

// output:
someUnit addItem "NVGoogles";
```

Where the first brackets contain the parameters used in front of SQF command and the second ones behind SQF command. If more than one parameter is passed, it will be converted to an array. This syntax can be used for **all** build in functions (also spawn and so on).

```
foo(x, y, z)(1, 2, 3);

// output:
[x, y, z] foo [1, 2, 3];
```

If the build in function does not accept parameters or only on one side (unary function), it can be called with a single pair of brackets:

```
hint("your text");
shownWatch();
```

### Special functions

There are some special functions in SQF, which also require special syntax in ASL. The examples presented here show how they are written in ASL and what the output will look like. Remember that ASL is case sensitive!

**exitWith**

```
exitwith { // NOT exitWith!
    // your code
}

// output:
if (true) exitWith {
    // your code
};
```

**waitUntil**

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

The code function is used to compile inline code. This does **not** replace SQF compile buildin function, but will return the contained ASL code as SQF.

```
// input:
var x = code("var y = 5;"); // pass as string

// output:
x = {y = 5;};
```

## Preprocessor

The preprocessor works like the original one, with some limitations.
Please visit the link at the bottom, to read about the preprocessor and how to use it. Generally, preprocessor lines must start with the hash key (#) and must stay in their own line. They are always printed as seperate lines. These features are not supported:

* replacing parts of words
* multi line preprocessor commands
* __EXEC (not used in SQF anyway)

If you use *__EXEC*, it will be replaced by a function call to it ([] call __EXEC).
*__LINE__* and *__FILE__* can be used, since they are identifiers:

```
if __LINE__ == 22 {
    // ...
}

if __FILE__ == "myScript.sqf" {
    // ...
}
```

## List of all keywords

Keywords should not be used as identifiers. Here is a full list of all keywords in ASL. Remember that build in function names should not be used neither.

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

The following features are not implemented yet, but will be in 1.3.0 or a future version:

* scopes
* else if
* selector in expression

scopes won't be supported, since they are a stupid concept and can be replaced by functions.

Selectors in expressions do not work (yet):

```
var x = ([1, 2, 3]-[1, 2])[0]; // should result in 3, but does not work
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
