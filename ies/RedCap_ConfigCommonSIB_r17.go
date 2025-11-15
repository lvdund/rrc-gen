package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RedCap_ConfigCommonSIB_r17 struct {
	halfDuplexRedCapAllowed_r17 *RedCap_ConfigCommonSIB_r17_halfDuplexRedCapAllowed_r17 `optional`
	cellBarredRedCap_r17        *RedCap_ConfigCommonSIB_r17_cellBarredRedCap_r17        `optional`
}

func (ie *RedCap_ConfigCommonSIB_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.halfDuplexRedCapAllowed_r17 != nil, ie.cellBarredRedCap_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.halfDuplexRedCapAllowed_r17 != nil {
		if err = ie.halfDuplexRedCapAllowed_r17.Encode(w); err != nil {
			return utils.WrapError("Encode halfDuplexRedCapAllowed_r17", err)
		}
	}
	if ie.cellBarredRedCap_r17 != nil {
		if err = ie.cellBarredRedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cellBarredRedCap_r17", err)
		}
	}
	return nil
}

func (ie *RedCap_ConfigCommonSIB_r17) Decode(r *uper.UperReader) error {
	var err error
	var halfDuplexRedCapAllowed_r17Present bool
	if halfDuplexRedCapAllowed_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cellBarredRedCap_r17Present bool
	if cellBarredRedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if halfDuplexRedCapAllowed_r17Present {
		ie.halfDuplexRedCapAllowed_r17 = new(RedCap_ConfigCommonSIB_r17_halfDuplexRedCapAllowed_r17)
		if err = ie.halfDuplexRedCapAllowed_r17.Decode(r); err != nil {
			return utils.WrapError("Decode halfDuplexRedCapAllowed_r17", err)
		}
	}
	if cellBarredRedCap_r17Present {
		ie.cellBarredRedCap_r17 = new(RedCap_ConfigCommonSIB_r17_cellBarredRedCap_r17)
		if err = ie.cellBarredRedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cellBarredRedCap_r17", err)
		}
	}
	return nil
}
