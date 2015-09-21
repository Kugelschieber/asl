var _a = 1;
var _b = 2;

func somefunc(_x, _y, _z) {
    if _x < _y {
        hint(_z);
    }
}

if _a < _b {
    somefunc(1, "two", 3);
}
else{
    _a = 3;
}

hint("this is a hint");
