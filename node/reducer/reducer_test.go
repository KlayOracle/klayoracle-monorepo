package reducer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var jsons = []string{
	`{
  "jsonrpc": "2.0",
  "result": [
    [
      1679338800000,
      71.82608036579228
    ],
    [
      1679338800000,
      71.82608036579228
    ]
  ]
}`,
	`
	{
		"result" : 123
	}`,
	`
	{"RAW":{"KLAY":{"USD":{"TYPE":"5","MARKET":"CCCAGG","FROMSYMBOL":"KLAY","TOSYMBOL":"USD","FLAGS":"1026","PRICE":0.2336,"LASTUPDATE":1680695267,"MEDIAN":0.2336,"LASTVOLUME":2445,"LASTVOLUMETO":571.152,"LASTTRADEID":"13831285","VOLUMEDAY":32608716.966410466,"VOLUMEDAYTO":7617396.283353485,"VOLUME24HOUR":41816082.965217054,"VOLUME24HOURTO":9768236.980674705,"OPENDAY":0.2222,"HIGHDAY":0.2397,"LOWDAY":0.221,"OPEN24HOUR":0.22,"HIGH24HOUR":0.2406804,"LOW24HOUR":0.21755648,"LASTMARKET":"Binance","VOLUMEHOUR":3695860.7178783054,"VOLUMEHOURTO":863353.0636963721,"OPENHOUR":0.235,"HIGHHOUR":0.2351,"LOWHOUR":0.2326,"TOPTIERVOLUME24HOUR":41816082.965217054,"TOPTIERVOLUME24HOURTO":9768236.980674705,"CHANGE24HOUR":0.013600000000000001,"CHANGEPCT24HOUR":6.181818181818182,"CHANGEDAY":0.011399999999999993,"CHANGEPCTDAY":5.130513051305127,"CHANGEHOUR":-0.0013999999999999846,"CHANGEPCTHOUR":-0.5957446808510573,"CONVERSIONTYPE":"multiply","CONVERSIONSYMBOL":"USDT","CONVERSIONLASTUPDATE":1680695271,"SUPPLY":11023323494.7529,"MKTCAP":2575048368.374277,"MKTCAPPENALTY":0,"CIRCULATINGSUPPLY":3083014688.15289,"CIRCULATINGSUPPLYMKTCAP":720192231.1525152,"TOTALVOLUME24H":70181684.79203492,"TOTALVOLUME24HTO":16394441.56741936,"TOTALTOPTIERVOLUME24H":70181684.79203492,"TOTALTOPTIERVOLUME24HTO":16394441.56741936,"IMAGEURL":"/media/40484639/klay.png"}}},"DISPLAY":{"KLAY":{"USD":{"FROMSYMBOL":"KLAY","TOSYMBOL":"$","MARKET":"CryptoCompare Index","PRICE":"$ 0.2336","LASTUPDATE":"Just now","LASTVOLUME":"KLAY 2,445.00","LASTVOLUMETO":"$ 571.15","LASTTRADEID":"13831285","VOLUMEDAY":"KLAY 32,608,717.0","VOLUMEDAYTO":"$ 7,617,396.3","VOLUME24HOUR":"KLAY 41,816,083.0","VOLUME24HOURTO":"$ 9,768,237.0","OPENDAY":"$ 0.2222","HIGHDAY":"$ 0.2397","LOWDAY":"$ 0.2210","OPEN24HOUR":"$ 0.2200","HIGH24HOUR":"$ 0.2407","LOW24HOUR":"$ 0.2176","LASTMARKET":"Binance","VOLUMEHOUR":"KLAY 3,695,860.7","VOLUMEHOURTO":"$ 863,353.1","OPENHOUR":"$ 0.2350","HIGHHOUR":"$ 0.2351","LOWHOUR":"$ 0.2326","TOPTIERVOLUME24HOUR":"KLAY 41,816,083.0","TOPTIERVOLUME24HOURTO":"$ 9,768,237.0","CHANGE24HOUR":"$ 0.014","CHANGEPCT24HOUR":"6.18","CHANGEDAY":"$ 0.011","CHANGEPCTDAY":"5.13","CHANGEHOUR":"$ -0.0014","CHANGEPCTHOUR":"-0.60","CONVERSIONTYPE":"multiply","CONVERSIONSYMBOL":"USDT","CONVERSIONLASTUPDATE":"Just now","SUPPLY":"KLAY 11,023,323,494.8","MKTCAP":"$ 2.58 B","MKTCAPPENALTY":"0 %","CIRCULATINGSUPPLY":"KLAY 3,083,014,688.2","CIRCULATINGSUPPLYMKTCAP":"$ 720.19 M","TOTALVOLUME24H":"KLAY 70.18 M","TOTALVOLUME24HTO":"$ 16.39 M","TOTALTOPTIERVOLUME24H":"KLAY 70.18 M","TOTALTOPTIERVOLUME24HTO":"$ 16.39 M","IMAGEURL":"/media/40484639/klay.png"}}}}
	`,
}

var reducers = []string{"$.result[-1:][1]", "$.result", "$.RAW.KLAY.USD.PRICE"}

var results = []float64{
	71.82608036579228,
	123,
	0.2336,
}

func TestPARSE(t *testing.T) {

	for index, jsonResponse := range jsons {

		result, err := PARSE(jsonResponse, reducers[index])

		if err != nil {
			assert.Fail(t, err.Error())
		}

		assert.Equal(t, result, results[index])
	}
}
