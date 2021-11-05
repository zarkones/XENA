rm -r components
mkdir components

# Build Xena-Service-Atila.
echo "Building XENA_SERVICE_ATILA."
sudo yarn --cwd ../../services/xena-service-atila/ build
echo "Done building XENA_SERVICE_ATILA."
sudo mv ../../services/xena-service-atila/build/build.js components/atila.js
echo "Imported XENA_SERVICE_ATILA."

# Build Xena-Bot-Apep.
echo "Building XENA_SERVICE_APEP."
cd ../xena-bot-apep && go build -o ../xena-bot-genesis/components
echo "Done building XENA_SERVICE_APEP."
cd ../xena-bot-genesis
mv components/xena-apep components/apep
strip components/apep
echo "Imported XENA_SERVICE_APEP."

# Grab the "Botchain".
echo "Building XENA_CONTRACT_BOTCHAIN."
cp ../../contracts/xena-contract-botchain.sol components/chain.sol
solcjs -o . --bin --abi components/xena-contract-botchain.sol
echo "Imported XENA_CONTRACT_BOTCHAIN recorded."

# Build the bot itself.
echo "Building XENA_BOT_GENESIS."
go build
strip main
echo "Done building XENA_BOT_GENESIS."
