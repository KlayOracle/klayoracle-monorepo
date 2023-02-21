const { expect, assert } = require('chai');
const {
    arrayify,
    randomBytes,
    hexlify,
    hexZeroPad,
    hashMessage
} = require('ethers/lib/utils');
const { ethers } = require('hardhat');

before(async function () {
    const [deployer, node] = await ethers.getSigners();
    const adapterId = randomBytes(32);

    const OracleProviderSample = await ethers.getContractFactory(
        'OracleProviderSample'
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
});

describe('KlayOracle', function () {
    it('should verify verison is v1.0.0', async function () {
        assert.equal(await this.oracleProviderSample.VERSION(), 'v1.0.0');
    });

    it('should confirm defaults', async function () {
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

    it('node should update latest answer with correct signature', async function () {
        const latestAnswer = hexZeroPad(arrayify(100), 32);
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

        assert.equal(round.answer, latestAnswer);
        assert.equal(round.roundTime, roundTime);
        assert.equal(round.timestamp, blockTimestamp);
    });

    it('node should not update latest answer with incorrect signature', async function () {
        const latestAnswer = hexZeroPad(arrayify(100), 32);
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
        ).to.be.revertedWith('Oracle: Invalid data');
    });

    it('wrong node should not update latest answer with correct signature', async function () {
        const latestAnswer = hexZeroPad(arrayify(100), 32);
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
        ).to.be.revertedWith('Oracle: node unknown');
    });
});
