package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PathlossReferenceRS_r16 struct {
	srs_PathlossReferenceRS_Id_r16 SRS_PathlossReferenceRS_Id_r16 `madatory`
	pathlossReferenceRS_r16        PathlossReferenceRS_Config     `madatory`
}

func (ie *PathlossReferenceRS_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.srs_PathlossReferenceRS_Id_r16.Encode(w); err != nil {
		return utils.WrapError("Encode srs_PathlossReferenceRS_Id_r16", err)
	}
	if err = ie.pathlossReferenceRS_r16.Encode(w); err != nil {
		return utils.WrapError("Encode pathlossReferenceRS_r16", err)
	}
	return nil
}

func (ie *PathlossReferenceRS_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.srs_PathlossReferenceRS_Id_r16.Decode(r); err != nil {
		return utils.WrapError("Decode srs_PathlossReferenceRS_Id_r16", err)
	}
	if err = ie.pathlossReferenceRS_r16.Decode(r); err != nil {
		return utils.WrapError("Decode pathlossReferenceRS_r16", err)
	}
	return nil
}
