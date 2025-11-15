package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_resourceMapping_r17 struct {
	startPosition_r17    int64                                                 `lb:0,ub:13,madatory`
	nrofSymbols_r17      SRS_Resource_resourceMapping_r17_nrofSymbols_r17      `madatory`
	repetitionFactor_r17 SRS_Resource_resourceMapping_r17_repetitionFactor_r17 `madatory`
}

func (ie *SRS_Resource_resourceMapping_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.startPosition_r17, &uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("WriteInteger startPosition_r17", err)
	}
	if err = ie.nrofSymbols_r17.Encode(w); err != nil {
		return utils.WrapError("Encode nrofSymbols_r17", err)
	}
	if err = ie.repetitionFactor_r17.Encode(w); err != nil {
		return utils.WrapError("Encode repetitionFactor_r17", err)
	}
	return nil
}

func (ie *SRS_Resource_resourceMapping_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_startPosition_r17 int64
	if tmp_int_startPosition_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("ReadInteger startPosition_r17", err)
	}
	ie.startPosition_r17 = tmp_int_startPosition_r17
	if err = ie.nrofSymbols_r17.Decode(r); err != nil {
		return utils.WrapError("Decode nrofSymbols_r17", err)
	}
	if err = ie.repetitionFactor_r17.Decode(r); err != nil {
		return utils.WrapError("Decode repetitionFactor_r17", err)
	}
	return nil
}
