package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SDAP_Config_sdap_HeaderDL_Enum_present uper.Enumerated = 0
	SDAP_Config_sdap_HeaderDL_Enum_absent  uper.Enumerated = 1
)

type SDAP_Config_sdap_HeaderDL struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SDAP_Config_sdap_HeaderDL) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SDAP_Config_sdap_HeaderDL", err)
	}
	return nil
}

func (ie *SDAP_Config_sdap_HeaderDL) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SDAP_Config_sdap_HeaderDL", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
