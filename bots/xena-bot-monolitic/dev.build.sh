rm -r build
mkdir build
g++ -static main.cpp -o build/dev_main_static
g++ main.cpp -o build/dev_main_dynamic
strip build/dev_main_static
strip build/dev_main_dynamic

