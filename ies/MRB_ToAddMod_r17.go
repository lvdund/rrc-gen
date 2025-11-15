package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRB_ToAddMod_r17 struct {
	mbs_SessionId_r17   *TMGI_r17                             `optional`
	mrb_Identity_r17    MRB_Identity_r17                      `madatory`
	mrb_IdentityNew_r17 *MRB_Identity_r17                     `optional`
	reestablishPDCP_r17 *MRB_ToAddMod_r17_reestablishPDCP_r17 `optional`
	recoverPDCP_r17     *MRB_ToAddMod_r17_recoverPDCP_r17     `optional`
	pdcp_Config_r17     *PDCP_Config                          `optional`
}

func (ie *MRB_ToAddMod_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mbs_SessionId_r17 != nil, ie.mrb_IdentityNew_r17 != nil, ie.reestablishPDCP_r17 != nil, ie.recoverPDCP_r17 != nil, ie.pdcp_Config_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mbs_SessionId_r17 != nil {
		if err = ie.mbs_SessionId_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mbs_SessionId_r17", err)
		}
	}
	if err = ie.mrb_Identity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mrb_Identity_r17", err)
	}
	if ie.mrb_IdentityNew_r17 != nil {
		if err = ie.mrb_IdentityNew_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mrb_IdentityNew_r17", err)
		}
	}
	if ie.reestablishPDCP_r17 != nil {
		if err = ie.reestablishPDCP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode reestablishPDCP_r17", err)
		}
	}
	if ie.recoverPDCP_r17 != nil {
		if err = ie.recoverPDCP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode recoverPDCP_r17", err)
		}
	}
	if ie.pdcp_Config_r17 != nil {
		if err = ie.pdcp_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdcp_Config_r17", err)
		}
	}
	return nil
}

func (ie *MRB_ToAddMod_r17) Decode(r *uper.UperReader) error {
	var err error
	var mbs_SessionId_r17Present bool
	if mbs_SessionId_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mrb_IdentityNew_r17Present bool
	if mrb_IdentityNew_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reestablishPDCP_r17Present bool
	if reestablishPDCP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var recoverPDCP_r17Present bool
	if recoverPDCP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcp_Config_r17Present bool
	if pdcp_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if mbs_SessionId_r17Present {
		ie.mbs_SessionId_r17 = new(TMGI_r17)
		if err = ie.mbs_SessionId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mbs_SessionId_r17", err)
		}
	}
	if err = ie.mrb_Identity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mrb_Identity_r17", err)
	}
	if mrb_IdentityNew_r17Present {
		ie.mrb_IdentityNew_r17 = new(MRB_Identity_r17)
		if err = ie.mrb_IdentityNew_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mrb_IdentityNew_r17", err)
		}
	}
	if reestablishPDCP_r17Present {
		ie.reestablishPDCP_r17 = new(MRB_ToAddMod_r17_reestablishPDCP_r17)
		if err = ie.reestablishPDCP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode reestablishPDCP_r17", err)
		}
	}
	if recoverPDCP_r17Present {
		ie.recoverPDCP_r17 = new(MRB_ToAddMod_r17_recoverPDCP_r17)
		if err = ie.recoverPDCP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode recoverPDCP_r17", err)
		}
	}
	if pdcp_Config_r17Present {
		ie.pdcp_Config_r17 = new(PDCP_Config)
		if err = ie.pdcp_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdcp_Config_r17", err)
		}
	}
	return nil
}
