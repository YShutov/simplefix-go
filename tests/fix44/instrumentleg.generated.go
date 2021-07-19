package fix44

import (
	"github.com/b2broker/simplefix-go/fix"
)

type InstrumentLeg struct {
	*fix.Component
}

func makeInstrumentLeg() *InstrumentLeg {
	return &InstrumentLeg{fix.NewComponent(
		fix.NewKeyValue(FieldLegSymbol, &fix.String{}),
		fix.NewKeyValue(FieldLegSymbolSfx, &fix.String{}),
		fix.NewKeyValue(FieldLegSecurityID, &fix.String{}),
		fix.NewKeyValue(FieldLegSecurityIDSource, &fix.String{}),
		NewLegSecurityAltIDGrp().Group,
		fix.NewKeyValue(FieldLegProduct, &fix.Int{}),
		fix.NewKeyValue(FieldLegCFICode, &fix.String{}),
		fix.NewKeyValue(FieldLegSecurityType, &fix.String{}),
		fix.NewKeyValue(FieldLegSecuritySubType, &fix.String{}),
		fix.NewKeyValue(FieldLegMaturityMonthYear, &fix.String{}),
		fix.NewKeyValue(FieldLegMaturityDate, &fix.String{}),
		fix.NewKeyValue(FieldLegCouponPaymentDate, &fix.String{}),
		fix.NewKeyValue(FieldLegIssueDate, &fix.String{}),
		fix.NewKeyValue(FieldLegRepoCollateralSecurityType, &fix.Int{}),
		fix.NewKeyValue(FieldLegRepurchaseTerm, &fix.Int{}),
		fix.NewKeyValue(FieldLegRepurchaseRate, &fix.Float{}),
		fix.NewKeyValue(FieldLegFactor, &fix.Float{}),
		fix.NewKeyValue(FieldLegCreditRating, &fix.String{}),
		fix.NewKeyValue(FieldLegInstrRegistry, &fix.String{}),
		fix.NewKeyValue(FieldLegCountryOfIssue, &fix.String{}),
		fix.NewKeyValue(FieldLegStateOrProvinceOfIssue, &fix.String{}),
		fix.NewKeyValue(FieldLegLocaleOfIssue, &fix.String{}),
		fix.NewKeyValue(FieldLegRedemptionDate, &fix.String{}),
		fix.NewKeyValue(FieldLegStrikePrice, &fix.Float{}),
		fix.NewKeyValue(FieldLegStrikeCurrency, &fix.String{}),
		fix.NewKeyValue(FieldLegOptAttribute, &fix.String{}),
		fix.NewKeyValue(FieldLegContractMultiplier, &fix.Float{}),
		fix.NewKeyValue(FieldLegCouponRate, &fix.Float{}),
		fix.NewKeyValue(FieldLegSecurityExchange, &fix.String{}),
		fix.NewKeyValue(FieldLegIssuer, &fix.String{}),
		fix.NewKeyValue(FieldEncodedLegIssuerLen, &fix.Int{}),
		fix.NewKeyValue(FieldEncodedLegIssuer, &fix.String{}),
		fix.NewKeyValue(FieldLegSecurityDesc, &fix.String{}),
		fix.NewKeyValue(FieldEncodedLegSecurityDescLen, &fix.Int{}),
		fix.NewKeyValue(FieldEncodedLegSecurityDesc, &fix.String{}),
		fix.NewKeyValue(FieldLegRatioQty, &fix.Float{}),
		fix.NewKeyValue(FieldLegSide, &fix.String{}),
		fix.NewKeyValue(FieldLegCurrency, &fix.String{}),
		fix.NewKeyValue(FieldLegPool, &fix.String{}),
		fix.NewKeyValue(FieldLegDatedDate, &fix.String{}),
		fix.NewKeyValue(FieldLegContractSettlMonth, &fix.String{}),
		fix.NewKeyValue(FieldLegInterestAccrualDate, &fix.String{}),
	)}
}

func NewInstrumentLeg() *InstrumentLeg {
	return makeInstrumentLeg()
}

func (instrumentLeg *InstrumentLeg) LegSymbol() string {
	kv := instrumentLeg.Get(0)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegSymbol(legSymbol string) *InstrumentLeg {
	kv := instrumentLeg.Get(0).(*fix.KeyValue)
	_ = kv.Load().Set(legSymbol)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegSymbolSfx() string {
	kv := instrumentLeg.Get(1)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegSymbolSfx(legSymbolSfx string) *InstrumentLeg {
	kv := instrumentLeg.Get(1).(*fix.KeyValue)
	_ = kv.Load().Set(legSymbolSfx)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegSecurityID() string {
	kv := instrumentLeg.Get(2)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegSecurityID(legSecurityID string) *InstrumentLeg {
	kv := instrumentLeg.Get(2).(*fix.KeyValue)
	_ = kv.Load().Set(legSecurityID)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegSecurityIDSource() string {
	kv := instrumentLeg.Get(3)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegSecurityIDSource(legSecurityIDSource string) *InstrumentLeg {
	kv := instrumentLeg.Get(3).(*fix.KeyValue)
	_ = kv.Load().Set(legSecurityIDSource)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegSecurityAltIDGrp() *LegSecurityAltIDGrp {
	group := instrumentLeg.Get(4).(*fix.Group)

	return &LegSecurityAltIDGrp{group}
}

func (instrumentLeg *InstrumentLeg) SetLegSecurityAltIDGrp(noLegSecurityAltID *LegSecurityAltIDGrp) *InstrumentLeg {
	instrumentLeg.Set(4, noLegSecurityAltID.Group)

	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegProduct() int {
	kv := instrumentLeg.Get(5)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (instrumentLeg *InstrumentLeg) SetLegProduct(legProduct int) *InstrumentLeg {
	kv := instrumentLeg.Get(5).(*fix.KeyValue)
	_ = kv.Load().Set(legProduct)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegCFICode() string {
	kv := instrumentLeg.Get(6)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegCFICode(legCFICode string) *InstrumentLeg {
	kv := instrumentLeg.Get(6).(*fix.KeyValue)
	_ = kv.Load().Set(legCFICode)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegSecurityType() string {
	kv := instrumentLeg.Get(7)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegSecurityType(legSecurityType string) *InstrumentLeg {
	kv := instrumentLeg.Get(7).(*fix.KeyValue)
	_ = kv.Load().Set(legSecurityType)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegSecuritySubType() string {
	kv := instrumentLeg.Get(8)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegSecuritySubType(legSecuritySubType string) *InstrumentLeg {
	kv := instrumentLeg.Get(8).(*fix.KeyValue)
	_ = kv.Load().Set(legSecuritySubType)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegMaturityMonthYear() string {
	kv := instrumentLeg.Get(9)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegMaturityMonthYear(legMaturityMonthYear string) *InstrumentLeg {
	kv := instrumentLeg.Get(9).(*fix.KeyValue)
	_ = kv.Load().Set(legMaturityMonthYear)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegMaturityDate() string {
	kv := instrumentLeg.Get(10)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegMaturityDate(legMaturityDate string) *InstrumentLeg {
	kv := instrumentLeg.Get(10).(*fix.KeyValue)
	_ = kv.Load().Set(legMaturityDate)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegCouponPaymentDate() string {
	kv := instrumentLeg.Get(11)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegCouponPaymentDate(legCouponPaymentDate string) *InstrumentLeg {
	kv := instrumentLeg.Get(11).(*fix.KeyValue)
	_ = kv.Load().Set(legCouponPaymentDate)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegIssueDate() string {
	kv := instrumentLeg.Get(12)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegIssueDate(legIssueDate string) *InstrumentLeg {
	kv := instrumentLeg.Get(12).(*fix.KeyValue)
	_ = kv.Load().Set(legIssueDate)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegRepoCollateralSecurityType() int {
	kv := instrumentLeg.Get(13)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (instrumentLeg *InstrumentLeg) SetLegRepoCollateralSecurityType(legRepoCollateralSecurityType int) *InstrumentLeg {
	kv := instrumentLeg.Get(13).(*fix.KeyValue)
	_ = kv.Load().Set(legRepoCollateralSecurityType)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegRepurchaseTerm() int {
	kv := instrumentLeg.Get(14)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (instrumentLeg *InstrumentLeg) SetLegRepurchaseTerm(legRepurchaseTerm int) *InstrumentLeg {
	kv := instrumentLeg.Get(14).(*fix.KeyValue)
	_ = kv.Load().Set(legRepurchaseTerm)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegRepurchaseRate() float64 {
	kv := instrumentLeg.Get(15)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(float64)
}

func (instrumentLeg *InstrumentLeg) SetLegRepurchaseRate(legRepurchaseRate float64) *InstrumentLeg {
	kv := instrumentLeg.Get(15).(*fix.KeyValue)
	_ = kv.Load().Set(legRepurchaseRate)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegFactor() float64 {
	kv := instrumentLeg.Get(16)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(float64)
}

func (instrumentLeg *InstrumentLeg) SetLegFactor(legFactor float64) *InstrumentLeg {
	kv := instrumentLeg.Get(16).(*fix.KeyValue)
	_ = kv.Load().Set(legFactor)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegCreditRating() string {
	kv := instrumentLeg.Get(17)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegCreditRating(legCreditRating string) *InstrumentLeg {
	kv := instrumentLeg.Get(17).(*fix.KeyValue)
	_ = kv.Load().Set(legCreditRating)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegInstrRegistry() string {
	kv := instrumentLeg.Get(18)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegInstrRegistry(legInstrRegistry string) *InstrumentLeg {
	kv := instrumentLeg.Get(18).(*fix.KeyValue)
	_ = kv.Load().Set(legInstrRegistry)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegCountryOfIssue() string {
	kv := instrumentLeg.Get(19)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegCountryOfIssue(legCountryOfIssue string) *InstrumentLeg {
	kv := instrumentLeg.Get(19).(*fix.KeyValue)
	_ = kv.Load().Set(legCountryOfIssue)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegStateOrProvinceOfIssue() string {
	kv := instrumentLeg.Get(20)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegStateOrProvinceOfIssue(legStateOrProvinceOfIssue string) *InstrumentLeg {
	kv := instrumentLeg.Get(20).(*fix.KeyValue)
	_ = kv.Load().Set(legStateOrProvinceOfIssue)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegLocaleOfIssue() string {
	kv := instrumentLeg.Get(21)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegLocaleOfIssue(legLocaleOfIssue string) *InstrumentLeg {
	kv := instrumentLeg.Get(21).(*fix.KeyValue)
	_ = kv.Load().Set(legLocaleOfIssue)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegRedemptionDate() string {
	kv := instrumentLeg.Get(22)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegRedemptionDate(legRedemptionDate string) *InstrumentLeg {
	kv := instrumentLeg.Get(22).(*fix.KeyValue)
	_ = kv.Load().Set(legRedemptionDate)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegStrikePrice() float64 {
	kv := instrumentLeg.Get(23)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(float64)
}

func (instrumentLeg *InstrumentLeg) SetLegStrikePrice(legStrikePrice float64) *InstrumentLeg {
	kv := instrumentLeg.Get(23).(*fix.KeyValue)
	_ = kv.Load().Set(legStrikePrice)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegStrikeCurrency() string {
	kv := instrumentLeg.Get(24)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegStrikeCurrency(legStrikeCurrency string) *InstrumentLeg {
	kv := instrumentLeg.Get(24).(*fix.KeyValue)
	_ = kv.Load().Set(legStrikeCurrency)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegOptAttribute() string {
	kv := instrumentLeg.Get(25)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegOptAttribute(legOptAttribute string) *InstrumentLeg {
	kv := instrumentLeg.Get(25).(*fix.KeyValue)
	_ = kv.Load().Set(legOptAttribute)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegContractMultiplier() float64 {
	kv := instrumentLeg.Get(26)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(float64)
}

func (instrumentLeg *InstrumentLeg) SetLegContractMultiplier(legContractMultiplier float64) *InstrumentLeg {
	kv := instrumentLeg.Get(26).(*fix.KeyValue)
	_ = kv.Load().Set(legContractMultiplier)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegCouponRate() float64 {
	kv := instrumentLeg.Get(27)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(float64)
}

func (instrumentLeg *InstrumentLeg) SetLegCouponRate(legCouponRate float64) *InstrumentLeg {
	kv := instrumentLeg.Get(27).(*fix.KeyValue)
	_ = kv.Load().Set(legCouponRate)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegSecurityExchange() string {
	kv := instrumentLeg.Get(28)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegSecurityExchange(legSecurityExchange string) *InstrumentLeg {
	kv := instrumentLeg.Get(28).(*fix.KeyValue)
	_ = kv.Load().Set(legSecurityExchange)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegIssuer() string {
	kv := instrumentLeg.Get(29)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegIssuer(legIssuer string) *InstrumentLeg {
	kv := instrumentLeg.Get(29).(*fix.KeyValue)
	_ = kv.Load().Set(legIssuer)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) EncodedLegIssuerLen() int {
	kv := instrumentLeg.Get(30)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (instrumentLeg *InstrumentLeg) SetEncodedLegIssuerLen(encodedLegIssuerLen int) *InstrumentLeg {
	kv := instrumentLeg.Get(30).(*fix.KeyValue)
	_ = kv.Load().Set(encodedLegIssuerLen)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) EncodedLegIssuer() string {
	kv := instrumentLeg.Get(31)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetEncodedLegIssuer(encodedLegIssuer string) *InstrumentLeg {
	kv := instrumentLeg.Get(31).(*fix.KeyValue)
	_ = kv.Load().Set(encodedLegIssuer)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegSecurityDesc() string {
	kv := instrumentLeg.Get(32)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegSecurityDesc(legSecurityDesc string) *InstrumentLeg {
	kv := instrumentLeg.Get(32).(*fix.KeyValue)
	_ = kv.Load().Set(legSecurityDesc)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) EncodedLegSecurityDescLen() int {
	kv := instrumentLeg.Get(33)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (instrumentLeg *InstrumentLeg) SetEncodedLegSecurityDescLen(encodedLegSecurityDescLen int) *InstrumentLeg {
	kv := instrumentLeg.Get(33).(*fix.KeyValue)
	_ = kv.Load().Set(encodedLegSecurityDescLen)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) EncodedLegSecurityDesc() string {
	kv := instrumentLeg.Get(34)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetEncodedLegSecurityDesc(encodedLegSecurityDesc string) *InstrumentLeg {
	kv := instrumentLeg.Get(34).(*fix.KeyValue)
	_ = kv.Load().Set(encodedLegSecurityDesc)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegRatioQty() float64 {
	kv := instrumentLeg.Get(35)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(float64)
}

func (instrumentLeg *InstrumentLeg) SetLegRatioQty(legRatioQty float64) *InstrumentLeg {
	kv := instrumentLeg.Get(35).(*fix.KeyValue)
	_ = kv.Load().Set(legRatioQty)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegSide() string {
	kv := instrumentLeg.Get(36)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegSide(legSide string) *InstrumentLeg {
	kv := instrumentLeg.Get(36).(*fix.KeyValue)
	_ = kv.Load().Set(legSide)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegCurrency() string {
	kv := instrumentLeg.Get(37)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegCurrency(legCurrency string) *InstrumentLeg {
	kv := instrumentLeg.Get(37).(*fix.KeyValue)
	_ = kv.Load().Set(legCurrency)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegPool() string {
	kv := instrumentLeg.Get(38)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegPool(legPool string) *InstrumentLeg {
	kv := instrumentLeg.Get(38).(*fix.KeyValue)
	_ = kv.Load().Set(legPool)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegDatedDate() string {
	kv := instrumentLeg.Get(39)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegDatedDate(legDatedDate string) *InstrumentLeg {
	kv := instrumentLeg.Get(39).(*fix.KeyValue)
	_ = kv.Load().Set(legDatedDate)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegContractSettlMonth() string {
	kv := instrumentLeg.Get(40)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegContractSettlMonth(legContractSettlMonth string) *InstrumentLeg {
	kv := instrumentLeg.Get(40).(*fix.KeyValue)
	_ = kv.Load().Set(legContractSettlMonth)
	return instrumentLeg
}

func (instrumentLeg *InstrumentLeg) LegInterestAccrualDate() string {
	kv := instrumentLeg.Get(41)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (instrumentLeg *InstrumentLeg) SetLegInterestAccrualDate(legInterestAccrualDate string) *InstrumentLeg {
	kv := instrumentLeg.Get(41).(*fix.KeyValue)
	_ = kv.Load().Set(legInterestAccrualDate)
	return instrumentLeg
}
