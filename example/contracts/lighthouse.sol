pragma solidity ^0.4.24;

// Searcher is an interface for contracts that want to be notified of incoming data
//
contract Searcher {

    // poke is called when new data arrives
    //
    function poke() public;

    // this is called to ensure that only valid Searchers can be added to the Lighthouse - returns an arbitrarily chosen number
    //
    function identify() external pure returns(uint) {
        return 0xda4b055; 
    }
}

// for operation of this contract see the readme file.
//
contract Lighthouse {

    uint value;                       // holds all the data bit fiddled into a single 32 byte word.

    uint maxAge;                      // if non zero, sets a limit to data validity

    function setMaxAge(uint newMaxAge) public  {
        maxAge = newMaxAge;
    }
    
    function notTooLongSinceUpdated() public view returns (bool) {
        uint since = now - ((value >> 128) & 
        0x000000000000000000000000000000000000000000000000ffffffffffffffff);
        return (since < maxAge) || (maxAge == 0);
    }
    
    function peekData() external view returns (uint128 v,bool b) {
        v = uint128(value);
        b = notTooLongSinceUpdated() && value != 0;
        return;
    }
    
    function peekUpdated()  external view returns (uint32 v,bool b) {
        uint v2 = value >> 128;
        v = uint32(v2);
        b = notTooLongSinceUpdated() && value != 0;
        return;
    }
    


    function peek() external view returns (bytes32 v ,bool ok) {
        v = bytes32(value & 0x00000000000000000000000000000000ffffffffffffffffffffffffffffffff);
        ok = notTooLongSinceUpdated() && value != 0;
        return;
    }
    
    function read() external view returns (bytes32 x) {
        require(notTooLongSinceUpdated() && value != 0, "Invalid data stored");
        return bytes32(value & 0x00000000000000000000000000000000ffffffffffffffffffffffffffffffff);
    }
    
    function write(uint  DataValue) external  {
        require ((DataValue >> 128) == 0, "Value too large");
        
        value = DataValue  + (now << 128) ;
     }
}