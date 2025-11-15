package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultSRS_RSRP_r16 struct {
	srs_ResourceId_r16  SRS_ResourceId     `madatory`
	srs_RSRP_Result_r16 SRS_RSRP_Range_r16 `madatory`
}

func (ie *MeasResultSRS_RSRP_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.srs_ResourceId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode srs_ResourceId_r16", err)
	}
	if err = ie.srs_RSRP_Result_r16.Encode(w); err != nil {
		return utils.WrapError("Encode srs_RSRP_Result_r16", err)
	}
	return nil
}

func (ie *MeasResultSRS_RSRP_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.srs_ResourceId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode srs_ResourceId_r16", err)
	}
	if err = ie.srs_RSRP_Result_r16.Decode(r); err != nil {
		return utils.WrapError("Decode srs_RSRP_Result_r16", err)
	}
	return nil
}
