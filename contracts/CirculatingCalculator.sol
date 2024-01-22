// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface BulkBalanceCalculator {
    function calcualte(address[] calldata addresses) external view returns (uint256 total);
}

contract CirculatingCalculator {
    function calculate(uint256 supply, address[] calldata addresses) public view returns (uint256 total) {
        BulkBalanceCalculator calculator = BulkBalanceCalculator(0x0000000000E162430e50852F2654052BDEccB9D6);
        uint256 subTotal = calculator.calcualte(addresses);
        return (supply - subTotal) / 1 ether;
    }
}