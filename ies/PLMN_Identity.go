package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PLMN_Identity struct {
	mcc *MCC `optional`
	mnc MNC  `madatory`
}

func (ie *PLMN_Identity) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mcc != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mcc != nil {
		if err = ie.mcc.Encode(w); err != nil {
			return utils.WrapError("Encode mcc", err)
		}
	}
	if err = ie.mnc.Encode(w); err != nil {
		return utils.WrapError("Encode mnc", err)
	}
	return nil
}

func (ie *PLMN_Identity) Decode(r *uper.UperReader) error {
	var err error
	var mccPresent bool
	if mccPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if mccPresent {
		ie.mcc = new(MCC)
		if err = ie.mcc.Decode(r); err != nil {
			return utils.WrapError("Decode mcc", err)
		}
	}
	if err = ie.mnc.Decode(r); err != nil {
		return utils.WrapError("Decode mnc", err)
	}
	return nil
}
