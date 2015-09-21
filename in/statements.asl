func foo(_out) {
    hint(_out);
}

func bar(_out) {
    hint(reverse(_out));
}

func inlineFunction(_multiple, _parameters, _should, _work) {
    _a = "Not working yet...";
}

var _array;

each array {
    if _x > 200 {
        foo(_x);
    } else {
        bar("nope", _x);
    }
}

var _x = 123;

for var _i = 0; _i < 10; _i = _i+1; {
    _x = _x+1;
    _foo = "Foo";
    _bar = "Bar";
    fun(_foo, _bar);
}

switch _x {
    case 1:
    case 2:
        somefunc(2);
    default:
        somefunc(3);
        _x = -1;
}

func myFunc(_x, _y) {
    if _x < _y {
        myFunc(_y, _x);
        hint("Rekursiv");
    }
}

if _a < _b {
    if _b < _a {
        myFunc(_a, _b+9-(2/2));
        myFunc(_a, _b);
        myFunc(_a, _b);
    }
}

myFunc(_a, _b);
myFunc(_a, _b);
