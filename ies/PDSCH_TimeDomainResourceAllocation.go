package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_TimeDomainResourceAllocation struct {
	k0                   *int64                                         `lb:0,ub:32,optional`
	mappingType          PDSCH_TimeDomainResourceAllocation_mappingType `madatory`
	startSymbolAndLength int64                                          `lb:0,ub:127,madatory`
}

func (ie *PDSCH_TimeDomainResourceAllocation) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.k0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.k0 != nil {
		if err = w.WriteInteger(*ie.k0, &uper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode k0", err)
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

func (ie *PDSCH_TimeDomainResourceAllocation) Decode(r *uper.UperReader) error {
	var err error
	var k0Present bool
	if k0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if k0Present {
		var tmp_int_k0 int64
		if tmp_int_k0, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode k0", err)
		}
		ie.k0 = &tmp_int_k0
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
