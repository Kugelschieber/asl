if (!isServer&&player!=player) then {
waitUntil {player==player};
};
_getunit = (_this select 0);
_file = (_this select 1);
if (!(isNil _getunit)) then {
_unit = objNull;
call (compile (format ["_unit = %1", _getunit]));
if ((local _unit)) then {
try {
_unit execVM _file;
} catch {
};
};
};
