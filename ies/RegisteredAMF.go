package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RegisteredAMF struct {
	plmn_Identity  *PLMN_Identity `optional`
	amf_Identifier AMF_Identifier `madatory`
}

func (ie *RegisteredAMF) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.plmn_Identity != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.plmn_Identity != nil {
		if err = ie.plmn_Identity.Encode(w); err != nil {
			return utils.WrapError("Encode plmn_Identity", err)
		}
	}
	if err = ie.amf_Identifier.Encode(w); err != nil {
		return utils.WrapError("Encode amf_Identifier", err)
	}
	return nil
}

func (ie *RegisteredAMF) Decode(r *uper.UperReader) error {
	var err error
	var plmn_IdentityPresent bool
	if plmn_IdentityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if plmn_IdentityPresent {
		ie.plmn_Identity = new(PLMN_Identity)
		if err = ie.plmn_Identity.Decode(r); err != nil {
			return utils.WrapError("Decode plmn_Identity", err)
		}
	}
	if err = ie.amf_Identifier.Decode(r); err != nil {
		return utils.WrapError("Decode amf_Identifier", err)
	}
	return nil
}
