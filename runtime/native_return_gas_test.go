package runtime

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/BestSilverTiger/thor/builtin"
	"github.com/BestSilverTiger/thor/muxdb"
	"github.com/BestSilverTiger/thor/state"
	"github.com/BestSilverTiger/thor/thor"
	"github.com/BestSilverTiger/thor/tx"
	"github.com/BestSilverTiger/thor/xenv"
)

func TestNativeCallReturnGas(t *testing.T) {
	db := muxdb.NewMem()
	state := state.New(db, thor.Bytes32{})
	state.SetCode(builtin.Measure.Address, builtin.Measure.RuntimeBytecodes())

	inner, _ := builtin.Measure.ABI.MethodByName("inner")
	innerData, _ := inner.EncodeInput()
	outer, _ := builtin.Measure.ABI.MethodByName("outer")
	outerData, _ := outer.EncodeInput()

	exec, _ := New(nil, state, &xenv.BlockContext{}, thor.NoFork).PrepareClause(
		tx.NewClause(&builtin.Measure.Address).WithData(innerData),
		0,
		math.MaxUint64,
		&xenv.TransactionContext{})
	innerOutput, _, err := exec()

	assert.Nil(t, err)
	assert.Nil(t, innerOutput.VMErr)

	exec, _ = New(nil, state, &xenv.BlockContext{}, thor.NoFork).PrepareClause(
		tx.NewClause(&builtin.Measure.Address).WithData(outerData),
		0,
		math.MaxUint64,
		&xenv.TransactionContext{})

	outerOutput, _, err := exec()

	assert.Nil(t, err)
	assert.Nil(t, outerOutput.VMErr)

	innerGasUsed := math.MaxUint64 - innerOutput.LeftOverGas
	outerGasUsed := math.MaxUint64 - outerOutput.LeftOverGas

	// gas = enter1 + prepare2 + enter2 + leave2 + leave1
	// here returns prepare2
	assert.Equal(t, uint64(1562), outerGasUsed-innerGasUsed*2)
}
