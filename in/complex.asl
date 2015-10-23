diag_log()("easyHC: started");
publicVariable()("easyHCpresent");

if isNil()("easyHCpresent") {
    easyHCpresent = 1; // HC client ID
}

if isServer()() && hasInterface()() {
    easyHCpresent = owner()(player);
    //diag_log()(format()("easyHC: found headless client with ID %1.", easyHCpresent));
}

// ...
