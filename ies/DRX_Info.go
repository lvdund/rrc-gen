package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRX_Info struct {
	drx_LongCycleStartOffset DRX_Info_drx_LongCycleStartOffset `lb:0,ub:9,madatory`
	shortDRX                 *DRX_Info_shortDRX                `lb:1,ub:16,optional`
}

func (ie *DRX_Info) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.shortDRX != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.drx_LongCycleStartOffset.Encode(w); err != nil {
		return utils.WrapError("Encode drx_LongCycleStartOffset", err)
	}
	if ie.shortDRX != nil {
		if err = ie.shortDRX.Encode(w); err != nil {
			return utils.WrapError("Encode shortDRX", err)
		}
	}
	return nil
}

func (ie *DRX_Info) Decode(r *uper.UperReader) error {
	var err error
	var shortDRXPresent bool
	if shortDRXPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.drx_LongCycleStartOffset.Decode(r); err != nil {
		return utils.WrapError("Decode drx_LongCycleStartOffset", err)
	}
	if shortDRXPresent {
		ie.shortDRX = new(DRX_Info_shortDRX)
		if err = ie.shortDRX.Decode(r); err != nil {
			return utils.WrapError("Decode shortDRX", err)
		}
	}
	return nil
}
