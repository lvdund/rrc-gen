package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConfigRestrictInfoSCG_servCellIndexRangeSCG struct {
	lowBound ServCellIndex `madatory`
	upBound  ServCellIndex `madatory`
}

func (ie *ConfigRestrictInfoSCG_servCellIndexRangeSCG) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.lowBound.Encode(w); err != nil {
		return utils.WrapError("Encode lowBound", err)
	}
	if err = ie.upBound.Encode(w); err != nil {
		return utils.WrapError("Encode upBound", err)
	}
	return nil
}

func (ie *ConfigRestrictInfoSCG_servCellIndexRangeSCG) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.lowBound.Decode(r); err != nil {
		return utils.WrapError("Decode lowBound", err)
	}
	if err = ie.upBound.Decode(r); err != nil {
		return utils.WrapError("Decode upBound", err)
	}
	return nil
}
