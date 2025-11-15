package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SidelinkParametersEUTRA_r16 struct {
	sl_ParametersEUTRA1_r16            *[]byte                 `optional`
	sl_ParametersEUTRA2_r16            *[]byte                 `optional`
	sl_ParametersEUTRA3_r16            *[]byte                 `optional`
	supportedBandListSidelinkEUTRA_r16 []BandSidelinkEUTRA_r16 `lb:1,ub:maxBandsEUTRA,optional`
}

func (ie *SidelinkParametersEUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_ParametersEUTRA1_r16 != nil, ie.sl_ParametersEUTRA2_r16 != nil, ie.sl_ParametersEUTRA3_r16 != nil, len(ie.supportedBandListSidelinkEUTRA_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_ParametersEUTRA1_r16 != nil {
		if err = w.WriteOctetString(*ie.sl_ParametersEUTRA1_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sl_ParametersEUTRA1_r16", err)
		}
	}
	if ie.sl_ParametersEUTRA2_r16 != nil {
		if err = w.WriteOctetString(*ie.sl_ParametersEUTRA2_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sl_ParametersEUTRA2_r16", err)
		}
	}
	if ie.sl_ParametersEUTRA3_r16 != nil {
		if err = w.WriteOctetString(*ie.sl_ParametersEUTRA3_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sl_ParametersEUTRA3_r16", err)
		}
	}
	if len(ie.supportedBandListSidelinkEUTRA_r16) > 0 {
		tmp_supportedBandListSidelinkEUTRA_r16 := utils.NewSequence[*BandSidelinkEUTRA_r16]([]*BandSidelinkEUTRA_r16{}, uper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
		for _, i := range ie.supportedBandListSidelinkEUTRA_r16 {
			tmp_supportedBandListSidelinkEUTRA_r16.Value = append(tmp_supportedBandListSidelinkEUTRA_r16.Value, &i)
		}
		if err = tmp_supportedBandListSidelinkEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandListSidelinkEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *SidelinkParametersEUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_ParametersEUTRA1_r16Present bool
	if sl_ParametersEUTRA1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ParametersEUTRA2_r16Present bool
	if sl_ParametersEUTRA2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ParametersEUTRA3_r16Present bool
	if sl_ParametersEUTRA3_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandListSidelinkEUTRA_r16Present bool
	if supportedBandListSidelinkEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_ParametersEUTRA1_r16Present {
		var tmp_os_sl_ParametersEUTRA1_r16 []byte
		if tmp_os_sl_ParametersEUTRA1_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sl_ParametersEUTRA1_r16", err)
		}
		ie.sl_ParametersEUTRA1_r16 = &tmp_os_sl_ParametersEUTRA1_r16
	}
	if sl_ParametersEUTRA2_r16Present {
		var tmp_os_sl_ParametersEUTRA2_r16 []byte
		if tmp_os_sl_ParametersEUTRA2_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sl_ParametersEUTRA2_r16", err)
		}
		ie.sl_ParametersEUTRA2_r16 = &tmp_os_sl_ParametersEUTRA2_r16
	}
	if sl_ParametersEUTRA3_r16Present {
		var tmp_os_sl_ParametersEUTRA3_r16 []byte
		if tmp_os_sl_ParametersEUTRA3_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sl_ParametersEUTRA3_r16", err)
		}
		ie.sl_ParametersEUTRA3_r16 = &tmp_os_sl_ParametersEUTRA3_r16
	}
	if supportedBandListSidelinkEUTRA_r16Present {
		tmp_supportedBandListSidelinkEUTRA_r16 := utils.NewSequence[*BandSidelinkEUTRA_r16]([]*BandSidelinkEUTRA_r16{}, uper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
		fn_supportedBandListSidelinkEUTRA_r16 := func() *BandSidelinkEUTRA_r16 {
			return new(BandSidelinkEUTRA_r16)
		}
		if err = tmp_supportedBandListSidelinkEUTRA_r16.Decode(r, fn_supportedBandListSidelinkEUTRA_r16); err != nil {
			return utils.WrapError("Decode supportedBandListSidelinkEUTRA_r16", err)
		}
		ie.supportedBandListSidelinkEUTRA_r16 = []BandSidelinkEUTRA_r16{}
		for _, i := range tmp_supportedBandListSidelinkEUTRA_r16.Value {
			ie.supportedBandListSidelinkEUTRA_r16 = append(ie.supportedBandListSidelinkEUTRA_r16, *i)
		}
	}
	return nil
}
