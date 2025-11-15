package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRB_CountInfo struct {
	drb_Identity   DRB_Identity `madatory`
	count_Uplink   int64        `lb:0,ub:4294967295,madatory`
	count_Downlink int64        `lb:0,ub:4294967295,madatory`
}

func (ie *DRB_CountInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.drb_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode drb_Identity", err)
	}
	if err = w.WriteInteger(ie.count_Uplink, &uper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return utils.WrapError("WriteInteger count_Uplink", err)
	}
	if err = w.WriteInteger(ie.count_Downlink, &uper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return utils.WrapError("WriteInteger count_Downlink", err)
	}
	return nil
}

func (ie *DRB_CountInfo) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.drb_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode drb_Identity", err)
	}
	var tmp_int_count_Uplink int64
	if tmp_int_count_Uplink, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return utils.WrapError("ReadInteger count_Uplink", err)
	}
	ie.count_Uplink = tmp_int_count_Uplink
	var tmp_int_count_Downlink int64
	if tmp_int_count_Downlink, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return utils.WrapError("ReadInteger count_Downlink", err)
	}
	ie.count_Downlink = tmp_int_count_Downlink
	return nil
}
