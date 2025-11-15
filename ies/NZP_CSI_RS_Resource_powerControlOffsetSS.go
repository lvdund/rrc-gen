package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NZP_CSI_RS_Resource_powerControlOffsetSS_Enum_db_3 uper.Enumerated = 0
	NZP_CSI_RS_Resource_powerControlOffsetSS_Enum_db0  uper.Enumerated = 1
	NZP_CSI_RS_Resource_powerControlOffsetSS_Enum_db3  uper.Enumerated = 2
	NZP_CSI_RS_Resource_powerControlOffsetSS_Enum_db6  uper.Enumerated = 3
)

type NZP_CSI_RS_Resource_powerControlOffsetSS struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *NZP_CSI_RS_Resource_powerControlOffsetSS) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode NZP_CSI_RS_Resource_powerControlOffsetSS", err)
	}
	return nil
}

func (ie *NZP_CSI_RS_Resource_powerControlOffsetSS) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode NZP_CSI_RS_Resource_powerControlOffsetSS", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
