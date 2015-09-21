if isServer {
    exitWith();
}

func ZeusGrpPlaced(_curators, _group) {
    each allCurators-_curators {
        $addCuratorEditableObjects(_x)(_group, true);
    }
}

func ZeusObjPlaced(_curators, _unit) {
    each allCurators-_curators {
        $addCuratorEditableObjects(_x)([_unit], true);
    }
}

each allCurators {
    $addCuratorEditableObjects(_x)(allUnits, true);
    $addCuratorEditableObjects(_x)(allMissionObjects("All"), false);
    
    _curator = _x;
    
    each allUnits {
        if vehicle(_x) != _x {
            $addCuratorEditableObjects(_curator)([vehicle(_x)], true);
        }
    }
}
