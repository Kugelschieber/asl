/*
if (!isServer && player != player) then { // wait until player is on the server
    waitUntil {player == player};
};

_getunit = _this select 0;
_file = _this select 1;

if (!isNil _getunit) then { // unit exists?
    _unit = objNull;
    call compile format ["_unit = %1", _getunit];
    
    if (local _unit) then { // test if local (player character)
        try {
            [_unit] execVM _file;
        }
        catch {
            // do nothing
        }
    }
};
*/

if !isServer && player != player {
    exit()(); // does not work for SQF, need to implement exitWith...
}

var _getunit = select(_this)(0);
var _file = select(_this)(1);

if !(isNil()(_getunit)) {
    var _unit = objNull;
    call()(compile()(format()("_unit = %1", _getunit)));
    
    if local()(_unit) {
        // try...catch
        execVM(_unit)(_file);
    }
}
