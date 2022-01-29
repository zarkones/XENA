#ifndef FILE_SYSTEM_HPP
#define FILE_SYSTEM_HPP

#include <stdio.h>
#include <dirent.h>
#include <string>
#include <vector>
#include <stdlib.h>

class FileSystem {
  public:
    static std::vector<std::string> ls (std::string path) {
      struct dirent * entry = nullptr;
      DIR * directory = nullptr;
      std::vector<std::string> result;

      directory = opendir(path.c_str());

      if (directory != nullptr) {
        while ((entry = readdir(directory)))
          result.push_back(entry->d_name);
      }

      closedir(directory);

      return result;
    }

    static std::string process (std::string command) {
      std::string result = "";

      // Start a new process and open a pipe.
      FILE* pipe = popen(command.c_str(), "r");
      #if defined(TALK)
      if (!pipe) {
        std::cout << "Failed to open pipe to file." << std::endl;
      }
      #endif

      char buffer[512];

      // Wait for the process to finish running.
      while (!feof(pipe)) {
        // Read the output bit by bit.
        if (fgets(buffer, 512, pipe) != NULL)
          result += buffer;
      }

      // Exit the process.
      pclose(pipe);

      // Return the output of the process.
      return result;
    }
};

#endif // FILE_SYSTEM_HPP
