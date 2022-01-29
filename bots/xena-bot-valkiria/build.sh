# C++ (works, tested on Debian)

# - Stage.

# With traces.
# haxe -m Main -cpp build/cpp -D no-debug -D linux -D staticStdLibs

# - Production.

# It does not produce any outpout.
# Bare in mind that the Haxe compiler needs to be configured in order to statically link binaries.
haxe -m Main -cpp build/cpp -D no-debug -D linux -D no_traces -D staticStdLibs



# Python (works, tested on Debian)

# - Stage.

# With traces.
# haxe -m Main -python build/python.py -D no-debug -D linux

# - Production.

# It does not produce any outpout.
haxe -m Main -python build/python.py -D no_traces -D no-debug -D linux



# C# (works, tested on Debian)

# Requires the Mono run-time.
# It produces .exe binary, but it runs on Linux. Probably because of the Mono run-time.
# Probably with slight modifications this can drop binaries from Windows into the Windows Subsystem for Linux.

# - Stage.

# With traces.
# haxe -m Main -cs build/cs -D no-debug -D linux

# - Production.

# It does not produce any outpout.
haxe -m Main -cs build/cs -D no-debug -D no_traces  -D linux



# PHP (not tested)
# haxe -m Main -php build/php -D no-debug -D linux
