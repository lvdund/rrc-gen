package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FDM_TDM_r16 struct {
	repetitionScheme_r16      FDM_TDM_r16_repetitionScheme_r16 `madatory`
	startingSymbolOffsetK_r16 *int64                           `lb:0,ub:7,optional`
}

func (ie *FDM_TDM_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.startingSymbolOffsetK_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.repetitionScheme_r16.Encode(w); err != nil {
		return utils.WrapError("Encode repetitionScheme_r16", err)
	}
	if ie.startingSymbolOffsetK_r16 != nil {
		if err = w.WriteInteger(*ie.startingSymbolOffsetK_r16, &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Encode startingSymbolOffsetK_r16", err)
		}
	}
	return nil
}

func (ie *FDM_TDM_r16) Decode(r *uper.UperReader) error {
	var err error
	var startingSymbolOffsetK_r16Present bool
	if startingSymbolOffsetK_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.repetitionScheme_r16.Decode(r); err != nil {
		return utils.WrapError("Decode repetitionScheme_r16", err)
	}
	if startingSymbolOffsetK_r16Present {
		var tmp_int_startingSymbolOffsetK_r16 int64
		if tmp_int_startingSymbolOffsetK_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode startingSymbolOffsetK_r16", err)
		}
		ie.startingSymbolOffsetK_r16 = &tmp_int_startingSymbolOffsetK_r16
	}
	return nil
}
