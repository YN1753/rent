package model

import "time"

type RentOrder struct {
	Id         int        //数据库ID
	OrderId    int        //订单ID
	Price      int        //租金
	OwnerId    string     //房主ID
	TenantId   string     //租客ID
	HouseId    string     //房源ID
	Deposit    int        //押金
	LeaseStart time.Time  //租期开始
	LeaseEnd   time.Time  //租期结束
	Status     int        //0待确认1已确认2已签约3已入住4已退租5取消
	PayTime    *time.Time //支付时间,为什么是指针类型呢？因为当未支付的情况下，该值会有一个默认值，如果是指针类型的话那么在数据库里面就是NULL否则是一个时间
	CreatedAt  time.Time  //订单创建时间
}
