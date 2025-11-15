package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LogMeasResultBT_r16 struct {
	bt_Addr_r16 uper.BitString `lb:48,ub:48,madatory`
	rssi_BT_r16 *int64         `lb:-128,ub:127,optional`
}

func (ie *LogMeasResultBT_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rssi_BT_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBitString(ie.bt_Addr_r16.Bytes, uint(ie.bt_Addr_r16.NumBits), &uper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
		return utils.WrapError("WriteBitString bt_Addr_r16", err)
	}
	if ie.rssi_BT_r16 != nil {
		if err = w.WriteInteger(*ie.rssi_BT_r16, &uper.Constraint{Lb: -128, Ub: 127}, false); err != nil {
			return utils.WrapError("Encode rssi_BT_r16", err)
		}
	}
	return nil
}

func (ie *LogMeasResultBT_r16) Decode(r *uper.UperReader) error {
	var err error
	var rssi_BT_r16Present bool
	if rssi_BT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bs_bt_Addr_r16 []byte
	var n_bt_Addr_r16 uint
	if tmp_bs_bt_Addr_r16, n_bt_Addr_r16, err = r.ReadBitString(&uper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
		return utils.WrapError("ReadBitString bt_Addr_r16", err)
	}
	ie.bt_Addr_r16 = uper.BitString{
		Bytes:   tmp_bs_bt_Addr_r16,
		NumBits: uint64(n_bt_Addr_r16),
	}
	if rssi_BT_r16Present {
		var tmp_int_rssi_BT_r16 int64
		if tmp_int_rssi_BT_r16, err = r.ReadInteger(&uper.Constraint{Lb: -128, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode rssi_BT_r16", err)
		}
		ie.rssi_BT_r16 = &tmp_int_rssi_BT_r16
	}
	return nil
}
