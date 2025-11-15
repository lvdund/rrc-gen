package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DMRS_DownlinkConfig_dmrs_AdditionalPosition_Enum_pos0 uper.Enumerated = 0
	DMRS_DownlinkConfig_dmrs_AdditionalPosition_Enum_pos1 uper.Enumerated = 1
	DMRS_DownlinkConfig_dmrs_AdditionalPosition_Enum_pos3 uper.Enumerated = 2
)

type DMRS_DownlinkConfig_dmrs_AdditionalPosition struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *DMRS_DownlinkConfig_dmrs_AdditionalPosition) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode DMRS_DownlinkConfig_dmrs_AdditionalPosition", err)
	}
	return nil
}

func (ie *DMRS_DownlinkConfig_dmrs_AdditionalPosition) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode DMRS_DownlinkConfig_dmrs_AdditionalPosition", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
