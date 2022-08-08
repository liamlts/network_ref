#include <stdio.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>

int main(void) {
    int network_socket = socket(AF_INET, SOCK_STREAM, 0); 

    struct sockaddr_in server_addr; 
    server_addr.sin_family = AF_INET; 
    server_addr.sin_port = htons(7777); 
    server_addr.sin_addr.s_addr = INADDR_ANY; 

    int connection_staus = connect(network_socket, (struct sockaddr *)&server_addr, sizeof(server_addr));
    if (connection_staus == -1) {
        perror("There was an issue connecting\n"); 
    }
    char server_resp[256]; 
    recv(network_socket, &server_resp, sizeof(server_resp), 0);
    printf("Server send the data: %s", server_resp);
    return 0; 
}