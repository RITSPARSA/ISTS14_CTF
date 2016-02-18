### == Scan This! ==


##### Scenario
You have a network with a super sloppy Systems Administrator. He's left for the day, and you have news that something bad might be targeting your online presence.
As you frantically search for documentation, which doesn't exist, you're pager starts going off. You need to know what your exposure is, and you need to know now.
You try to SSH into the host. Of course it's not running on port 22, that would be too easy. What do you do now?

##### Mission:
Create a simple portscanner in a programming language of your choice.
Identify what ports are open to the world, and what services are listening behind them.
It should take a hostname as a parameter and vend a report to stdout that describes all the open ports, and if possible their protocols.
Not all ports will have a service listening on them, some may just be there open

Hint:
* Ports above 1025 are fair game.
* Some protocols you might want to scan for:
* SSH
* FTP
* HTTP
