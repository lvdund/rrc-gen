package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxDirectCurrentBWP struct {
	bwp_Id                  BWP_Id `madatory`
	shift7dot5kHz           bool   `madatory`
	txDirectCurrentLocation int64  `lb:0,ub:3301,madatory`
}

func (ie *UplinkTxDirectCurrentBWP) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.bwp_Id.Encode(w); err != nil {
		return utils.WrapError("Encode bwp_Id", err)
	}
	if err = w.WriteBoolean(ie.shift7dot5kHz); err != nil {
		return utils.WrapError("WriteBoolean shift7dot5kHz", err)
	}
	if err = w.WriteInteger(ie.txDirectCurrentLocation, &uper.Constraint{Lb: 0, Ub: 3301}, false); err != nil {
		return utils.WrapError("WriteInteger txDirectCurrentLocation", err)
	}
	return nil
}

func (ie *UplinkTxDirectCurrentBWP) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.bwp_Id.Decode(r); err != nil {
		return utils.WrapError("Decode bwp_Id", err)
	}
	var tmp_bool_shift7dot5kHz bool
	if tmp_bool_shift7dot5kHz, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean shift7dot5kHz", err)
	}
	ie.shift7dot5kHz = tmp_bool_shift7dot5kHz
	var tmp_int_txDirectCurrentLocation int64
	if tmp_int_txDirectCurrentLocation, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3301}, false); err != nil {
		return utils.WrapError("ReadInteger txDirectCurrentLocation", err)
	}
	ie.txDirectCurrentLocation = tmp_int_txDirectCurrentLocation
	return nil
}
