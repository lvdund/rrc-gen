package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MulticastConfig_r17 struct {
	pdsch_HARQ_ACK_CodebookListMulticast_r17 *PDSCH_HARQ_ACK_CodebookList_r16                       `optional,setuprelease`
	type1_Codebook_GenerationMode_r17        *MulticastConfig_r17_type1_Codebook_GenerationMode_r17 `optional`
}

func (ie *MulticastConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pdsch_HARQ_ACK_CodebookListMulticast_r17 != nil, ie.type1_Codebook_GenerationMode_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pdsch_HARQ_ACK_CodebookListMulticast_r17 != nil {
		tmp_pdsch_HARQ_ACK_CodebookListMulticast_r17 := utils.SetupRelease[*PDSCH_HARQ_ACK_CodebookList_r16]{
			Setup: ie.pdsch_HARQ_ACK_CodebookListMulticast_r17,
		}
		if err = tmp_pdsch_HARQ_ACK_CodebookListMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_HARQ_ACK_CodebookListMulticast_r17", err)
		}
	}
	if ie.type1_Codebook_GenerationMode_r17 != nil {
		if err = ie.type1_Codebook_GenerationMode_r17.Encode(w); err != nil {
			return utils.WrapError("Encode type1_Codebook_GenerationMode_r17", err)
		}
	}
	return nil
}

func (ie *MulticastConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var pdsch_HARQ_ACK_CodebookListMulticast_r17Present bool
	if pdsch_HARQ_ACK_CodebookListMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1_Codebook_GenerationMode_r17Present bool
	if type1_Codebook_GenerationMode_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pdsch_HARQ_ACK_CodebookListMulticast_r17Present {
		tmp_pdsch_HARQ_ACK_CodebookListMulticast_r17 := utils.SetupRelease[*PDSCH_HARQ_ACK_CodebookList_r16]{}
		if err = tmp_pdsch_HARQ_ACK_CodebookListMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_HARQ_ACK_CodebookListMulticast_r17", err)
		}
		ie.pdsch_HARQ_ACK_CodebookListMulticast_r17 = tmp_pdsch_HARQ_ACK_CodebookListMulticast_r17.Setup
	}
	if type1_Codebook_GenerationMode_r17Present {
		ie.type1_Codebook_GenerationMode_r17 = new(MulticastConfig_r17_type1_Codebook_GenerationMode_r17)
		if err = ie.type1_Codebook_GenerationMode_r17.Decode(r); err != nil {
			return utils.WrapError("Decode type1_Codebook_GenerationMode_r17", err)
		}
	}
	return nil
}
