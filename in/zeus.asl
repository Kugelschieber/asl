if isServer {
    exitWith();
}

func ZeusGrpPlaced(_group) {
    
}

func ZeusObjPlaced(_unit) {

}

foreach curator => allCurators {
    addCuratorEditableObjects(curator, allUnits, true);
    addCuratorEditableObjects(curator, allMissionObjects("All"), false);
    
    foreach unit => allUnits {
        var _vehicle = vehicle(unit);
    
        if _vehicle != unit {
            addCuratorEditableObjects(curator, _vehicle, true);
        }
    }
    
    addEventHandler(curator, "CuratorGroupPlaced", BIS_fnc_MP());
}
