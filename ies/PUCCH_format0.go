package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_format0 struct {
	initialCyclicShift  int64 `lb:0,ub:11,madatory`
	nrofSymbols         int64 `lb:1,ub:2,madatory`
	startingSymbolIndex int64 `lb:0,ub:13,madatory`
}

func (ie *PUCCH_format0) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.initialCyclicShift, &uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("WriteInteger initialCyclicShift", err)
	}
	if err = w.WriteInteger(ie.nrofSymbols, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger nrofSymbols", err)
	}
	if err = w.WriteInteger(ie.startingSymbolIndex, &uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("WriteInteger startingSymbolIndex", err)
	}
	return nil
}

func (ie *PUCCH_format0) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_initialCyclicShift int64
	if tmp_int_initialCyclicShift, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("ReadInteger initialCyclicShift", err)
	}
	ie.initialCyclicShift = tmp_int_initialCyclicShift
	var tmp_int_nrofSymbols int64
	if tmp_int_nrofSymbols, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger nrofSymbols", err)
	}
	ie.nrofSymbols = tmp_int_nrofSymbols
	var tmp_int_startingSymbolIndex int64
	if tmp_int_startingSymbolIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("ReadInteger startingSymbolIndex", err)
	}
	ie.startingSymbolIndex = tmp_int_startingSymbolIndex
	return nil
}
