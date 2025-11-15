package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_resourceMapping struct {
	startPosition    int64                                         `lb:0,ub:5,madatory`
	nrofSymbols      SRS_Resource_resourceMapping_nrofSymbols      `madatory`
	repetitionFactor SRS_Resource_resourceMapping_repetitionFactor `madatory`
}

func (ie *SRS_Resource_resourceMapping) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.startPosition, &uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("WriteInteger startPosition", err)
	}
	if err = ie.nrofSymbols.Encode(w); err != nil {
		return utils.WrapError("Encode nrofSymbols", err)
	}
	if err = ie.repetitionFactor.Encode(w); err != nil {
		return utils.WrapError("Encode repetitionFactor", err)
	}
	return nil
}

func (ie *SRS_Resource_resourceMapping) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_startPosition int64
	if tmp_int_startPosition, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("ReadInteger startPosition", err)
	}
	ie.startPosition = tmp_int_startPosition
	if err = ie.nrofSymbols.Decode(r); err != nil {
		return utils.WrapError("Decode nrofSymbols", err)
	}
	if err = ie.repetitionFactor.Decode(r); err != nil {
		return utils.WrapError("Decode repetitionFactor", err)
	}
	return nil
}
