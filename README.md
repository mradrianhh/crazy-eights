# crazy-eights

Desired implementation:

	Structure:

-crazy-eights/
	-Crazyeights/
		-Program
	-Client/
		-Program
	-Server/
		-Program

	Roles:

Main.

When the user launches the game, it is presented with the "Main" program. In this program, the user can choose to
host or join a game, or quit.

Client.

When the user joins a game, it is through their client connecting to a server. The client acts as the interface with which the user can interact with the "dealer"(server).
It is in the client that the user is presented with their hands, the topcard, and their options.

Server.

The server hosts and controls the game, acting as the dealer. Players connect to the server, verifies themselves to the server, gets dealt through the server,
communicates with the server, and wins through the server. The players do not interact directly with the server, they interact with the clients which in turn
communicates with the server.

	Flow:

The user launches the program and can choose between three options:
	1. Join a game.
	2. Host a game.
	3. Exit.

Join a game.

The user must provide an ip-address for the server, and a username for the server to store them as and identify them by. Multiple users can not share the same username.
Given that the ip-address is correct, they will connect to the server, and given that their username is viable and available, they will be presented with the client-interface.

Host a game.

The program launches a server terminal-window. In the server terminal-window the user are presented with basic server tools and the server ip-address.
After the server is launched, the program auto-launches the client-program. The user only has to provide his username, 
and given that it is viable and available (which it should be because he just started the server), the client will auto-connect to his server.
He can share the ip-address presented to him in the server window with his friends so they can join him.

Exit.

The program terminates.