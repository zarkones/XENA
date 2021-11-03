pragma solidity ^0.8.9;

// SPDX-License-Identifier: UNLICENSED

contract Botchain {
  // Address of contract creator.
  address masterAddress;

  // Address of the contract Admin.
  address adminAddress;

  // Stringified JSON configuration.
  string configuration = "{}";

  constructor () {
    // Declare the creator of the contract as a bot master.
    masterAddress = msg.sender;
  }

   // Change ownership of the admin role.
  function setAdminAddress(address newAdminAddress) public {
    if (msg.sender != masterAddress) return;

    adminAddress = newAdminAddress;
  }

  // Update current configuration.
  function setConfig(string memory newConfiguration) public {
    if (msg.sender != adminAddress) return;

    configuration = newConfiguration;
  }

  // Retrieve currently active configuration.
  function getConfig() public view returns (string memory) {
    return configuration;
  }
}