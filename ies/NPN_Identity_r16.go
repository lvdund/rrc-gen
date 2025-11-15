package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NPN_Identity_r16_Choice_nothing uint64 = iota
	NPN_Identity_r16_Choice_pni_npn_r16
	NPN_Identity_r16_Choice_snpn_r16
)

type NPN_Identity_r16 struct {
	Choice      uint64
	pni_npn_r16 *NPN_Identity_r16_pni_npn_r16
	snpn_r16    *NPN_Identity_r16_snpn_r16
}

func (ie *NPN_Identity_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NPN_Identity_r16_Choice_pni_npn_r16:
		if err = ie.pni_npn_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode pni_npn_r16", err)
		}
	case NPN_Identity_r16_Choice_snpn_r16:
		if err = ie.snpn_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode snpn_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *NPN_Identity_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NPN_Identity_r16_Choice_pni_npn_r16:
		ie.pni_npn_r16 = new(NPN_Identity_r16_pni_npn_r16)
		if err = ie.pni_npn_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pni_npn_r16", err)
		}
	case NPN_Identity_r16_Choice_snpn_r16:
		ie.snpn_r16 = new(NPN_Identity_r16_snpn_r16)
		if err = ie.snpn_r16.Decode(r); err != nil {
			return utils.WrapError("Decode snpn_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
