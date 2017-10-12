// g++ -o report_packet_stats report_packet_stats.cpp

#include <fstream>
#include <iostream>
#include <string>
#include <sstream>
#include <vector>
#include <errno.h>

#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

#include <inttypes.h>
#include <time.h>
#include <string.h>
#include <unistd.h>
#include <signal.h>

int sock = -1;

void usage(char* argv[])
{
    printf("usage: %s MONITOR_IP_ADDR", argv[0]);
    exit(1);
}

// see http://faithandbrave.hateblo.jp/entry/2014/05/01/171631
std::vector<std::string> split(const std::string& input, char delimiter)
{
    std::istringstream stream(input);

    std::string field;
    std::vector<std::string> result;
    while (std::getline(stream, field, delimiter)) {
        result.push_back(field);
    }
    return result;
}

int find(std::string haystack,std::string niddle)
{
    
    for(int i=0;i<haystack.size();i++){
	//for(int j=0;j<strlen(niddle);j++){
	if(niddle.compare(haystack.substr(i,niddle.size()))==0){
	    return i; 
	}
	//}
    }
    return -1;

}
// see http://www.wabiapp.com/WabiSampleSource/windows/replace_string.html
std::string replaceString(
      std::string String1   // 置き換え対象
    , std::string String2   // 検索対象
    , std::string String3   // 置き換える内容
)
{
    std::string::size_type  Pos( String1.find( String2 ) );
 
    while( Pos != std::string::npos )
    {
        String1.replace( Pos, String2.length(), String3 );
        Pos = String1.find( String2, Pos + String3.length() );
    }
 
    return String1;
}

void connect_socket(char* destaddr, int port_no)
{

    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock < 0) {
        printf("socket failed (errno: %d)", errno);
        exit(1);
    }
    struct sockaddr_in server;
    server.sin_family = AF_INET; 
    server.sin_port = htons(port_no); // etime
    server.sin_addr.s_addr = inet_addr(destaddr);

    connection:
    int res;
    res = connect(sock, (struct sockaddr *)&server, sizeof(server));
    if(res < 0){
        if(errno == ECONNREFUSED){
            fprintf(stderr, "connection refused. retrying...\n");
        }
        else{
            fprintf(stderr, "connect failed (errno: %d). retrying...\n", errno);
        }
        sleep(10);
        goto connection;
    }
    // printf("connection established\n");
}

std::string  getNumTransmittedPackets()
{
    std::ifstream ifs("/proc/net/dev");
    std::string str;
    std::vector<std::string> splitted;
    std::string num_packets_str = "";
    if (ifs.fail())
    {
        std::cerr << "cannot read proc-fs" << std::endl;
        return NULL;
    }
    while (getline(ifs, str))
    {
	std::cout << str << std::endl;
        if(find(str.c_str(),"enp")>=0||find(str.c_str(),"eth")>=0){//if(str.find("enp") >= 0 || str.find("eth") >= 0){ // starts with 'e'
            str = replaceString(str, "       ", " ");
            str = replaceString(str, "      ", " ");
            str = replaceString(str, "     ", " ");
            str = replaceString(str, "    ", " ");
            str = replaceString(str, "   ", " ");
            str = replaceString(str, "  ", " ");
            std::cout << str << std::endl;
            splitted = split(str, ' ');
            //for(int i; i < splitted.size(); i++)
            //     std::cout << splitted[i] << std::endl;
            //num_packets = std::stoi(splitted[10]);
             num_packets_str = splitted[10];
            break;
        }
    }
    return num_packets_str;
}

double getTimeStamp()
{
    struct timespec spec;
    clock_gettime(CLOCK_REALTIME, &spec);
    return (double) spec.tv_sec + (double) spec.tv_nsec / 1.0e9;
}

void close_socket(int param)
{
    close(sock);
    printf("closed socket. exit\n");
    exit(0);
}

int main(int argc, char* argv[])
{
    //printf("main¥n");
    if(argc != 2){
        usage(argv);
    }
    //printf("SIGINT handler¥n");
    // SIGINT handler
    signal(SIGINT, close_socket);
    //printf("corect time¥n");
    // correct time
    connect_socket(argv[1], 37133);
    char buf[128];
    double t1 = getTimeStamp();
    sprintf(buf, "%.6f\n", t1);
    write(sock, buf, strlen(buf));
    read(sock, buf, strlen(buf));
    double t3 = getTimeStamp();
    close(sock); 
    double t2 = std::stod(buf);
    double dn = (t3 - t1) / 2;
    double dt = t2 - t1 / 2 - t3 / 2;
    printf("dt = %.6f\n", dt);

    // open socket 
    connect_socket(argv[1], 31337);
    while(1){    
        std::string num_packets = getNumTransmittedPackets();
        //printf("num transmitted packets: %d\n", num_packets);        
        double timestamp = getTimeStamp() + dt; // fit to Monitor's current time
        //printf("timestamp: %.3f\n", timestamp);

        char buf[128];
	printf("%.6f %s\n", timestamp, num_packets.c_str());
        sprintf(buf, "%.6f %s\n", timestamp, num_packets.c_str());
        write(sock, buf, strlen(buf));
        usleep(1000 * 1000);
    }

    return 0;
}
