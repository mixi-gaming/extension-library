package main

import (
	"fmt"

	"github.com/mixi-gaming/extension-library/core/vda"
	"github.com/mixi-gaming/extension-library/transport"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	transport.Constructor(
		viper.GetString("nats.url"),
		viper.GetString("domain"),
	)

	var resp map[string]interface{}

	// VDD test
	// vddApikey := viper.GetString("api_key.vdd")
	// resp = vdd.NatsRetrieveAllDataInBucket(vddApikey, "location_site", "")

	// VDD test
	vdpApikey := viper.GetString("api_key.vdp")
	resp = vda.GetAllRolesInDomain(vdpApikey)
	fmt.Println(resp)
}
