// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract BulkBalanceCalculator {
    function calcualte(address[] calldata addresses) public view returns (uint256 total) {
        for(uint8 i = 0; i < addresses.length; i++) {
            total += addresses[i].balance;
        }

        return total;
    }
}