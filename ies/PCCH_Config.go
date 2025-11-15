package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PCCH_Config struct {
	defaultPagingCycle                          PagingCycle                        `madatory`
	nAndPagingFrameOffset                       PCCH_Config_nAndPagingFrameOffset  `lb:0,ub:1,madatory`
	ns                                          PCCH_Config_ns                     `madatory`
	firstPDCCH_MonitoringOccasionOfPO           []int64                            `lb:1,ub:maxPO_perPF,e_lb:0,e_ub:139,optional`
	nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 *int64                             `lb:2,ub:4,optional,ext-1`
	ranPagingInIdlePO_r17                       *PCCH_Config_ranPagingInIdlePO_r17 `optional,ext-2`
	firstPDCCH_MonitoringOccasionOfPO_v1710     []int64                            `lb:1,ub:maxPO_perPF,e_lb:0,e_ub:35839,optional,ext-2`
}

func (ie *PCCH_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 != nil || ie.ranPagingInIdlePO_r17 != nil || len(ie.firstPDCCH_MonitoringOccasionOfPO_v1710) > 0
	preambleBits := []bool{hasExtensions, len(ie.firstPDCCH_MonitoringOccasionOfPO) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.defaultPagingCycle.Encode(w); err != nil {
		return utils.WrapError("Encode defaultPagingCycle", err)
	}
	if err = ie.nAndPagingFrameOffset.Encode(w); err != nil {
		return utils.WrapError("Encode nAndPagingFrameOffset", err)
	}
	if err = ie.ns.Encode(w); err != nil {
		return utils.WrapError("Encode ns", err)
	}
	if len(ie.firstPDCCH_MonitoringOccasionOfPO) > 0 {
		tmp_firstPDCCH_MonitoringOccasionOfPO := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxPO_perPF}, false)
		for _, i := range ie.firstPDCCH_MonitoringOccasionOfPO {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 139}, false)
			tmp_firstPDCCH_MonitoringOccasionOfPO.Value = append(tmp_firstPDCCH_MonitoringOccasionOfPO.Value, &tmp_ie)
		}
		if err = tmp_firstPDCCH_MonitoringOccasionOfPO.Encode(w); err != nil {
			return utils.WrapError("Encode firstPDCCH_MonitoringOccasionOfPO", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 != nil, ie.ranPagingInIdlePO_r17 != nil || len(ie.firstPDCCH_MonitoringOccasionOfPO_v1710) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PCCH_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 optional
			if ie.nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 != nil {
				if err = extWriter.WriteInteger(*ie.nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16, &uper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16", err)
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
			optionals_ext_2 := []bool{ie.ranPagingInIdlePO_r17 != nil, len(ie.firstPDCCH_MonitoringOccasionOfPO_v1710) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ranPagingInIdlePO_r17 optional
			if ie.ranPagingInIdlePO_r17 != nil {
				if err = ie.ranPagingInIdlePO_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ranPagingInIdlePO_r17", err)
				}
			}
			// encode firstPDCCH_MonitoringOccasionOfPO_v1710 optional
			if len(ie.firstPDCCH_MonitoringOccasionOfPO_v1710) > 0 {
				tmp_firstPDCCH_MonitoringOccasionOfPO_v1710 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxPO_perPF}, false)
				for _, i := range ie.firstPDCCH_MonitoringOccasionOfPO_v1710 {
					tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 35839}, false)
					tmp_firstPDCCH_MonitoringOccasionOfPO_v1710.Value = append(tmp_firstPDCCH_MonitoringOccasionOfPO_v1710.Value, &tmp_ie)
				}
				if err = tmp_firstPDCCH_MonitoringOccasionOfPO_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode firstPDCCH_MonitoringOccasionOfPO_v1710", err)
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

func (ie *PCCH_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var firstPDCCH_MonitoringOccasionOfPOPresent bool
	if firstPDCCH_MonitoringOccasionOfPOPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.defaultPagingCycle.Decode(r); err != nil {
		return utils.WrapError("Decode defaultPagingCycle", err)
	}
	if err = ie.nAndPagingFrameOffset.Decode(r); err != nil {
		return utils.WrapError("Decode nAndPagingFrameOffset", err)
	}
	if err = ie.ns.Decode(r); err != nil {
		return utils.WrapError("Decode ns", err)
	}
	if firstPDCCH_MonitoringOccasionOfPOPresent {
		tmp_firstPDCCH_MonitoringOccasionOfPO := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxPO_perPF}, false)
		fn_firstPDCCH_MonitoringOccasionOfPO := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 139}, false)
			return &ie
		}
		if err = tmp_firstPDCCH_MonitoringOccasionOfPO.Decode(r, fn_firstPDCCH_MonitoringOccasionOfPO); err != nil {
			return utils.WrapError("Decode firstPDCCH_MonitoringOccasionOfPO", err)
		}
		ie.firstPDCCH_MonitoringOccasionOfPO = []int64{}
		for _, i := range tmp_firstPDCCH_MonitoringOccasionOfPO.Value {
			ie.firstPDCCH_MonitoringOccasionOfPO = append(ie.firstPDCCH_MonitoringOccasionOfPO, int64(i.Value))
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

			nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 optional
			if nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16Present {
				var tmp_int_nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 int64
				if tmp_int_nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16", err)
				}
				ie.nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 = &tmp_int_nrofPDCCH_MonitoringOccasionPerSSB_InPO_r16
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ranPagingInIdlePO_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			firstPDCCH_MonitoringOccasionOfPO_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ranPagingInIdlePO_r17 optional
			if ranPagingInIdlePO_r17Present {
				ie.ranPagingInIdlePO_r17 = new(PCCH_Config_ranPagingInIdlePO_r17)
				if err = ie.ranPagingInIdlePO_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ranPagingInIdlePO_r17", err)
				}
			}
			// decode firstPDCCH_MonitoringOccasionOfPO_v1710 optional
			if firstPDCCH_MonitoringOccasionOfPO_v1710Present {
				tmp_firstPDCCH_MonitoringOccasionOfPO_v1710 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxPO_perPF}, false)
				fn_firstPDCCH_MonitoringOccasionOfPO_v1710 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 35839}, false)
					return &ie
				}
				if err = tmp_firstPDCCH_MonitoringOccasionOfPO_v1710.Decode(extReader, fn_firstPDCCH_MonitoringOccasionOfPO_v1710); err != nil {
					return utils.WrapError("Decode firstPDCCH_MonitoringOccasionOfPO_v1710", err)
				}
				ie.firstPDCCH_MonitoringOccasionOfPO_v1710 = []int64{}
				for _, i := range tmp_firstPDCCH_MonitoringOccasionOfPO_v1710.Value {
					ie.firstPDCCH_MonitoringOccasionOfPO_v1710 = append(ie.firstPDCCH_MonitoringOccasionOfPO_v1710, int64(i.Value))
				}
			}
		}
	}
	return nil
}
