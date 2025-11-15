package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CGI_Info_Logging_r16 struct {
	plmn_Identity_r16    PLMN_Identity     `madatory`
	cellIdentity_r16     CellIdentity      `madatory`
	trackingAreaCode_r16 *TrackingAreaCode `optional`
}

func (ie *CGI_Info_Logging_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.trackingAreaCode_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.plmn_Identity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_Identity_r16", err)
	}
	if err = ie.cellIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode cellIdentity_r16", err)
	}
	if ie.trackingAreaCode_r16 != nil {
		if err = ie.trackingAreaCode_r16.Encode(w); err != nil {
			return utils.WrapError("Encode trackingAreaCode_r16", err)
		}
	}
	return nil
}

func (ie *CGI_Info_Logging_r16) Decode(r *uper.UperReader) error {
	var err error
	var trackingAreaCode_r16Present bool
	if trackingAreaCode_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.plmn_Identity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode plmn_Identity_r16", err)
	}
	if err = ie.cellIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode cellIdentity_r16", err)
	}
	if trackingAreaCode_r16Present {
		ie.trackingAreaCode_r16 = new(TrackingAreaCode)
		if err = ie.trackingAreaCode_r16.Decode(r); err != nil {
			return utils.WrapError("Decode trackingAreaCode_r16", err)
		}
	}
	return nil
}
