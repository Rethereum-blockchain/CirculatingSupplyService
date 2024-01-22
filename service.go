package main

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"net/http"
)

type Config struct {
	RpcUrl          string
	ContractAddress string
	ApiUrl          string
	Addresses       []common.Address
	Port            string
	Mode            string
}

var (
	abi string = `
[
   {
      "type":"function",
      "stateMutability":"view",
      "outputs":[
         {
            "type":"uint256",
            "name":"total",
            "internalType":"uint256"
         }
      ],
      "name":"calculate",
      "inputs":[
         {
            "type":"uint256",
            "name":"supply",
            "internalType":"uint256"
         },
         {
            "type":"address[]",
            "name":"addresses",
            "internalType":"address[]"
         }
      ]
   }
]`
	conf     Config
	provider *web3.Web3
	contract *eth.Contract
)

func init() {
	_, err := toml.DecodeFile("config.toml", &conf)
	if err != nil {
		log.Fatal("Error decoding config file")
	}
	provider, err = web3.NewWeb3(conf.RpcUrl)
	if err != nil {
		log.Fatal("Error connecting to RPC")
	}
	contract, err = provider.Eth.NewContract(abi, conf.ContractAddress)
	if err != nil {
		log.Fatal("Error initializing contract")
	}
	gin.SetMode(conf.Mode)
}

func main() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the Hypra Supply API")
	})
	server.GET("/circulatingsupply", getCirculatingSupply)
	err := server.Run(":" + conf.Port)
	if err != nil {
		log.Fatal("Error, failed to start server", err)
		return
	}
}

type Total struct {
	Total string `json:"result"`
}

func getCirculatingSupply(c *gin.Context) {
	body, err := http.Get(conf.ApiUrl)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error calculating circulating supply")
		log.Print(err)
		return
	}

	total := Total{}
	err = json.NewDecoder(body.Body).Decode(&total)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error calculating circulating supply")
		log.Print(err)
		return
	}

	bigTotal, _ := new(big.Int).SetString(total.Total, 10)

	circulating, err := contract.Call("calculate", bigTotal, conf.Addresses)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error calculating circulating supply")
		log.Print(err)
		return
	}

	c.JSON(http.StatusOK, circulating.(*big.Int).Uint64())
}
