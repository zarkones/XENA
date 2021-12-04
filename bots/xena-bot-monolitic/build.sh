rm -r build
mkdir build
g++ -static -DNDEBUG -g0 -Os main.cpp -o build/main_static
g++ -DNDEBUG -g0 -Os main.cpp -o build/main_dynamic
strip build/main_static
strip build/main_dynamic

