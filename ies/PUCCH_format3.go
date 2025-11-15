package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_format3 struct {
	nrofPRBs            int64 `lb:1,ub:16,madatory`
	nrofSymbols         int64 `lb:4,ub:14,madatory`
	startingSymbolIndex int64 `lb:0,ub:10,madatory`
}

func (ie *PUCCH_format3) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.nrofPRBs, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger nrofPRBs", err)
	}
	if err = w.WriteInteger(ie.nrofSymbols, &uper.Constraint{Lb: 4, Ub: 14}, false); err != nil {
		return utils.WrapError("WriteInteger nrofSymbols", err)
	}
	if err = w.WriteInteger(ie.startingSymbolIndex, &uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("WriteInteger startingSymbolIndex", err)
	}
	return nil
}

func (ie *PUCCH_format3) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_nrofPRBs int64
	if tmp_int_nrofPRBs, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger nrofPRBs", err)
	}
	ie.nrofPRBs = tmp_int_nrofPRBs
	var tmp_int_nrofSymbols int64
	if tmp_int_nrofSymbols, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: 14}, false); err != nil {
		return utils.WrapError("ReadInteger nrofSymbols", err)
	}
	ie.nrofSymbols = tmp_int_nrofSymbols
	var tmp_int_startingSymbolIndex int64
	if tmp_int_startingSymbolIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("ReadInteger startingSymbolIndex", err)
	}
	ie.startingSymbolIndex = tmp_int_startingSymbolIndex
	return nil
}
