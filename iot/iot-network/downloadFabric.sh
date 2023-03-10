

# Exit on first error, print all commands.
set -e
# Set ARCH
ARCH=`uname -m`

# Grab the current directory.
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Pull and tag the latest Hyperledger Fabric base image.
docker pull hyperledger/fabric-peer:$ARCH-1.1.0
docker pull hyperledger/fabric-ca:$ARCH-1.1.0
docker pull hyperledger/fabric-ccenv:$ARCH-1.1.0
docker pull hyperledger/fabric-orderer:$ARCH-1.1.0
docker pull hyperledger/fabric-couchdb:$ARCH-0.4.6