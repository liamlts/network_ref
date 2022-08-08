#include <stdio.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <unistd.h>

int main(void) {
    char server_message[] = "This is the server message.\n";

    int server_socket = socket(AF_INET, SOCK_STREAM, 0); 

    struct sockaddr_in server_address;
    server_address.sin_family = AF_INET; 
    server_address.sin_port = htons(7777); 
    server_address.sin_addr.s_addr = INADDR_ANY; 

    bind(server_socket, (struct sockaddr *) &server_address, sizeof(server_address)); 
    listen(server_socket, 6); 
    int client_socket = accept(server_socket, NULL, NULL); 
    send(client_socket, server_message, sizeof(server_message), 0); 
    close(client_socket);
    return 0; 
}