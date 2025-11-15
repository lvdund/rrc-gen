package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResumeRequest1_IEs struct {
	resumeIdentity I_RNTI_Value   `madatory`
	resumeMAC_I    uper.BitString `lb:16,ub:16,madatory`
	resumeCause    ResumeCause    `madatory`
	spare          uper.BitString `lb:1,ub:1,madatory`
}

func (ie *RRCResumeRequest1_IEs) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.resumeIdentity.Encode(w); err != nil {
		return utils.WrapError("Encode resumeIdentity", err)
	}
	if err = w.WriteBitString(ie.resumeMAC_I.Bytes, uint(ie.resumeMAC_I.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteBitString resumeMAC_I", err)
	}
	if err = ie.resumeCause.Encode(w); err != nil {
		return utils.WrapError("Encode resumeCause", err)
	}
	if err = w.WriteBitString(ie.spare.Bytes, uint(ie.spare.NumBits), &uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteBitString spare", err)
	}
	return nil
}

func (ie *RRCResumeRequest1_IEs) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.resumeIdentity.Decode(r); err != nil {
		return utils.WrapError("Decode resumeIdentity", err)
	}
	var tmp_bs_resumeMAC_I []byte
	var n_resumeMAC_I uint
	if tmp_bs_resumeMAC_I, n_resumeMAC_I, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadBitString resumeMAC_I", err)
	}
	ie.resumeMAC_I = uper.BitString{
		Bytes:   tmp_bs_resumeMAC_I,
		NumBits: uint64(n_resumeMAC_I),
	}
	if err = ie.resumeCause.Decode(r); err != nil {
		return utils.WrapError("Decode resumeCause", err)
	}
	var tmp_bs_spare []byte
	var n_spare uint
	if tmp_bs_spare, n_spare, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadBitString spare", err)
	}
	ie.spare = uper.BitString{
		Bytes:   tmp_bs_spare,
		NumBits: uint64(n_spare),
	}
	return nil
}
