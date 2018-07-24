var abc = artifacts.require("./ABCtoken.sol");

module.exports = function(deployer) {
  deployer.deploy(abc);
};
