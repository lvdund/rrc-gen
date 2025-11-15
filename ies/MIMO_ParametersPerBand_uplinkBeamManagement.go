package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_uplinkBeamManagement struct {
	maxNumberSRS_ResourcePerSet_BM MIMO_ParametersPerBand_uplinkBeamManagement_maxNumberSRS_ResourcePerSet_BM `madatory`
	maxNumberSRS_ResourceSet       int64                                                                      `lb:1,ub:8,madatory`
}

func (ie *MIMO_ParametersPerBand_uplinkBeamManagement) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxNumberSRS_ResourcePerSet_BM.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSRS_ResourcePerSet_BM", err)
	}
	if err = w.WriteInteger(ie.maxNumberSRS_ResourceSet, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberSRS_ResourceSet", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_uplinkBeamManagement) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxNumberSRS_ResourcePerSet_BM.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSRS_ResourcePerSet_BM", err)
	}
	var tmp_int_maxNumberSRS_ResourceSet int64
	if tmp_int_maxNumberSRS_ResourceSet, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberSRS_ResourceSet", err)
	}
	ie.maxNumberSRS_ResourceSet = tmp_int_maxNumberSRS_ResourceSet
	return nil
}
