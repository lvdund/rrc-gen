package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRX_Info_shortDRX struct {
	drx_ShortCycle      DRX_Info_shortDRX_drx_ShortCycle `madatory`
	drx_ShortCycleTimer int64                            `lb:1,ub:16,madatory`
}

func (ie *DRX_Info_shortDRX) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.drx_ShortCycle.Encode(w); err != nil {
		return utils.WrapError("Encode drx_ShortCycle", err)
	}
	if err = w.WriteInteger(ie.drx_ShortCycleTimer, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger drx_ShortCycleTimer", err)
	}
	return nil
}

func (ie *DRX_Info_shortDRX) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.drx_ShortCycle.Decode(r); err != nil {
		return utils.WrapError("Decode drx_ShortCycle", err)
	}
	var tmp_int_drx_ShortCycleTimer int64
	if tmp_int_drx_ShortCycleTimer, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger drx_ShortCycleTimer", err)
	}
	ie.drx_ShortCycleTimer = tmp_int_drx_ShortCycleTimer
	return nil
}
