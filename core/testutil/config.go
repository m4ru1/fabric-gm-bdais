/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package testutil

import (
	"flag"
	"fmt"
<<<<<<< HEAD
	"io/ioutil"
	"strings"
=======
	"strings"
	"testing"
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19

	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/core/config/configtest"
	"github.com/hyperledger/fabric/msp"
	"github.com/spf13/viper"
)

// SetupTestConfig setup the config during test execution
<<<<<<< HEAD
func SetupTestConfig() {
=======
func SetupTestConfig(t *testing.T) {
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	flag.Parse()

	// Now set the configuration file
	viper.SetEnvPrefix("CORE")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName("core") // name of config file (without extension)
	configtest.AddDevConfigPath(nil)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Init the BCCSP
	var bccspConfig *factory.FactoryOpts
	err = viper.UnmarshalKey("peer.BCCSP", &bccspConfig)
	if err != nil {
		bccspConfig = nil
	}

<<<<<<< HEAD
	tmpKeyStore, err := ioutil.TempDir("/tmp", "msp-keystore")
	if err != nil {
		panic(fmt.Errorf("Could not create temporary directory: %s\n", tmpKeyStore))
	}
=======
	tmpKeyStore := t.TempDir()
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19

	msp.SetupBCCSPKeystoreConfig(bccspConfig, tmpKeyStore)

	err = factory.InitFactories(bccspConfig)
	if err != nil {
		panic(fmt.Errorf("Could not initialize BCCSP Factories [%s]", err))
	}
}
