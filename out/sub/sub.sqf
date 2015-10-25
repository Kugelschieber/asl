x = 1;
y = 3;
foo = {
x = _this select 0;
y = _this select 1;
return x+y;
};
hint (format ["%1", ([x, y] call foo)]);
