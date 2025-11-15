package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SimulSRS_ForAntennaSwitching_r16 struct {
	supportSRS_xTyR_xLessThanY_r16  *SimulSRS_ForAntennaSwitching_r16_supportSRS_xTyR_xLessThanY_r16  `optional`
	supportSRS_xTyR_xEqualToY_r16   *SimulSRS_ForAntennaSwitching_r16_supportSRS_xTyR_xEqualToY_r16   `optional`
	supportSRS_AntennaSwitching_r16 *SimulSRS_ForAntennaSwitching_r16_supportSRS_AntennaSwitching_r16 `optional`
}

func (ie *SimulSRS_ForAntennaSwitching_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportSRS_xTyR_xLessThanY_r16 != nil, ie.supportSRS_xTyR_xEqualToY_r16 != nil, ie.supportSRS_AntennaSwitching_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportSRS_xTyR_xLessThanY_r16 != nil {
		if err = ie.supportSRS_xTyR_xLessThanY_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportSRS_xTyR_xLessThanY_r16", err)
		}
	}
	if ie.supportSRS_xTyR_xEqualToY_r16 != nil {
		if err = ie.supportSRS_xTyR_xEqualToY_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportSRS_xTyR_xEqualToY_r16", err)
		}
	}
	if ie.supportSRS_AntennaSwitching_r16 != nil {
		if err = ie.supportSRS_AntennaSwitching_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportSRS_AntennaSwitching_r16", err)
		}
	}
	return nil
}

func (ie *SimulSRS_ForAntennaSwitching_r16) Decode(r *uper.UperReader) error {
	var err error
	var supportSRS_xTyR_xLessThanY_r16Present bool
	if supportSRS_xTyR_xLessThanY_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportSRS_xTyR_xEqualToY_r16Present bool
	if supportSRS_xTyR_xEqualToY_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportSRS_AntennaSwitching_r16Present bool
	if supportSRS_AntennaSwitching_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportSRS_xTyR_xLessThanY_r16Present {
		ie.supportSRS_xTyR_xLessThanY_r16 = new(SimulSRS_ForAntennaSwitching_r16_supportSRS_xTyR_xLessThanY_r16)
		if err = ie.supportSRS_xTyR_xLessThanY_r16.Decode(r); err != nil {
			return utils.WrapError("Decode supportSRS_xTyR_xLessThanY_r16", err)
		}
	}
	if supportSRS_xTyR_xEqualToY_r16Present {
		ie.supportSRS_xTyR_xEqualToY_r16 = new(SimulSRS_ForAntennaSwitching_r16_supportSRS_xTyR_xEqualToY_r16)
		if err = ie.supportSRS_xTyR_xEqualToY_r16.Decode(r); err != nil {
			return utils.WrapError("Decode supportSRS_xTyR_xEqualToY_r16", err)
		}
	}
	if supportSRS_AntennaSwitching_r16Present {
		ie.supportSRS_AntennaSwitching_r16 = new(SimulSRS_ForAntennaSwitching_r16_supportSRS_AntennaSwitching_r16)
		if err = ie.supportSRS_AntennaSwitching_r16.Decode(r); err != nil {
			return utils.WrapError("Decode supportSRS_AntennaSwitching_r16", err)
		}
	}
	return nil
}
