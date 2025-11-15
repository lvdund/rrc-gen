package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Uplink_powerControl_r17 struct {
	ul_powercontrolId_r17  Uplink_powerControlId_r17 `madatory`
	p0AlphaSetforPUSCH_r17 *P0AlphaSet_r17           `optional`
	p0AlphaSetforPUCCH_r17 *P0AlphaSet_r17           `optional`
	p0AlphaSetforSRS_r17   *P0AlphaSet_r17           `optional`
}

func (ie *Uplink_powerControl_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.p0AlphaSetforPUSCH_r17 != nil, ie.p0AlphaSetforPUCCH_r17 != nil, ie.p0AlphaSetforSRS_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ul_powercontrolId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ul_powercontrolId_r17", err)
	}
	if ie.p0AlphaSetforPUSCH_r17 != nil {
		if err = ie.p0AlphaSetforPUSCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode p0AlphaSetforPUSCH_r17", err)
		}
	}
	if ie.p0AlphaSetforPUCCH_r17 != nil {
		if err = ie.p0AlphaSetforPUCCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode p0AlphaSetforPUCCH_r17", err)
		}
	}
	if ie.p0AlphaSetforSRS_r17 != nil {
		if err = ie.p0AlphaSetforSRS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode p0AlphaSetforSRS_r17", err)
		}
	}
	return nil
}

func (ie *Uplink_powerControl_r17) Decode(r *uper.UperReader) error {
	var err error
	var p0AlphaSetforPUSCH_r17Present bool
	if p0AlphaSetforPUSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var p0AlphaSetforPUCCH_r17Present bool
	if p0AlphaSetforPUCCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var p0AlphaSetforSRS_r17Present bool
	if p0AlphaSetforSRS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ul_powercontrolId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ul_powercontrolId_r17", err)
	}
	if p0AlphaSetforPUSCH_r17Present {
		ie.p0AlphaSetforPUSCH_r17 = new(P0AlphaSet_r17)
		if err = ie.p0AlphaSetforPUSCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode p0AlphaSetforPUSCH_r17", err)
		}
	}
	if p0AlphaSetforPUCCH_r17Present {
		ie.p0AlphaSetforPUCCH_r17 = new(P0AlphaSet_r17)
		if err = ie.p0AlphaSetforPUCCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode p0AlphaSetforPUCCH_r17", err)
		}
	}
	if p0AlphaSetforSRS_r17Present {
		ie.p0AlphaSetforSRS_r17 = new(P0AlphaSet_r17)
		if err = ie.p0AlphaSetforSRS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode p0AlphaSetforSRS_r17", err)
		}
	}
	return nil
}
