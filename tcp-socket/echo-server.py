import socket

#ECHO SERVER

HOST = "127.0.0.1" #standard loopback interface address(localhost)
PORT = 65432    #port to listen on (non-privileged ports are > 1023)

#socket.socket creates object with context manager and so need to s.close()
#arguments are constants: AF_INET - Internet Address family for IPv4, SOCK_STREAM - TCP
with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    #associates socket with specific network interface and port number. accepts two tuple as IPv4
    s.bind((HOST,PORT))
    #.listen() enables server to accept connections
    s.listen()
    conn, addr = s.accept() #blicks execution and waiting for new connection
    with conn:
        print(f"Connected by {addr}")
        while True:
            data  = conn.recv(1024) # if sends empty bytes data that indicates client closed and loop terminated
            if not data:
                break
            conn.sendall(data)
            
