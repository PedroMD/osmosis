package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/figment-networks/osmosis/v8/x/pool-incentives/client/cli"
	"github.com/figment-networks/osmosis/v8/x/pool-incentives/client/rest"
)

var UpdatePoolIncentivesHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpdatePoolIncentivesProposal, rest.ProposalUpdatePoolIncentivesRESTHandler)
