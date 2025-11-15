package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCP_Config_drb_pdcp_SN_SizeUL_Enum_len12bits uper.Enumerated = 0
	PDCP_Config_drb_pdcp_SN_SizeUL_Enum_len18bits uper.Enumerated = 1
)

type PDCP_Config_drb_pdcp_SN_SizeUL struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PDCP_Config_drb_pdcp_SN_SizeUL) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PDCP_Config_drb_pdcp_SN_SizeUL", err)
	}
	return nil
}

func (ie *PDCP_Config_drb_pdcp_SN_SizeUL) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PDCP_Config_drb_pdcp_SN_SizeUL", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
