package cfg

import (
	"github.com/urfave/cli"
)

// Vars hold the vars
type Vars struct {
	PostgresURL         string
	ETH_NodeURL         string
	ETH_PrivateKey      string
	ETH_ContractAddress string
	C3_NodeURL          string
	ImageHash           string
	Gensis              bool
}

func (v *Vars) validate() error {
	//if v.PostgresURL == "" {
	//return errors.New("postgres url is required")
	//}
	//if v.ETH_NodeURL == "" {
	//return errors.New("ETH node URL is required")
	//}
	//if v.ETH_PrivateKey == "" {
	//return errors.New("ETH private key is required")
	//}
	//if v.ETH_ContractAddress == "" {
	//return errors.New("ETH contract address is required")
	//}

	return nil
}

// Constants hold the vars
type Constants struct {
	vars *Vars
}

// Get returns the constants
func (c Constants) Get() Vars {
	return *c.vars
}

// New gathers command line args and env vars
func New(args []string) (*Constants, error) {
	vars := &Vars{}
	constants := Constants{
		vars: vars,
	}

	app := cli.NewApp()
	app.Action = func(ctx *cli.Context) error {
		vars = &Vars{
			PostgresURL:         ctx.GlobalString("postgres-url"),
			ETH_PrivateKey:      ctx.GlobalString("eth-private-key"),
			ETH_NodeURL:         ctx.GlobalString("eth-node-url"),
			ETH_ContractAddress: ctx.GlobalString("eth-contract-address"),
			C3_NodeURL:          ctx.GlobalString("c3-node-url"),
			ImageHash:           ctx.GlobalString("image-hash"),
			Genesis:             ctx.GlobalBool("genesis"),
		}
		constants.vars = vars
		return vars.validate()
	}

	app.Authors = []cli.Author{
		{
			Name: "C3 Labs <hello@c3labs.io>",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "postgres-url, p",
			Usage:  "The url of the postgres database",
			Value:  "postgres://postgres:@localhost:5432/db?sslmode=disable",
			EnvVar: "POSTGRES_URL",
		},
		cli.StringFlag{
			Name:   "eth-private-key",
			Usage:  "Ethereum private key",
			Value:  "",
			EnvVar: "ETH_PRIVATE_KEY",
		},
		cli.StringFlag{
			Name:   "eth-node-url",
			Usage:  "URL to the ETH node",
			Value:  "",
			EnvVar: "ETH_NODE_URL",
		},
		cli.StringFlag{
			Name:   "eth-contract-address",
			Usage:  "The ethereum contract address that will hold the tokens",
			Value:  "",
			EnvVar: "ETH_CONTRACT_ADDRESS",
		},
		cli.StringFlag{
			Name:   "c3-node-url",
			Usage:  "URL to the C3 node",
			Value:  "",
			EnvVar: "C3_NODE_URL",
		},
		cli.StringFlag{
			Name:   "image-hash",
			Usage:  "The C3 image hash",
			Value:  "",
			EnvVar: "C3_IMAGE_HASH",
		},
		cli.BoolFlag{
			Name:   "genesis",
			Usage:  "Send a blank genesis block?",
			Value:  false,
			EnvVar: "GENESIS",
		},
	}
	app.Name = "C3 DEX"

	app.Version = "0.0.1"
	return &constants, app.Run(args)
}
