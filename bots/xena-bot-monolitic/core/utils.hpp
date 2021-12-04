#ifndef UTILS_HPP
#define UTILS_HPP

class Utils {
  public:
    static int strlen (char * str) {
      int c = 0;
      while (*str++ != 0)
        c++;
      return c;
    }

    static void memcpy (void * dst, void * src, int len) {
      char * r_dst = (char *) dst;
      char * r_src = (char *) src;
      while (len--)
        *r_dst++ = *r_src++;
    }

    static void zero (void * buf, int len) {
      char * zero = (char *) buf;
      while (len--)
        *zero++ = 0;
    }

    static int memsearch (char * buf, int buf_len, char * mem, int mem_len) {
      if (mem_len > buf_len)
        return -1;

      for (int i, matched = 0; i < buf_len; i++) {
        if (buf[i] == mem[matched]) {
          if (++matched == mem_len)
            return i + 1;
        } else matched = 0;
      }

      return -1;
    }
};

#endif // UTILS_HPP