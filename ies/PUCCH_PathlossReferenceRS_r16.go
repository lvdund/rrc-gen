package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_PathlossReferenceRS_r16 struct {
	pucch_PathlossReferenceRS_Id_r16 PUCCH_PathlossReferenceRS_Id_v1610                `madatory`
	referenceSignal_r16              PUCCH_PathlossReferenceRS_r16_referenceSignal_r16 `madatory`
}

func (ie *PUCCH_PathlossReferenceRS_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.pucch_PathlossReferenceRS_Id_r16.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_PathlossReferenceRS_Id_r16", err)
	}
	if err = ie.referenceSignal_r16.Encode(w); err != nil {
		return utils.WrapError("Encode referenceSignal_r16", err)
	}
	return nil
}

func (ie *PUCCH_PathlossReferenceRS_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.pucch_PathlossReferenceRS_Id_r16.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_PathlossReferenceRS_Id_r16", err)
	}
	if err = ie.referenceSignal_r16.Decode(r); err != nil {
		return utils.WrapError("Decode referenceSignal_r16", err)
	}
	return nil
}
