require('dotenv').config();

var HDWalletProvider = require("truffle-hdwallet-provider");

module.exports = {
  networks: {
    // Ganache GUI
    development: {
      host: "localhost",
      port: 7545,
      network_id: "*"
    },
    // Test Network
    rinkeby: {
      provider: function() {
        return new HDWalletProvider(process.env.MNENOMIC, "https://rinkeby.infura.io/" + process.env.INFURA_API_KEY, 0) // <-- account index, zero based
      },
      network_id: 4
    }
  }
};
