const { expect, assert } = require("chai");
const {
  arrayify,
  randomBytes,
  hexlify,
  hexZeroPad,
  hashMessage,
  parseEther,
} = require("ethers/lib/utils");
const { ethers } = require("hardhat");

before(async function () {
  const [deployer, node] = await ethers.getSigners();
  const adapterId = randomBytes(32);

  const OracleProviderSample = await ethers.getContractFactory(
    "OracleProviderSample"
  );
  const oracleProviderSample = await OracleProviderSample.deploy(
    node.address,
    adapterId
  );
  await oracleProviderSample.deployed();

  this.oracleProviderSample = oracleProviderSample;
  this.node = node;
  this.deployer = deployer;
  this.adapterId = adapterId;

  const OracleConsumerSample = await ethers.getContractFactory(
    "OracleConsumerSample"
  );
  const oracleConsumerSample = await OracleConsumerSample.deploy(
    oracleProviderSample.address
  );
  await oracleConsumerSample.deployed();

  this.oracleConsumerSample = oracleConsumerSample;
});

describe("KlayOracle", function () {
  it("should verify verison is v1.0.0", async function () {
    assert.equal(await this.oracleProviderSample.VERSION(), "v1.0.0");
  });

  it("should confirm defaults", async function () {
    assert.equal(
      await this.oracleProviderSample.nodeAddress(),
      this.node.address
    );
    assert.equal(await this.oracleProviderSample.fulfilledCount(), 0);
    assert.equal(
      await this.oracleProviderSample.adapterId(),
      hexlify(this.adapterId)
    );
  });

  it("node should update latest answer with correct signature", async function () {
    const latestAnswer = hexZeroPad(parseEther("100").toHexString(), 32);
    const roundTime = 3000;

    const signature = await this.node.signMessage(latestAnswer);

    await this.oracleProviderSample
      .connect(this.node)
      .newRoundData(
        roundTime,
        latestAnswer,
        hashMessage(latestAnswer),
        signature
      );

    const round = await this.oracleProviderSample.latestRound();
    const blockTimestamp = (await ethers.provider.getBlock()).timestamp;

    assert.equal(round.answer, 100e18);
    assert.equal(round.roundTime, roundTime);
    assert.equal(round.timestamp, blockTimestamp);
  });

  it("node should not update latest answer with incorrect signature", async function () {
    const latestAnswer = hexZeroPad(parseEther("100").toHexString(), 32);
    const roundTime = 3000;

    const signature = await this.deployer.signMessage(latestAnswer);

    await expect(
      this.oracleProviderSample
        .connect(this.node)
        .newRoundData(
          roundTime,
          latestAnswer,
          hashMessage(latestAnswer),
          signature
        )
    ).to.be.revertedWith("Oracle: Invalid data");
  });

  it("wrong node should not update latest answer with correct signature", async function () {
    const latestAnswer = hexZeroPad(parseEther("100").toHexString(), 32);
    const roundTime = 3000;

    const signature = await this.node.signMessage(latestAnswer);

    await expect(
      this.oracleProviderSample
        .connect(this.deployer)
        .newRoundData(
          roundTime,
          latestAnswer,
          hashMessage(latestAnswer),
          signature
        )
    ).to.be.revertedWith("Oracle: node unknown");
  });

  it("should return latest answer to consumer", async function () {
    await this.oracleConsumerSample.swapEthtoKlay();

    const latestAnswer = await this.oracleConsumerSample.klayOutput();

    assert.equal(latestAnswer, 100e18);
  });

  it("node should update latest answer to 101.2", async function () {
    const latestAnswer = hexZeroPad(parseEther("100.2").toHexString(), 32);
    const roundTime = 3000;

    const signature = await this.node.signMessage(latestAnswer);

    await this.oracleProviderSample
      .connect(this.node)
      .newRoundData(
        roundTime,
        latestAnswer,
        hashMessage(latestAnswer),
        signature
      );

    await this.oracleConsumerSample.swapEthtoKlay();

    assert.equal(await this.oracleConsumerSample.klayOutput(), 100.2e18);
  });
});
