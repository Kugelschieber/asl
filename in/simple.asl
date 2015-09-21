var _a = 1;
var _b = 2;

func myFunc(_x, _y) {
    if _x < _y {
        myFunc(_y, _x);
    }
}

if _a < _b {
    if _b < _a {
        myFunc(_a, _b+9-(2/2));
    }
}

if (_a+_b)/2 > 10 {
    hint("a");
} else {
    myFunc("multiple", "statements");
    hint("b");
}

myFunc(_a, _b);
myFunc(_a, _b);

var _x = 3;

switch _x {
    case 1:
    case 2:
        somefunc(2);
    default:
        somefunc(3);
        _x = -1;
}
