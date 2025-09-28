package maintainer

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/mapprotocol/atlas/contracts"
	"github.com/mapprotocol/atlas/contracts/abis"
	"github.com/mapprotocol/atlas/core/vm"
	"github.com/mapprotocol/atlas/params"
)

var (
	orchestrateMethod = contracts.NewRegisteredContractMethod(params.MaintainerId, abis.Maintainer, "orchestrate", params.MaxGasForOrchestrate)
)

func Orchestrate(vmRunner vm.EVMRunner) error {
	return orchestrateMethod.Execute(vmRunner, nil, common.Big0)
}
