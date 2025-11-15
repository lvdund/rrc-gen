package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NSAG_IdentityInfo_r17 struct {
	nsag_ID_r17          NSAG_ID_r17       `madatory`
	trackingAreaCode_r17 *TrackingAreaCode `optional`
}

func (ie *NSAG_IdentityInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.trackingAreaCode_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.nsag_ID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode nsag_ID_r17", err)
	}
	if ie.trackingAreaCode_r17 != nil {
		if err = ie.trackingAreaCode_r17.Encode(w); err != nil {
			return utils.WrapError("Encode trackingAreaCode_r17", err)
		}
	}
	return nil
}

func (ie *NSAG_IdentityInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var trackingAreaCode_r17Present bool
	if trackingAreaCode_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.nsag_ID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode nsag_ID_r17", err)
	}
	if trackingAreaCode_r17Present {
		ie.trackingAreaCode_r17 = new(TrackingAreaCode)
		if err = ie.trackingAreaCode_r17.Decode(r); err != nil {
			return utils.WrapError("Decode trackingAreaCode_r17", err)
		}
	}
	return nil
}
