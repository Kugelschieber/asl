# ASL

ASL stands for Arma Scripting Language, a C-style scripting language compiled to SQF.
ASL is intended to simplify Arma 3 mod and mission development and eliminate the pain of SQF's bad syntax.

Main reasons for ASL:

* consistent and clean syntax
* less writing
* easier to read and maintain
* easy to learn and understand
* full replacement of SQF
* compatible with Arma wiki and commands

The compiler is written in Go and implemented as a simple recursive decent parser and uses concurrency to compile multiple files at once, which makes it really fast.

## Usage

ASL is a command line tool. After you have downloaded it, execute it using the terminal/"cmd" by navigating to the binary or setting PATH variable and executing the following statement:

```
asl [-v|-r|-pretty|--help] <input directory> <output directory>
```

| Parameter | Optional/Required | Meaning |
| --------- | ----------------- | ------- |
| -v | optional | Shows ASL version. |
| -r | optional | Read input directory recursively. |
| -pretty | optional | Enable pretty printing to SQF. |
| --help | optional | Show usage. |
| input directory | required | Directory to read ASL files from. |
| output directory | required | Directory for SQF output. Can be the same as input directory. |

**Example:**

```
asl ./missions/myMission/myScripts ./missions/myMission/compiledScripts
```

## Syntax

### Comments

Comments are declared exactly like in SQF:

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
```

### Control structures

Controll structure syntax is C-like. Notice they are all using the same brackets and do not require to set a semicolon at the end.

```
if 1 < 2 {
    // ...
}

while 1 < 2 {
    // ...
}

for var _i = 0; _i < 100; _i = _i+1 { // var before identifier is optional
    // ...
}

each allUnits { // foreach, iterates over all units in this case
    // element is available as _x here
}

switch x {
    case 1:
        // ...
        break;
    case 2:
        // ...
        break;
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
// result is 3
```

### Call build in commands

To call SQF build in commands (like hint, getDir, addItem, ...) we have to call them differently.
Since SQF commands can take arguments in the front and behind the keyword, we use a special syntax. So this SQF code:

```
someUnit addItem "NVGoogles";
```

is equivalent to:

```
addItem(someUnit)("NVGoogles");
```

Where the first brackets contain the parameters used in front of SQF command and the second ones behind SQF command. If more than one parameter is passed, it will be converted to an array.

```
foo(x, y, z)(1, 2, 3);
// will be:
[x, y, z] foo [1, 2, 3];
```

### Special functions

There are some special functions in SQF, which also require special syntax in ASL. The examples presented here show how they are written in ASL and what the output will look like. Remember that ASL is case sensitive!

**exitWith**

```
exitwith {
    // your code
}

// output:
if (true) exitWith {
    // your code
};
```

**waitUntil**

```
waituntil(condition);
// or
waituntil(expression;condition);

// output:
waitUntil {condition};
// or
waitUntil {expression;condition};
```

## List of all keywords

Keywords should not be used as identifiers. Here is a full list of all keywords in ASL. Remember that build in functions should not be used neither.

| Keyword |
| ------- |
| var |
| if |
| while |
| witch |
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

## Contribute

To contribute please create pull requests or explain your ideas in the issue section on GitHub. Report any bugs or incompatible ASL <-> SQF syntax you can find.

## Further information

For further information you can read the SQF tutorial and documentation of scripting commands on the Arma wiki.

* [Arma Wiki](https://community.bistudio.com/wiki/Main_Page)
* [Scripting commands](https://community.bistudio.com/wiki/Category:Scripting_Commands_Arma_3)

## License

MIT
