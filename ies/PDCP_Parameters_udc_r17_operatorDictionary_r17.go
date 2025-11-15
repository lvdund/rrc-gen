package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Parameters_udc_r17_operatorDictionary_r17 struct {
	versionOfDictionary_r17 int64         `lb:0,ub:15,madatory`
	associatedPLMN_ID_r17   PLMN_Identity `madatory`
}

func (ie *PDCP_Parameters_udc_r17_operatorDictionary_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.versionOfDictionary_r17, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger versionOfDictionary_r17", err)
	}
	if err = ie.associatedPLMN_ID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode associatedPLMN_ID_r17", err)
	}
	return nil
}

func (ie *PDCP_Parameters_udc_r17_operatorDictionary_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_versionOfDictionary_r17 int64
	if tmp_int_versionOfDictionary_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger versionOfDictionary_r17", err)
	}
	ie.versionOfDictionary_r17 = tmp_int_versionOfDictionary_r17
	if err = ie.associatedPLMN_ID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode associatedPLMN_ID_r17", err)
	}
	return nil
}
