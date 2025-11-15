package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_format4 struct {
	nrofSymbols         int64                    `lb:4,ub:14,madatory`
	occ_Length          PUCCH_format4_occ_Length `madatory`
	occ_Index           PUCCH_format4_occ_Index  `madatory`
	startingSymbolIndex int64                    `lb:0,ub:10,madatory`
}

func (ie *PUCCH_format4) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.nrofSymbols, &uper.Constraint{Lb: 4, Ub: 14}, false); err != nil {
		return utils.WrapError("WriteInteger nrofSymbols", err)
	}
	if err = ie.occ_Length.Encode(w); err != nil {
		return utils.WrapError("Encode occ_Length", err)
	}
	if err = ie.occ_Index.Encode(w); err != nil {
		return utils.WrapError("Encode occ_Index", err)
	}
	if err = w.WriteInteger(ie.startingSymbolIndex, &uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("WriteInteger startingSymbolIndex", err)
	}
	return nil
}

func (ie *PUCCH_format4) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_nrofSymbols int64
	if tmp_int_nrofSymbols, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: 14}, false); err != nil {
		return utils.WrapError("ReadInteger nrofSymbols", err)
	}
	ie.nrofSymbols = tmp_int_nrofSymbols
	if err = ie.occ_Length.Decode(r); err != nil {
		return utils.WrapError("Decode occ_Length", err)
	}
	if err = ie.occ_Index.Decode(r); err != nil {
		return utils.WrapError("Decode occ_Index", err)
	}
	var tmp_int_startingSymbolIndex int64
	if tmp_int_startingSymbolIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("ReadInteger startingSymbolIndex", err)
	}
	ie.startingSymbolIndex = tmp_int_startingSymbolIndex
	return nil
}
