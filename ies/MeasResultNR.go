package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultNR struct {
	physCellId         *PhysCellId                      `optional`
	measResult         *MeasResultNR_measResult         `optional`
	cgi_Info           *CGI_InfoNR                      `optional,ext-1`
	choCandidate_r17   *MeasResultNR_choCandidate_r17   `optional,ext-2`
	choConfig_r17      []CondTriggerConfig_r16          `lb:1,ub:2,optional,ext-2`
	triggeredEvent_r17 *MeasResultNR_triggeredEvent_r17 `optional,ext-2`
}

func (ie *MeasResultNR) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.cgi_Info != nil || ie.choCandidate_r17 != nil || len(ie.choConfig_r17) > 0 || ie.triggeredEvent_r17 != nil
	preambleBits := []bool{hasExtensions, ie.physCellId != nil, ie.measResult != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.physCellId != nil {
		if err = ie.physCellId.Encode(w); err != nil {
			return utils.WrapError("Encode physCellId", err)
		}
	}
	if ie.measResult != nil {
		if err = ie.measResult.Encode(w); err != nil {
			return utils.WrapError("Encode measResult", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.cgi_Info != nil, ie.choCandidate_r17 != nil || len(ie.choConfig_r17) > 0 || ie.triggeredEvent_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasResultNR", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.cgi_Info != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode cgi_Info optional
			if ie.cgi_Info != nil {
				if err = ie.cgi_Info.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cgi_Info", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.choCandidate_r17 != nil, len(ie.choConfig_r17) > 0, ie.triggeredEvent_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode choCandidate_r17 optional
			if ie.choCandidate_r17 != nil {
				if err = ie.choCandidate_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode choCandidate_r17", err)
				}
			}
			// encode choConfig_r17 optional
			if len(ie.choConfig_r17) > 0 {
				tmp_choConfig_r17 := utils.NewSequence[*CondTriggerConfig_r16]([]*CondTriggerConfig_r16{}, uper.Constraint{Lb: 1, Ub: 2}, false)
				for _, i := range ie.choConfig_r17 {
					tmp_choConfig_r17.Value = append(tmp_choConfig_r17.Value, &i)
				}
				if err = tmp_choConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode choConfig_r17", err)
				}
			}
			// encode triggeredEvent_r17 optional
			if ie.triggeredEvent_r17 != nil {
				if err = ie.triggeredEvent_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode triggeredEvent_r17", err)
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

func (ie *MeasResultNR) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var physCellIdPresent bool
	if physCellIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultPresent bool
	if measResultPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if physCellIdPresent {
		ie.physCellId = new(PhysCellId)
		if err = ie.physCellId.Decode(r); err != nil {
			return utils.WrapError("Decode physCellId", err)
		}
	}
	if measResultPresent {
		ie.measResult = new(MeasResultNR_measResult)
		if err = ie.measResult.Decode(r); err != nil {
			return utils.WrapError("Decode measResult", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			cgi_InfoPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode cgi_Info optional
			if cgi_InfoPresent {
				ie.cgi_Info = new(CGI_InfoNR)
				if err = ie.cgi_Info.Decode(extReader); err != nil {
					return utils.WrapError("Decode cgi_Info", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			choCandidate_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			choConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			triggeredEvent_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode choCandidate_r17 optional
			if choCandidate_r17Present {
				ie.choCandidate_r17 = new(MeasResultNR_choCandidate_r17)
				if err = ie.choCandidate_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode choCandidate_r17", err)
				}
			}
			// decode choConfig_r17 optional
			if choConfig_r17Present {
				tmp_choConfig_r17 := utils.NewSequence[*CondTriggerConfig_r16]([]*CondTriggerConfig_r16{}, uper.Constraint{Lb: 1, Ub: 2}, false)
				fn_choConfig_r17 := func() *CondTriggerConfig_r16 {
					return new(CondTriggerConfig_r16)
				}
				if err = tmp_choConfig_r17.Decode(extReader, fn_choConfig_r17); err != nil {
					return utils.WrapError("Decode choConfig_r17", err)
				}
				ie.choConfig_r17 = []CondTriggerConfig_r16{}
				for _, i := range tmp_choConfig_r17.Value {
					ie.choConfig_r17 = append(ie.choConfig_r17, *i)
				}
			}
			// decode triggeredEvent_r17 optional
			if triggeredEvent_r17Present {
				ie.triggeredEvent_r17 = new(MeasResultNR_triggeredEvent_r17)
				if err = ie.triggeredEvent_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode triggeredEvent_r17", err)
				}
			}
		}
	}
	return nil
}
