rm -r components
mkdir components

# Build Xena-Service-Atila.
sudo yarn --cwd ../../services/xena-service-atila/ build
sudo mv ../../services/xena-service-atila/build/build.js components/atila.js
echo "XENA_SERVICE_ATILA has been built."

# Build Xena-Bot-Apep.
cd ../xena-bot-apep && go build -o ../xena-bot-genesis/components
cd ../xena-bot-genesis
mv components/xena-apep components/apep
echo "XENA_SERVICE_APEP has been built."

# Grab the "Botchain".
cp ../../contracts/botchain.sol components/chain.sol
echo "XENA_CONTRACT_BOTCHAIN recorded."

# Build the bot itself.
go build
strip main
echo "XENA_BOT_GENESIS has been built."
