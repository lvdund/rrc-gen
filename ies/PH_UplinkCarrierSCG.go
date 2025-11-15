package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PH_UplinkCarrierSCG struct {
	ph_Type1or3 PH_UplinkCarrierSCG_ph_Type1or3 `madatory`
}

func (ie *PH_UplinkCarrierSCG) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ph_Type1or3.Encode(w); err != nil {
		return utils.WrapError("Encode ph_Type1or3", err)
	}
	return nil
}

func (ie *PH_UplinkCarrierSCG) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ph_Type1or3.Decode(r); err != nil {
		return utils.WrapError("Decode ph_Type1or3", err)
	}
	return nil
}
