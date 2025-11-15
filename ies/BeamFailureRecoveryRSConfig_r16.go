package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BeamFailureRecoveryRSConfig_r16 struct {
	rsrp_ThresholdBFR_r16     *RSRP_Range           `optional`
	candidateBeamRS_List_r16  []CandidateBeamRS_r16 `lb:1,ub:maxNrofCandidateBeams_r16,optional`
	candidateBeamRS_List2_r17 []CandidateBeamRS_r16 `lb:1,ub:maxNrofCandidateBeams_r16,optional,ext-1`
}

func (ie *BeamFailureRecoveryRSConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.candidateBeamRS_List2_r17) > 0
	preambleBits := []bool{hasExtensions, ie.rsrp_ThresholdBFR_r16 != nil, len(ie.candidateBeamRS_List_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rsrp_ThresholdBFR_r16 != nil {
		if err = ie.rsrp_ThresholdBFR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rsrp_ThresholdBFR_r16", err)
		}
	}
	if len(ie.candidateBeamRS_List_r16) > 0 {
		tmp_candidateBeamRS_List_r16 := utils.NewSequence[*CandidateBeamRS_r16]([]*CandidateBeamRS_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCandidateBeams_r16}, false)
		for _, i := range ie.candidateBeamRS_List_r16 {
			tmp_candidateBeamRS_List_r16.Value = append(tmp_candidateBeamRS_List_r16.Value, &i)
		}
		if err = tmp_candidateBeamRS_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode candidateBeamRS_List_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{len(ie.candidateBeamRS_List2_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BeamFailureRecoveryRSConfig_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.candidateBeamRS_List2_r17) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode candidateBeamRS_List2_r17 optional
			if len(ie.candidateBeamRS_List2_r17) > 0 {
				tmp_candidateBeamRS_List2_r17 := utils.NewSequence[*CandidateBeamRS_r16]([]*CandidateBeamRS_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCandidateBeams_r16}, false)
				for _, i := range ie.candidateBeamRS_List2_r17 {
					tmp_candidateBeamRS_List2_r17.Value = append(tmp_candidateBeamRS_List2_r17.Value, &i)
				}
				if err = tmp_candidateBeamRS_List2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode candidateBeamRS_List2_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *BeamFailureRecoveryRSConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var rsrp_ThresholdBFR_r16Present bool
	if rsrp_ThresholdBFR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var candidateBeamRS_List_r16Present bool
	if candidateBeamRS_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if rsrp_ThresholdBFR_r16Present {
		ie.rsrp_ThresholdBFR_r16 = new(RSRP_Range)
		if err = ie.rsrp_ThresholdBFR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rsrp_ThresholdBFR_r16", err)
		}
	}
	if candidateBeamRS_List_r16Present {
		tmp_candidateBeamRS_List_r16 := utils.NewSequence[*CandidateBeamRS_r16]([]*CandidateBeamRS_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCandidateBeams_r16}, false)
		fn_candidateBeamRS_List_r16 := func() *CandidateBeamRS_r16 {
			return new(CandidateBeamRS_r16)
		}
		if err = tmp_candidateBeamRS_List_r16.Decode(r, fn_candidateBeamRS_List_r16); err != nil {
			return utils.WrapError("Decode candidateBeamRS_List_r16", err)
		}
		ie.candidateBeamRS_List_r16 = []CandidateBeamRS_r16{}
		for _, i := range tmp_candidateBeamRS_List_r16.Value {
			ie.candidateBeamRS_List_r16 = append(ie.candidateBeamRS_List_r16, *i)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			candidateBeamRS_List2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode candidateBeamRS_List2_r17 optional
			if candidateBeamRS_List2_r17Present {
				tmp_candidateBeamRS_List2_r17 := utils.NewSequence[*CandidateBeamRS_r16]([]*CandidateBeamRS_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCandidateBeams_r16}, false)
				fn_candidateBeamRS_List2_r17 := func() *CandidateBeamRS_r16 {
					return new(CandidateBeamRS_r16)
				}
				if err = tmp_candidateBeamRS_List2_r17.Decode(extReader, fn_candidateBeamRS_List2_r17); err != nil {
					return utils.WrapError("Decode candidateBeamRS_List2_r17", err)
				}
				ie.candidateBeamRS_List2_r17 = []CandidateBeamRS_r16{}
				for _, i := range tmp_candidateBeamRS_List2_r17.Value {
					ie.candidateBeamRS_List2_r17 = append(ie.candidateBeamRS_List2_r17, *i)
				}
			}
		}
	}
	return nil
}
