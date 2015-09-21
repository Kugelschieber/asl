// comment
func Foo(_a, _b) {
    var _i = 0;

    // yeah...
    while _a < _b {
        /*
        So this will just compile to plain SQF.
        */
        sqf:
            sleep 3;
            hint format ["%1", _i];
        sqf
        
        _i = _i+1;
    }
}

/* comment */
var _a = 1;
var _b = 2;

Foo(_a, _b);
