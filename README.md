# NetRogue
A multiplayer rogue-like in the spirit of TOAG, but implemented in termbox-go.


## Concept
Every instance should be capable of hosting as a server (thus also broadcasting for viewing) and acting as a respective client to connect to said server. The choice of mode or interaction should be performed through a definitive main-menu which is displayed prior to the beginning of the game.

For the sake of compact-ness and efficiency all maps should be manually encoded through the encMaps utility to export/encode all of the textually-composed maps as arrays of sprites (with provided colors and descriptions) capable of being loaded into the primary server/client upon start.

Due to interactions of clients events within the game will automatically take place after a pre-determined n (30?) second timer in which all clients/server will be synced regardless of action. Order of operations of actions should be determined by the (millisecond) difference between actions as performed by players in real-time.

All sprites should have a foreground and background color encoded into them, when a player moves over an item which is traversable, the foreground color of the "covered" item should become the background color of the player.




## Usage
Run: `go generate` to build all binaries and place them in the current working directory.
All binaries should be executed from the base NetRogue/ directory.
