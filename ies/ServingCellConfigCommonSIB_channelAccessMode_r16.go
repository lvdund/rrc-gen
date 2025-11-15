package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfigCommonSIB_channelAccessMode_r16_Choice_nothing uint64 = iota
	ServingCellConfigCommonSIB_channelAccessMode_r16_Choice_dynamic
	ServingCellConfigCommonSIB_channelAccessMode_r16_Choice_semiStatic
)

type ServingCellConfigCommonSIB_channelAccessMode_r16 struct {
	Choice     uint64
	dynamic    uper.NULL `madatory`
	semiStatic *SemiStaticChannelAccessConfig_r16
}

func (ie *ServingCellConfigCommonSIB_channelAccessMode_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ServingCellConfigCommonSIB_channelAccessMode_r16_Choice_dynamic:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode dynamic", err)
		}
	case ServingCellConfigCommonSIB_channelAccessMode_r16_Choice_semiStatic:
		if err = ie.semiStatic.Encode(w); err != nil {
			err = utils.WrapError("Encode semiStatic", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ServingCellConfigCommonSIB_channelAccessMode_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ServingCellConfigCommonSIB_channelAccessMode_r16_Choice_dynamic:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode dynamic", err)
		}
	case ServingCellConfigCommonSIB_channelAccessMode_r16_Choice_semiStatic:
		ie.semiStatic = new(SemiStaticChannelAccessConfig_r16)
		if err = ie.semiStatic.Decode(r); err != nil {
			return utils.WrapError("Decode semiStatic", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
