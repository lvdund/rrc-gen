package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_PathlossReferenceRS_r16 struct {
	pusch_PathlossReferenceRS_Id_r16 PUSCH_PathlossReferenceRS_Id_v1610                `madatory`
	referenceSignal_r16              PUSCH_PathlossReferenceRS_r16_referenceSignal_r16 `madatory`
}

func (ie *PUSCH_PathlossReferenceRS_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.pusch_PathlossReferenceRS_Id_r16.Encode(w); err != nil {
		return utils.WrapError("Encode pusch_PathlossReferenceRS_Id_r16", err)
	}
	if err = ie.referenceSignal_r16.Encode(w); err != nil {
		return utils.WrapError("Encode referenceSignal_r16", err)
	}
	return nil
}

func (ie *PUSCH_PathlossReferenceRS_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.pusch_PathlossReferenceRS_Id_r16.Decode(r); err != nil {
		return utils.WrapError("Decode pusch_PathlossReferenceRS_Id_r16", err)
	}
	if err = ie.referenceSignal_r16.Decode(r); err != nil {
		return utils.WrapError("Decode referenceSignal_r16", err)
	}
	return nil
}
