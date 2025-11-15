package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_resourceMapping_r16 struct {
	startPosition_r16    int64                                                 `lb:0,ub:13,madatory`
	nrofSymbols_r16      SRS_Resource_resourceMapping_r16_nrofSymbols_r16      `madatory`
	repetitionFactor_r16 SRS_Resource_resourceMapping_r16_repetitionFactor_r16 `madatory`
}

func (ie *SRS_Resource_resourceMapping_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.startPosition_r16, &uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("WriteInteger startPosition_r16", err)
	}
	if err = ie.nrofSymbols_r16.Encode(w); err != nil {
		return utils.WrapError("Encode nrofSymbols_r16", err)
	}
	if err = ie.repetitionFactor_r16.Encode(w); err != nil {
		return utils.WrapError("Encode repetitionFactor_r16", err)
	}
	return nil
}

func (ie *SRS_Resource_resourceMapping_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_startPosition_r16 int64
	if tmp_int_startPosition_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("ReadInteger startPosition_r16", err)
	}
	ie.startPosition_r16 = tmp_int_startPosition_r16
	if err = ie.nrofSymbols_r16.Decode(r); err != nil {
		return utils.WrapError("Decode nrofSymbols_r16", err)
	}
	if err = ie.repetitionFactor_r16.Decode(r); err != nil {
		return utils.WrapError("Decode repetitionFactor_r16", err)
	}
	return nil
}
