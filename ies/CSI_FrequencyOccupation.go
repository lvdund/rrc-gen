package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_FrequencyOccupation struct {
	startingRB int64 `lb:0,ub:maxNrofPhysicalResourceBlocks_1,madatory`
	nrofRBs    int64 `lb:24,ub:maxNrofPhysicalResourceBlocksPlus1,madatory`
}

func (ie *CSI_FrequencyOccupation) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.startingRB, &uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("WriteInteger startingRB", err)
	}
	if err = w.WriteInteger(ie.nrofRBs, &uper.Constraint{Lb: 24, Ub: maxNrofPhysicalResourceBlocksPlus1}, false); err != nil {
		return utils.WrapError("WriteInteger nrofRBs", err)
	}
	return nil
}

func (ie *CSI_FrequencyOccupation) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_startingRB int64
	if tmp_int_startingRB, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("ReadInteger startingRB", err)
	}
	ie.startingRB = tmp_int_startingRB
	var tmp_int_nrofRBs int64
	if tmp_int_nrofRBs, err = r.ReadInteger(&uper.Constraint{Lb: 24, Ub: maxNrofPhysicalResourceBlocksPlus1}, false); err != nil {
		return utils.WrapError("ReadInteger nrofRBs", err)
	}
	ie.nrofRBs = tmp_int_nrofRBs
	return nil
}
