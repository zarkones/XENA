rm -r components
mkdir components

# Build Xena-Service-Atila.
echo "BUILDING XENA_SERVICE_ATILA."
sudo yarn --cwd ../../services/xena-service-atila/ build
echo "DONE BUILDING XENA_SERVICE_ATILA."
sudo mv ../../services/xena-service-atila/build/build.js components/atila.js
echo "IMPORTED XENA_SERVICE_ATILA.\n"

# Build Xena-Service-Pyramid.
echo "BUILDING XENA_SERVICE_PYRAMID."
sudo yarn --cwd ../../services/xena-service-pyramid/ build
echo "DONE BUILDING XENA_SERVICE_PYRAMID."
sudo mv ../../services/xena-service-pyramid/build/build.js components/pyramid.js
echo "IMPORTED XENA_SERVICE_PYRAMID.\n"

# Build Xena-Service-Ra.
echo "BUILDING XENA_SERVICE_RA."
sudo yarn --cwd ../../services/xena-service-ra/ build
echo "DONE BUILDING XENA_SERVICE_RA."
sudo mv ../../services/xena-service-ra/build/build.js components/ra.js
echo "IMPORTED XENA_SERVICE_RA.\n"

# Build Xena-Bot-Apep.
echo "BUILDING XENA_BOT_APEP."
cd ../xena-bot-apep && go build -o ../xena-bot-genesis/components
echo "DONE BUILDING XENA_BOT_APEP."
cd ../xena-bot-genesis
mv components/xena-apep components/apep
strip components/apep
echo "IMPORTED XENA_BOT_APEP.\n"

# Build Xena-Bot-Varvara. (Python3)
echo "BUILDING XENA_BOT_VARVARA_PYTHON."
sh ../xena-bot-varvara/build.sh
echo "DONE BUILDING XENA_BOT_VARVARA_PYTHON."
cp ../xena-bot-varvara/build/python.py components/varvara_python.py
echo "IMPORTED XENA_BOT_VARVARA_PYTHON.\n"

# Build Xena-Bot-Varvara. (DotNet)
echo "BUILDING XENA_BOT_VARVARA_DOTNET."
sh ../xena-bot-varvara/build.sh
echo "DONE BUILDING XENA_BOT_VARVARA_DOTNET."
cp ../xena-bot-varvara/build/cs/bin/Main.exe components/varvara_dotnet.exe
strip components/varvara_dotnet.exe
echo "IMPORTED XENA_BOT_VARVARA_DOTNET.\n"

# Build Xena-Bot-Varvara. (C++)
echo "BUILDING XENA_BOT_VARVARA_CPP."
sh ../xena-bot-varvara/build.sh
echo "DONE BUILDING XENA_BOT_VARVARA_CPP."
cp ../xena-bot-varvara/build/cpp/Main components/varvara_cpp
strip components/varvara_cpp
echo "IMPORTED XENA_BOT_VARVARA_CPP.\n"


# Grab the "Botchain".
echo "BUILDING XENA_CONTRACT_BOTCHAIN."
cp ../../contracts/xena-contract-botchain.sol components/chain.sol
solcjs -o components --bin --abi components/chain.sol
rm components/chain.sol
echo "IMPORTED XENA_CONTRACT_BOTCHAIN.\n"

# Build the bot itself.
echo "BUILDING XENA_BOT_GENESIS."
go build
strip main
echo "DONE BUILDING XENA_BOT_GENESIS."
