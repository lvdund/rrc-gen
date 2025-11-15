package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_TimeDomainResourceAllocation struct {
	k2                   *int64                                         `lb:0,ub:32,optional`
	mappingType          PUSCH_TimeDomainResourceAllocation_mappingType `madatory`
	startSymbolAndLength int64                                          `lb:0,ub:127,madatory`
}

func (ie *PUSCH_TimeDomainResourceAllocation) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.k2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.k2 != nil {
		if err = w.WriteInteger(*ie.k2, &uper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode k2", err)
		}
	}
	if err = ie.mappingType.Encode(w); err != nil {
		return utils.WrapError("Encode mappingType", err)
	}
	if err = w.WriteInteger(ie.startSymbolAndLength, &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
		return utils.WrapError("WriteInteger startSymbolAndLength", err)
	}
	return nil
}

func (ie *PUSCH_TimeDomainResourceAllocation) Decode(r *uper.UperReader) error {
	var err error
	var k2Present bool
	if k2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if k2Present {
		var tmp_int_k2 int64
		if tmp_int_k2, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode k2", err)
		}
		ie.k2 = &tmp_int_k2
	}
	if err = ie.mappingType.Decode(r); err != nil {
		return utils.WrapError("Decode mappingType", err)
	}
	var tmp_int_startSymbolAndLength int64
	if tmp_int_startSymbolAndLength, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
		return utils.WrapError("ReadInteger startSymbolAndLength", err)
	}
	ie.startSymbolAndLength = tmp_int_startSymbolAndLength
	return nil
}
