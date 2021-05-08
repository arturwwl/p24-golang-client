package paymentMethod

type PaymentMethodList struct {
	Data         []*PaymentMethod
	ResponseCode uint64
}

type PaymentMethod struct {
	Name              string
	ID                uint64
	Status            bool
	ImgUrl            string
	MobileImgUrl      string
	Mobile            bool
	AvailabilityHours *AvailabilityHours
}

type AvailabilityHours struct {
	MondayToFriday *string //sometimes it is true, sometimes it is 0-24
	Saturday       *string
	Sunday         *string
}
