import os

if __name__ == '__main__':
  # Remove data from previous compilation.
  os.system('rm build.zip')
  os.system('rm app.go')
  
  # Boundle the python project into a self executing binary.
  os.system('zip -r  build . --exclude @exclude.lst')
  os.system('''echo '#!/usr/bin/env python3' | cat - build.zip > app''')
  
  # Clean up build leftovers.
  os.system('rm build.zip')
