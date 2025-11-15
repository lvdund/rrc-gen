package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UL_UM_RLC struct {
	sn_FieldLength *SN_FieldLengthUM `optional`
}

func (ie *UL_UM_RLC) Encode(w *uper.UperWriter) error {
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
	return nil
}

func (ie *UL_UM_RLC) Decode(r *uper.UperReader) error {
	var err error
	var sn_FieldLengthPresent bool
	if sn_FieldLengthPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if sn_FieldLengthPresent {
		ie.sn_FieldLength = new(SN_FieldLengthUM)
		if err = ie.sn_FieldLength.Decode(r); err != nil {
			return utils.WrapError("Decode sn_FieldLength", err)
		}
	}
	return nil
}
