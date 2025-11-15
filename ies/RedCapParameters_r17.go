package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RedCapParameters_r17 struct {
	supportOfRedCap_r17       *RedCapParameters_r17_supportOfRedCap_r17       `optional`
	supportOf16DRB_RedCap_r17 *RedCapParameters_r17_supportOf16DRB_RedCap_r17 `optional`
}

func (ie *RedCapParameters_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportOfRedCap_r17 != nil, ie.supportOf16DRB_RedCap_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportOfRedCap_r17 != nil {
		if err = ie.supportOfRedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode supportOfRedCap_r17", err)
		}
	}
	if ie.supportOf16DRB_RedCap_r17 != nil {
		if err = ie.supportOf16DRB_RedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode supportOf16DRB_RedCap_r17", err)
		}
	}
	return nil
}

func (ie *RedCapParameters_r17) Decode(r *uper.UperReader) error {
	var err error
	var supportOfRedCap_r17Present bool
	if supportOfRedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportOf16DRB_RedCap_r17Present bool
	if supportOf16DRB_RedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportOfRedCap_r17Present {
		ie.supportOfRedCap_r17 = new(RedCapParameters_r17_supportOfRedCap_r17)
		if err = ie.supportOfRedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode supportOfRedCap_r17", err)
		}
	}
	if supportOf16DRB_RedCap_r17Present {
		ie.supportOf16DRB_RedCap_r17 = new(RedCapParameters_r17_supportOf16DRB_RedCap_r17)
		if err = ie.supportOf16DRB_RedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode supportOf16DRB_RedCap_r17", err)
		}
	}
	return nil
}
