// Package environment get environment & app config, all the public field must after init()
// finished and flag.Parse().
package environment

import (
	"flag"
	"os"
	"strconv"
	"time"
)

// deploy environment.
const (
	DeployEnvDev  = "dev"
	DeployEnvFat  = "fat"
	DeployEnvUat  = "uat"
	DeployEnvPre  = "pre"
	DeployEnvProd = "prod"
)

// environment default value.
const (
	// environment
	_region    = "region01"
	_zone      = "zone01"
	_deployEnv = "dev"
)

// environment configuration.
var (
	// Region available region where app at.
	Region string
	// Zone available zone where app at.
	Zone string
	// Hostname machine hostname.
	Hostname string
	// DeployEnv deploy environment where app at.
	DeployEnv string
	// AppID is global unique application id, register by service tree.
	// such as main.arch.disocvery.
	AppID string
	// Color is the identification of different experimental group in one caster cluster.
	Color string
	// DiscoveryNodes is seed nodes.
	DiscoveryNodes string
	// ConfigType [apollo, nacos, files]
	ConfDriverType string
)

func Setup() {
	var err error
	Hostname = os.Getenv("HOSTNAME")
	if Hostname == "" {
		Hostname, err = os.Hostname()
		if err != nil {
			Hostname = strconv.Itoa(int(time.Now().UnixNano()))
		}
	}
	addFlag(flag.CommandLine)
}

func addFlag(fs *flag.FlagSet) {
	// environment
	fs.StringVar(&Region, "region", defaultString("REGION", _region), "available region. or use REGION environment variable, value: sh etc.")
	fs.StringVar(&Zone, "zone", defaultString("ZONE", _zone), "available zone. or use ZONE environment variable, value: sh001/sh002 etc.")
	fs.StringVar(&AppID, "appid", os.Getenv("APP_ID"), "appid is global unique application id, register by service tree. or use APP_ID environment variable.")
	fs.StringVar(&DeployEnv, "deploy.environment", defaultString("DEPLOY_ENV", _deployEnv), "deploy environment. or use DEPLOY_ENV environment variable, value: dev/fat1/uat/pre/prod etc.")
	fs.StringVar(&Color, "deploy.color", os.Getenv("DEPLOY_COLOR"), "deploy.color is the identification of different experimental group.")
	fs.StringVar(&DiscoveryNodes, "discovery.nodes", os.Getenv("DISCOVERY_NODES"), "discovery.nodes is seed nodes. value: 127.0.0.1:7171,127.0.0.2:7171 etc.")
	fs.StringVar(&ConfDriverType, "config.driver.type", os.Getenv("CONFIG_DRIVER_TYPE"), "config driver type. value: apollo,files,nacos,consul etc.")
}

func defaultString(env, value string) string {
	v := os.Getenv(env)
	if v == "" {
		return value
	}
	return v
}
