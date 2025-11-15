package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BWP_Downlink struct {
	bwp_Id        BWP_Id                 `madatory`
	bwp_Common    *BWP_DownlinkCommon    `optional`
	bwp_Dedicated *BWP_DownlinkDedicated `optional`
}

func (ie *BWP_Downlink) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bwp_Common != nil, ie.bwp_Dedicated != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.bwp_Id.Encode(w); err != nil {
		return utils.WrapError("Encode bwp_Id", err)
	}
	if ie.bwp_Common != nil {
		if err = ie.bwp_Common.Encode(w); err != nil {
			return utils.WrapError("Encode bwp_Common", err)
		}
	}
	if ie.bwp_Dedicated != nil {
		if err = ie.bwp_Dedicated.Encode(w); err != nil {
			return utils.WrapError("Encode bwp_Dedicated", err)
		}
	}
	return nil
}

func (ie *BWP_Downlink) Decode(r *uper.UperReader) error {
	var err error
	var bwp_CommonPresent bool
	if bwp_CommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var bwp_DedicatedPresent bool
	if bwp_DedicatedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.bwp_Id.Decode(r); err != nil {
		return utils.WrapError("Decode bwp_Id", err)
	}
	if bwp_CommonPresent {
		ie.bwp_Common = new(BWP_DownlinkCommon)
		if err = ie.bwp_Common.Decode(r); err != nil {
			return utils.WrapError("Decode bwp_Common", err)
		}
	}
	if bwp_DedicatedPresent {
		ie.bwp_Dedicated = new(BWP_DownlinkDedicated)
		if err = ie.bwp_Dedicated.Decode(r); err != nil {
			return utils.WrapError("Decode bwp_Dedicated", err)
		}
	}
	return nil
}
