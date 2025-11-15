package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_Configuration_r16_sfn0_Offset_r16 struct {
	sfn_Offset_r16            int64  `lb:0,ub:1023,madatory`
	integerSubframeOffset_r16 *int64 `lb:0,ub:9,optional`
}

func (ie *SSB_Configuration_r16_sfn0_Offset_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.integerSubframeOffset_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.sfn_Offset_r16, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("WriteInteger sfn_Offset_r16", err)
	}
	if ie.integerSubframeOffset_r16 != nil {
		if err = w.WriteInteger(*ie.integerSubframeOffset_r16, &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Encode integerSubframeOffset_r16", err)
		}
	}
	return nil
}

func (ie *SSB_Configuration_r16_sfn0_Offset_r16) Decode(r *uper.UperReader) error {
	var err error
	var integerSubframeOffset_r16Present bool
	if integerSubframeOffset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_sfn_Offset_r16 int64
	if tmp_int_sfn_Offset_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("ReadInteger sfn_Offset_r16", err)
	}
	ie.sfn_Offset_r16 = tmp_int_sfn_Offset_r16
	if integerSubframeOffset_r16Present {
		var tmp_int_integerSubframeOffset_r16 int64
		if tmp_int_integerSubframeOffset_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode integerSubframeOffset_r16", err)
		}
		ie.integerSubframeOffset_r16 = &tmp_int_integerSubframeOffset_r16
	}
	return nil
}
