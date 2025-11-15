package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UAC_BarringPerPLMN struct {
	plmn_IdentityIndex    int64                                     `lb:1,ub:maxPLMN,madatory`
	uac_ACBarringListType *UAC_BarringPerPLMN_uac_ACBarringListType `lb:maxAccessCat_1,ub:maxAccessCat_1,optional`
}

func (ie *UAC_BarringPerPLMN) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.uac_ACBarringListType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.plmn_IdentityIndex, &uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
		return utils.WrapError("WriteInteger plmn_IdentityIndex", err)
	}
	if ie.uac_ACBarringListType != nil {
		if err = ie.uac_ACBarringListType.Encode(w); err != nil {
			return utils.WrapError("Encode uac_ACBarringListType", err)
		}
	}
	return nil
}

func (ie *UAC_BarringPerPLMN) Decode(r *uper.UperReader) error {
	var err error
	var uac_ACBarringListTypePresent bool
	if uac_ACBarringListTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_plmn_IdentityIndex int64
	if tmp_int_plmn_IdentityIndex, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
		return utils.WrapError("ReadInteger plmn_IdentityIndex", err)
	}
	ie.plmn_IdentityIndex = tmp_int_plmn_IdentityIndex
	if uac_ACBarringListTypePresent {
		ie.uac_ACBarringListType = new(UAC_BarringPerPLMN_uac_ACBarringListType)
		if err = ie.uac_ACBarringListType.Decode(r); err != nil {
			return utils.WrapError("Decode uac_ACBarringListType", err)
		}
	}
	return nil
}
