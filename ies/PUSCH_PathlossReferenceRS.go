package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_PathlossReferenceRS struct {
	pusch_PathlossReferenceRS_Id PUSCH_PathlossReferenceRS_Id              `madatory`
	referenceSignal              PUSCH_PathlossReferenceRS_referenceSignal `madatory`
}

func (ie *PUSCH_PathlossReferenceRS) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.pusch_PathlossReferenceRS_Id.Encode(w); err != nil {
		return utils.WrapError("Encode pusch_PathlossReferenceRS_Id", err)
	}
	if err = ie.referenceSignal.Encode(w); err != nil {
		return utils.WrapError("Encode referenceSignal", err)
	}
	return nil
}

func (ie *PUSCH_PathlossReferenceRS) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.pusch_PathlossReferenceRS_Id.Decode(r); err != nil {
		return utils.WrapError("Decode pusch_PathlossReferenceRS_Id", err)
	}
	if err = ie.referenceSignal.Decode(r); err != nil {
		return utils.WrapError("Decode referenceSignal", err)
	}
	return nil
}
