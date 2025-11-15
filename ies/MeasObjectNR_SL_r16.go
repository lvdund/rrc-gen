package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectNR_SL_r16 struct {
	tx_PoolMeasToRemoveList_r16 *Tx_PoolMeasList_r16 `optional`
	tx_PoolMeasToAddModList_r16 *Tx_PoolMeasList_r16 `optional`
}

func (ie *MeasObjectNR_SL_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.tx_PoolMeasToRemoveList_r16 != nil, ie.tx_PoolMeasToAddModList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.tx_PoolMeasToRemoveList_r16 != nil {
		if err = ie.tx_PoolMeasToRemoveList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode tx_PoolMeasToRemoveList_r16", err)
		}
	}
	if ie.tx_PoolMeasToAddModList_r16 != nil {
		if err = ie.tx_PoolMeasToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode tx_PoolMeasToAddModList_r16", err)
		}
	}
	return nil
}

func (ie *MeasObjectNR_SL_r16) Decode(r *uper.UperReader) error {
	var err error
	var tx_PoolMeasToRemoveList_r16Present bool
	if tx_PoolMeasToRemoveList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tx_PoolMeasToAddModList_r16Present bool
	if tx_PoolMeasToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if tx_PoolMeasToRemoveList_r16Present {
		ie.tx_PoolMeasToRemoveList_r16 = new(Tx_PoolMeasList_r16)
		if err = ie.tx_PoolMeasToRemoveList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode tx_PoolMeasToRemoveList_r16", err)
		}
	}
	if tx_PoolMeasToAddModList_r16Present {
		ie.tx_PoolMeasToAddModList_r16 = new(Tx_PoolMeasList_r16)
		if err = ie.tx_PoolMeasToAddModList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode tx_PoolMeasToAddModList_r16", err)
		}
	}
	return nil
}
