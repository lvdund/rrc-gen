package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DL_AM_RLC struct {
	sn_FieldLength   *SN_FieldLengthAM `optional`
	t_Reassembly     T_Reassembly      `madatory`
	t_StatusProhibit T_StatusProhibit  `madatory`
}

func (ie *DL_AM_RLC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sn_FieldLength != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sn_FieldLength != nil {
		if err = ie.sn_FieldLength.Encode(w); err != nil {
			return utils.WrapError("Encode sn_FieldLength", err)
		}
	}
	if err = ie.t_Reassembly.Encode(w); err != nil {
		return utils.WrapError("Encode t_Reassembly", err)
	}
	if err = ie.t_StatusProhibit.Encode(w); err != nil {
		return utils.WrapError("Encode t_StatusProhibit", err)
	}
	return nil
}

func (ie *DL_AM_RLC) Decode(r *uper.UperReader) error {
	var err error
	var sn_FieldLengthPresent bool
	if sn_FieldLengthPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if sn_FieldLengthPresent {
		ie.sn_FieldLength = new(SN_FieldLengthAM)
		if err = ie.sn_FieldLength.Decode(r); err != nil {
			return utils.WrapError("Decode sn_FieldLength", err)
		}
	}
	if err = ie.t_Reassembly.Decode(r); err != nil {
		return utils.WrapError("Decode t_Reassembly", err)
	}
	if err = ie.t_StatusProhibit.Decode(r); err != nil {
		return utils.WrapError("Decode t_StatusProhibit", err)
	}
	return nil
}
