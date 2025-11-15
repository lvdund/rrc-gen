package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_PathlossReferenceRS struct {
	pucch_PathlossReferenceRS_Id PUCCH_PathlossReferenceRS_Id              `madatory`
	referenceSignal              PUCCH_PathlossReferenceRS_referenceSignal `madatory`
}

func (ie *PUCCH_PathlossReferenceRS) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.pucch_PathlossReferenceRS_Id.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_PathlossReferenceRS_Id", err)
	}
	if err = ie.referenceSignal.Encode(w); err != nil {
		return utils.WrapError("Encode referenceSignal", err)
	}
	return nil
}

func (ie *PUCCH_PathlossReferenceRS) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.pucch_PathlossReferenceRS_Id.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_PathlossReferenceRS_Id", err)
	}
	if err = ie.referenceSignal.Decode(r); err != nil {
		return utils.WrapError("Decode referenceSignal", err)
	}
	return nil
}
