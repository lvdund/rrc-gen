package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SpCellConfig_rlmInSyncOutOfSyncThreshold_Enum_n1 uper.Enumerated = 0
)

type SpCellConfig_rlmInSyncOutOfSyncThreshold struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *SpCellConfig_rlmInSyncOutOfSyncThreshold) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode SpCellConfig_rlmInSyncOutOfSyncThreshold", err)
	}
	return nil
}

func (ie *SpCellConfig_rlmInSyncOutOfSyncThreshold) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode SpCellConfig_rlmInSyncOutOfSyncThreshold", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
