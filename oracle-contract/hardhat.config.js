require('@nomicfoundation/hardhat-toolbox');
require('dotenv').config({ path: '.env' });

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
    networks: {
        baobab: {
            url: 'https://api.baobab.klaytn.net:8651',
            accounts: [process.env.PRIVATE_KEY],
            chainId: 1001
        }
    },
    solidity: {
        version: '0.8.16',
        settings: {
            optimizer: {
                enabled: true,
                runs: 1000
            }
        }
    }
};
