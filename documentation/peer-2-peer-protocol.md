# Network Modes

Based on setup, botnets of different owners should be able to share the P2P layer, the aim being improving resiliency.

# Peer Behavior

Bot should create a loop in which at random intervals it reaches out to random peer from its list. So that it can check if the peer is online and are there any new messages.

If a peer has not been online for 30 days and was pinged 6 times it should be removed from the peer list.

If a peer has less than 25 peers it should reach out to other peers and ask for a peer list. The responer peer should not send more than 15 peers.

# Peer Ranking

A bot should hold about 1000 peers ranked by:

1. How many valid messages they exchanged.
2. Bots which share a message that was previously unseen and is valid get ranked better.

# Blacklisting

If a bot issues a message with a invalid signature it should be blacklisted.

# Peer Message Specification

{
  // If true, the responder should send a peer list if it has it populated.
  needPeers: bool

  // Slice of messages from bot herder which are supposed to be executed.
  messages: []Message
}