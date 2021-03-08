package model

import "time"

type AWBDetailPartner struct {
	ID              string
	AWBNumber       string
	OriginCode      string
	DestinationCode string
	ServiceTypeID   string
	SLAMin          int
	SLAMax          int
	Weight          int
	PartnerID       uint
	Sender          string
	SenderAddress   string
	Receiver        string
	ReceiverAddress string
	PricePerKg      int
	TotalPrice      int
	OrderID         string
	// note public or status detail
	PODReceiver string
	// date when package has received and status delivered
	PODReceiverTime time.Time
	// date when package has picked ip by courier
	SendDate time.Time
	// last status, e.g: delivered, loss, etc.
	LastStatus string
	// city code, eg: bogor = BOO, etc.
	RepresentativeCode string
	IsHandover         bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
	AWBHistories       []*AWBHistory
	Partner            *Partner
	ServiceType        *ServiceType
}

type AWBHistory struct {
	ID           string    `bson:"_id"`
	AWBID        string    `bson:"awbNumber"`
	CourierID    int64     `bson:"courierID"`
	CourierName  string    `bson:"courierName`
	CourierNIK   string    `bson:"courierNIK"`
	Status       string    `bson:"status"`
	StatusTime   time.Time `bson:"statusTime"`
	StatusDetail string    `bson:"statusDetail"`
	BranchName   string    `bson:"branch_name"`
	City         string    `bson:"city"`
	Latitude     float64   `bson:"latitude"`
	Longitude    float64   `bson:"longitude"`
	CreatedAt    time.Time `bson:"createdAt"`
}

type Partner struct {
	PartnerID  string    `bson:"_id,omitempty"`
	ID         uint      `bson:"masterdataID,omitempty"`
	Name       string    `bson:"name"`
	IsActive   bool      `bson:"isActive"`
	IsDeleted  bool      `bson:"isDeleted"`
	Code       string    `bson:"code,omitempty"`
	CategoryID uint      `bson:"categoryId,omitempty"`
	CreatedAt  time.Time `bson:"createdAt"`
	UpdatedAt  time.Time `bson:"updatedAt"`
}

type ServiceType struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Code        string    `json:"code" bson:"code"`
	Description string    `json:"description" bson:"description"`
	IsActive    bool      `json:"is_active" bson:"isActive"`
	IsDeleted   bool      `json:"is_deleted" bson:"isDeleted"`
	CreatedAt   time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updatedAt"`
}
