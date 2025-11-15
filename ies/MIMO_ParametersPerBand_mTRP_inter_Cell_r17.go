package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_mTRP_inter_Cell_r17 struct {
	maxNumAdditionalPCI_Case1_r17 int64 `lb:1,ub:7,madatory`
	maxNumAdditionalPCI_Case2_r17 int64 `lb:0,ub:7,madatory`
}

func (ie *MIMO_ParametersPerBand_mTRP_inter_Cell_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxNumAdditionalPCI_Case1_r17, &uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumAdditionalPCI_Case1_r17", err)
	}
	if err = w.WriteInteger(ie.maxNumAdditionalPCI_Case2_r17, &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumAdditionalPCI_Case2_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_mTRP_inter_Cell_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxNumAdditionalPCI_Case1_r17 int64
	if tmp_int_maxNumAdditionalPCI_Case1_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumAdditionalPCI_Case1_r17", err)
	}
	ie.maxNumAdditionalPCI_Case1_r17 = tmp_int_maxNumAdditionalPCI_Case1_r17
	var tmp_int_maxNumAdditionalPCI_Case2_r17 int64
	if tmp_int_maxNumAdditionalPCI_Case2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumAdditionalPCI_Case2_r17", err)
	}
	ie.maxNumAdditionalPCI_Case2_r17 = tmp_int_maxNumAdditionalPCI_Case2_r17
	return nil
}
