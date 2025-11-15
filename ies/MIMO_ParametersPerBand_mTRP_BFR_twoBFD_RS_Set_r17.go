package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_mTRP_BFR_twoBFD_RS_Set_r17 struct {
	maxBFD_RS_resourcesPerSetPerBWP_r17     MIMO_ParametersPerBand_mTRP_BFR_twoBFD_RS_Set_r17_maxBFD_RS_resourcesPerSetPerBWP_r17     `madatory`
	maxBFR_r17                              int64                                                                                     `lb:1,ub:9,madatory`
	maxBFD_RS_resourcesAcrossSetsPerBWP_r17 MIMO_ParametersPerBand_mTRP_BFR_twoBFD_RS_Set_r17_maxBFD_RS_resourcesAcrossSetsPerBWP_r17 `madatory`
}

func (ie *MIMO_ParametersPerBand_mTRP_BFR_twoBFD_RS_Set_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxBFD_RS_resourcesPerSetPerBWP_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxBFD_RS_resourcesPerSetPerBWP_r17", err)
	}
	if err = w.WriteInteger(ie.maxBFR_r17, &uper.Constraint{Lb: 1, Ub: 9}, false); err != nil {
		return utils.WrapError("WriteInteger maxBFR_r17", err)
	}
	if err = ie.maxBFD_RS_resourcesAcrossSetsPerBWP_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxBFD_RS_resourcesAcrossSetsPerBWP_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_mTRP_BFR_twoBFD_RS_Set_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxBFD_RS_resourcesPerSetPerBWP_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxBFD_RS_resourcesPerSetPerBWP_r17", err)
	}
	var tmp_int_maxBFR_r17 int64
	if tmp_int_maxBFR_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 9}, false); err != nil {
		return utils.WrapError("ReadInteger maxBFR_r17", err)
	}
	ie.maxBFR_r17 = tmp_int_maxBFR_r17
	if err = ie.maxBFD_RS_resourcesAcrossSetsPerBWP_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxBFD_RS_resourcesAcrossSetsPerBWP_r17", err)
	}
	return nil
}
