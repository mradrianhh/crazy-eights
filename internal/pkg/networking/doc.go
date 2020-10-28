package network

/*
Communication protocol.

Detailed layout of the communication protocol between client and server.

When using the term "requests" it is implied that the client or server sends a message of messagetype "request".
Similarly, when using the term "responds", it is implied that the client or server sends a message of messagetype "response".
One request followed by one response is one dialogue. You can not send only a request, or only a response. The dialogue must be complete.

Basic actions.

Client:
	- Give server username.
	- Draw a card.
	- Place a card.
	- Pass.

Server:
	- Inform clients know whose turn it is.
	- Inform client if placed card is legal.
	- Inform client of his current hand.
	- Inform clients that game terminates.

Communication flow.

Server launches and listens.
Client connects.
Server accepts connection, and requests username.
Client receives request, and confirms that it received the request by sending a response. It reads the requesttype, and acts accordingly by asking the user to enter a username.
Client sends a request with the username, asking for validation.
Server receives request, analyses the username. If appropriate, it adds the user to the session and responds with a VALID-response.
	If inappropriate, it does not add the user to the session, and responds with a INVALID-response.
Client receives response. If valid, the client-interface updates to the in-game interface. If invalid, the user is returned to the main-menu(main program).

Assuming user provided valid username and has connected to session.

Client requests the server to deal a hand for the user. It sends a request-message of requesttype DEAL along with the username.
Server receives request, reads the username, looks up the player-object associated with the provided username and deals a hand to it. The server confirms that it received the request.

Each turn, server sends a request to each player with their hand and the current top card.
Client receives request, and responds with a confirmation.
If the current player is the clients player, it updates its user interface so the user can choose an action.

Messagetypes:
	- Request. When receiving a request, the client or server must send a response.
	- Response. When receiving a response, it marks the end of the dialogue.

Requesttypes:
	- Provide username.
	- Username validation.
	- Deal user a hand.
	- Informational. I.e sending each client their current hand and the top card.

Responsetypes:
	- Confirmation.
	- Valid.
	- Invalid.
*/
