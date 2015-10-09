# ASL

**ASL is under heavy development and not production ready, the features listed here are not fully implemented. Please visit again when the final version is releaesed. If you like to contribute or if you are just interested, go on.**

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

### Controll structures

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
```

### Functions

Functions are declared using the keyword *func*. The parameters will be available by their specified identifier.

```
func add(_a, _b) {
    return _a+_b;
}
```

Call it:

```
var _x = add(1, 2);
// result is 3
```

### Embed SQF Code

To allow copy and paste of SQF code and reduce migration costs, it is possible (but not recommended) to embed SQF code.
To start a SQF block, use *SQF:* or *sqf:*. To end it use *SQF* or *sqf*.

```
var _a = 1;
var _b = 2;

if _a < _b {
    SQF:
        hint format ["%1 is lower than %2", _a, _b];    
    SQF
}
```

This will compile to:

```
_a = 1;
_b = 2;

if(_a < _b) then {
    hint format ["%1 is lower than %2", _a, _b];
};
```

### Call build in commands

To call SQF build in commands (like hint, getDir, addItem, ...) we have to call them differently.
Since SQF commands can take arguments in the front and behind the keyword, we use a special syntax. So this SQF code:

```
someUnit addItem "NVGoogles";
```

is equivalent to:

```
$addItem(someUnit)("NVGoogles");
```

## Contribute

To contribute please create pull requests or explain your ideas in the issue section on GitHub. Report any bugs or incompatible ASL <-> SQF syntax you can find.

## Further information

For further information you can read the SQF tutorial and documentation of scripting commands on the Arma wiki.

* [Arma Wiki](https://community.bistudio.com/wiki/Main_Page)
* [Scripting commands](https://community.bistudio.com/wiki/Category:Scripting_Commands_Arma_3)

## License

MIT
