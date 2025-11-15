package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRB_CountMSB_Info struct {
	drb_Identity      DRB_Identity `madatory`
	countMSB_Uplink   int64        `lb:0,ub:33554431,madatory`
	countMSB_Downlink int64        `lb:0,ub:33554431,madatory`
}

func (ie *DRB_CountMSB_Info) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.drb_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode drb_Identity", err)
	}
	if err = w.WriteInteger(ie.countMSB_Uplink, &uper.Constraint{Lb: 0, Ub: 33554431}, false); err != nil {
		return utils.WrapError("WriteInteger countMSB_Uplink", err)
	}
	if err = w.WriteInteger(ie.countMSB_Downlink, &uper.Constraint{Lb: 0, Ub: 33554431}, false); err != nil {
		return utils.WrapError("WriteInteger countMSB_Downlink", err)
	}
	return nil
}

func (ie *DRB_CountMSB_Info) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.drb_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode drb_Identity", err)
	}
	var tmp_int_countMSB_Uplink int64
	if tmp_int_countMSB_Uplink, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 33554431}, false); err != nil {
		return utils.WrapError("ReadInteger countMSB_Uplink", err)
	}
	ie.countMSB_Uplink = tmp_int_countMSB_Uplink
	var tmp_int_countMSB_Downlink int64
	if tmp_int_countMSB_Downlink, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 33554431}, false); err != nil {
		return utils.WrapError("ReadInteger countMSB_Downlink", err)
	}
	ie.countMSB_Downlink = tmp_int_countMSB_Downlink
	return nil
}
