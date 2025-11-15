package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFRA_occasions struct {
	rach_ConfigGeneric   RACH_ConfigGeneric                   `madatory`
	ssb_perRACH_Occasion *CFRA_occasions_ssb_perRACH_Occasion `optional`
}

func (ie *CFRA_occasions) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ssb_perRACH_Occasion != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.rach_ConfigGeneric.Encode(w); err != nil {
		return utils.WrapError("Encode rach_ConfigGeneric", err)
	}
	if ie.ssb_perRACH_Occasion != nil {
		if err = ie.ssb_perRACH_Occasion.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_perRACH_Occasion", err)
		}
	}
	return nil
}

func (ie *CFRA_occasions) Decode(r *uper.UperReader) error {
	var err error
	var ssb_perRACH_OccasionPresent bool
	if ssb_perRACH_OccasionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.rach_ConfigGeneric.Decode(r); err != nil {
		return utils.WrapError("Decode rach_ConfigGeneric", err)
	}
	if ssb_perRACH_OccasionPresent {
		ie.ssb_perRACH_Occasion = new(CFRA_occasions_ssb_perRACH_Occasion)
		if err = ie.ssb_perRACH_Occasion.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_perRACH_Occasion", err)
		}
	}
	return nil
}
