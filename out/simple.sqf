waitUntil {x=x+1;x<100};
if (timeIsOver) then {
if (true) exitWith {
[] call foo;
[] call bar;
};
};
